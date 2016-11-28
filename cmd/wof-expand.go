package main

import (
	"flag"
	"fmt"
	"github.com/whosonfirst/go-whosonfirst-uri"
	"log"
	"path"
	"strconv"
)

func main() {

	var root = flag.String("root", "", "A root directory for absolute paths")
	var prefix = flag.String("prefix", "", "Prepend this prefix to all paths")

	var alt = flag.Bool("alternate", false, "...")
	var source = flag.String("source", "", "...")
	var function = flag.String("function", "", "...")
	var strict = flag.Bool("strict", false, "...")

	flag.Parse()

	for _, str_id := range flag.Args() {

		id, err := strconv.Atoi(str_id)

		if err != nil {
			log.Fatal("Unable to parse %s, because %v", str_id, err)
		}

		var args *uri.URIArgs

		if *alt {

			args = uri.NewAlternateURIArgs(*source, *function)
			args.Strict = *strict

			// to do: extras

		} else {
			args = uri.NewDefaultURIArgs()
		}

		wof_path, err := uri.Id2RelPath(id, args)

		if err != nil {
			log.Printf("failed to generate a URI for %s, because '%v'\n", str_id, err)
			continue
		}

		if *prefix != "" {
			wof_path = path.Join(*prefix, wof_path)
		}

		if *root != "" {
			wof_path = path.Join(*root, wof_path)
		}

		fmt.Println(wof_path)
	}
}
