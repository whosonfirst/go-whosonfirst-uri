package main

import (
	"flag"
	"fmt"
	"github.com/whosonfirst/go-whosonfirst-uri"	
	"log"	
)

func main() {

	flag.Parse()

	for _, path := range flag.Args() {

		ok, err := uri.IsAltFile(path)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(path, ok)
		
		alt, err := uri.AltGeomFromPath(path)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(alt)
	}
}
