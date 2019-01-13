package tree

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"testing"
)

type AssertFn func(name string, err error)

type unwritable struct{}

func (u *unwritable) Write(p []byte) (n int, err error) {
	return 0, fmt.Errorf("this io.Write is unwritable")
}

const testdataDirCase1 = "testdata/case1"
const testdataDirCase2 = "testdata/case2"

func TestIsValidInput(t *testing.T) {

	noError := func(name string, err error) {
		if err != nil {
			t.Errorf("%v: expected no error. error:%v", name, err)
		}
	}
	withError := func(name string, err error) {
		if err == nil {
			t.Errorf("%v: expected returning error", name)
		}
	}

	home, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	err = os.Chdir(home + "/../")
	if err != nil {
		panic(err)
	}
	fmt.Println(os.Getwd())

	output := testdataDirCase1
	output += `
├── dir1
│   ├── dir11
│   │   ├── file2
│   │   └── file3
│   └── dir12
│       ├── file4
│       └── file5
└── file1
`
	outputdeps1 := testdataDirCase1
	outputdeps1 += `
├── dir1
└── file1
`

	outputdeps2 := testdataDirCase1
	outputdeps2 += `
├── dir1
│   ├── dir11
│   └── dir12
└── file1
`

	if err := os.Chmod(testdataDirCase2, 0700); err != nil {
		fmt.Println(err)
	}

	cases := []struct {
		name      string
		dir       string
		output    string
		deps      int
		outStream interface{}
		assertFn  AssertFn
	}{
		{
			name:      "default",
			dir:       testdataDirCase1,
			output:    output,
			outStream: new(bytes.Buffer),
			assertFn:  noError,
		},
		{
			name:      "deps1",
			dir:       testdataDirCase1,
			output:    outputdeps1,
			deps:      1,
			outStream: new(bytes.Buffer),
			assertFn:  noError,
		},
		{
			name:      "deps2",
			dir:       testdataDirCase1,
			output:    outputdeps2,
			deps:      2,
			outStream: new(bytes.Buffer),
			assertFn:  noError,
		},
		{
			name:      "deps error",
			dir:       testdataDirCase1,
			deps:      -1,
			outStream: new(bytes.Buffer),
			assertFn:  withError,
		},
		{
			name:      "dir_not_exist",
			dir:       "dir_not_exist",
			outStream: new(bytes.Buffer),
			assertFn:  withError,
		},
		{
			name:      "outstream error",
			dir:       testdataDirCase1,
			outStream: new(unwritable),
			assertFn:  withError,
		},
		{
			name:      "directory_permission_error",
			dir:       testdataDirCase2,
			outStream: new(unwritable),
			assertFn:  withError,
		},
	}

	for _, v := range cases {
		t.Run(v.name, func(t *testing.T) {

			tree, err := NewTree(v.deps)
			if err != nil {
				v.assertFn(v.name, err)
				return
			}

			err = tree.OutputTree(v.dir, v.outStream.(io.Writer))
			if err != nil {
				v.assertFn(v.name, err)
				return
			}
			buffer := v.outStream.(*bytes.Buffer)
			if buffer.String() != v.output {
				t.Errorf("%v\nexpected \n%v", buffer.String(), v.output)
			}
			v.assertFn(v.name, err)
		})
	}

	if err := os.Chmod(testdataDirCase2, 0755); err != nil {
		fmt.Println(err)
	}
}
