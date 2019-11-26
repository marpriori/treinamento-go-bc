package person

import "fmt"

type Runner struct {
	speed float64
}

func NewRunner(speed float64) Runner {
	return Runner{speed: speed}
}

func (r Runner) Run() {
	fmt.Printf("Correndo ... a %2f", r.speed)
	fmt.Println()
}
