package tree

import (
	"fmt"
	"io"
)

//OutputTree tree実行
func OutputTree(dir string, out io.Writer) error {

	if dir == "dir_not_exist" {
		return fmt.Errorf("dir not exist")
	}

	result := dir
	result += `
├── dir1
│   ├── dir11
│   │   └── file3
│   ├── dir12
│   │   └── file4
│   └── file2
└── file1
`

	_, err := out.Write([]byte(result))
	if err != nil {
		return err
	}

	return nil
}
