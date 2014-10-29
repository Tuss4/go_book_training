package main

import (
    "fmt"
    "strings"
    "os"
    "io/ioutil"
    "path/filepath"
    "errors"
    "container/list"
    "sort"
    "hash/crc32"
)

func main() {
    fmt.Println(
        // strings
        strings.Contains("test", "es"),
        strings.Count("test", "t"),
        strings.HasPrefix("test", "te"),
        strings.HasSuffix("test", "st"),
        strings.Index("test", "1"),
        strings.Join([]string{"a","b"}, "-"),
        strings.Repeat("a", 5),
        strings.Replace("aaaa", "a", "b", 2),
        strings.Split("a-b-c-d-e", "-"),
        strings.ToLower("TEST"),
        strings.ToUpper("test"),
    )
    // Convert a string to bytes print out the bytes in binary.
    arr := []byte("test")
    fmt.Println(arr)
    for _, b := range arr {
        fmt.Printf("%b\n", b)
    }

    // Files
    bs, err := ioutil.ReadFile("test.txt")
    if err != nil {
        return
    }
    str := string(bs)
    fmt.Println(str)

    file, err := os.Create("test_2.txt")
    if err != nil {
        return
    }
    defer file.Close()

    file.WriteString("whimmy wham wham wazzle")

    dir, err := os.Open(".")
    if err != nil {
        return
    }
    defer dir.Close()

    fileInfos, err := dir.Readdir(-1)
    if err != nil {
        return
    }

    for _, fi := range fileInfos {
        fmt.Println(fi.Name())
    }

    filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
        fmt.Println(path)
        return nil
        })

    da_err := errors.New("I'm not really upset.")
    fmt.Println(da_err)

    // Containers and sort
    var x list.List
    x.PushBack(1)
    x.PushBack(2)
    x.PushBack(3)

    for e := x.Front(); e != nil; e=e.Next() {
        fmt.Println(e.Value.(int))
    }

    kids := []Person{
        {"Jill", 9},
        {"Jack", 10},
    }

    sort.Sort(ByName(kids))
    fmt.Println(kids)
    sort.Sort(ByAge(kids))
    fmt.Println(kids)

    h1, err := getHash("test1.txt")
    if err != nil {
        return
    }
    h2, err := getHash("test2.txt")
    if err != nil {
        return
    }
    fmt.Println(h1, h2, h1 == h2)
}

// Sort
type Person struct {
    Name string
    Age int
}

type ByName []Person

func (this ByName) Len() int {
    return len(this)
}

func (this ByName) Less(i, j int) bool {
    return this[i].Name < this[j].Name
}

func (this ByName) Swap(i, j int) {
    this[i], this[j] = this[j], this[i]
}

type ByAge []Person
func (this ByAge) Len() int {
    return len(this)
}

func (this ByAge) Less(i, j int) bool {
    return this[i].Age < this[j].Age
}

func (this ByAge) Swap(i, j int) {
    this[i], this[j] = this[j], this[i]
}

func getHash(filename string) (uint32, error) {
    bs, err := ioutil.ReadFile(filename)
    if err != nil {
        return 0, err
    }
    h := crc32.NewIEEE()
    h.Write(bs)
    return h.Sum32(), nil
}
