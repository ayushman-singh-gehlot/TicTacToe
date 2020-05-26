package components

type Player struct {
	name string
	mark string
}

func NewPlayer(name, mark string) *Player {
	return &Player{
		name: name,
		mark: mark,
	}
}
