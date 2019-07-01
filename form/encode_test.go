package form_test

import (
	"strings"
	"testing"

	"github.com/moltin/gomo/form"
)

func boundary(s string) (b string) {
	bits := strings.SplitN(s, "=", 2)
	if len(bits) == 2 {
		b = "--" + bits[1]
	}
	return
}

func TestEncode(t *testing.T) {
	for _, test := range []struct {
		name     string
		object   interface{}
		expected string
	}{
		{
			"simple",
			struct {
				Name  string `form:"name"`
				Skip  string `form:"-"`
				Empty string `form:"empty,omitempty"`
			}{
				Name: "test",
				Skip: "foo",
			},
			`BOUNDARY
Content-Disposition: form-data; name="name"

test
BOUNDARY--
`,
		},
		{
			"reader",
			struct {
				Name *form.File `form:"name"`
			}{
				Name: &form.File{
					Name:    "test.txt",
					Content: strings.NewReader("test"),
				},
			},
			`BOUNDARY
Content-Disposition: form-data; name="name"; filename="test.txt"
Content-Type: application/octet-stream

test
BOUNDARY--
`,
		},
		{
			"bad type",
			true,
			`
BOUNDARY--
`,
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			body, contentType, err := form.Encode(test.object)
			if err != nil {
				t.Fatal(err)
			}
			expected := strings.ReplaceAll(
				test.expected,
				"BOUNDARY",
				boundary(contentType),
			)
			bodyS := strings.ReplaceAll(string(body), "\r", "")
			if string(bodyS) != expected {
				t.Errorf(
					"\nexpected:\n>%s<\ngot:\n>%s<",
					expected,
					bodyS,
				)
			}
		})
	}
}
