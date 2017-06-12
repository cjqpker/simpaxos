package main

const (
	acceptorAmount = 10
	proposerAmount = 2
	learnerAmount  = 1

	sleepDuration     = 1000 // loop wait
	lotteryDifficulty = 5    // rate for node's random recovery, more offen when smaller
)

var (
	acceptors [acceptorAmount]*acceptor
	proposers [proposerAmount]*proposer
	learners  [learnerAmount]*learner
)

func main() {
	initS()

	for i := 0; i < acceptorAmount; i++ {
		acceptors[i] = &acceptor{id: i}
		go acceptors[i].run()
	}

	for i := 0; i < proposerAmount; i++ {
		proposers[i] = &proposer{id: i}
		go proposers[i].run()
	}

	for i := 0; i < learnerAmount; i++ {
		learners[i] = &learner{id: i}
		go learners[i].run()
	}

	for {
		//print()
		sleepLong()
	}
}
