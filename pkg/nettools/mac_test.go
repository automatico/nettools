package nettools_test

import (
	"testing"

	"github.com/automatico/nettools/pkg/nettools"
)

func TestMACStripper(t *testing.T) {
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
		got := nettools.MACStripper(tc.have)
		if tc.want != got {
			t.Errorf("want: %s, got: %s", tc.want, got)
		}
	}
}

func TestMACFormatter(t *testing.T) {
	type testCase struct {
		have       string
		want       string
		format     string
		expectFail bool
	}
	testCases := []testCase{
		{have: "1111.2222.3333", want: "111122223333", format: "raw", expectFail: false},
		{have: "111122223333", want: "1111.2222.3333", format: "dot", expectFail: false},
		{have: "11-11-22-22-33-33", want: "11:11:22:22:33:33", format: "unix", expectFail: false},
		{have: "11:11:22:22:33:33", want: "11-11-22-22-33-33", format: "eui", expectFail: false},
		{have: "99999", want: "1111.2222.3333", format: "dot", expectFail: true},
		{have: "11:11:22:22:33:33", want: "1111.2222.3333", format: "UNKOWN-FORMAT", expectFail: true},
	}
	for _, tc := range testCases {
		got, err := nettools.MACFormatter(tc.have, tc.format)
		if err != nil && !tc.expectFail {
			t.Fatal(err)
		}
		if tc.want != got && !tc.expectFail {
			t.Errorf("want: %s, got: %s", tc.want, got)
		}
	}
}
