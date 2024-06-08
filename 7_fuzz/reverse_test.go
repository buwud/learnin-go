package main

import (
	"testing"
	"unicode/utf8"
)

//UNIT TEST
// func TestReverse(t *testing.T) {
// 	testcases := []struct {
// 		in, want string
// 	}{
// 		{"Buwu", "uwuB"},
// 		{" ", " "},
// 		{"!1234", "4321!"},
// 	}
// 	for _, tc := range testcases {
// 		rev := Reverse(tc.in)
// 		if rev != tc.want {
// 			t.Errorf("Reverse : %q, want %q", rev, tc.want)
// 		}
// 	}
// }

// When fuzzing, you can’t predict the expected output, since you don’t have control over the inputs.
func FuzzReverse(f *testing.F) {
	testcases := []string{"Buwu", " ", "!1234"}
	for _, tc := range testcases {
		f.Add(tc)
	}
	f.Fuzz(func(t *testing.T, original string) {
		rev, err1 := Reverse(original)
		if err1 != nil {
			return
		}

		doubleRev, doubleRevErr := Reverse(rev)
		if doubleRevErr != nil {
			return
		}

		if original != doubleRev {
			t.Errorf("Before: %q, after: %q", original, doubleRev)
		}
		if utf8.ValidString(original) && !utf8.ValidString(rev) {
			t.Errorf("Reverse produced invalid UTF-8 string %q", rev)
		}
	})
}
