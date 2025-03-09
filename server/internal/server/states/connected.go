package states

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strings"

	"golang.org/x/crypto/bcrypt"

	"github.com/sijoma/godot-mmo/internal/server"
	"github.com/sijoma/godot-mmo/internal/server/db"
	"github.com/sijoma/godot-mmo/pkg/packets"
)

type Connected struct {
	client server.ClientInterfacer

	queries *db.Queries
	dbCtx   context.Context

	logger *log.Logger
}

func (c *Connected) Name() string {
	return "Connected"
}

func (c *Connected) SetClient(client server.ClientInterfacer) {
	c.client = client
	c.queries = client.DbTx().Queries
	c.dbCtx = client.DbTx().Ctx

	loggingPrefix := fmt.Sprintf("Client %d [%s]: ", client.Id(), c.Name())
	c.logger = log.New(log.Writer(), loggingPrefix, log.LstdFlags)
}

func (c *Connected) OnEnter() {
	// A newly connected client will want to know its own ID first
	c.client.SocketSend(packets.NewId(c.client.Id()))
}

func (c *Connected) HandleMessage(senderId uint64, message packets.Msg) {
	switch msg := message.(type) {
	case *packets.Packet_LoginRequest:
		c.handleLoginRequest(senderId, msg)
	case *packets.Packet_RegisterRequest:
		c.handleRegisterRequest(senderId, msg)
	}
}

func (c *Connected) OnExit() {
}

func (c *Connected) handleLoginRequest(senderId uint64, msg *packets.Packet_LoginRequest) {
	if senderId != c.client.Id() {
		c.logger.Printf("Received login message from another client: %d\n", senderId)
		return
	}

	username := msg.LoginRequest.GetUsername()

	genericFailMessage := packets.NewDenyResponse("Incorrect username or password")

	user, err := c.queries.GetUserByUsername(c.dbCtx, strings.ToLower(username))
	if err != nil {
		c.logger.Printf("Error getting user %s: %v", username, err)
		c.client.SocketSend(genericFailMessage)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(msg.LoginRequest.GetPassword()))
	if err != nil {
		c.logger.Printf("User entered wrong password: %s", username)
		c.client.SocketSend(genericFailMessage)
		return
	}

	c.logger.Printf("User %s logged in successfully", username)
	c.client.SocketSend(packets.NewOkResponse())
}

func (c *Connected) handleRegisterRequest(senderId uint64, msg *packets.Packet_RegisterRequest) {
	if senderId != c.client.Id() {
		c.logger.Printf("Received register message from another client: %d\n", senderId)
		return
	}

	username := strings.ToLower(msg.RegisterRequest.GetUsername())
	err := validateUsername(msg.RegisterRequest.GetUsername())
	if err != nil {
		reason := fmt.Sprintf("Invalid username: %v", err)
		c.logger.Printf("handleRegisterRequest: %s", reason)
		c.client.SocketSend(packets.NewDenyResponse(reason))
	}

	_, err = c.queries.GetUserByUsername(c.dbCtx, username)
	if err == nil {
		c.logger.Printf("User %s already registered", username)
		c.client.SocketSend(packets.NewDenyResponse("User already registered"))
		return
	}

	genericFailMessage := packets.NewDenyResponse("Error registering user (internal server error) - please try again later")

	// Add new user
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(msg.RegisterRequest.Password), bcrypt.DefaultCost)
	if err != nil {
		c.logger.Printf("Failed to hash password: %s", username)
		c.client.SocketSend(genericFailMessage)
		return
	}

	_, err = c.queries.CreateUser(c.dbCtx, db.CreateUserParams{
		Username:     username,
		PasswordHash: string(passwordHash),
	})

	if err != nil {
		c.logger.Printf("Failed to create user %s: %v", username, err)
		c.client.SocketSend(genericFailMessage)
		return
	}

	c.client.SocketSend(packets.NewOkResponse())

	c.logger.Printf("User %s registered successfully", username)

}

func validateUsername(username string) error {
	if len(username) <= 0 {
		return errors.New("empty")
	}
	if len(username) > 20 {
		return errors.New("too long")
	}
	if username != strings.TrimSpace(username) {
		return errors.New("leading or trailing whitespace")
	}
	return nil
}
