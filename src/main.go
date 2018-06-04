package main

import "./db"
import "fmt"

func main() {
    fmt.Println(db.Read("a"))
}
