package sources

import (
	"encoding/json"
	_ "errors"
	"github.com/whosonfirst/go-whosonfirst-sources/sources"
)

// please to be returning an actual thingy and not just an interface

func Spec() (interface{}, error) {

	var d interface{}
	err := json.Unmarshal([]byte(sources.Specification), &d)

	if err != nil {
		return nil, err
	}

	return d, nil
}
