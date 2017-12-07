package core

import (
    "math"
)

//Dive Table Calculator (based on US DIVING MANUAL REV.7
//created by: 18215036 / Ahmad Fawwaz Zuhdi

//NO DECO DIVE HANDLING at the moment, will be implemented
//Add another dive (y/n)?Add another dive (y/n)? //problem? brought from C

var T1 [24]DATA1
var T2 [16]DATA2
var T3 [24]DATA3

//convert meter to feet
func mToF(m int) int {
    return int(math.Ceil(3.28084*(float64(m))))
}

//convert group number to group letter
func nToL(n int) string {
    switch n {
    case 1:
        return "A"
    case 2:
        return "B"
    case 3:
        return "C"
    case 4:
        return "D"
    case 5:
        return "E"
    case 6:
        return "F"
    case 7:
        return "G"
    case 8:
        return "H"
    case 9:
        return "I"
    case 10:
        return "J"
    case 11:
        return "K"
    case 12:
        return "L"
    case 13:
        return "M"
    case 14:
        return "N"
    case 15:
        return "O"
    case 16:
        return "Z"
    default:
        return "0"
    }
}

//TABLE 1 FUNCTION //Dive Group after Dive# //-6 = deco dive
//depth in meters, total in minutes
func F1(depth int, total int) string {
    i := 0
    ft := mToF(depth)

    for ft > T1[i].d3.DE {
        i++
    }

    switch {
    case total > T1[i].NSL :
        return "-6"
    case total <= T1[i].d3.d2.A :
        return "A"
    case total <= T1[i].d3.d2.B :
        return "B"
    case total <= T1[i].d3.d2.C :
        return "C"
    case total <= T1[i].d3.d2.D :
        return "D"
    case total <= T1[i].d3.d2.E :
        return "E"
    case total <= T1[i].d3.d2.F :
        return "F"
    case total <= T1[i].d3.d2.G :
        return "G"
    case total <= T1[i].d3.d2.H :
        return "H"
    case total <= T1[i].d3.d2.I :
        return "I"
    case total <= T1[i].d3.d2.J :
        return "J"
    case total <= T1[i].d3.d2.K :
        return "K"
    case total <= T1[i].d3.d2.L :
        return "L"
    case total <= T1[i].d3.d2.M :
        return "M"
    case total <= T1[i].d3.d2.N :
        return "N"
    case total <= T1[i].d3.d2.O :
        return "O"
    default :
        return "Z"
    }
}

//TABLE 2 FUNCTION //Dive Group after Surface Interval //-4 = not a repetitive dive //-5 = SI<10
func F2(pGroup string, S int) string {
    if S < 10 {
        return "-5"
    }

    i := 0
    for nToL(i+1) != pGroup {
        i++
    }

    nGroup := "-4"
    if S <= T2[i].A {
        if nGroup = "A"; S <= T2[i].B {
            if nGroup = "B"; S <= T2[i].C {
                if nGroup = "C"; S <= T2[i].D {
                    if nGroup = "D"; S <= T2[i].E {
                        if nGroup = "E"; S <= T2[i].F {
                            if nGroup = "F"; S <= T2[i].G {
                                if nGroup = "G"; S <= T2[i].H {
                                    if nGroup = "H"; S <= T2[i].I {
                                        if nGroup = "I"; S <= T2[i].J {
                                            if nGroup = "J"; S <= T2[i].K {
                                                if nGroup = "K"; S <= T2[i].L {
                                                    if nGroup = "L"; S <= T2[i].M {
                                                        if nGroup = "M"; S <= T2[i].N {
                                                            if nGroup = "N"; S <= T2[i].O {
                                                                if nGroup = "O"; S <= T2[i].Z {
                                                                    nGroup = "Z"
                                                                }
                                                            }
                                                        }
                                                    }
                                                }
                                            }
                                        }
                                    }
                                }
                            }
                        }
                    }
                }
            }
        }
    }
    return nGroup
}

