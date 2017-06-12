package main

import (
	"fmt"
	"sync"
)

type learner struct {
	id    int
	round int
	lock  sync.Mutex
}

func (l *learner) reset() {
	l.round = nextRound()

	print()
}

func (l *learner) check() {
	l.lock.Lock()
	defer l.lock.Unlock()

	if amt, v := countAccepted(); amt > acceptorAmount/2 {
		fmt.Printf("Consensus achieved: %d\n", v)
		l.reset()
	}
}

func (l *learner) run() {
	l.round = getRound()

	for {
		sleepRandomShort()

		l.check()
	}
}
