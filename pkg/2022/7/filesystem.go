package day7

import (
	"fmt"
	"strconv"
	"strings"
)

type fstype int

const (
	FSTypeFile fstype = iota
	FSTypeDirectory
)

type node struct {
	kind     fstype
	name     string
	size     int
	path     string
	parent   *node
	children map[string]*node
}

type Filesystem struct {
	root           *node
	currentDir     *node
	currentCommand command
}

type command string

const (
	CD command = "cd"
	LS command = "ls"
)

func NewFS() *Filesystem {
	fs := &Filesystem{
		root: &node{
			kind:     FSTypeDirectory,
			name:     "",
			path:     "/",
			parent:   nil,
			children: make(map[string]*node),
		},
	}

	fs.currentDir = fs.root

	return fs
}

func newNode(name, parentPath string, kind fstype, parent *node) *node {
	return &node{
		kind:     kind,
		name:     name,
		path:     normalize(fmt.Sprintf("%s/%s", parentPath, name)),
		parent:   parent,
		children: make(map[string]*node),
	}
}

func normalize(path string) string {
	return strings.ReplaceAll(path, "//", "/")
}

func (n *node) IsDir() bool {
	return n.kind == FSTypeDirectory
}

func (f *Filesystem) HandleOutput(s string) {
	if f.isCommand(s) {
		f.parseCommand(s)
		return
	}
	if f.currentCommand == LS {
		f.updateCurrentDir(s)
		return
	}
}

func (f *Filesystem) isCommand(s string) bool {
	return s[0] == '$'
}

func (f *Filesystem) parseCommand(s string) {
	ss := strings.Split(s, " ")[1:]

	switch command(ss[0]) {
	case CD:
		f.currentCommand = CD
		f.setCurrentDir(ss[1])
		return
	case LS:
		f.currentCommand = LS
		return
	}
}

func (f *Filesystem) CurrentDir() *node {
	return f.currentDir
}

func (f *Filesystem) List() []*node {
	values := make([]*node, 0)
	for _, v := range f.currentDir.children {
		values = append(values, v)
	}

	return values
}

func (f *Filesystem) setCurrentDir(s string) {
	if s == ".." {
		f.currentDir = f.currentDir.parent
		return
	}

	if _, ok := f.currentDir.children[s]; ok {
		f.currentDir = f.currentDir.children[s]
		return
	}

	newDir := newNode(s, f.currentDir.path, FSTypeDirectory, f.currentDir)
	f.currentDir.children[s] = newDir
	f.currentDir = newDir

	return
}

func (f *Filesystem) updateCurrentDir(s string) error {
	ss := strings.Split(s, " ")

	if ss[0] == "dir" {
		n := newNode(ss[1], f.currentDir.path, FSTypeDirectory, f.currentDir)
		f.currentDir.children[ss[1]] = n

		return nil
	} else {
		filesize, err := strconv.Atoi(ss[0])
		if err != nil {
			return err
		}
		n := newNode(ss[1], f.currentDir.path, FSTypeFile, f.currentDir)
		n.size = filesize
		f.updateSize(filesize)

		return nil
	}
}

func (f *Filesystem) updateSize(size int) {
	cd := f.currentDir

	for cd != nil {
		cd.size += size
		cd = cd.parent
	}
}
