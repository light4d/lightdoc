package main

import (

"fmt"
"io/ioutil"
"strings"

 
)
//定义接口
type Searcher interface {
		chazhao()
}

type User struct {

}
func (p User) chazhao(){
	var file, err = ioutil.ReadFile("README.md")
    if err!=nil{
        fmt.Println(err)
        return
    }
	//这里的file一定要加上string，否则会打印一个int数组
	var result =string(file)
    fmt.Println(result)  
	var str string
	fmt.Println("请输入你想查找的")
	fmt.Scanf("%s",&str)
	//Contains 是判断第一个result里面是否有str的值，若有就为真，否则为假
	if (strings.Contains(result, str))==true {
		fmt.Println("文章中有您查找的快去查看吧，您查找到的数据是:",str)
  }else{
		fmt.Println("抱歉,没有您查找的内容!\n")
  }

}

func main(){
	
	
 
	var zhangsan = new (User)
	zhangsan.chazhao()
	
	 
}
