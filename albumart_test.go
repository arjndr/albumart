package albumart

import (
	"testing"
	"strings"
	"github.com/arjndr/albumart"
)

func TestNotFound(t *testing.T) {
	artwork := albumart.GetArtwork("func", "1337")
	if artwork != "Not Found" {
		t.Fail()
	}
}

func TestNormal(t *testing.T) {
	artwork := albumart.GetArtwork("MK", "17")
	if !strings.Contains(artwork, "https://") {
		t.Fail()
	}
}
