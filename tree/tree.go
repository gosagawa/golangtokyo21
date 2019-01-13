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

//Tree treeコマンド
type Tree struct {
	Deps int //取得する深さ
}

//NewTree treeコマンド作成
func NewTree(deps int) (*Tree, error) {

	tree := Tree{
		Deps: deps,
	}
	err := tree.validate()
	if err != nil {
		return nil, err
	}

	return &tree, nil
}

//OutputTree tree実行
func (t *Tree) OutputTree(dir string, out io.Writer) error {

	result := fmt.Sprintf("%v\n", dir)
	err := t.searchDir(dir, "", &result, 0)
	if err != nil {
		return err
	}

	_, err = out.Write([]byte(result))
	if err != nil {
		return err
	}

	return nil
}

func (t *Tree) searchDir(dir string, base string, result *string, deps int) error {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return fmt.Errorf("cannot read %v : %v", dir, err)
	}

	if t.isOverDeps(deps) {
		return nil
	}
	deps++

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
			err = t.searchDir(nextDir, nextBase, result, deps)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (t *Tree) validate() error {
	if t.Deps < 0 {
		return fmt.Errorf("deps must be over 0")
	}
	return nil
}

func (t *Tree) isOverDeps(deps int) bool {

	if t.Deps == 0 {
		return false
	}
	if t.Deps > deps {
		return false
	}
	return true
}
