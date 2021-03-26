package nettools_test

import (
	"testing"

	"github.com/automatico/nettools/pkg/nettools"
)

func TestMacStripper(t *testing.T) {
	type testCase struct {
		have string
		want string
	}
	testCases := []testCase{
		{have: "1111.2222.3333", want: "111122223333"},
		{have: "111122223333", want: "111122223333"},
		{have: "11-11-22-22-33-33", want: "111122223333"},
		{have: "11:11:22:22:33:33", want: "111122223333"},
	}
	for _, tc := range testCases {
		got := nettools.MacStripper(tc.have)
		if tc.want != got {
			t.Errorf("want: %s, got: %s", tc.want, got)
		}
	}
}

func TestMacFormatter(t *testing.T) {
	have := "11112222333"
	want := "1111.2222.3333"
	got, err := nettools.MacFormatter(have, "dot")
	if err != nil {
		t.Errorf("invalid mac address %s", have)
	}
	if want != got {
		t.Errorf("want: %s, got: %s", want, got)
	}
}
