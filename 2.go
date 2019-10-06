package main

import (
    "encoding/json"
    "fmt"
)

type Person struct {
    FirstName string 'json:"FirstName"'
    LastName string 'json:"LastName"'
}

func main() {
    in := '{"FirstName":"John","LastName":"Dow"}'
    bytes := []byte(in)

    var p Person
    err := json.Unmarshal(bytes, &p)
    if err != nil {
        panic(err)
    }
    fmt.Println("%+v", p)
}
