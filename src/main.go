package main

import (
	"fmt"
	"math/rand"
	"time"
)

var ops = [...]string{"+", "-"}

const max = 10
const min = 1
const total = 60

func main() {

	for i := 1; i <= total; i++ {

		random := getRandom()

		op := randomOp(random)

		item := ""

		switch op {
		case "+":
			item = randomAdd(random, min, max)
		case "-":
			item = randomSub(random, min, max)
		}

		fmt.Printf("%2d) %s\n", i, item)
	}
}

func randomAdd(random *rand.Rand, min int, max int) string {
	a := random.Intn(max) + min
	b := random.Intn(max) + min

	//a + b
	for {
		if a+b <= max {
			break
		}

		if b > a {
			b = random.Intn(max) + min
		} else if b < a {
			a = random.Intn(max) + min
		} else {
			a = random.Intn(max) + min
			b = random.Intn(max) + min
		}
	}

	return render(a, b, "+")
}

func randomSub(random *rand.Rand, min int, max int) string {
	a := random.Intn(max) + min
	b := random.Intn(max) + min

	//a - b
	for {
		if a > b {
			break
		} else if a < b {
			temp := b
			b = a
			a = temp
		} else {
			a = random.Intn(max) + min
			b = random.Intn(max) + min
		}
	}

	return render(a, b, "-")
}

func render(a int, b int, op string) string {
	return fmt.Sprintf("%2d %s %d", a, op, b)
}

func randomOp(random *rand.Rand) string {

	index := random.Intn(len(ops))

	return ops[index]
}

func getRandom() *rand.Rand {
	now := time.Now().UnixNano()
	seed := rand.NewSource(now)
	random := rand.New(seed)
	// fmt.Print(now, ":")

	return random
}
