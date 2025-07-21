package interfaces

import (
	"fmt"
	"math/rand"
)

type Player interface {
	KickBall()
	Name() string
}

type FootballPlayer struct {
	name    string
	stamina int
	power   int
}

func (f FootballPlayer) KickBall() {
	shot := f.stamina + f.power
	fmt.Printf("%s is kicking the ball:  %d\n", f.Name(), shot)
}

func (f FootballPlayer) Name() string {
	return f.name
}

type CR7 struct {
	name    string
	stamina int
	power   int
	SUIII   int
}

func (f CR7) KickBall() {
	shot := f.stamina + f.power*f.SUIII
	fmt.Println("CR7's kicking the ball: ", shot)
}

func (f CR7) Name() string {
	return f.name
}

func Interface_Sample() {
	team := make([]Player, 11)

	for i := 0; i < len(team)-1; i++ {
		team[i] = FootballPlayer{
			name:    fmt.Sprintf("random - %d", rand.Intn(2)),
			stamina: rand.Intn(10),
			power:   rand.Intn(10),
		}
	}

	team[len(team)-1] = CR7{
		name:    "Ronaldo",
		stamina: 10,
		power:   10,
		SUIII:   1000000,
	}

	for i := 0; i < len(team); i++ {
		team[i].KickBall()
	}
}
