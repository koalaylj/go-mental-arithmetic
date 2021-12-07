package config

type OP_ADD struct {
	Min    int
	Max    int
	Bounds []int
	Carry  bool
}
type OP_SUB struct {
	Min    int
	Max    int
	Bounds []int
	Borrow bool
}
