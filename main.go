package main

import (
	"fmt"
	"math/rand"
)

const setSize = 100
const targetNum = 30
const numOfTest = 100000000

func main() {
	if targetNum > setSize {
		fmt.Println("Invalid targetNum")
		return
	}
	experiment(123369)
}

func experiment(seed int64) {
	rand.Seed(seed)
	naiveApproach()
	prisonersApproach()
}

func naiveApproach() {
	var totalNum uint64
	totalNum = 0
	for i := 0; i < numOfTest; i++ {
		source := genRandomSet()
		sequence := genRandomSet()
		round := whenFinish(source, sequence)
		totalNum += uint64(round)
	}
	expectedRound := float64(totalNum) / numOfTest
	fmt.Println("Expected Round num =", expectedRound)
}

func prisonersApproach() {
	var totalNum uint64
	totalNum = 0
	for i := 0; i < numOfTest; i++ {
		source := genRandomSet()
		round := whenFinishPrisoner(source)
		totalNum += uint64(round)
	}
	expectedRound := float64(totalNum) / numOfTest
	fmt.Println("Expected Round num =", expectedRound)
}

func genRandomSet() []uint32 {
	var ret = make([]uint32, setSize)
	var i uint32
	for i = 0; i < setSize; i++ {
		ret[i] = i + 1
	}
	// start to shuffle
	rand.Shuffle(setSize, func(i, j int) {
		ret[i], ret[j] = ret[j], ret[i]
	})
	return ret
}

func whenFinish(source, sequence []uint32) int {
	for i := 0; i < setSize; i++ {
		target := sequence[i]
		if source[target-1] == targetNum {
			return i + 1
		}
	}
	return setSize
}

func whenFinishPrisoner(source []uint32) int {
	target := targetNum
	for i := 0; i < setSize; i++ {
		if source[target-1] == targetNum {
			return i + 1
		}
		target = int(source[target-1])
	}
	return setSize
}