//TABLE 3 FUNCTION //RNT for Next Dive //-2? //-3?
func F3(pGroup string, nDepth int) int {
    i := 0
    ft := mToF(nDepth)
    
    for ft > T3[i].DE {
        i++
    }

    switch pGroup {
    case "A":
        return T3[i].d2.A
    case "B":
        return T3[i].d2.B
    case "C":
        return T3[i].d2.C
    case "D":
        return T3[i].d2.D
    case "E":
        return T3[i].d2.E
    case "F":
        return T3[i].d2.F
    case "G":
        return T3[i].d2.G
    case "H":
        return T3[i].d2.H
    case "I":
        return T3[i].d2.I
    case "J":
        return T3[i].d2.J
    case "K":
        return T3[i].d2.K
    case "L":
        return T3[i].d2.L
    case "M":
        return T3[i].d2.M
    case "N":
        return T3[i].d2.N
    case "O":
        return T3[i].d2.O
    case "Z":
        return T3[i].d2.Z
    default:
        return 0
    }
}

func Calculate(sequence DiveSeq) (DiveSeq, string) {
    d := []Dive{sequence.Dive1, sequence.Dive2, sequence.Dive3, sequence.Dive4, sequence.Dive5}
    i := 0
    firstDive := true
    decoDive := false

    for !decoDive && i < 5 && &(d[i]) != nil {
    	if firstDive {
    		d[i].P = "0"
    		d[i].R = 0
    		firstDive = !firstDive
    	} else {
            d[i].P = F2(d[i-1].C, d[i-1].S)
            if d[i].P == "-4" {
            	d[i].P = "0"
            	d[i].R = 0
            } else {
            	if d[i].P == "-5" {
            		if d[i].D <= d[i-1].D { d[i].D = d[i-1].D }
					d[i].A += d[i-1].A
					d[i].S += d[i-1].S	//added Surface Interval
					d[i].P = d[i-1].P
            	}
            	d[i].R = F3(d[i].P, d[i].D)
            }
    	}

    	if d[i].R == -2 {
            d[i].T = d[i].A
            d[i].C = d[i].P
    	} else if d[i].R == -3 {
            d[i].T = -6  //deco
            d[i].C = "-6"
    	} else {
            d[i].T = d[i].R + d[i].A
            d[i].C = F1(d[i].D, d[i].T)
    	}

    	if d[i].C == "-6" { decoDive = true }
    	i++
    }
    sequence.Dive1 = d[0]
    sequence.Dive2 = d[1]
    sequence.Dive3 = d[2]
    sequence.Dive4 = d[3]
    sequence.Dive5 = d[4]

    return sequence, unparseJSON(d)
}

/* -1 : Unlimited No-Stop Limit */
/* -2 : Residual Nitrogen Time cannot be determined using this table (see paragraph 9-9.1 subparagraph 8 for instructions).
 * At depths of 10, 15, and 20 fsw, some of the higher repetitive groups do not
have a defined residual nitrogen time. These groups are marked with a double
asterisk in the lower half of Table 9-8. The RNT is undefined because the tissue
nitrogen loading associated with those repetitive groups is higher than the
nitrogen loading that could be achieved even if the diver were to remain at
those depths for an infinite period of time. A diver entering the dive in one of
those higher groups marked by a double asterisk can still perform a repetitive
dive at 10, 15 or 20 fsw because the no-decompression time at those depths
is unlimited. An RNT time is not required to make the dive. If a subsequent
repetitive dive to a deeper depth is planned, however, the diver will need a
repetitive group at the end of the shallow dive in order to continue using the
RNT table. If a double asterisk is encountered in Table 9-8, assume that the
repetitive group remains unchanged during the course of the dive at 10, 15, or
20 fsw.*/
/* -3 : Read vertically downward to the 30 fsw repetitive dive depth. Use the corresponding residual nitrogen times to compute the
equivalent single dive time. Decompress using the 30 fsw air decompression table. */
