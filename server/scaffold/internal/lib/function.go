package lib

import "fmt"

func HelloWorld(name string) (welcome string) {
	welcome = fmt.Sprintf("Hello world to %v" , name)
    return welcome
}
