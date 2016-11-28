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

	flag.Parse()

	for _, str_id := range flag.Args() {
		id, err := strconv.Atoi(str_id)

		if err != nil {
			log.Fatal("Unable to parse %s, because %v", str_id, err)
		}

		args := uri.NewDefaultURIArgs()

		wof_path := uri.Id2RelPath(id, args)

		if *prefix != "" {
			wof_path = path.Join(*prefix, wof_path)
		}

		if *root != "" {
			wof_path = path.Join(*root, wof_path)
		}

		fmt.Println(wof_path)
	}
}
