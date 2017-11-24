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

func init() { //dataOne Initializer
    data1, err := readLines("./D1")
    for i := 0; i < 24; i++ {
        dataSlice := strings.Split(data1[i], ",")
        j := 0
        T1[i].d3.DE = dataSlice[j++]
        T1[i].NSL = dataSlice[j++]
        T1[i].d3.d2.A = dataSlice[j++]
        T1[i].d3.d2.B  = dataSlice[j++]
        T1[i].d3.d2.C  = dataSlice[j++]
        T1[i].d3.d2.D  = dataSlice[j++]
        T1[i].d3.d2.E  = dataSlice[j++]
        T1[i].d3.d2.F  = dataSlice[j++]
        T1[i].d3.d2.G  = dataSlice[j++]
        T1[i].d3.d2.H  = dataSlice[j++]
        T1[i].d3.d2.I  = dataSlice[j++]
        T1[i].d3.d2.J  = dataSlice[j++]
        T1[i].d3.d2.K  = dataSlice[j++]
        T1[i].d3.d2.L  = dataSlice[j++]
        T1[i].d3.d2.M  = dataSlice[j++]
        T1[i].d3.d2.N  = dataSlice[j++]
        T1[i].d3.d2.O  = dataSlice[j++]
        T1[i].d3.d2.Z  = dataSlice[j++]
    }
}

func init() { //dataTwo Initializer
}

func init() { //dataThree Initializer
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

func parseTables (path string) []