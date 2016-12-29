package cmd

import (
	"reflect"
	"testing"
)

func TestListing(t *testing.T) {
	var directory = "../vendor/github.com/wtolson/go-taglib"
	meta := []*Audio{}
	fl := new(Audio)
	fl.Artist = "The Artist"
	fl.Album = "The Album"

	meta = append(meta, fl)
	//file, _ := tags.Read(directory)
	// if err != nil {
	// 	t.Fatalf("Read returned error: %s", err)
	// }

	//defer file.Close()

	res := listMetadata(directory)
	if reflect.DeepEqual(res, meta) == false {
		t.Errorf("Metadata is not equal between %v and %v", res[0], meta[0])
	}
}
