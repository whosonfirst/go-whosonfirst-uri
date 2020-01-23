package uri

import (
	"errors"
	"github.com/whosonfirst/go-whosonfirst-sources"
	_ "log"
	"net/url"
	"path/filepath"
	"strconv"
	"strings"
)

type URIArgs struct {
	// PLEASE UPDATE THIS TO USE/EXPECT AN *AltGeom KTHXBYE (20190501/thisisaaronland)
	Alternate bool
	Source    string
	Function  string
	Extras    []string
	Strict    bool
}

type AltGeom struct {
	Source   string
	Function string
	Extras   []string
}

func (a *AltGeom) String() string {

	parts := []string{
		a.Source,
	}

	if a.Function != "" {
		parts = append(parts, a.Function)
	}

	for _, ex := range a.Extras {
		parts = append(parts, ex)
	}

	return strings.Join(parts, "-")
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

func NewAlternateURIArgs(source string, function string, extras ...string) *URIArgs {

	u := URIArgs{
		Alternate: true,
		Source:    source,
		Function:  function,
		Extras:    extras,
		Strict:    false,
	}

	return &u
}

// See also: https://github.com/whosonfirst/whosonfirst-cookbook/blob/master/how_to/creating_alt_geometries.md

func Id2Fname(id int64, args ...*URIArgs) (string, error) {

	str_id := strconv.FormatInt(id, 10)
	parts := []string{str_id}

	if len(args) == 1 {

		uri_args := args[0]

		if uri_args.Alternate {

			if uri_args.Source == "" && uri_args.Strict {
				return "", errors.New("Missing source argument for alternate geometry")
			}

			if uri_args.Source == "" {
				uri_args.Source = "unknown"

			}

			if uri_args.Strict && !sources.IsValidSource(uri_args.Source) {
				return "", errors.New("Invalid or unknown source argument for alternate geometry")
			}

			parts = append(parts, "alt")
			parts = append(parts, uri_args.Source)

			if uri_args.Function != "" {
				parts = append(parts, uri_args.Function)
			}

			for _, e := range uri_args.Extras {
				parts = append(parts, e)
			}
		}

	}

	str_parts := strings.Join(parts, "-")

	fname := str_parts + ".geojson"
	return fname, nil
}

func Id2Path(id int64) (string, error) {

	parts := []string{}
	input := strconv.FormatInt(id, 10)

	for len(input) > 3 {

		chunk := input[0:3]
		input = input[3:]
		parts = append(parts, chunk)
	}

	if len(input) > 0 {
		parts = append(parts, input)
	}

	path := filepath.Join(parts...)
	return path, nil
}

func Id2RelPath(id int64, args ...*URIArgs) (string, error) {

	fname, err := Id2Fname(id, args...)

	if err != nil {
		return "", err
	}

	root, err := Id2Path(id)

	if err != nil {
		return "", err
	}

	rel_path := filepath.Join(root, fname)
	return rel_path, nil
}

func Id2AbsPath(root string, id int64, args ...*URIArgs) (string, error) {

	rel, err := Id2RelPath(id, args...)

	if err != nil {
		return "", err
	}

	var abs_path string

	// because filepath.Join will screw up scheme URIs
	// (20170124/thisisaaronland)

	_, err = url.Parse(root)

	if err == nil {

		if !strings.HasSuffix(root, "/") {
			root += "/"
		}

		abs_path = root + rel

	} else {
		abs_path = filepath.Join(root, rel)
	}

	return abs_path, nil
}
