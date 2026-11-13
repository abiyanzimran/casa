package api_test

import (
	"testing"

	"github.com/abiyanzimran/casa/server/internal/api"
)

func TestParseVersion(t *testing.T) {
	valid := map[string]api.Version{
		"0.0.0":      {},
		"1.9.0":      {Major: 1, Minor: 9},
		"1.10.0":     {Major: 1, Minor: 10},
		"999999.0.1": {Major: 999999, Patch: 1},
	}
	for s, want := range valid {
		got, err := api.ParseVersion(s)
		if err != nil || got != want {
			t.Errorf("ParseVersion(%q) = %v, %v; want %v, nil", s, got, err, want)
		}
	}
	for _, s := range []string{"", "1.2", "1.2.3.4", "1.2.3-beta", "v1.2.3", "01.2.3", "1.2.x", " 1.2.3", "1000000.0.0"} {
		if got, err := api.ParseVersion(s); err == nil {
			t.Errorf("ParseVersion(%q) = %v, want an error", s, got)
		}
	}
}

func TestVersionLessComparesNumbers(t *testing.T) {
	cases := []struct {
		v, w api.Version
		want bool
	}{
		{api.Version{Major: 1, Minor: 9}, api.Version{Major: 1, Minor: 10}, true},
		{api.Version{Major: 1, Minor: 10}, api.Version{Major: 1, Minor: 9}, false},
		{api.Version{Major: 2}, api.Version{Major: 2}, false},
		{api.Version{Major: 1, Minor: 9, Patch: 9}, api.Version{Major: 2}, true},
		{api.Version{Major: 1, Minor: 9, Patch: 1}, api.Version{Major: 1, Minor: 9, Patch: 2}, true},
	}
	for _, c := range cases {
		if got := c.v.Less(c.w); got != c.want {
			t.Errorf("%v.Less(%v) = %v, want %v", c.v, c.w, got, c.want)
		}
	}
}
