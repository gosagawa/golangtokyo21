package tree

import (
	"fmt"
	"io"
	"io/ioutil"
)

//RelationLine ツリー表示がな文字列
type RelationLine string

const (
	// RelationLineT T形
	RelationLineT RelationLine = "├── "

	// RelationLineI I形
	RelationLineI RelationLine = "│   "

	// RelationLineL L形
	RelationLineL RelationLine = "└── "

	// RelationLineEmpty ない場合
	RelationLineEmpty RelationLine = "    "
)

//OutputTree tree実行
func OutputTree(dir string, out io.Writer) error {

	result := fmt.Sprintf("%v\n", dir)
	err := searchDir(dir, "", &result)
	if err != nil {
		return err
	}

	_, err = out.Write([]byte(result))
	if err != nil {
		return err
	}

	return nil
}

func searchDir(dir string, base string, result *string) error {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return fmt.Errorf("cannnot read %v : %v", dir, err)
	}
	lastIndex := len(files) - 1
	for i, file := range files {

		var relationLineParent, relationLineChild RelationLine
		if i == lastIndex {
			relationLineParent = RelationLineL
			relationLineChild = RelationLineEmpty
		} else {
			relationLineParent = RelationLineT
			relationLineChild = RelationLineI
		}
		*result += fmt.Sprintf("%v%v%v\n", base, relationLineParent, file.Name())
		if file.IsDir() {
			nextDir := fmt.Sprintf("%v/%v", dir, file.Name())
			nextBase := fmt.Sprintf("%v%v", base, relationLineChild)
			err = searchDir(nextDir, nextBase, result)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
