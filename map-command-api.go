//go:generate go run github.com/bytecodealliance/wasm-tools-go/cmd/wit-bindgen-go generate --world map-command-api --out gen ./wit
package main

import (
	"log/slog"

	"github.com/Mattilsynet/map-command-api/pkg/nats"
	"github.com/Mattilsynet/map-command-api/pkg/subject"
	"github.com/Mattilsynet/mapis/gen/go/command/v1"
	"github.com/google/uuid"
	"go.wasmcloud.dev/component/log/wasilog"
)

var (
	conn   *nats.Conn
	js     *nats.JetStreamContext
	logger *slog.Logger
)

func init() {
	logger = wasilog.ContextLogger("map-command-api")
	conn = nats.NewConn()
	var err error
	js, err = conn.Jetstream()
	if err != nil {
		logger.Error("error getting jetstreamcontext", "err", err)
		return
	}
	conn.RegisterRequestReply(MapCommandApi)
}

func MapCommandApi(msg *nats.Msg) *nats.Msg {
	logger.Info("MapCommandApi: got msg: ", "data", string(msg.Data))
	uuid := uuid.NewString()
	replyMsg := &nats.Msg{
		Subject: msg.Reply,
		Data:    msg.Data,
	}

	cmd := &command.Command{}
	err := cmd.UnmarshalVT(msg.Data)
	if err != nil {
		logger.Error("error unmarshalling", "err", err)
		replyMsg.Data = []byte(err.Error())
		return replyMsg
	}
	status := command.CommandStatus{}
	status.Id = uuid
	cmd.Status = &status
	logger.Info("MapCommandApi:", "cmd:", string(cmd.Spec.GetTypePayload()))
	subj := subject.NewCommandSubject(cmd)
	logger.Info("MapCommandApi: subj:", "subj", subj.ToCommand(cmd))
	bytes, err := cmd.MarshalVT()
	if err != nil {
		logger.Error("error marshalling", "err", err)
		replyMsg.Data = []byte(err.Error())
		return replyMsg
	}
	logger.Info("MapCommandApi: marshalled:", "bytes", string(bytes))
	replyMsg.Data = bytes
	err = js.Publish(subj.ToCommand(cmd), bytes)
	if err != nil {
		logger.Error("error publishing", "err", err)
		replyMsg.Data = []byte(err.Error())
	}
	return replyMsg
}

//go:generate go run github.com/bytecodealliance/wasm-tools-go/cmd/wit-bindgen-go generate --world starter-kit --out gen ./wit
func main() {}
