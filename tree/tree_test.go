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
	cases := []struct {
		name      string
		dir       string
		output    string
		outStream interface{}
		assertFn  AssertFn
	}{
		{name: "all correct", dir: testdataDirCase1, output: output, outStream: new(bytes.Buffer), assertFn: noError},
		{name: "dir_not_exist", dir: "dir_not_exist", output: "", outStream: new(bytes.Buffer), assertFn: withError},
		{name: "output error", dir: testdataDirCase1, output: output, outStream: new(unwritable), assertFn: withError},
		{name: "directory_permission_error", dir: testdataDirCase2, output: output, outStream: new(unwritable), assertFn: withError},
	}

	for _, v := range cases {
		t.Run(v.name, func(t *testing.T) {
			err := OutputTree(v.dir, v.outStream.(io.Writer))
			v.assertFn(v.name, err)
			if err != nil {
				return
			}
			buffer := v.outStream.(*bytes.Buffer)
			if buffer.String() != v.output {
				t.Errorf("%v\nexpected \n%v", buffer.String(), v.output)
			}
		})
	}
}
