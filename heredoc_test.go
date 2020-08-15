package here_test

import (
	"testing"

	"github.com/h2so5/here"
)

type testCase struct {
	raw, expect string
}

var tests = []testCase{
	{"", ""},
	{`
		Foo
		Bar
		`,
		"Foo\nBar\n"},
	{`Foo
		Bar`,
		"Foo\nBar"},
	{`Foo
			
		Bar
		`,
		"Foo\n\t\nBar\n"}, // Second line contains two tabs.
	{`
		Foo
			Bar
				Hoge
					`,
		"Foo\n\tBar\n\t\tHoge\n\t\t\t"},
	{`Foo Bar`, "Foo Bar"},
	{
		`
		Foo
		Bar
	`, "Foo\nBar\n"},
	{"\n\u3000zenkaku space", "\u3000zenkaku space"},
}

func TestDoc(t *testing.T) {
	for i, test := range tests {
		result := here.Doc(test.raw)
		if result != test.expect {
			t.Errorf("tests[%d] failed: expected=> %#v, result=> %#v", i, test.expect, result)
		}
	}
}

func TestDocf(t *testing.T) {
	tc := `
		int: %3d
		string: %s
	`
	i := 42
	s := "Hello"
	expect := "int:  42\nstring: Hello\n"

	result := here.Docf(tc, i, s)
	if result != expect {
		t.Errorf("test failed: expected=> %#v, result=> %#v", expect, result)
	}
}
