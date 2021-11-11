package op

import (
	"fmt"
	"math/rand"
)

func RandomAdd(random *rand.Rand, max int, min int) string {
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

func RandomSub(random *rand.Rand, max int, min int) string {
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
	return fmt.Sprintf("%2d %s %d =", a, op, b)
}
