package states

import (
	"fmt"
	"log"
	"math/rand/v2"

	"github.com/sijoma/godot-mmo/internal/server"
	"github.com/sijoma/godot-mmo/internal/server/objects"
	"github.com/sijoma/godot-mmo/pkg/packets"
)

type InGame struct {
	client server.ClientInterfacer
	player *objects.Player
	logger *log.Logger
}

func (g *InGame) Name() string {
	return "InGame"
}

func (g *InGame) SetClient(client server.ClientInterfacer) {
	g.client = client
	loggingPrefix := fmt.Sprintf("Client %d [%s]: ", client.Id(), g.Name())
	g.logger = log.New(log.Writer(), loggingPrefix, log.LstdFlags)
}

func (g *InGame) OnEnter() {
	g.logger.Printf("Adding player %s to the shared collection", g.player.Name)
	go g.client.SharedGameObjects().Players.Add(g.player, g.client.Id())

	// Set the initial properties of the player
	g.player.X = rand.Float64() * 1000
	g.player.Y = rand.Float64() * 1000
	g.player.Speed = 150.0
	g.player.Radius = 20.0

	// Send the player's initial state to the client
	g.client.SocketSend(packets.NewPlayer(g.client.Id(), g.player))
}

func (g *InGame) HandleMessage(senderId uint64, message packets.Msg) {
}

func (g *InGame) OnExit() {
	g.client.SharedGameObjects().Players.Remove(g.client.Id())
}
