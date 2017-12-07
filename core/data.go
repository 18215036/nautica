package core

type DiveSeq struct {
    ID     string //Dive Sequence ID
    Dive1  Dive
    Dive2  Dive
    Dive3  Dive
    Dive4  Dive
    Dive5  Dive
}

type Dive struct { //individual dive data
    D   int    //depth
    R   int    //residual nitrogen time
    A   int    //actual bottom time
    T   int    //total bottom time
    S   int    //surface interval
    P   string //previous dive group
    C   string //current dive group
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