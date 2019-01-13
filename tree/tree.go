package tree

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
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

//GopherString Gopherの顔文字
const GopherString = "ʕ◔ϖ◔ʔ"

//Tree treeコマンド
type Tree struct {
	Option     Option //オプション
	DirAmount  int    //ディレクトリ数
	FileAmount int    //ファイル数
	Result     string //実行結果文字列
}

//Option treeコマンドオプション
type Option struct {
	Deps       int  //取得する深さ
	DirOnly    bool //ディレクトリのみ表示
	MaskGopher bool //Gopherでマスクするか？
}

//NewTree treeコマンド作成
func NewTree(option Option) (*Tree, error) {

	tree := Tree{
		Option: option,
	}
	err := tree.validate()
	if err != nil {
		return nil, err
	}

	return &tree, nil
}

//OutputTree tree実行
func (t *Tree) OutputTree(dir string, out io.Writer) error {

	t.Result += fmt.Sprintf("%v\n", dir)
	err := t.searchDir(dir, "", 0)
	if err != nil {
		return err
	}
	t.Result += fmt.Sprintf("\n%v\n", t.getFileCountString())

	_, err = out.Write([]byte(t.Result))
	if err != nil {
		return err
	}

	return nil
}

func (t *Tree) searchDir(dir string, base string, deps int) error {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return fmt.Errorf("cannot read %v : %v", dir, err)
	}

	if t.isOverDeps(deps) {
		return nil
	}
	deps++

	for i, file := range files {

		lineParent, lineChild := t.getRelationLine(i, files)
		if file.IsDir() {
			t.Result += fmt.Sprintf("%v%v%v\n", base, lineParent, t.getFileName(file.Name()))
			t.DirAmount++

			nextDir := fmt.Sprintf("%v/%v", dir, file.Name())
			nextBase := fmt.Sprintf("%v%v", base, lineChild)

			err = t.searchDir(nextDir, nextBase, deps)
			if err != nil {
				return err
			}
		} else if t.Option.DirOnly == false {

			t.Result += fmt.Sprintf("%v%v%v\n", base, lineParent, t.getFileName(file.Name()))
			t.FileAmount++
		}
	}
	return nil
}

func (t *Tree) validate() error {
	if t.Option.Deps < 0 {
		return fmt.Errorf("deps must be over 0")
	}
	return nil
}

func (t *Tree) isOverDeps(deps int) bool {

	if t.Option.Deps == 0 {
		return false
	}
	if t.Option.Deps > deps {
		return false
	}
	return true
}

func (t *Tree) getRelationLine(i int, files []os.FileInfo) (RelationLine, RelationLine) {
	var relationLineParent, relationLineChild RelationLine
	lastIndex := len(files) - 1

	//option.DirOnlyの時は、次がファイルだったらその階層の最終項目扱いにする
	if t.Option.DirOnly && i != lastIndex && !files[i+1].IsDir() {
		lastIndex = i
	}

	if i == lastIndex {
		relationLineParent = RelationLineL
		relationLineChild = RelationLineEmpty
	} else {
		relationLineParent = RelationLineT
		relationLineChild = RelationLineI
	}
	return relationLineParent, relationLineChild
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
	result := ""
	if t.Option.DirOnly {
		result = fmt.Sprintf("%v %v", t.DirAmount, dirUnit)
	} else {
		result = fmt.Sprintf("%v %v, %v %v", t.DirAmount, dirUnit, t.FileAmount, fileUnit)
	}
	return result
}
func (t *Tree) getFileName(name string) string {

	if t.Option.MaskGopher {
		name = GopherString
	}
	return name
}
