package main

import (
	"math/rand"
)

type proposer struct {
	id          int
	n           int // current n
	v           int // current proposal
	iv          int // initial proposal
	round       int
	promiseList map[int]int
}

//reset reset state
func (p *proposer) reset() {
	p.n = getN()
	p.round = getRound()
	p.iv = rand.Int() % 1000
	p.v = p.iv
	p.promiseList = make(map[int]int)
}

// run proposer's processing thread
func (p *proposer) run() {
	for {
		// wait for a while
		sleepRandomShort()

		//random go die
		if getLottery() {
			p.reset()
		}

		// check round
		if getRound() != p.round {
			p.reset()
			sleepLong()
			continue
		}

		// check promise list
		if len(p.promiseList) > acceptorAmount/2 {
			p.sendProposal()
			continue
		}

		// prepare
		p.sendPrepare()
	}
}

// sendPrepare send prepare
// restart when get no promise from only one the acceptors
func (p *proposer) sendPrepare() {
	randomIndex := randomIndex(acceptorAmount)
	totalBroadcast := float32(acceptorAmount) * broadcastRate
	for i := 0; i < int(totalBroadcast); i++ {
		a := acceptors[randomIndex[i]]
		sleepRandomShort()

		promised, _, cv := a.prepare(p.round, p.n)
		if !promised {
			p.n = getN()
			p.promiseList = make(map[int]int)
			print()
			sleepRandomLong()
			return
		}

		if cv != 0 {
			p.v = cv
		}
		p.promiseList[a.id] = a.id
		print()
	}
}

// sendProposal send proposal to acceptors that promised to current proposer
// restart to prepare stage when rejected by only one of the acceptors
func (p *proposer) sendProposal() {
	var rejected bool
	for _, a := range p.promiseList {
		sleepRandomShort()
		accepted := acceptors[a].accept(p.round, p.n, p.v)
		if !accepted {
			rejected = true
			continue
		}
	}

	if rejected {
		p.n = getN()
		p.promiseList = make(map[int]int)
		print()
		sleepRandomLong()
	}
}
