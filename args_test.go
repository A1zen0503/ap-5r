package main

import "testing"

func TestParseArgs(t *testing.T) {
	testCases := []struct {
		in  string
		out Args
	}{
		{in: "", out: Args{}},
		{in: "/mods tfp ",
			out: Args{Command: "/mods", Name: "tfp"}},
		{in: "/mods tie fighter pilot +small",
			out: Args{Command: "/mods", Name: "tie fighter pilot", Flags: []string{"+small"}}},
		{in: "/mods tie fighter pilot @ronoaldo",
			out: Args{Command: "/mods", Name: "tie fighter pilot", Profile: "ronoaldo"}},
	}

	for i := range testCases {
		tc := testCases[i]
		o := ParseArgs(tc.in)
		t.Logf("Test case #%d: `%s` ->\n%#v", i, tc.in, o)
		if o.Command != tc.out.Command {
			t.Errorf("Unexpected command: '%v', expected '%v'", o.Command, tc.out.Command)
		}
		if o.Name != tc.out.Name {
			t.Errorf("Unexpected name: '%v', expected '%v'", o.Name, tc.out.Name)
		}
		if o.Profile != tc.out.Profile {
			t.Errorf("Unexpected profile: '%v', expected '%v'", o.Profile, tc.out.Profile)
		}
		for _, f := range o.Flags {
			if !tc.out.ContainsFlag(f) {
				t.Errorf("> Unexpected flag: '%v'", f)
			}
		}
		for _, f := range tc.out.Flags {
			if !o.ContainsFlag(f) {
				t.Errorf("> Missing expected flag: '%v'", f)
			}
		}
	}
}