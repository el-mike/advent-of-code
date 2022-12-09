package day7

import (
	"el-mike/advent-of-code/common/ds"
	"strconv"
	"strings"
)

const (
	CommandIndicator = "$ "
	DirIndicator     = "dir"
	LsCommand        = "ls"
	CdCommand        = "cd"
	RootToken        = "/"
	GoBackToken      = ".."
)

const (
	DirElementType uint8 = iota
	FileElementType
)

type FileSystemElement struct {
	Type uint8
	Name string
	Path string
	Size int
}

func NewFileSystemElement(elementType uint8, name, path string, size int) *FileSystemElement {
	return &FileSystemElement{
		Type: elementType,
		Name: name,
		Path: path,
		Size: size,
	}
}

type FileSystemModel struct {
	Tree *ds.Tree[*FileSystemElement]

	CurrentDir *ds.TreeNode[*FileSystemElement]
}

func NewFileSystemModel() *FileSystemModel {
	return &FileSystemModel{
		Tree: ds.NewTree[*FileSystemElement](nil),
	}
}

func (fs *FileSystemModel) parseLine(line string) {
	if fs.isCommand(line) {
		command := strings.ReplaceAll(line, CommandIndicator, "")

		fs.parseCommand(command)
	} else {
		fs.parseFileSystemElement(line)
	}
}

func (fs *FileSystemModel) parseCommand(command string) {
	if strings.Contains(command, CdCommand) {
		dirName := strings.ReplaceAll(command, CdCommand+" ", "")

		// If CurrentDir is null, we don't have the root element yet, therefore we
		// parse the initial command and build root element from that.
		if fs.Tree.Root == nil {
			element := NewFileSystemElement(DirElementType, dirName, dirName, 0)

			fs.Tree.Root = ds.NewTreeNode(element, nil)
			fs.CurrentDir = fs.Tree.Root

			return
		}

		// In case of going back, we just reassign CurrentDir to its parent.
		if dirName == GoBackToken {
			fs.CurrentDir = fs.CurrentDir.Parent
			return
		}

		// "$ cd /" should redirect terminal to the root directory.
		if dirName == RootToken {
			fs.CurrentDir = fs.Tree.Root
			return
		}

		var nextDir *ds.TreeNode[*FileSystemElement]

		if !fs.exists(fs.CurrentDir, dirName) {
			nextDir = fs.addElement(fs.CurrentDir, DirElementType, dirName, 0)
		} else {
			nextDir = fs.findByName(fs.CurrentDir, dirName)
		}

		fs.CurrentDir = nextDir
	}
}

func (fs *FileSystemModel) parseFileSystemElement(elementStr string) {
	if fs.isDir(elementStr) {
		dirName := strings.Split(elementStr, " ")[1]

		if !fs.exists(fs.CurrentDir, dirName) {
			fs.addElement(fs.CurrentDir, DirElementType, dirName, 0)
		}
	} else {
		parts := strings.Split(elementStr, " ")
		sizeStr, fileName := parts[0], parts[1]

		size, err := strconv.Atoi(sizeStr)
		if err != nil {
			panic(err)
		}

		fs.addElement(fs.CurrentDir, FileElementType, fileName, size)
	}
}

func (fs *FileSystemModel) findByName(
	node *ds.TreeNode[*FileSystemElement],
	name string,
) *ds.TreeNode[*FileSystemElement] {
	for _, child := range node.Children {
		if child.Data.Name == name {
			return child
		}
	}

	return nil
}

func (fs *FileSystemModel) addElement(
	parent *ds.TreeNode[*FileSystemElement],
	elementType uint8,
	name string,
	size int,
) *ds.TreeNode[*FileSystemElement] {
	path := fs.getPath(parent, name)
	element := NewFileSystemElement(elementType, name, path, size)

	return parent.Insert(element)
}

func (fs *FileSystemModel) exists(node *ds.TreeNode[*FileSystemElement], name string) bool {
	if node == nil {
		return false
	}

	found := fs.findByName(node, name)

	return found != nil
}

func (fs *FileSystemModel) getPath(parent *ds.TreeNode[*FileSystemElement], name string) string {
	parentPathLen := len(parent.Data.Path)
	if string(parent.Data.Path[parentPathLen-1]) == RootToken {
		return parent.Data.Path + name
	} else {
		return parent.Data.Path + RootToken + name
	}
}

func (fs *FileSystemModel) isCommand(line string) bool {
	if strings.HasPrefix(line, CommandIndicator) {
		return true
	}

	return false
}

func (fs *FileSystemModel) isDir(line string) bool {
	if strings.Contains(line, DirIndicator) {
		return true
	}

	return false
}
