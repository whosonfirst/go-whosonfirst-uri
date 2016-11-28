package uri

import (
	"log"
	"path/filepath"
	"strconv"
	"strings"
)

type URIArgs struct {
	Alternate bool
	Source    string
	Function  string
	Extras    []string
	Strict    bool
}

func NewDefaultURIArgs() *URIArgs {

	u := URIArgs{
		Alternate: false,
		Source:    "",
		Function:  "",
		Extras:    make([]string, 0),
		Strict:    false,
	}

	return &u
}

type URI struct {
	Root string
	Args *URIArgs
}

func NewDefaultURI(root string) *URI {

	args := NewDefaultURIArgs()

	u := URI{
		Args: args,
		Root: root,
	}

	return &u
}

# See also: https://github.com/whosonfirst/whosonfirst-cookbook/blob/master/how_to/creating_alt_geometries.md

func Id2Fname(id int, args ...*URIArgs) string {

	str_id := strconv.Itoa(id)
	parts := []string{str_id}

	if len(args) == 1 {

		uri_args := args[0]

		log.Println(uri_args.Alternate)
	}

	str_parts := strings.Join(parts, "-")

	fname := str_parts + ".geojson"
	return fname
}

func Id2Path(id int) string {

	parts := []string{}
	input := strconv.Itoa(id)

	for len(input) > 3 {

		chunk := input[0:3]
		input = input[3:]
		parts = append(parts, chunk)
	}

	if len(input) > 0 {
		parts = append(parts, input)
	}

	path := filepath.Join(parts...)
	return path
}

func Id2RelPath(id int, args ...*URIArgs) string {

	fname := Id2Fname(id, args...)
	root := Id2Path(id)

	rel_path := filepath.Join(root, fname)
	return rel_path
}

func Id2AbsPath(root string, id int, args ...*URIArgs) string {

	rel := Id2RelPath(id, args...)

	abs_path := filepath.Join(root, rel)
	return abs_path
}
