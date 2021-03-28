package nettools_test

import (
	"testing"

	"github.com/automatico/nettools/pkg/nettools"
)

func TestASNPlainToDot(t *testing.T) {
	type testCase struct {
		have       string
		want       string
		expectFail bool
	}
	testCases := []testCase{
		{have: "1", want: "0.1", expectFail: false},
		{have: "65535", want: "0.65535", expectFail: false},
		{have: "65536", want: "1.0", expectFail: false},
		{have: "65432", want: "0.65432", expectFail: false},
		{have: "65546", want: "1.10", expectFail: false},
		{have: "131073", want: "2.1", expectFail: false},
		{have: "769672", want: "11.48776", expectFail: false},
		{have: "4294967295", want: "65535.65535", expectFail: false},
		{have: "4294967296", want: "", expectFail: true},
		{have: "-1", want: "", expectFail: true},
	}
	for _, tc := range testCases {
		got, err := nettools.ASNPlainToDot(tc.have)
		if tc.want != got {
			t.Errorf("want: %s, got: %s", tc.want, got)
		}
		if err != nil && !tc.expectFail {
			t.Fatal(err)
		}
	}
}

func TestASNDotToPlain(t *testing.T) {
	type testCase struct {
		have       string
		want       string
		expectFail bool
	}
	testCases := []testCase{
		{have: "1.0", want: "65536", expectFail: false},
		{have: "1.10", want: "65546", expectFail: false},
		{have: "2.10", want: "131082", expectFail: false},
	}
	for _, tc := range testCases {
		got := nettools.ASNDotToPlain(tc.have)
		if tc.want != got {
			t.Errorf("want: %s, got: %s", tc.want, got)
		}
	}
}
