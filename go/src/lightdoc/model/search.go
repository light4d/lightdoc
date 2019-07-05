package model

import (
	"io/ioutil"
	"lightdoc/config"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

type Searchresult struct {
	Path       string
	Key        string
	Isfilename bool
}
type Allsearchresult []Searchresult

func SerchKeyWords(key string) Allsearchresult {
	fileindex := SerchTextfile(config.Root)
	allresult := Allsearchresult{}
	for i, v := range fileindex {
		if Isinthere(key, i) {
			newpath := i[len(config.Root):]
			keywords := SearchText(key, i)
			webpage := Searchresult{newpath, keywords, false}
			allresult = append(allresult, webpage)
			continue
		}
		if IsName(key, v) {
			newpath := i[len(config.Root):]
			webpage := Searchresult{newpath, "", true}
			allresult = append(allresult, webpage)
		}
	}
	return allresult
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

//查询key是否与文件名匹配

func Isinthere(k string, path string) bool {
	ed, _ := ioutil.ReadFile(path)
	a := string(ed)
	re, _ := regexp.Compile("!\\[[^\\]]+\\]\\([^\\)]+\\)")
	a = re.ReplaceAllString(a, "")
	re, _ = regexp.Compile("(\\[.+\\]\\([^\\)]+\\))|(<.+>)")
	a = re.ReplaceAllString(a, "")
	re, _ = regexp.Compile("[\\\\\\`\\*\\_\\[\\]\\#\\+\\-\\!\\>]")
	a = re.ReplaceAllString(a, "")
	re, _ = regexp.Compile("\n")
	a = re.ReplaceAllString(a, "")

	return strings.Contains(a, k)
}

//查询key是否存在于文件内

func SearchText(key string, path string) string {
	ed, _ := ioutil.ReadFile(path)
	a := string(ed)
	re, _ := regexp.Compile("!\\[[^\\]]+\\]\\([^\\)]+\\)")
	a = re.ReplaceAllString(a, "")
	re, _ = regexp.Compile("(\\[.+\\]\\([^\\)]+\\))|(<.+>)")
	a = re.ReplaceAllString(a, "")
	re, _ = regexp.Compile("[\\\\\\`\\*\\_\\[\\]\\#\\+\\-\\!\\>]")
	a = re.ReplaceAllString(a, "")
	re, _ = regexp.Compile("\n")
	a = re.ReplaceAllString(a, "")

	end, nk := []rune(a), []rune(key)

	index := SearchIndex(end, nk)
	if index == -1 {
		return "Not Found"
	}
	indexstart := 0
	indexend := len(a) - 1
	if index-30 >= 0 {
		indexstart = index - 30
	}
	if indexstart+30 <= len(a)-1 {
		indexend = index + 30
	}
	//避免越界
	return string(end[indexstart:indexend])
}

func SearchIndex(body, key []rune) int {
	for i := range body {
		found := true
		for j := range key {
			if body[i+j] != key[j] {
				found = false
				break
			}
		}
		if found {
			return i
		}
	}
	return -1
}
