package interfaces

import (
	"fmt"
	"math/rand"
)

type Player interface {
	KickBall()
}

type FootballPlayer struct {
	stamina int
	power   int
}

func (f FootballPlayer) KickBall() {
	shot := f.stamina + f.power
	fmt.Println("I'm kicking the ball: ", shot)
}

type CR7 struct {
	stamina int
	power   int
	SUIII   int
}

func (f CR7) KickBall() {
	shot := f.stamina + f.power*f.SUIII
	fmt.Println("CR7's kicking the ball: ", shot)
}

func Interface_Sample() {
	team := make([]Player, 11)

	for i := 0; i < len(team)-1; i++ {
		team[i] = FootballPlayer{
			stamina: rand.Intn(10),
			power:   rand.Intn(10),
		}
	}

	team[len(team)-1] = CR7{
		stamina: 10,
		power:   10,
		SUIII:   1000000,
	}

	for i := 0; i < len(team); i++ {
		team[i].KickBall()
	}
}
