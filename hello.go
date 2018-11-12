package main

import (
	"fmt"
	"github.com/you/hello/bar"
	"github.com/you/hello/foo"
	"rsc.io/quote"
)

func main() {
	fmt.Println(quote.Hello())
	fmt.Println(foo.Hello())
	fmt.Println(bar.Hello())
}
