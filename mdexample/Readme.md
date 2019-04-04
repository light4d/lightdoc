# markdown示例


# 一级标题

## 二级标题

### 三级标题

## line
* * * 
line
 
* * * 
line
 
* * * 
line

## 代码


```  go
package main

import (
	"lightdoc/config"
	"lightdoc/route"
	"os"
)

func main() {

	if len(os.Args) > 1 {
		config.Root = os.Args[1]
		config.Dist = os.Args[2]
	}
	err := route.Init()
	if err != nil {
		panic(err)
	}
}
```

## 表格

| Tables        | Are           | Cool  |
| ------------- |:-------------:| -----:|
| col 3 is      | right-aligned | $1600 |
| col 2 is      | centered      |   $12 |
| zebra stripes | are neat      |    $1 |
