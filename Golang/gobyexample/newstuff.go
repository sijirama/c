package main

import "fmt"

type Human struct {
	complexion string
}

func conat (first *string , second string) {
    *first = *first + second
}

func cool() {
    fmt.Println("Hello world")
    var str string=  "Ilesanmi"
    var copy string
	for i := 0; i < len(str); i++ {
        defer conat(&copy , string(str[i]))
	}
    fmt.Println(copy)
}
