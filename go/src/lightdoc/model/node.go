package model

import (
	"io/ioutil"
	"lightdoc/config"
)

type Node struct {
	//Node 为文件夹
	Nodes []Node
	//Path为相对路径
	Path, Name string
	//Name
	IsDir bool
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
	files, _ := ioutil.ReadDir(path)
	for _, f := range files {
		childdir := path + "/" + f.Name()

		childnode := new(Node)
		childnode.IsDir = f.IsDir()
		childnode.Path = childdir[len(config.Root):]
		childnode.Name = f.Name()
		childnode.Filelist(childdir)

		tree.Nodes = append(tree.Nodes, *childnode)

	}

}
