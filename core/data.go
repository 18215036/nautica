package core

type Dive struct { //individual dive data
	D int
	R int
	A int
	T int
	S int
	P string
	C string
}

type DATA2 struct { //surface interval data
	A int
	B int
	C int
	D int
	E int
	F int
	G int
	H int
	I int
	J int
	K int
	L int
	M int
	N int
	O int
	Z int
}

type DATA3 struct { //residual nitrogen data
	d2 DATA2
	DE int    //next dive depth
}

type DATA1 struct { //no-deco data
	d3 DATA3
	NSL int    //no-stop limit
}