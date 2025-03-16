package entity

import (
	"os"
	"path/filepath"
	"strings"
)

type Entity struct {
	Type     EntityType
	Path     string
	Name     string
	Children []*Entity
}

func (e *Entity) IsDir() bool {
	return e.Type == Directory
}

func ReadFileStructure(rootPath string) (*Entity, error) {
	rootNode, err := nodeFromPath(rootPath)

	if err != nil {
		return nil, err
	}

	if !rootNode.IsDir() {
		return rootNode, nil
	}

	if err := readNode(rootNode); err != nil {
		return nil, err
	}

	return rootNode, nil
}

func nodeFromPath(rootPath string) (*Entity, error) {
	info, err := os.Stat(rootPath)
	if err != nil {
		return nil, err
	}

	entityTpe := File
	if info.IsDir() {
		entityTpe = Directory
	}

	rootNode := &Entity{
		Type:     entityTpe,
		Path:     rootPath,
		Name:     filepath.Base(rootPath),
		Children: nil,
	}

	return rootNode, nil
}

func readNode(node *Entity) error {
	entries, err := os.ReadDir(node.Path)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		childPath := filepath.Join(node.Path, entry.Name())

		entityTpe := File
		if entry.IsDir() {
			entityTpe = Directory
		}

		childNode := &Entity{
			Type:     entityTpe,
			Path:     childPath,
			Name:     entry.Name(),
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

func (node *Entity) Print() {
	formattedString := buildNodeString(node, "", true, true)
	println(formattedString)
}

func buildNodeString(node *Entity, prefix string, isLast bool, isRoot bool) string {
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
