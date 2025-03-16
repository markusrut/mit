package fs

import (
	"os"
	"path/filepath"
	"strings"
)

type FileNode struct {
	Path     string
	Name     string
	IsDir    bool
	Children []*FileNode
}

func ReadFileStructure(rootPath string) (*FileNode, error) {
	rootNode, err := nodeFromPath(rootPath)

	if err != nil {
		return nil, err
	}

	if !rootNode.IsDir {
		return rootNode, nil
	}

	if err := readNode(rootNode); err != nil {
		return nil, err
	}

	return rootNode, nil
}

func PrintStructure(node *FileNode) {
	formattedString := buildNodeString(node, "", true, true)
	println(formattedString)
}

func buildNodeString(node *FileNode, prefix string, isLast bool, isRoot bool) string {
	var sb strings.Builder

	connector := "├── "
	if isLast {
		connector = "└── "
	}

	if isRoot {
		sb.WriteString(node.Name)
	} else {
		sb.WriteString(prefix + connector + node.Name)
	}

	childPrefix := prefix
	if isRoot {
		childPrefix += ""
	} else if isLast {
		childPrefix += "    "
	} else {
		childPrefix += "│   "
	}

	for i, child := range node.Children {
		isLastChild := i == len(node.Children)-1
		childString := buildNodeString(child, childPrefix, isLastChild, false)
		sb.WriteString("\n" + childString)
	}

	return sb.String()
}

func nodeFromPath(rootPath string) (*FileNode, error) {
	info, err := os.Stat(rootPath)
	if err != nil {
		return nil, err
	}

	rootNode := &FileNode{
		Path:     rootPath,
		Name:     filepath.Base(rootPath),
		IsDir:    info.IsDir(),
		Children: nil,
	}

	return rootNode, nil
}

func readNode(node *FileNode) error {
	entries, err := os.ReadDir(node.Path)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		childPath := filepath.Join(node.Path, entry.Name())

		childNode := &FileNode{
			Path:     childPath,
			Name:     entry.Name(),
			IsDir:    entry.IsDir(),
			Children: nil,
		}

		if entry.IsDir() {
			if err := readNode(childNode); err != nil {
				return err
			}
		}

		node.Children = append(node.Children, childNode)
	}

	return nil
}
