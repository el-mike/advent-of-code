package day_07

import (
	"el-mike/advent-of-code/go/common"
	"el-mike/advent-of-code/go/common/ds"
)

const (
	MaxSize       = 100000
	TotalSpace    = 70000000
	RequiredSpace = 30000000
)

func NoSpaceLeftOnDevice() int {
	scanner, err := common.GetFileScanner("./year_2022/day_07/" + common.InputFilename)
	if err != nil {
		panic(err)
	}

	fileSystemModel := NewFileSystemModel()

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			continue
		}

		fileSystemModel.parseLine(line)
	}

	// With the file system tree built, we can traverse it Depth-first to calculate
	// the sizes of all directories.
	fileSystemModel.Tree.TraverseDF(func(node *ds.TreeNode[*FileSystemElement]) {
		if node.Parent != nil {
			node.Parent.Data.Size += node.Data.Size
		}
	})

	totalUsage := fileSystemModel.Tree.Root.Data.Size

	unusedSize := TotalSpace - totalUsage
	minSize := RequiredSpace - unusedSize

	var candidate *ds.TreeNode[*FileSystemElement]

	fileSystemModel.Tree.TraverseBF(func(node *ds.TreeNode[*FileSystemElement]) {
		if node.Data.Type == DirElementType && node.Data.Size >= minSize {
			if candidate == nil || node.Data.Size < candidate.Data.Size {
				candidate = node
			}
		}
	})

	return candidate.Data.Size

}
