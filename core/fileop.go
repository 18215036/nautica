package core

//dive tables initializer
//includes functions for file read/write operations with parsing

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strings"
)

func init() { //NDT Table Initializer
    data1, err := readLines("./D1")
    for i := 0; i < 24; i++ {
        dataSlice := strings.Split(data1[i], ",")
        j := 0
        T1[i].d3.DE = dataSlice[j]
        j++
        T1[i].NSL = dataSlice[j]
        j++
        T1[i].d3.d2.A = dataSlice[j]
        j++
        T1[i].d3.d2.B  = dataSlice[j]
        j++
        T1[i].d3.d2.C  = dataSlice[j]
        j++
        T1[i].d3.d2.D  = dataSlice[j]
        j++
        T1[i].d3.d2.E  = dataSlice[j]
        j++
        T1[i].d3.d2.F  = dataSlice[j]
        j++
        T1[i].d3.d2.G  = dataSlice[j]
        j++
        T1[i].d3.d2.H  = dataSlice[j]
        j++
        T1[i].d3.d2.I  = dataSlice[j]
        j++
        T1[i].d3.d2.J  = dataSlice[j]
        j++
        T1[i].d3.d2.K  = dataSlice[j]
        j++
        T1[i].d3.d2.L  = dataSlice[j]
        j++
        T1[i].d3.d2.M  = dataSlice[j]
        j++
        T1[i].d3.d2.N  = dataSlice[j]
        j++
        T1[i].d3.d2.O  = dataSlice[j]
        j++
        T1[i].d3.d2.Z  = dataSlice[j]
    }
}

func init() { //SI Table Initializer
    data2, err := readLines("./D2")
    for i := 0; i < 16; i++ {
        dataSlice := strings.Split(data2[i], ",")
        j := 0
        T2[i].A = dataSlice[j]
        j++
        T2[i].B  = dataSlice[j]
        j++
        T2[i].C  = dataSlice[j]
        j++
        T2[i].D  = dataSlice[j]
        j++
        T2[i].E  = dataSlice[j]
        j++
        T2[i].F  = dataSlice[j]
        j++
        T2[i].G  = dataSlice[j]
        j++
        T2[i].H  = dataSlice[j]
        j++
        T2[i].I  = dataSlice[j]
        j++
        T2[i].J  = dataSlice[j]
        j++
        T2[i].K  = dataSlice[j]
        j++
        T2[i].L  = dataSlice[j]
        j++
        T2[i].M  = dataSlice[j]
        j++
        T2[i].N  = dataSlice[j]
        j++
        T2[i].O  = dataSlice[j]
        j++
        T2[i].Z  = dataSlice[j]
    }
}

func init() { //RNT Table Initializer
    data3, err := readLines("./D3")
    for i := 0; i < 24; i++ {
        dataSlice := strings.Split(data3[i], ",")
        j := 0
        T3[i].DE = dataSlice[j]
        j++
        T3[i].d2.Z = dataSlice[j]
        j++
        T3[i].d2.O  = dataSlice[j]
        j++
        T3[i].d2.N  = dataSlice[j]
        j++
        T3[i].d2.M  = dataSlice[j]
        j++
        T3[i].d2.L  = dataSlice[j]
        j++
        T3[i].d2.K  = dataSlice[j]
        j++
        T3[i].d2.J  = dataSlice[j]
        j++
        T3[i].d2.I  = dataSlice[j]
        j++
        T3[i].d2.H  = dataSlice[j]
        j++
        T3[i].d2.G  = dataSlice[j]
        j++
        T3[i].d2.F  = dataSlice[j]
        j++
        T3[i].d2.E  = dataSlice[j]
        j++
        T3[i].d2.D  = dataSlice[j]
        j++
        T3[i].d2.C  = dataSlice[j]
        j++
        T3[i].d2.B  = dataSlice[j]
        j++
        T3[i].d2.A  = dataSlice[j]
    }
}

// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(path string) ([]string, error) {
    file, err := os.Open(path)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    var lines []string
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        lines = append(lines, scanner.Text())
    }
    return lines, scanner.Err()
}

// writeLines writes the lines to the given file.
func writeLines(lines []string, path string) error {
    file, err := os.Create(path)
    if err != nil {
        return err
    }
    defer file.Close()

    w := bufio.NewWriter(file)
    for _, line := range lines {
        fmt.Fprintln(w, line)
    }
    return w.Flush()
}