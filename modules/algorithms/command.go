package algorithms

import (
	"errors"
	"fmt"
)

// Command Pattern in Go
// useful when needing to create and execute commands with a task management queue that is separate from the execution of the task itself.

type command interface {
	execute() error
}

// tomagachi acts like a factory
type tomagachi struct {
	energy int
	poop   int
}

func newTomagachi() *tomagachi {
	return &tomagachi{energy: 35, poop: 0}
}

func (t *tomagachi) feedPet(n int) command {
	return &feed{n, t}
}
func (t *tomagachi) playPet(n int) command {
	return &play{n, t}
}
func (t *tomagachi) poopPet(n int) command {
	return &poop{n, t}
}

type simulate struct {
	commands []command
}

func (sim *simulate) executeAll() {
	var err error
	for _, commands := range sim.commands {
		err = commands.execute()
		// note the implementation of the tomagachi dying will continuously keep being called because the tasks will not stop until the end.
		if err != nil {
			fmt.Println("Tomagachi died.")
		}
	}
}

type feed struct {
	food int
	t    *tomagachi
}

func (f *feed) execute() error {
	f.t.energy += f.food
	fmt.Println("fed energy:", f.t.energy)
	return nil
}

type play struct {
	activity int
	t        *tomagachi
}

func (p *play) execute() error {
	p.t.energy -= p.activity
	fmt.Println("spent energy:", p.t.energy)
	return nil
}

type poop struct {
	amount int
	t      *tomagachi
}

func (p *poop) execute() error {
	p.t.poop += p.amount
	fmt.Println("total poop:", p.t.poop)
	if p.t.poop > 75 {
		return errors.New("too much poop.")
	}
	return nil
}

func ExampleCommandPattern() {
	t := newTomagachi()

	tasks := []command{
		t.feedPet(3),
		t.feedPet(3),
		t.playPet(6),
		t.playPet(6),
		t.poopPet(33),
		t.poopPet(3),
		t.feedPet(3),
		t.feedPet(11),
		t.playPet(6),
		t.playPet(22),
		t.poopPet(33),
		t.poopPet(3),
		t.feedPet(3),
		t.feedPet(11),
		t.playPet(6),
		t.playPet(22),
		t.poopPet(33),
		t.poopPet(3),
	}

	sims := []*simulate{
		&simulate{},
		&simulate{},
	}

	for i, task := range tasks {
		pet := sims[i%len(sims)]
		pet.commands = append(pet.commands, task)
	}

	for i, p := range sims {
		fmt.Println("pet", i, " did ")
		p.executeAll()
	}
}
