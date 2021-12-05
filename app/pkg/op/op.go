package op

import (
	"fmt"
	"math/rand"
)

func RandomAdd(random *rand.Rand, max int, min int, carry bool, limit []int) string {
	a := random.Intn(max) + min
	b := random.Intn(max) + min

	//a + b
	for {

		if a+b > limit[0] && a+b <= limit[1] {
			if carry {
				break
			} else {
				if a > 10 || b > 10 {
					break
				} else {
					a = random.Intn(max) + min
					b = random.Intn(max) + min
				}
			}
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
	fmt.Printf("%2d + %2d = %2v\n", a, b, a+b)
	return render(a, b, "+")
}

func RandomSub(random *rand.Rand, max int, min int, carry bool) string {
	a := random.Intn(max) + min
	b := random.Intn(max) + min

	//a - b
	for {
		if a > b && a > 10 {
			if carry {
				break
			} else {
				if b > 10 || (b < 10 && a-b > 10) {
					break
				} else {
					a = random.Intn(max) + min
					b = random.Intn(max) + min
				}
			}
		} else if a < b {
			temp := b
			b = a
			a = temp
		} else {
			a = random.Intn(max) + min
			b = random.Intn(max) + min
		}
	}

	fmt.Printf("%2d - %2d = %2v\n", a, b, a-b)
	return render(a, b, "-")
}

func render(a int, b int, op string) string {
	return fmt.Sprintf("%2d %s %d =", a, op, b)
}
