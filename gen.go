package main

import (
	"bufio"
	"bytes"
	"io"
	"os"
	"strings"
)

type gen struct {
	projectName    string
	oldProjectName string
	path           string
	data           []byte
}

func (g *gen) RenameDir() error {
	panic("implement me")
}

func (g *gen) ResetImport() error {
	panic("implement me")
}

func (g *gen) ResetGoMod() error {
	panic("implement me")
}

func NewGentor() *gen {
	return &gen{}
}

func (g *gen) SetProjectTemplate(path string) error {
	g.path = path
	return nil
}

func (g *gen) SetProjectName(projectName string) error {
	g.projectName = projectName
	return nil
}

func (g *gen) SetOldProjectName(projectName string) {
	g.oldProjectName = projectName
}

func (g *gen) ProjectName() string {
	return g.projectName
}

func (g *gen) Gen(filePath string) (io.Reader, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	result := bytes.NewBuffer(nil)
	buf := bufio.NewReader(f)
	for {
		str, err := buf.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
		if strings.Contains(str, g.oldProjectName) {
			str = strings.Replace(str, g.oldProjectName, g.projectName, 1)
		}
		result.WriteString(str)
	}
	return result, err
}
