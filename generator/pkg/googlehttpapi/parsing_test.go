package googlehttpapi

import "testing"

func TestParsing(t *testing.T) {
	tts := []struct {
		name string
		in   string
		out  string
	}{{
		name: "no replacement",
		in:   "asdf/asdf/asdf",
		out:  "asdf/asdf/asdf",
	}, {
		name: "single replacement",
		in:   "asdf/{id}/asdf",
		out:  "\"asdf/\" + arg.id + \"asdf/\"",
	}}
	for _, tt := range tts {
		t.Run(tt.name, func(t *testing.T) {
			out, err := Parsing(tt.in)
			if err != nil {
				t.Fatal(err)
			}
			if out != tt.out {
				t.Errorf("want %s, got %s", tt.out, out)
			}
		})
	}
}
