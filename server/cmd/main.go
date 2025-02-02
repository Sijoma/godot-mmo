package main

import (
	"fmt"

	"github.com/sijoma/godot-mmo/pkg/packets"
)

//go:generate protoc --proto_path=../../shared/packets --go_out=.. --go-grpc_out=.. ../../shared/packets/packets.proto

func main() {
	fmt.Println("Hello World")
	packet := &packets.Packet{
		SenderId: 60,
		Msg:      packets.NewChat("Hello World"),
	}

	fmt.Println(packet)
}
