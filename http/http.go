// package http provides net/http handlers for Who's On First URI-related HTTP requests.
package http

import (
	"github.com/whosonfirst/go-whosonfirst-uri"
	_ "log"
	gohttp "net/http"
	"strconv"
)

// The name of the HTTP response header where the Who's On First relative path
// derived from a request URI will be stored.
const HEADER_RELPATH string = "X-WhosOnFirst-Rel-Path"

// Parse a request URI (or ?id query string) in to a valid Who's On First ID relative
// path and assign the value to the 'X-WhosOnFirst-Rel-Path' response header. If 'next'
// is not nil then delegate to that handler, otherwise print the relative path to the
// response handler.
func ParseURIHandler(next gohttp.Handler) gohttp.HandlerFunc {

	fn := func(rsp gohttp.ResponseWriter, req *gohttp.Request) {

		path := req.URL.Path

		wofid, uri_args, err := uri.ParseURI(path)

		if err != nil || wofid == -1 {

			q := req.URL.Query()
			str_id := q.Get("id")

			if str_id == "" {
				gohttp.Error(rsp, err.Error(), gohttp.StatusNotFound)
				return
			}

			id, err := strconv.ParseInt(str_id, 10, 64)

			if err != nil {
				gohttp.Error(rsp, err.Error(), gohttp.StatusBadRequest)
				return
			}

			wofid = id

			uri_args = &uri.URIArgs{
				IsAlternate: false,
			}
		}

		rel_path, err := uri.Id2RelPath(wofid, uri_args)

		if err != nil {
			gohttp.Error(rsp, err.Error(), gohttp.StatusInternalServerError)
			return
		}

		rsp.Header().Set(HEADER_RELPATH, rel_path)

		if next != nil {
			next.ServeHTTP(rsp, req)
			return
		}

		rsp.Header().Set("Content-Type", "text/plain")
		rsp.Write([]byte(rel_path))
		return
	}

	return gohttp.HandlerFunc(fn)
}
