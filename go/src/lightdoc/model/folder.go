package model

import (
	"io/ioutil"
	"lightdoc/config"
)

type Folder struct {
	//Folder 为文件夹
	Folder []Folder
	//Path为相对路径
	Path string
	//Files 是文件名
	Files []string
	//Name
	Name string
}

func InitTree(path string) (tree *Folder) {
	tree = new(Folder)
	tree.Filelist(path)
	return
}
func (tree *Folder) Filelist(path string) {
	if tree.Folder == nil {
		tree.Folder = make([]Folder, 0)
	}
	if tree.Files == nil {
		tree.Files = make([]string, 0)
	}
	files, _ := ioutil.ReadDir(path)
	for _, f := range files {
		if f.IsDir() {
			childdir := path + "/" + f.Name()
			childnode := new(Folder)
			childnode.Path = childdir[len(config.Root):]
			childnode.Name = f.Name()
			childnode.Filelist(childdir)
			tree.Folder = append(tree.Folder, *childnode)
		} else {
			tree.Files = append(tree.Files, f.Name())
		}

	}

}
