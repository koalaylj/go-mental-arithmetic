package op

import (
	"fmt"
	"math/rand"

	"github.com/koalaylj/go-mental-arithmetic/app/pkg/config"
)

func RandomAdd(random *rand.Rand, options config.OP_ADD) string {
	a := randomBetween(random, options.Min, options.Max)
	b := randomBetween(random, options.Min, options.Max)

	//a + b
	for {
		if a+b > options.Bounds[0] && a+b <= options.Bounds[1] {
			if options.Carry {
				break
			} else {
				if a > 10 || b > 10 {
					break
				} else {
					a = randomBetween(random, options.Min, options.Max)
					b = randomBetween(random, options.Min, options.Max)
				}
			}
		}

		if b > a {
			b = randomBetween(random, options.Min, options.Max)
		} else if b < a {
			a = randomBetween(random, options.Min, options.Max)
		} else {
			a = randomBetween(random, options.Min, options.Max)
			b = randomBetween(random, options.Min, options.Max)
		}
	}
	fmt.Printf("%2d + %2d = %2v\n", a, b, a+b)
	return render(a, b, "+")
}

func RandomSub(random *rand.Rand, options config.OP_SUB) string {
	a := randomBetween(random, options.Min, options.Max)
	b := randomBetween(random, options.Min, options.Max)

	//a - b
	for {
		if a > 10 && a-b > options.Bounds[0] && a-b <= options.Bounds[1] {
			if options.Borrow {
				break
			} else {
				if b > 10 || (b < 10 && a-b > 10) {
					break
				} else {
					a = randomBetween(random, options.Min, options.Max)
					b = randomBetween(random, options.Min, options.Max)
				}
			}
		} else {
			a = randomBetween(random, options.Min, options.Max)
			b = randomBetween(random, options.Min, options.Max)
		}
	}

	fmt.Printf("%2d - %2d = %2v\n", a, b, a-b)
	return render(a, b, "-")
}

func render(a int, b int, op string) string {
	return fmt.Sprintf("%2d %s %d =", a, op, b)
}

func randomBetween(random *rand.Rand, min int, max int) int {
	return random.Intn(max-min) + min
}
