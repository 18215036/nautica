<<<<<<< HEAD
//18215036/Ahmad Fawwaz Zuhdi
package core

//dive tables initializer
//includes functions for file read/write operations with parsing

import (
    //"log"
    "bufio"
    "fmt"
    "os"
    "strings"
    "strconv"
    "encoding/json"
)

func init1() { //NDT Table Initializer
    data1, _  := ReadLines("core/D1")
    for i, line := range data1 {
        dataSlice := strings.Split(line, ",")
        j := 0
        T1[i].d3.DE, _ = strconv.Atoi(dataSlice[j])
        j++
        T1[i].NSL, _ = strconv.Atoi(dataSlice[j])
        j++
        T1[i].d3.d2.A, _ = strconv.Atoi(dataSlice[j])
        j++
        T1[i].d3.d2.B, _ = strconv.Atoi(dataSlice[j])
        j++
        T1[i].d3.d2.C, _ = strconv.Atoi(dataSlice[j])
        j++
        T1[i].d3.d2.D, _ = strconv.Atoi(dataSlice[j])
        j++
        T1[i].d3.d2.E, _ = strconv.Atoi(dataSlice[j])
        j++
        T1[i].d3.d2.F, _ = strconv.Atoi(dataSlice[j])
        j++
        T1[i].d3.d2.G, _ = strconv.Atoi(dataSlice[j])
        j++
        T1[i].d3.d2.H, _ = strconv.Atoi(dataSlice[j])
        j++
        T1[i].d3.d2.I, _ = strconv.Atoi(dataSlice[j])
        j++
        T1[i].d3.d2.J, _ = strconv.Atoi(dataSlice[j])
        j++
        T1[i].d3.d2.K, _ = strconv.Atoi(dataSlice[j])
        j++
        T1[i].d3.d2.L, _ = strconv.Atoi(dataSlice[j])
        j++
        T1[i].d3.d2.M, _ = strconv.Atoi(dataSlice[j])
        j++
        T1[i].d3.d2.N, _ = strconv.Atoi(dataSlice[j])
        j++
        T1[i].d3.d2.O, _ = strconv.Atoi(dataSlice[j])
        j++
        T1[i].d3.d2.Z, _ = strconv.Atoi(dataSlice[j])
    }
}

func init2() { //SI Table Initializer
    data2, _ := ReadLines("core/D2")
   for i, line := range data2 {
        dataSlice := strings.Split(line, ",")
        j := 0
        T2[i].A, _ = strconv.Atoi(dataSlice[j])
        j++
        T2[i].B, _ = strconv.Atoi(dataSlice[j])
        j++
        T2[i].C, _ = strconv.Atoi(dataSlice[j])
        j++
        T2[i].D, _ = strconv.Atoi(dataSlice[j])
        j++
        T2[i].E, _ = strconv.Atoi(dataSlice[j])
        j++
        T2[i].F, _ = strconv.Atoi(dataSlice[j])
        j++
        T2[i].G, _ = strconv.Atoi(dataSlice[j])
        j++
        T2[i].H, _ = strconv.Atoi(dataSlice[j])
        j++
        T2[i].I, _ = strconv.Atoi(dataSlice[j])
        j++
        T2[i].J, _ = strconv.Atoi(dataSlice[j])
        j++
        T2[i].K, _ = strconv.Atoi(dataSlice[j])
        j++
        T2[i].L, _ = strconv.Atoi(dataSlice[j])
        j++
        T2[i].M, _ = strconv.Atoi(dataSlice[j])
        j++
        T2[i].N, _ = strconv.Atoi(dataSlice[j])
        j++
        T2[i].O, _ = strconv.Atoi(dataSlice[j])
        j++
        T2[i].Z, _ = strconv.Atoi(dataSlice[j])
    }
}

