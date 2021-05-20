package hello

import (
    "fmt"
)

type person struct{
    Name string
    Age int
}

func hello(){
    p1:= person{
        Name: "PK",
        Age:36,
    }
    fmt.Println(p1)
}