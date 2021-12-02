package day2

type Command struct {
	movement Movement
	units    int
}

type Movement string

const (
	Forward Movement = "forward"
	Back    Movement = "back"
	Down    Movement = "down"
	Up      Movement = "up"
)
