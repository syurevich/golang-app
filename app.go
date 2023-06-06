package main

import (
	"fmt"
)

func main() {

    painters := GetPainters()

    for _, painter := range painters {
        fmt.Println(painter)
    }
}

