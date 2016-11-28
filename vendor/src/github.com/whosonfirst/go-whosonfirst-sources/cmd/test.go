package main

import (
	"fmt"
	"github.com/whosonfirst/go-whosonfirst-sources"
)

func main() {

	s, _ := sources.Spec()
	fmt.Printf("%v\n", s)
}
