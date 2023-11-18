/*
	Copyright NetFoundry Inc.

	Licensed under the Apache License, Version 2.0 (the "License");
	you may not use this file except in compliance with the License.
	You may obtain a copy of the License at

	https://www.apache.org/licenses/LICENSE-2.0

	Unless required by applicable law or agreed to in writing, software
	distributed under the License is distributed on an "AS IS" BASIS,
	WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
	See the License for the specific language governing permissions and
	limitations under the License.
*/

package handler_ctrl

import (
	"fmt"
	"github.com/michaelquigley/pfxlog"
	"github.com/openziti/channel/v2"
	"github.com/openziti/channel/v2/protobufs"
	"github.com/openziti/foundation/v2/concurrenz"
	"github.com/openziti/ziti/common/pb/ctrl_pb"
	"google.golang.org/protobuf/proto"
	"io"
	"net"
	"time"
)

type pipeRegistry struct {
	pipes concurrenz.CopyOnWriteMap[uint32, *ctrlChanPipe]
}

type ctrlPipeHandler struct {
	registry *pipeRegistry
	ch       channel.Channel
}

func newCtrlPipeHandler(registry *pipeRegistry, ch channel.Channel) *ctrlPipeHandler {
	return &ctrlPipeHandler{
		registry: registry,
		ch:       ch,
	}
}

func (*ctrlPipeHandler) ContentType() int32 {
	return int32(ctrl_pb.ContentType_CtrlPipeRequestType)
}

func (handler *ctrlPipeHandler) HandleReceive(msg *channel.Message, ch channel.Channel) {
	log := pfxlog.ContextLogger(ch.Label()).Entry

	req := &ctrl_pb.CtrlPipeRequest{}
	if err := proto.Unmarshal(msg.Body, req); err != nil {
		log.WithError(err).Error("unable to unmarshal ssh tunnel request")
		return
	}

	conn, err := net.Dial("tcp", "localhost:22")
	if err != nil {
		log.WithError(err).Error("failed to dial ssh")
		handler.respondError(msg, err.Error())
		return
	}

	tunnel := &ctrlChanPipe{
		conn: conn,
		ch:   ch,
		id:   req.ConnId,
	}

	handler.registry.pipes.Put(tunnel.id, tunnel)

	log = log.WithField("connId", tunnel.id)
	log.Info("registered ctrl channel pipe connection")

	response := &ctrl_pb.CtrlPipeResponse{
		Success: true,
	}

	if sendErr := protobufs.MarshalTyped(response).ReplyTo(msg).WithTimeout(5 * time.Second).SendAndWaitForWire(ch); sendErr != nil {
		log.WithError(sendErr).Error("unable to send ctrl channel pipe response for successful pipe")
		tunnel.close(sendErr)
		return
	}

	log.Info("started ssh tunnel")

	go tunnel.readLoop()
}

func (handler *ctrlPipeHandler) respondError(request *channel.Message, msg string) {
	response := &ctrl_pb.CtrlPipeResponse{
		Success: false,
		Msg:     msg,
	}

	if sendErr := protobufs.MarshalTyped(response).ReplyTo(request).WithTimeout(5 * time.Second).SendAndWaitForWire(handler.ch); sendErr != nil {
		log := pfxlog.ContextLogger(handler.ch.Label()).Entry
		log.WithError(sendErr).Error("unable to send ctrl channel pipe response for failed pipe")
	}
}

type ctrlChanPipe struct {
	id   uint32
	conn net.Conn
	ch   channel.Channel
}

func (self *ctrlChanPipe) readLoop() {
	for {
		buf := make([]byte, 10240)
		n, err := self.conn.Read(buf)
		if err != nil {
			self.close(err)
			return
		}
		buf = buf[:n]
		msg := channel.NewMessage(int32(ctrl_pb.ContentType_CtrlPipeDataType), buf)
		msg.PutUint32Header(int32(ctrl_pb.ControlHeaders_CtrlPipeIdHeader), self.id)
		if err = self.ch.Send(msg); err != nil {
			self.close(err)
			return
		}
	}
}

func (self *ctrlChanPipe) close(err error) {
	log := pfxlog.ContextLogger(self.ch.Label()).WithField("connId", self.id)

	log.WithError(err).Info("closing ctrl channel pipe connection")

	if closeErr := self.conn.Close(); closeErr != nil {
		log.WithError(closeErr).Error("failed closing ctrl channel pipe connection")
	}

	if err != io.EOF && err != nil {
		msg := channel.NewMessage(int32(ctrl_pb.ContentType_CtrlPipeCloseType), []byte(err.Error()))
		msg.PutUint32Header(int32(ctrl_pb.ControlHeaders_CtrlPipeIdHeader), self.id)
		if sendErr := self.ch.Send(msg); sendErr != nil {
			log.WithError(sendErr).Error("failed sending ctrl channel pipe close message")
		}
	}
}

func newCtrlPipeDataHandler(registry *pipeRegistry) *ctrlPipeDataHandler {
	return &ctrlPipeDataHandler{
		registry: registry,
	}
}

type ctrlPipeDataHandler struct {
	registry *pipeRegistry
}

func (*ctrlPipeDataHandler) ContentType() int32 {
	return int32(ctrl_pb.ContentType_CtrlPipeDataType)
}

func (handler *ctrlPipeDataHandler) HandleReceive(msg *channel.Message, ch channel.Channel) {
	connId, _ := msg.GetUint32Header(int32(ctrl_pb.ControlHeaders_CtrlPipeIdHeader))
	tunnel := handler.registry.pipes.Get(connId)

	if tunnel == nil {
		pfxlog.ContextLogger(ch.Label()).
			WithField("connId", connId).
			Error("no ctrl channel pipe found for given id")

		go func() {
			errorMsg := fmt.Sprintf("invalid conn id '%v", connId)
			replyMsg := channel.NewMessage(int32(ctrl_pb.ContentType_CtrlPipeCloseType), []byte(errorMsg))
			replyMsg.PutUint32Header(int32(ctrl_pb.ControlHeaders_CtrlPipeIdHeader), connId)
			if sendErr := ch.Send(msg); sendErr != nil {
				pfxlog.ContextLogger(ch.Label()).
					WithField("connId", connId).
					WithError(sendErr).
					Error("failed sending ctrl channel pipe close message after data with invalid conn")
			}
		}()
		return
	}

	if _, err := tunnel.conn.Write(msg.Body); err != nil {
		tunnel.close(err)
	}
}
