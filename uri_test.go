package uri

import (
	"encoding/json"
	_ "fmt"
	"testing"
)

func TestId2Fname(t *testing.T) {

	id := int64(1527911959)
	expected := "1527911959.geojson"

	fname, err := Id2Fname(id)

	if err != nil {
		t.Fatal(err)
	}

	if fname != expected {
		t.Fatalf("Invalid filename for '%d'. Expected '%s' but got '%s'", id, expected, fname)
	}
}

func TestId2FnameAlt(t *testing.T) {

	id := int64(1527911959)

	alt_geom := &AltGeom{
		Source:   "swim",
		Function: "path",
		Extras:   []string{},
	}

	uri_args := &URIArgs{
		IsAlternate: true,
		AltGeom:     alt_geom,
	}

	expected := "1527911959-alt-swim-path.geojson"

	fname, err := Id2Fname(id, uri_args)

	if err != nil {
		t.Fatal(err)
	}

	if fname != expected {
		t.Fatalf("Invalid alt filename for '%d'. Expected '%s' but got '%s'", id, expected, fname)
	}
}

func TestId2Path(t *testing.T) {

	id := int64(1527911959)
	expected := "152/791/195/9"

	path, err := Id2Path(id)

	if err != nil {
		t.Fatal(err)
	}

	if path != expected {
		t.Fatalf("Invalid path for '%d'. Expected '%s' but got '%s'", id, expected, path)
	}
}

func TestId2RelPath(t *testing.T) {

	id := int64(1527911959)
	expected := "152/791/195/9/1527911959.geojson"

	rel_path, err := Id2RelPath(id)

	if err != nil {
		t.Fatal(err)
	}

	if rel_path != expected {
		t.Fatalf("Invalid relative path for '%d'. Expected '%s' but got '%s'", id, expected, rel_path)
	}
}

func TestURIArgsJSON(t *testing.T) {

	enc_args := `{"is_alternate":true,"alternate_geometry":{"source":"mapzen","function":"display","extras":["1024"],"strict":true}}`

	var uri_args *URIArgs

	err := json.Unmarshal([]byte(enc_args), &uri_args)

	if err != nil {
		t.Fatalf("Failed to unmarshal args, %v", err)
	}

	if !uri_args.IsAlternate {
		t.Fatalf("Expected alternate geometry")
	}

	alt_geom := uri_args.AltGeom

	if alt_geom.Source != "mapzen" {
		t.Fatalf("Unexpected alternate geometry source")
	}

	if alt_geom.Function != "display" {
		t.Fatalf("Unexpected alternate geometry function")
	}

	if len(alt_geom.Extras) != 1 {
		t.Fatalf("Unexpected alternate geometry extras")
	}

	enc_args2, err := json.Marshal(uri_args)

	if string(enc_args2) != enc_args {
		t.Fatalf("Unexpected alternate geometry marshaling")
	}

}