func init3() { //RNT Table Initializer
    data3, _ := ReadLines("core/D3")
    for i, line := range data3 {
        dataSlice := strings.Split(line, ",")
        j := 0
        T3[i].DE, _ = strconv.Atoi(dataSlice[j])
        j++
        T3[i].d2.Z, _ = strconv.Atoi(dataSlice[j])
        j++
        T3[i].d2.O, _ = strconv.Atoi(dataSlice[j])
        j++
        T3[i].d2.N, _ = strconv.Atoi(dataSlice[j])
        j++
        T3[i].d2.M, _ = strconv.Atoi(dataSlice[j])
        j++
        T3[i].d2.L, _ = strconv.Atoi(dataSlice[j])
        j++
        T3[i].d2.K, _ = strconv.Atoi(dataSlice[j])
        j++
        T3[i].d2.J, _ = strconv.Atoi(dataSlice[j])
        j++
        T3[i].d2.I, _ = strconv.Atoi(dataSlice[j])
        j++
        T3[i].d2.H, _ = strconv.Atoi(dataSlice[j])
        j++
        T3[i].d2.G, _ = strconv.Atoi(dataSlice[j])
        j++
        T3[i].d2.F, _ = strconv.Atoi(dataSlice[j])
        j++
        T3[i].d2.E, _ = strconv.Atoi(dataSlice[j])
        j++
        T3[i].d2.D, _ = strconv.Atoi(dataSlice[j])
        j++
        T3[i].d2.C, _ = strconv.Atoi(dataSlice[j])
        j++
        T3[i].d2.B, _ = strconv.Atoi(dataSlice[j])
        j++
        T3[i].d2.A, _ = strconv.Atoi(dataSlice[j])
    }
}

func init() {
	init1()
	init2()
	init3()
}

