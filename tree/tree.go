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
	Deps       int //取得する深さ
	DirAmount  int //ディレクトリ数
	FileAmount int //ファイル数
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
	result += fmt.Sprintf("\n%v\n", t.getFileCountString())

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
			t.DirAmount++

			nextDir := fmt.Sprintf("%v/%v", dir, file.Name())
			nextBase := fmt.Sprintf("%v%v", base, relationLineChild)

			err = t.searchDir(nextDir, nextBase, result, deps)
			if err != nil {
				return err
			}
		} else {
			t.FileAmount++
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

func (t *Tree) getFileCountString() string {

	dirUnit := "directory"
	if t.DirAmount > 1 {
		dirUnit = "directories"
	}
	fileUnit := "file"
	if t.FileAmount > 1 {
		fileUnit = "files"
	}
	return fmt.Sprintf("%v %v, %v %v", t.DirAmount, dirUnit, t.FileAmount, fileUnit)
}
