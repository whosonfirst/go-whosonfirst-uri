# go-whosonfirst-uri

Go package for working with URIs for Who's On First documents

## Example

### Simple

```
import (
	"github.com/whosonfirst/go-whosonfirst-uri"
)

fname, _ := uri.Id2Fname(101736545)
rel_path, _ := uri.Id2RelPath(101736545)
abs_path, _ := uri.Id2AbsPath("/usr/local/data", 101736545)
```

Produces:

```
101736545.geojson
101/736/545/101736545.geojson
/usr/local/data/101/736/545/101736545.geojson
```

### Fancy

```
import (
	"github.com/whosonfirst/go-whosonfirst-uri"
)

source := "mapzen"
function := "display"
extras := []string{ "1024" }

args := uri.NewAlternateURIArgs(source, function, extras...)

fname, _ := uri.Id2Fname(101736545, args)
rel_path, _ := uri.Id2RelPath(101736545, args)
abs_path, _ := uri.Id2AbsPath("/usr/local/data", 101736545, args)
```

Produces:

```
101736545-alt-mapzen-display-1024.geojson
101/736/545/101736545-alt-mapzen-display-1024.geojson
/usr/local/data/101/736/545/101736545-alt-mapzen-display-1024.geojson
```

## The Long Version

Please read this: https://github.com/whosonfirst/whosonfirst-cookbook/blob/master/how_to/creating_alt_geometries.md

## Utilities

### wof-cat

Expand and concatenate Who's On First IDs and print them to `STDOUT`.

```
./bin/wof-cat -h
Usage of ./bin/wof-cat:
  -alternate
    	Encode URI as an alternate geometry
  -extras string
    	A comma-separated list of extra information to include with an alternate geometry (optional)
  -function string
    	The function of the alternate geometry (optional)
  -root string
    	If empty defaults to the current working directory + "/data".
  -source string
    	The source of the alternate geometry
  -strict
    	Ensure that the source for an alternate geometry is valid (see also: go-whosonfirst-sources)
```

For example, assuming you are in the `whosonfirst-data` repo:

```
$> wof-cat 0 | less

{
  "id": 0,
  "type": "Feature",
  "properties": {
    "edtf:cessation":"uuuu",
    "edtf:inception":"uuuu",
    "geom:area":64800.0,
    "geom:bbox":"-180.0,-90.0,180.0,90.0",
    "geom:latitude":-0.0,
    "geom:longitude":-0.0,
    "iso:country":"",
    "mz:hierarchy_label":1,
    "name:chi_x_preferred":[
        "\u5730\u7403"
    ],
    "name:chi_x_variant":[
        "\u4e16\u754c"
    ],
    "name:dut_x_preferred":[
        "Aarde"
    ],
    "name:dut_x_variant":[
        "Wereld"
    ],
    "name:eng_x_preferred":[
        "Earth"
    ],
    ...
  },
  "geometry": {"coordinates":[[[-180.0,-90.0],[-180.0,90.0],[180.0,90.0],[180.0,-90.0],[-180.0,-90.0]]],"type":"Polygon"}
}
```

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

* * https://github.com/whosonfirst/whosonfirst-cookbook/blob/master/how_to/creating_alt_geometries.md
* https://github.com/whosonfirst/py-mapzen-whosonfirst-uri
