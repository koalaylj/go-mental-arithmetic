package main

import (
	"fmt"
	"math/rand"
	"time"
)

func randomAdd(random *rand.Rand) string {
	a := random.Intn(ARG_MAX) + ARG_MIN
	b := random.Intn(ARG_MAX) + ARG_MIN

	//a + b
	for {
		if a+b <= ARG_MAX {
			break
		}

		if b > a {
			b = random.Intn(ARG_MAX) + ARG_MIN
		} else if b < a {
			a = random.Intn(ARG_MAX) + ARG_MIN
		} else {
			a = random.Intn(ARG_MAX) + ARG_MIN
			b = random.Intn(ARG_MAX) + ARG_MIN
		}
	}

	return render(a, b, "+")
}

func randomSub(random *rand.Rand) string {
	a := random.Intn(ARG_MAX) + ARG_MIN
	b := random.Intn(ARG_MAX) + ARG_MIN

	//a - b
	for {
		if a > b {
			break
		} else if a < b {
			temp := b
			b = a
			a = temp
		} else {
			a = random.Intn(ARG_MAX) + ARG_MIN
			b = random.Intn(ARG_MAX) + ARG_MIN
		}
	}

	return render(a, b, "-")
}

func render(a int, b int, op string) string {
	return fmt.Sprintf("%2d %s %d =", a, op, b)
}

func randomOp(random *rand.Rand) string {
	index := random.Intn(len(ops))
	return ops[index]
}

func getRandom() *rand.Rand {
	now := time.Now().UnixNano()
	seed := rand.NewSource(now)
	random := rand.New(seed)
	return random
}
