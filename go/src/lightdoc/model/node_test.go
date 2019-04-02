package model

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestNode_Filelist(t *testing.T) {
	path := "/home/zh/testforgo"
	tree := new(Node)

	tree.Filelist(path)

	fmt.Println(tree)
	bs, _ := json.Marshal(tree)
	fmt.Println(string(bs))
}
