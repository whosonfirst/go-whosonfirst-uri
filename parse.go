package uri

import (
	"errors"
	_ "log"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

type ParsedURI struct {
	Id          int64
	URI         string
	URIArgs     *URIArgs
	IsAlternate bool
}

var re_uri *regexp.Regexp

func init() {
	re_uri = regexp.MustCompile(`^(\d+)(?:\-alt(?:\-([a-zA-Z0-9_]+(?:\-[a-zA-Z0-9_]+(?:\-[a-zA-Z0-9_\-]+)?)?)))?(?:\.[^\.]+|\/)?$`)
}

func ParseURI(path string) (int64, *URIArgs, error) {

	abs_path, err := filepath.Abs(path)

	if err != nil {
		return -1, nil, err
	}

	fname := filepath.Base(abs_path)

	match := re_uri.FindStringSubmatch(fname)

	// log.Println(fname, match)

	if len(match) == 0 {
		return -1, nil, errors.New("Unable to parse WOF ID")
	}

	if len(match) < 2 {
		return -1, nil, errors.New("Unable to parse WOF ID")
	}

	str_id := match[1]
	str_alt := match[2]

	wofid, err := strconv.ParseInt(str_id, 10, 64)

	if err != nil {
		return -1, nil, err
	}

	args := &URIArgs{
		Alternate: false,
	}

	if str_alt != "" {

		args.Alternate = true

		alt := strings.Split(str_alt, "-")

		switch len(alt) {
		case 1:
			args.Source = alt[0]
		case 2:
			args.Source = alt[0]
			args.Function = alt[1]
		default:
			args.Source = alt[0]
			args.Function = alt[1]
			args.Extras = alt[2:]
		}
	}

	return wofid, args, nil
}

//

func IsWOFFile(path string) (bool, error) {

	_, _, err := ParseURI(path)

	if err != nil {
		return false, nil
	}

	return true, nil
}

func IsAltFile(path string) (bool, error) {

	_, uri_args, err := ParseURI(path)

	if err != nil {
		return false, err
	}

	is_alt := uri_args.Alternate
	return is_alt, nil
}

func AltGeomFromPath(path string) (*AltGeom, error) {

	_, uri_args, err := ParseURI(path)

	if err != nil {
		return nil, err
	}

	if !uri_args.Alternate {
		return nil, errors.New("Not an alternate geometry")
	}

	alt := &AltGeom{
		Source:   uri_args.Source,
		Function: uri_args.Function,
		Extras:   uri_args.Extras,
	}

	return alt, nil
}

func IdFromPath(path string) (int64, error) {

	id, _, err := ParseURI(path)
	return id, err
}

func RepoFromPath(path string) (string, error) {

	abs_path, err := filepath.Abs(path)

	if err != nil {
		return "", err
	}

	wofid, err := IdFromPath(abs_path)

	if err != nil {
		return "", err
	}

	rel_path, err := Id2RelPath(wofid)

	if err != nil {
		return "", err
	}

	root_path := strings.Replace(abs_path, rel_path, "", 1)
	root_path = strings.TrimRight(root_path, "/")

	repo := ""

	for {

		base := filepath.Base(root_path)
		root_path = filepath.Dir(root_path)

		if strings.HasPrefix(base, "whosonfirst-data") {
			repo = base
			break
		}

		if root_path == "/" {
			break
		}

		if root_path == "" {
			break
		}
	}

	if repo == "" {
		return "", errors.New("Unable to determine repo from path")
	}

	return repo, nil
}
