package common

import "testing"

func TestCopyFiles(t *testing.T) {
	dst := "./"
	src := "/Users/peng/qyspace/dev/tmpweb"

	err := CopyFiles(dst, src)
	if err != nil {
		t.Error(err)
	}

}

func Test_copy(t *testing.T) {
	dst := "./"
	src := "../../plugins/plugins.go"

	err := copy(dst, src, true)
	if err != nil {
		t.Error(err)
	}
}
