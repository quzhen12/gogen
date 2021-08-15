package gen

import (
	"fmt"
	"os"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_travles(t *testing.T) {
	g := NewGentor()
	gg := NewGen(g)
	name := "qiaoyu"
	gg.G.SetProjectName(name)
	gg.G.SetOldProjectName("github.com/quzhen12/tmpweb")
	err := os.Mkdir("./"+name, 0777)
	if err != nil {
		fmt.Println("mkdir err", err)
	}
	err = gg.Travels("/usr/local/gogen/plugins/tmpweb")
	if err != nil {
		t.Log(err)
	}
}

func Test_path(t *testing.T) {
	g := NewGentor()
	gg := NewGen(g)
	gg.G.SetProjectName("toucheart")
	Convey("", t, func() {
		targetPath := "toucheart/api"
		srcPath := "temp/tmpweb/api"
		result := gg.path(srcPath)
		So(targetPath, ShouldEqual, result)
	})
}
