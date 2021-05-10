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

// URIArgs is a struct indicating whether or not a URI is considered an "alternate" geometry and specific details if it is.
type URIArgs struct {
	// Boolean value indicating whether or not a URI is considered an "alternate" geometry
	IsAlternate bool
	// And *AltGeom struct containing details about an "alternate" geometry
	AltGeom *AltGeom
}

// AltGeom is a struct containing details about an "alternate" geometry
type AltGeom struct {
	// The source of the alternate geometry. This value is required and SHOULD match a corresponding entry in the whosonfirst/whosonfirst-sources repository.
	Source string
	// The labeled function for the alternate geometry. This value MAY be a controlled value relative to `Source`.
	Function string
	// A list of optional strings to append to the alternate geometry's URI.
	Extras []string
	// A boolean value used to indicate whether the `Source` value has a corresponding entry in the whosonfirst/whosonfirst-sources repository.
	Strict bool
}

// Return the string value for an "alternate" geometry.
func (a *AltGeom) String() (string, error) {

	source := a.Source

	if a.Strict && source == "" {
		return "", errors.New("Missing source argument for alternate geometry")
	}

	if source == "" {
		source = "unknown"

	}

	if a.Strict && !sources.IsValidSource(source) {
		return "", errors.New("Invalid or unknown source argument for alternate geometry")
	}

	parts := []string{
		source,
	}

	if a.Function != "" {
		parts = append(parts, a.Function)
	}

	for _, ex := range a.Extras {
		parts = append(parts, ex)
	}

	alt_str := strings.Join(parts, "-")

	return alt_str, nil
}

// Return a URIArgs struct whose IsAlternate flag is false.
func NewDefaultURIArgs() *URIArgs {

	alt_geom := &AltGeom{}

	u := URIArgs{
		IsAlternate: false,
		AltGeom:     alt_geom,
	}

	return &u
}

func NewAlternateURIArgs(source string, function string, extras ...string) *URIArgs {

	alt_geom := &AltGeom{
		Source:   source,
		Function: function,
		Extras:   extras,
	}

	u := URIArgs{
		IsAlternate: true,
		AltGeom:     alt_geom,
	}

	return &u
}

// See also: https://github.com/whosonfirst/whosonfirst-cookbook/blob/master/how_to/creating_alt_geometries.md

// Id2Fname parses a Who's On First
func Id2Fname(id int64, args ...*URIArgs) (string, error) {

	str_id := strconv.FormatInt(id, 10)
	parts := []string{str_id}

	if len(args) == 1 {

		uri_args := args[0]

		if uri_args.IsAlternate {

			alt_str, err := uri_args.AltGeom.String()

			if err != nil {
				return "", err
			}

			parts = append(parts, "alt")
			parts = append(parts, alt_str)
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
