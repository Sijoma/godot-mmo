package states

import (
	"fmt"
	"log"

	"github.com/sijoma/godot-mmo/internal/server"
	"github.com/sijoma/godot-mmo/pkg/packets"
)

type Connected struct {
	client server.ClientInterfacer
	logger *log.Logger
}

func (c *Connected) Name() string {
	return "Connected"
}

func (c *Connected) SetClient(client server.ClientInterfacer) {
	c.client = client
	loggingPrefix := fmt.Sprintf("Client %d [%s]: ", client.Id(), c.Name())
	c.logger = log.New(log.Writer(), loggingPrefix, log.LstdFlags)
}

func (c *Connected) OnEnter() {
	// A newly connected client will want to know its own ID first
	c.client.SocketSend(packets.NewId(c.client.Id()))
}

func (c *Connected) HandleMessage(senderId uint64, message packets.Msg) {
	if senderId == c.client.Id() {
		c.client.Broadcast(message)
	} else {
		c.client.SocketSendAs(message, senderId)
	}
}

func (c *Connected) OnExit() {
}
