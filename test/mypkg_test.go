package test

import "testing"

// START OMIT
func TestStringLen(t *testing.T) {
	testCases := map[string]struct {
		s    string
		want int
	}{
		"empty": {
			s:    "",
			want: 0,
		},
		"simple": {
			s:    "Avocode and Gophers",
			want: 19,
		},
	}

	for desc, tc := range testCases {
		t.Log(desc)
		got := StringLen(tc.s)
		if got != tc.want {
			t.Errorf("got %d, want %d", got, tc.want)
		}
	}
}

// END OMIT
