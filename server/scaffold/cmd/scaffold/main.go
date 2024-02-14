package main

import (
	"fmt"
	"scaffold/internal/helpers"
	"scaffold/internal/lib"
	//"github.com/sijirama/scaffold/internal/helpers"
)

func main() {
    //helpers.Helloworld("sijibomi")
    
    vara := lib.HelloWorld("sijibomi")
    va := lib.WhatTheActualFuck()
    fmt.Println(vara , va)
    helpers.Helloworld("sijibomi")
}
