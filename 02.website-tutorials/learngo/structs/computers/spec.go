package computers

type Spec struct { //exported
	Maker string // exported
	Price int    // exported
	model string //unexported
}
