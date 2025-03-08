extends Node

const packets := preload("res://packets.gd")

@onready var _log: Log = $Log

func _input(event: InputEvent) -> void:
	if event.is_action_pressed("ui_accept"):
		_log.info("got input %s" % event.as_text())
		var packet := packets.Packet.new()
		var chat_msg := packet.new_chat()
		chat_msg.set_msg("My Event "+event.as_text())
			
		var err := WS.send(packet)
		if err:
			_log.error("Error sending packet")
		else:
			_log.success("Sent packet")

func _ready() -> void:
	WS.connected_to_server.connect(_on_ws_connected_to_server)
	WS.connection_closed.connect(_on_ws_connection_closed)
	WS.packet_received.connect(_on_ws_packet_received)
	
	_log.info("Connecting to server...")
	WS.connect_to_url("ws://127.0.0.1:8080/ws")

func _on_ws_connected_to_server() -> void:
	var packet := packets.Packet.new()
	var chat_msg := packet.new_chat()
	chat_msg.set_msg("Hello, Golan1g!")
	
	var err := WS.send(packet)
	if err:
		_log.error("Error sending packet")
	else:
		_log.success("Sent packet")
	
func _on_ws_connection_closed() -> void:
	_log.error("Connection closed")
	
func _on_ws_packet_received(packet: packets.Packet) -> void:
	_log.info("Received packet from the server: %s" % packet)
