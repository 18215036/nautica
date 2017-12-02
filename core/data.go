package core

type DiveSeq struct {
	ID     string  'json:"id,omitempty"'     //Dive Sequence ID
	Dive1  Dive    'json:"dive1,omitempty"'
	Dive2  Dive    'json:"dive2,omitempty"'
	Dive3  Dive    'json:"dive3,omitempty"'
	Dive4  Dive    'json:"dive4,omitempty"'
	Dive5  Dive    'json:"dive5,omitempty"'
}

type Dive struct { //individual dive data
	D   int     'json:"depth,omitempty"'  //depth
	R   int     'json:"rnt,omitempty"'    //residual nitrogen time
	A   int     'json:"abt,omitempty"'    //actual bottom time
	T   int     'json:"tbt,omitempty"'    //total bottom time
	S   int     'json:"si,omitempty"'     //surface interval
	P   string  'json:"pg,omitempty"'     //previous dive group
	C   string  'json:"cg,omitempty"'     //current dive group
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
	d2 DATA2  //dive groups
	DE int    //next dive depth
}

type DATA1 struct { //no-deco data
	d3 DATA3   //depth & dive groups
	NSL int    //no-stop limit
}