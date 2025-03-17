class_name Spore
extends Area2D

const Scene := preload("res://objects/spore/spore.tscn")


var spore_id: int
var x: float
var y: float
var radius: float
var color: Color

@onready var _collision_shape: CircleShape2D = $CollisionShape2D.shape

static func instantiate(p_spore_id: int, p_x: float, p_y: float, p_radius: float) -> Spore:
	var spore := Scene.instantiate() as Spore
	spore.spore_id = p_spore_id
	spore.x = p_x
	spore.y = p_y
	spore.radius = p_radius
	
	return spore


func _ready() -> void:
	position.x = x
	position.y = y
	_collision_shape.radius = radius

	color = Color.from_hsv(randf(), 1, 1, 1)

func _draw() -> void:
	draw_circle(Vector2.ZERO, radius, color)
