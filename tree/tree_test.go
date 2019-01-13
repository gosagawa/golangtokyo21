package tree

import (
	"bytes"
	"testing"
)

type AssertFn func(name string, err error)

func TestIsValidInput(t *testing.T) {

	noError := func(name string, err error) {
		if err != nil {
			t.Errorf("%v: expected no error", name)
		}
	}
	withError := func(name string, err error) {
		if err == nil {
			t.Errorf("%v: expected returning error", name)
		}
	}

	dir := "testdata"
	output := dir
	output += `
├── dir1
│   ├── dir11
│   │   └── file3
│   ├── dir12
│   │   └── file4
│   └── file2
└── file1
`
	cases := []struct {
		name     string
		dir      string
		output   string
		assertFn AssertFn
	}{
		{name: "all correct", dir: dir, output: output, assertFn: noError},
		{name: "dir_not_exist", dir: "dir_not_exist", output: "", assertFn: withError},
	}

	for _, v := range cases {
		t.Run(v.name, func(t *testing.T) {
			outStream := new(bytes.Buffer)
			err := OutputTree(v.dir, outStream)
			v.assertFn(v.name, err)
			if err != nil {
				return
			}
			if outStream.String() != v.output {
				t.Errorf("%v: expected %v", outStream.String(), v.output)
			}
		})
	}
}
