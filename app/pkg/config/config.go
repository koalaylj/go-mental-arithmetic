package config

type OP_ADD struct {
	Min   int
	Max   int
	Limit []int
	Carry bool
}
type OP_SUB struct {
	Min    int
	Max    int
	Limit  []int
	Borrow bool
}