// readLines reads a whole file into memory
// and returns a slice of its lines.
func ReadLines(path string) ([]string, error) {
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
func WriteLines(lines []string, path string) error {
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

/* probably not needed
// parsing json to struct
func parseJSON(req string) []Dive {
    var dives []Dive
    err := json.Unmarshal(req, &dives)
    if err != nil {
        fmt.Println("error:", err)
        os.Exit(1)
    }
    return dives
}
*/

//unparsing struct to json
func unparseJSON(dive []Dive) string {
    res, err := json.Marshal(dive)
    if err != nil {
        fmt.Println("error:", err)
        os.Exit(1)
    }
    return string(res)
}
=======
package core

//dive tables initializer
//includes functions for file read/write operations with parsing

import (
    //"log"
    "bufio"
    "fmt"
    "os"
    "strings"
    "strconv"
    "encoding/json"
)

func init1() { //NDT Table Initializer
    data1, _  := ReadLines("core/D1")
    for i, line := range data1 {
        dataSlice := strings.Split(line, ",")
        j := 0
        T1[i].d3.DE, _ = strconv.Atoi(dataSlice[j])
        j++
        T1[i].NSL, _ = strconv.Atoi(dataSlice[j])
        j++
        T1[i].d3.d2.A, _ = strconv.Atoi(dataSlice[j])
        j++
        T1[i].d3.d2.B, _ = strconv.Atoi(dataSlice[j])
        j++
        T1[i].d3.d2.C, _ = strconv.Atoi(dataSlice[j])
        j++
        T1[i].d3.d2.D, _ = strconv.Atoi(dataSlice[j])
        j++
        T1[i].d3.d2.E, _ = strconv.Atoi(dataSlice[j])
        j++
        T1[i].d3.d2.F, _ = strconv.Atoi(dataSlice[j])
        j++
        T1[i].d3.d2.G, _ = strconv.Atoi(dataSlice[j])
        j++
        T1[i].d3.d2.H, _ = strconv.Atoi(dataSlice[j])
        j++
        T1[i].d3.d2.I, _ = strconv.Atoi(dataSlice[j])
        j++
        T1[i].d3.d2.J, _ = strconv.Atoi(dataSlice[j])
        j++
        T1[i].d3.d2.K, _ = strconv.Atoi(dataSlice[j])
        j++
        T1[i].d3.d2.L, _ = strconv.Atoi(dataSlice[j])
        j++
        T1[i].d3.d2.M, _ = strconv.Atoi(dataSlice[j])
        j++
        T1[i].d3.d2.N, _ = strconv.Atoi(dataSlice[j])
        j++
        T1[i].d3.d2.O, _ = strconv.Atoi(dataSlice[j])
        j++
        T1[i].d3.d2.Z, _ = strconv.Atoi(dataSlice[j])
    }
}

func init2() { //SI Table Initializer
    data2, _ := ReadLines("core/D2")
   for i, line := range data2 {
        dataSlice := strings.Split(line, ",")
        j := 0
        T2[i].A, _ = strconv.Atoi(dataSlice[j])
        j++
        T2[i].B, _ = strconv.Atoi(dataSlice[j])
        j++
        T2[i].C, _ = strconv.Atoi(dataSlice[j])
        j++
        T2[i].D, _ = strconv.Atoi(dataSlice[j])
        j++
        T2[i].E, _ = strconv.Atoi(dataSlice[j])
        j++
        T2[i].F, _ = strconv.Atoi(dataSlice[j])
        j++
        T2[i].G, _ = strconv.Atoi(dataSlice[j])
        j++
        T2[i].H, _ = strconv.Atoi(dataSlice[j])
        j++
        T2[i].I, _ = strconv.Atoi(dataSlice[j])
        j++
        T2[i].J, _ = strconv.Atoi(dataSlice[j])
        j++
        T2[i].K, _ = strconv.Atoi(dataSlice[j])
        j++
        T2[i].L, _ = strconv.Atoi(dataSlice[j])
        j++
        T2[i].M, _ = strconv.Atoi(dataSlice[j])
        j++
        T2[i].N, _ = strconv.Atoi(dataSlice[j])
        j++
        T2[i].O, _ = strconv.Atoi(dataSlice[j])
        j++
        T2[i].Z, _ = strconv.Atoi(dataSlice[j])
    }
}

func init3() { //RNT Table Initializer
    data3, _ := ReadLines("core/D3")
    for i, line := range data3 {
        dataSlice := strings.Split(line, ",")
        j := 0
        T3[i].DE, _ = strconv.Atoi(dataSlice[j])
        j++
        T3[i].d2.Z, _ = strconv.Atoi(dataSlice[j])
        j++
        T3[i].d2.O, _ = strconv.Atoi(dataSlice[j])
        j++
        T3[i].d2.N, _ = strconv.Atoi(dataSlice[j])
        j++
        T3[i].d2.M, _ = strconv.Atoi(dataSlice[j])
        j++
        T3[i].d2.L, _ = strconv.Atoi(dataSlice[j])
        j++
        T3[i].d2.K, _ = strconv.Atoi(dataSlice[j])
        j++
        T3[i].d2.J, _ = strconv.Atoi(dataSlice[j])
        j++
        T3[i].d2.I, _ = strconv.Atoi(dataSlice[j])
        j++
        T3[i].d2.H, _ = strconv.Atoi(dataSlice[j])
        j++
        T3[i].d2.G, _ = strconv.Atoi(dataSlice[j])
        j++
        T3[i].d2.F, _ = strconv.Atoi(dataSlice[j])
        j++
        T3[i].d2.E, _ = strconv.Atoi(dataSlice[j])
        j++
        T3[i].d2.D, _ = strconv.Atoi(dataSlice[j])
        j++
        T3[i].d2.C, _ = strconv.Atoi(dataSlice[j])
        j++
        T3[i].d2.B, _ = strconv.Atoi(dataSlice[j])
        j++
        T3[i].d2.A, _ = strconv.Atoi(dataSlice[j])
    }
}

func init() {
	init1()
	init2()
	init3()
}

// readLines reads a whole file into memory
// and returns a slice of its lines.
func ReadLines(path string) ([]string, error) {
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
func WriteLines(lines []string, path string) error {
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

/* probably not needed
// parsing json to struct
func parseJSON(req string) []Dive {
    var dives []Dive
    err := json.Unmarshal(req, &dives)
    if err != nil {
        fmt.Println("error:", err)
        os.Exit(1)
    }
    return dives
}
*/

//unparsing struct to json
func unparseJSON(dive []Dive) string {
    res, err := json.Marshal(dive)
    if err != nil {
        fmt.Println("error:", err)
        os.Exit(1)
    }
    return string(res)
}
>>>>>>> a0515de94cffc247399381688e840fcfac72213b
