extends Node

const packets := preload("res://packets.gd")

func _input(event: InputEvent) -> void:
	print("got input", event.as_text())
	var packet := packets.Packet.new()
	var chat_msg := packet.new_chat()
	chat_msg.set_msg("My Event "+event.as_text())
		
	var err := WS.send(packet)
	if err:
		print("error sending")
	else:
		print("sent packet")

func _ready() -> void:
	WS.connected_to_server.connect(_on_ws_connected_to_server)
	WS.connection_closed.connect(_on_ws_connection_closed)
	WS.packet_received.connect(_on_ws_packet_received)
	
	print("Connecting to server...")
	WS.connect_to_url("ws://127.0.0.1:8080/ws")

func _on_ws_connected_to_server() -> void:
	var packet := packets.Packet.new()
	var chat_msg := packet.new_chat()
	chat_msg.set_msg("Hello, Golan1g!")
	
	var err := WS.send(packet)
	if err:
		print("Error sending packet")
	else:
		print("Sent packet")
	
func _on_ws_connection_closed() -> void:
	print("Connection closed")
	
func _on_ws_packet_received(packet: packets.Packet) -> void:
	print("Received packet from the server: %s" % packet)
