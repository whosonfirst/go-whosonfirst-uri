# go-whosonfirst-uri

Go package for working with URIs for Who's On First documents

## tl;dr

Too soon. This will supersede similar functions in `go-whosonfirst-utils`.

## Utilities

### wof-expand

```
./bin/wof-expand -h
Usage of ./bin/wof-expand:
  -alternate
    	Encode URI as an alternate geometry
  -extras string
    	A comma-separated list of extra information to include with an alternate geometry (optional)
  -function string
    	The function of the alternate geometry (optional)
  -prefix string
    	Prepend this prefix to all paths
  -root string
    	A root directory for absolute paths
  -source string
    	The source of the alternate geometry
  -strict
    	Ensure that the source for an alternate geometry is valid (see also: go-whosonfirst-sources)
```

## See also

* https://github.com/whosonfirst/whosonfirst-cookbook/blob/master/how_to/creating_alt_geometries.md
* https://github.com/whosonfirst/py-mapzen-whosonfirst-uri
