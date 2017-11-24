package core

import{
    "math"
}

//Dive Table Calculator (based on US DIVING MANUAL REV.7
//created by: 18215036 / Ahmad Fawwaz Zuhdi

//NO DECO DIVE HANDLING at the moment, will be implemented
//Add another dive (y/n)?Add another dive (y/n)? //problem? brought from C

func mToF(m int) int {
    return Ceil(3.28084*m)
}

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