<<<<<<< HEAD
//18215036/Ahmad Fawwaz Zuhdi
package core

type DiveSeq struct {
    ID     string //Dive Sequence ID
    Dive1  Dive `json:",omitempty"`
    Dive2  Dive `json:",omitempty"`
    Dive3  Dive `json:",omitempty"`
    Dive4  Dive `json:",omitempty"`
    Dive5  Dive `json:",omitempty"`
}

type Dive struct { //individual dive data
    D   int    `json:",omitempty"` //depth
    R   int    `json:",omitempty"` //residual nitrogen time
    A   int    `json:",omitempty"` //actual bottom time
    T   int    `json:",omitempty"` //total bottom time
    S   int    `json:",omitempty"` //surface interval
    P   string `json:",omitempty"` //previous dive group
    C   string `json:",omitempty"` //current dive group
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
=======
package core

type DiveSeq struct {
    //ID     string //Dive Sequence ID
    Dive1  Dive `json:",omitempty"`
    Dive2  Dive `json:",omitempty"`
    Dive3  Dive `json:",omitempty"`
    Dive4  Dive `json:",omitempty"`
    Dive5  Dive `json:",omitempty"`
}

type Dive struct { //individual dive data
    D   int    `json:",omitempty"` //depth
    R   int    `json:",omitempty"` //residual nitrogen time
    A   int    `json:",omitempty"` //actual bottom time
    T   int    `json:",omitempty"` //total bottom time
    S   int    `json:",omitempty"` //surface interval
    P   string `json:",omitempty"` //previous dive group
    C   string `json:",omitempty"` //current dive group
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
>>>>>>> a0515de94cffc247399381688e840fcfac72213b
