extends Node

const packets := preload("res://packets.gd")

@onready var _log: Log = $UI/Log

func _ready() -> void:
	WS.connected_to_server.connect(_on_ws_connected_to_server)
	WS.connection_closed.connect(_on_ws_connection_closed)
	WS.packet_received.connect(_on_ws_packet_received)

	_log.info("Connecting to server...")
	WS.connect_to_url("wss://17ec-2a00-6020-a585-6600-7c88-b831-e9d9-432e.ngrok-free.app/ws")
	#WS.connect_to_url("ws://localhost:8080/ws")

func _on_ws_connected_to_server() -> void:
	_log.info("Connected to server")

func _on_ws_connection_closed() -> void:
	_log.info("Connection closed")

func _on_ws_packet_received(packet: packets.Packet) -> void:
	var sender_id := packet.get_sender_id()
	if packet.has_id():
		_handle_id_msg(sender_id, packet.get_id())

func _handle_id_msg(_sender_id: int, id_msg: packets.IdMessage) -> void:
	GameManager.client_id = id_msg.get_id()
	GameManager.set_state(GameManager.State.CONNECTED)
