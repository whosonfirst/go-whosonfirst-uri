package uri

import (
	"testing"
)

func TestParseURI(t *testing.T) {

	possible := []string{
		"152/791/195/9/1527911959.geojson",
		"/1527911959/",
		"1527911959.geojson",
	}

	expected_id := int64(1527911959)

	for _, p := range possible {

		id, uri_args, err := ParseURI(p)

		if err != nil {
			t.Fatal(err)
		}

		if id != expected_id {
			t.Fatalf("Invalid ID for URI '%s'. Expected '%d' but got '%d'.", p, expected_id, id)
		}

		if uri_args.IsAlternate {
			t.Fatal("Invalid URI. Reported as alternate but is not.")
		}
	}

}

func TestParseURIAlt(t *testing.T) {

	possible := []string{
		"152/791/195/9/1527911959-alt-swim-path.geojson",
		"/1527911959-alt-swim-path/",
		"1527911959-alt-swim-path",		
		"1527911959-alt-swim-path.geojson",
	}

	expected_id := int64(1527911959)
	expected_source := "swim"
	expected_function := "path"

	for _, p := range possible {

		id, uri_args, err := ParseURI(p)

		if err != nil {
			t.Fatal(err)
		}

		if id != expected_id {
			t.Fatalf("Invalid ID for URI '%s'. Expected '%d' but got '%d'.", p, expected_id, id)
		}

		if !uri_args.IsAlternate {
			t.Fatal("Invalid URI. Does not report as alternate but should.")
		}

		alt_geom := uri_args.AltGeom

		if alt_geom.Source != expected_source {
			t.Fatalf("Invalid alternate source for URI '%s'. Expected '%s' but got '%s'.", p, expected_source, alt_geom.Source)
		}

		if alt_geom.Function != expected_function {
			t.Fatalf("Invalid alternate function for URI '%s'. Expected '%s' but got '%s'.", p, expected_function, alt_geom.Function)
		}

	}

}

func TestIsWOFFile(t *testing.T) {

	is_wof := []string{
		"152/791/195/9/1527911959.geojson",
		"1527911959.geojson",
		"152/791/195/9/1527911959-alt-swim-path.geojson",
		"1527911959-alt-swim-path.geojson",
	}

	not_wof := []string{
		"1527911959.json",
		"1527911959.txt",
		"152/791/195/9/1527911959.md",
	}

	for _, uri := range is_wof {

		ok, err := IsWOFFile(uri)

		if err != nil {
			t.Fatal(err)
		}

		if !ok {
			t.Fatalf("'%s' reported as not a WOF file, but should be.", uri)
		}
	}

	for _, uri := range not_wof {

		ok, err := IsWOFFile(uri)

		if err != nil {
			t.Fatal(err)
		}

		if ok {
			t.Fatalf("'%s' reported as a WOF file, but should not be.", uri)
		}
	}

}

func TestIsAltFile(t *testing.T) {

	is_alt := []string{
		"152/791/195/9/1527911959-alt-swim-path.geojson",
		"1527911959-alt-swim-path.geojson",
	}

	not_alt := []string{
		"1527911959.json",
		"1527911959.txt",
		"152/791/195/9/1527911959.md",
		"152/791/195/9/1527911959.geojson",
		"1527911959.geojson",
	}

	for _, uri := range is_alt {

		ok, err := IsAltFile(uri)

		if err != nil {
			t.Fatal(err)
		}

		if !ok {
			t.Fatalf("'%s' reported as not an alt file, but should be.", uri)
		}
	}

	for _, uri := range not_alt {

		ok, err := IsAltFile(uri)

		if err != nil {
			t.Fatal(err)
		}

		if ok {
			t.Fatalf("'%s' reported as an alt file, but should not be.", uri)
		}
	}

}

func TestAltGeomFromPath(t *testing.T) {

	paths := []string{
		"152/791/195/9/1527911959-alt-swim-path.geojson",
	}

	expected_source := "swim"
	expected_function := "path"

	for _, p := range paths {

		alt_geom, err := AltGeomFromPath(p)

		if err != nil {
			t.Fatal(err)
		}

		if alt_geom.Source != expected_source {
			t.Fatalf("Invalid alternate source for URI '%s'. Expected '%s' but got '%s'.", p, expected_source, alt_geom.Source)
		}

		if alt_geom.Function != expected_function {
			t.Fatalf("Invalid alternate function for URI '%s'. Expected '%s' but got '%s'.", p, expected_function, alt_geom.Function)
		}

	}
}

func TestIdFromPath(t *testing.T) {

	possible := []string{
		"152/791/195/9/1527911959.geojson",
		"/1527911959/",
		"1527911959.geojson",
		"152/791/195/9/1527911959-alt-swim-path.geojson",
		"/1527911959-alt-swim-path/",
		"1527911959-alt-swim-path.geojson",
	}

	expected_id := int64(1527911959)

	for _, p := range possible {

		id, err := IdFromPath(p)

		if err != nil {
			t.Fatal(err)
		}

		if id != expected_id {
			t.Fatalf("Invalid ID for URI '%s'. Expected '%d' but got '%d'.", p, expected_id, id)
		}
	}

}
