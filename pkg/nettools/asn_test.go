package nettools_test

import (
	"testing"

	"github.com/automatico/nettools/pkg/nettools"
)

func TestASNPlainToDot(t *testing.T) {
	type testCase struct {
		have string
		want string
	}
	testCases := []testCase{
		{have: "65536", want: "1.0"},
		{have: "65432", want: "0.65432"},
		{have: "65546", want: "1.10"},
		{have: "131073", want: "2.1"},
		{have: "769672", want: "11.48776"},
	}
	for _, tc := range testCases {
		got := nettools.ASNPlainToDot(tc.have)
		if tc.want != got {
			t.Errorf("want: %s, got: %s", tc.want, got)
		}
	}
}

func TestASNDotToPlain(t *testing.T) {
	type testCase struct {
		have string
		want string
	}
	testCases := []testCase{
		{have: "1.0", want: "65536"},
		{have: "1.10", want: "65546"},
		{have: "2.10", want: "131082"},
	}
	for _, tc := range testCases {
		got := nettools.ASNDotToPlain(tc.have)
		if tc.want != got {
			t.Errorf("want: %s, got: %s", tc.want, got)
		}
	}
}
