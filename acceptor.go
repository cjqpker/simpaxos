package main

import (
	"log"
	"sync"
)

type acceptor struct {
	id    int
	cn    int
	cv    int
	round int
	lock  sync.Mutex
}

func (a *acceptor) reset() {
	a.cn = 0
	a.cv = 0
}

// prepare
func (a *acceptor) prepare(r, n int) (promised bool, cn, cv int) {
	a.lock.Lock()
	defer a.lock.Unlock()

	//random go die
	if getLottery() {
		a.reset()
	}

	// check round
	cr := getRound()
	if a.round != cr {
		a.reset()
		a.round = cr
	}
	if r != cr {
		promised = false
		return
	}

	// first prepare
	if a.cn == 0 {
		a.cn = n
		promised = true
		print()
		return
	}

	// reject
	if n < a.cn {
		cn = a.cn
		cv = a.cv
		promised = false
		return
	}

	// higher n
	if n > a.cn {
		a.cn = n
		promised = true
		cv = a.cv
		print()
		return
	}

	// should never happen
	if n == a.cn {
		log.Fatal("n = a.cn should never happen")
	}

	return
}

func (a *acceptor) accept(r, n, v int) (accepted bool) {
	a.lock.Lock()
	defer a.lock.Unlock()

	//random go die
	if getLottery() {
		a.reset()
	}

	// check round
	cr := getRound()
	if a.round != cr {
		a.reset()
		a.round = cr
	}
	if r != cr {
		accepted = false
		return
	}

	// check current n
	if n < a.cn {
		accepted = false
		return
	}

	// accept proposal
	if n == a.cn {
		a.cv = v
		accepted = true
		print()
		return
	}

	// might happen after die
	if n > a.cn {
		return false
	}

	return
}

// check whether shoud we start a new round
func (a *acceptor) check() {
	a.lock.Lock()
	defer a.lock.Unlock()

	if getRound() != a.round {
		a.reset()
	}
}

//accepted whether accepted a proposal or not
func (a *acceptor) accepted() (bool, int) {
	a.lock.Lock()
	a.lock.Unlock()

	if a.cv != 0 {
		return true, a.cv
	}

	return false, 0
}

func countAccepted() (int, int) {
	amt := 0
	vret := 0
	for _, a := range acceptors {
		if accepted, v := a.accepted(); accepted {
			amt++
			vret = v
		}
	}

	return amt, vret
}

func (a *acceptor) run() {
	for {
		sleepRandomShort()
		a.check()
	}
}
