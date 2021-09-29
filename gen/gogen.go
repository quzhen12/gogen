package gen

import (
	"bufio"
	"fmt"
	"gogen/config"
	"io"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

type goGen interface {
	SetProjectTemplate(path string)
	SetProjectName(projectName string)
	SetOldProjectName(projectName string)
	Gen(filePath string) (io.Reader, error)
	ProjectName() string
	RenameDir() error
	ResetImport() error
	ResetGoMod() error
}

type gGen struct {
	G       goGen
	dstPath string
	srcPath string
}

func (g *gGen) SetDstPath(path string) {
	g.dstPath = path
}

func (g *gGen) DstPath() string {
	return g.dstPath
}

func (g *gGen) SetSrcPath(path string) {
	g.srcPath = path
}

func (g *gGen) SrcPath() string {
	return g.srcPath
}

func NewGen(g goGen) *gGen {
	return &gGen{G: g}
}

func (g *gGen) Gen(projectName string, path string) error {
	g.G.SetProjectTemplate(path)
	g.G.SetProjectName(projectName)
	err := os.Mkdir("./"+projectName, 0777)
	if err != nil {
		fmt.Println("mkdir err", err)
	}
	return g.Travels(path)
}

func (g *gGen) Travels(pathname string) error {
	rd, err := ioutil.ReadDir(pathname)
	if err != nil {
		return err
	}
	for _, fi := range rd {
		if IsIgnored(fi.Name()) {
			continue
		}
		childPath := path.Join(pathname, fi.Name())
		g.SetSrcPath(childPath)
		g.SetDstPath(g.path(childPath))
		err := g.copyFile(fi)
		if err != nil {
			return err
		}
		if fi.IsDir() {
			err := g.Travels(childPath)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func IsIgnored(name string) bool {
	ignoreMap := map[string]bool{
		".idea":      true,
		"vendor":     true,
		".gitignore": true,
	}
	return ignoreMap[name]
}

func (g *gGen) createFile() error {
	data, err := g.G.Gen(g.srcPath)
	if err != nil {
		return err
	}
	buf := bufio.NewReader(data)
	f, err := os.OpenFile(g.DstPath(), os.O_CREATE|os.O_RDWR, 0777)
	_, err = buf.WriteTo(f)
	return err
}

func (g *gGen) createDir() error {
	_, err := os.Stat(g.DstPath())
	if os.IsNotExist(err) {
		return os.Mkdir(g.DstPath(), 0777)
	}
	return nil
}

func (g *gGen) copyFile(f os.FileInfo) error {
	if f.IsDir() {
		return g.createDir()
	}
	return g.createFile()
}

func (g *gGen) path(filePath string) string {
	dir := strings.Split(filePath, config.PluginsDir2)
	if len(dir) < 2 {
		return ""
	}
	list := strings.Split(dir[1], "/")
	if len(list) < 2 {
		return ""
	}
	return path.Join(g.G.ProjectName(), path.Join(list[2:]...))
}
