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
    bytes, err := json.Marshal(person{
        FirstName: "John",
        LasName: "Dow",
    })
    if err != nil {
        panic(err)
    }
    fmt.Println(string(bytes))
}
