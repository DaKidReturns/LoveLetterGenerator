package main

import (
    llc "lovelettercreator"
    "fmt"
)

func main(){
    loveletter := llc.CreateLetter(false, "","")
    fmt.Println(loveletter)
}
