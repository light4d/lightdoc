package model

import (
	"io/ioutil"
)

type Node struct {
	Nodes []Node
	Path  string
	Files []string
}

func InitTree(path string) (tree *Node) {
	tree = new(Node)
	tree.Filelist(path)
	return
}
func (tree *Node) Filelist(path string) {
	if tree.Nodes == nil {
		tree.Nodes = make([]Node, 0)
	}
	if tree.Files == nil {
		tree.Files = make([]string, 0)
	}
	files, _ := ioutil.ReadDir(path)
	for _, f := range files {
		if f.IsDir() {
			childdir := path + "/" + f.Name()
			childnode := new(Node)
			childnode.Path = childdir
			childnode.Filelist(childdir)
			tree.Nodes = append(tree.Nodes, *childnode)
		} else {
			tree.Files = append(tree.Files, f.Name())
		}

	}

}
