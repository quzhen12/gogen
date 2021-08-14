package gen

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

type goGen interface {
	SetProjectTemplate(path string) error
	SetProjectName(projectName string) error
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
	err := g.G.SetProjectTemplate(path)
	if err != nil {
		return err
	}
	err = g.G.SetProjectName(projectName)
	if err != nil {
		return err
	}
	err = os.Mkdir("./"+projectName, 0777)
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
	f, err := os.Create(g.DstPath())
	_, err = buf.WriteTo(f)
	return err
}

func (g *gGen) createDir() error {
	_, err := os.Stat(g.DstPath())
	if os.IsNotExist(err) {
		return os.Mkdir(g.DstPath(), 0700)
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
	list := strings.Split(filePath, "/")
	if len(list) < 3 {
		return ""
	}
	return path.Join(g.G.ProjectName(), path.Join(list[2:]...))
}
