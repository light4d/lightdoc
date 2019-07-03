package model

import (
	"io/ioutil"
	"lightdoc/config"
	"os"
	"path/filepath"
	"strings"
)

func SerchKeyWords(k string) []string {
	fileindex := SerchTextfile(config.Root)
	webpage := []string{}
	for i, v := range fileindex {
		if Isinthere(k, i) || IsName(k, v) {
			webpage = append(webpage, i)
		}
	}
	return webpage
}

func SerchTextfile(path string) map[string]string {
	textfxp := map[string]int{"md": 1, "txt": 2} //支持的文本格式
	target := map[string]string{}
	filepath.Walk(path,
		func(path string, f os.FileInfo, err error) error {
			if f == nil {
				return err
			}
			for i := len(path) - 1; i > 0; i-- {
				if path[i] == '.' {
					fxe := path[i+1:]
					if _, ok := textfxp[fxe]; ok {
						target[path] = f.Name()
					}
					break
				}
			}
			return nil
		})
	return target
}

//遍历config.Root下所有支持的文本文件并存入target

//遍历所有文本文件查找关键词(Key Words)
func IsName(k string, t string) bool {
	return strings.Contains(t, k)
}

//文件名匹配

func Isinthere(k string, path string) bool {
	d, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	str := string(d)
	return strings.Contains(str, k)
}

//内容匹配
