package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var (
	n     int
	nLock sync.Mutex

	round int
	rLock sync.Mutex

	pLock sync.Mutex // locker for print
)

func initS() {
	rand.Seed(time.Now().Unix())
	round = rand.Int()%10 + 1
}

// getN get n
func getN() int {
	nLock.Lock()
	defer nLock.Unlock()

	if n == 0 {
		n = rand.Int()%10 + 1
	}

	n++
	return n
}

//getRound get current Round
func getRound() int {
	rLock.Lock()
	defer rLock.Unlock()

	return round
}

//nextRound next round
func nextRound() int {
	rLock.Lock()
	defer rLock.Unlock()

	round++
	return round
}

func sleepRandomShort() {
	time.Sleep(getDuration())
}

func sleepRandomLong() {
	time.Sleep(getDuration() * 10)
}

func sleepLong() {
	time.Sleep(time.Second * 5)
}

// getDuration
func getDuration() time.Duration {
	return time.Millisecond * time.Duration(rand.Int()%sleepDuration)
}

func getLottery() bool {
	return rand.Int()%lotteryDifficulty == 0
}

func print() {
	pLock.Lock()
	defer pLock.Unlock()

	fmt.Printf("round:%2d\n", getRound())
	for _, p := range proposers {
		fmt.Printf("p%02d    [n: %3d]    [v: %3d -> %3d]\n", p.id+1, p.n, p.iv, p.v)
	}
	for _, a := range acceptors {
		fmt.Printf("a%02d    [n: %3d]    [v:        %3d]\n", a.id+1, a.cn, a.cv)
	}
}

//randomIndex make a index array of random sequence
func randomIndex(l int) []int {
	larr := make([]int, l)
	for i := 0; i < l; i++ {
		larr[i] = i
	}

	var ret []int
	for len(ret) < l {
		index := rand.Int() % l
		if larr[index] != -1 {
			ret = append(ret, larr[index])
			larr[index] = -1
		}
	}

	return ret
}
