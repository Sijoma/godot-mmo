extends Area2D

const packets := preload("res://packets.gd")

const Scene := preload("res://objects/actor/actor.tscn")
const Actor := preload("res://objects/actor/actor.gd")

var actor_id: int
var actor_name: String
var start_x: float
var start_y: float
var start_rad: float
var speed: float
var is_player: bool

var velocity: Vector2
var radius: float:
	set(new_radius):
		radius = new_radius
		_collision_shape.set_radius(radius)
		queue_redraw()

@onready var _nameplate: Label = $Label
@onready var _camera: Camera2D = $Camera2D
@onready var _collision_shape: CircleShape2D = $CollisionShape2D.shape

static func instantiate(p_actor_id: int, p_actor_name: String, p_x: float, p_y: float, p_radius: float, p_speed: float, p_is_player: bool) -> Actor:
	var actor := Scene.instantiate()
	actor.actor_id = p_actor_id
	actor.actor_name = p_actor_name
	actor.start_x = p_x
	actor.start_y = p_y
	actor.start_rad = p_radius
	actor.speed = p_speed
	actor.is_player = p_is_player

	return actor


func _ready():
	position.x = start_x
	position.y = start_y
	velocity = Vector2.RIGHT * speed
	radius = start_rad
	
	_collision_shape.radius = radius
	_nameplate.text = actor_name


func _input(event):
	if is_player and event is InputEventMouseButton and event.is_pressed():
		match event.button_index:
			MOUSE_BUTTON_WHEEL_UP:
				_camera.zoom.x = min(4, _camera.zoom.x + 0.1)
			MOUSE_BUTTON_WHEEL_DOWN:
				_camera.zoom.x = max(0.1, _camera.zoom.x - 0.1)

		_camera.zoom.y = _camera.zoom.x


func _physics_process(delta) -> void:
	position += velocity * delta
	
	if not is_player:
		return
	# Player-specific stuff below here
		
	var mouse_pos := get_global_mouse_position()
	
	var input_vec = position.direction_to(mouse_pos).normalized()
	if abs(velocity.angle_to(input_vec)) > TAU / 15: # 24 degrees
		velocity = input_vec * speed
		var packet := packets.Packet.new()
		var player_direction_message = packet.new_player_direction()
		player_direction_message.set_direction(velocity.angle())
		WS.send(packet)

func _draw() -> void:
	draw_circle(Vector2.ZERO, _collision_shape.radius, Color.DARK_ORCHID)
