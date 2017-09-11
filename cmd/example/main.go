package main

import (
	"fmt"

	"github.com/matzhouse/example-go-circle/pkg/hello"
)

func main() {

	greet := hello.Hello("mat")

	fmt.Println(greet)

}
