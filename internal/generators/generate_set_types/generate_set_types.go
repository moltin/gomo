package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type CoreType struct {
	TypeName string
	APIName  string
	Variable string
}

type CoreTypes []CoreType

func getTypes() (types CoreTypes) {
	f, err := os.Open("types.go")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	reader := bufio.NewReader(f)
	for {
		var c CoreType
		line, _, err := reader.ReadLine()
		if err != nil {
			break
		}
		bits := strings.SplitN(string(line), "=", 2)
		if len(bits) != 2 {
			continue
		}
		v := strings.TrimSuffix(strings.TrimSpace(bits[0]), "Type")
		c.Variable = v[0:1]
		c.TypeName = strings.Title(v)
		bits = strings.SplitN(bits[1], `"`, 3)
		if len(bits) != 3 {
			continue
		}
		c.APIName = bits[1]
		types = append(types, c)
	}
	return
}

func writeFile(types CoreTypes) {
	f, err := os.OpenFile(
		"set_types.go",
		os.O_WRONLY|os.O_CREATE|os.O_TRUNC,
		0644,
	)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	fmt.Fprintln(f, `package core`)
	fmt.Fprintln(f, ``)
	fmt.Fprintln(f, `//go:generate go run ../internal/generators/generate_set_types/`)
	fmt.Fprintln(f, ``)

	for _, c := range types {
		fmt.Fprintf(f, "func (%s *%s) SetType() {\n", c.Variable, c.TypeName)
		fmt.Fprintf(f, "	%s.Type = `%s`\n", c.Variable, c.APIName)
		fmt.Fprintln(f, `}`)
		fmt.Fprintln(f, ``)
	}
}

func writeTestFile(types CoreTypes) {
	f, err := os.OpenFile(
		"set_types_test.go",
		os.O_WRONLY|os.O_CREATE|os.O_TRUNC,
		0644,
	)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	fmt.Fprintln(f, `package core_test`)
	fmt.Fprintln(f, ``)
	fmt.Fprintln(f, `import (`)
	fmt.Fprintln(f, `	"testing"`)
	fmt.Fprintln(f, ``)
	fmt.Fprintln(f, `	"github.com/moltin/gomo/core"`)
	fmt.Fprintln(f, `)`)
	fmt.Fprintln(f, ``)
	for _, c := range types {
		fmt.Fprintf(f, "func Test%sType(t *testing.T) {\n", c.TypeName)
		fmt.Fprintf(f, "	var c%s core.%s\n", c.Variable, c.TypeName)
		fmt.Fprintf(f, "	c%s.SetType()\n", c.Variable)
		fmt.Fprintln(f, ``)
		fmt.Fprintf(f, "	if c%s.Type != `%s` {\n", c.Variable, c.APIName)
		fmt.Fprintln(f, `		t.Error("failed to set type")`)
		fmt.Fprintln(f, `	}`)
		fmt.Fprintln(f, `}`)
		fmt.Fprintln(f, ``)
	}
}

func main() {
	types := getTypes()
	writeFile(types)
	writeTestFile(types)
}
