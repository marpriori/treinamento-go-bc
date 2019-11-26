package person

type Person struct {
	Runner
	name string
	age int
}

func NewPerson(name string, age int, speed float64) Person {
	return Person {
		name:	name,
		age:	age,
		Runner: NewRunner(speed),
	}
}

func (p Person) Age() int {
	return p.age
}