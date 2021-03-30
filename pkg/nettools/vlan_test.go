package nettools_test

import (
	"testing"

	"github.com/automatico/nettools/pkg/nettools"
)

func TestExpandVLANs(t *testing.T) {
	have := "1-4"
	want := "1,2,3,4"
	got := nettools.ExpandVLANs(have)
	if want != got {
		t.Errorf("want: %s, got: %s", want, got)
	}

}
