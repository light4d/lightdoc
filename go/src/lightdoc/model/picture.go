package model

import (
	"io/ioutil"
	"path"
	"strings"
)

type Picture struct {
	Path string
	Name string
}

func InitPic(path string) (pic *Picture) {
	pic = new(Picture)
	pic.Piclist(path)
	return
}

func (pic *Picture) Piclist(p string) (picture *Picture) {
	files, _ := ioutil.ReadDir(p)

	for _, f := range files {
		if !f.IsDir() {
			sname := path.Ext(f.Name())
			pname := strings.TrimSuffix(f.Name(), sname)
			switch pname {
			case "logo":
				{
					pic.Name = f.Name()
					pic.Path = "/" + f.Name()
				}
			default:
				continue
			}
		}
	}
	return
}
