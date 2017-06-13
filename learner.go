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

	acceptedVMap := countAccepted(l.round)

	for k, v := range acceptedVMap {
		if v > acceptorAmount/2 {
			fmt.Printf("Consensus achieved: %d\n", k)
			l.reset()
			break
		}
	}
}

func (l *learner) run() {
	l.round = getRound()

	for {
		sleepRandomShort()

		l.check()
	}
}
