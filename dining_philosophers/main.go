package main

import (
	"fmt"
	"sync"
	"time"
)

type Philosopher struct {
	Name string
	LeftFork int
	RightFork int
}

type SafeFork struct {
	mux sync.Mutex
}

func (p *Philosopher) eat(safeForks []SafeFork, wg *sync.WaitGroup) {
	defer wg.Done()
	safeForks[p.LeftFork].mux.Lock()
	safeForks[p.RightFork].mux.Lock()
	fmt.Printf("%s is eating.\n", p.Name)
	time.Sleep(time.Second)
	fmt.Printf("%s is done eating.\n", p.Name)
	safeForks[p.LeftFork].mux.Unlock()
	safeForks[p.RightFork].mux.Unlock()
}

func main() {
	var wg sync.WaitGroup
	philosophers := []*Philosopher{
		&Philosopher{"Baruch Spinoza", 0, 1},
		&Philosopher{"Gilles Deleuze", 1, 2},
		&Philosopher{"Karl Marx", 2, 3},
		&Philosopher{"Friedrich Nietzsche", 3, 4},
		&Philosopher{"Michel Foucault", 0, 4}}
	safeForks := make([]SafeFork, 10)
	for _, p := range philosophers {
		wg.Add(1)
		go func (p *Philosopher) {
			p.eat(safeForks, &wg)
		}(p)

	}

	wg.Wait()
}
