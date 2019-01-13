package tree

import "io"

//OutputTree tree実行
func OutputTree(dir string, out io.Writer) error {
	out.Write([]byte(dir))

	return nil
}
