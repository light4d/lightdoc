# lightdoc
轻量级文档服务

## ![概述](doc/prd.md)
## 使用方法
```
sudo docker run  --name=lightdoc  -v hostmddir:/var/doc/  -p 8080:8080 -p 8000:8000 light4d/lightdoc
``` 
 ## 开放端口
 * 8000 
对外提供文档
 * 8080
 为前端提供的代理端口
 ## logo
 将logo.jpg(或logo.png,logo.gif)放入hostmddir目录下
 将hostmddir换成自己文件夹路径 例如：
  
 ```sudo docker run  --name=lightdoc  -v /media/zh/Data1/CODE/github.com/lightdoc/mdexample:/var/doc/  -p 8080:8080 -p 8000:8000 light4d/lightdoc```
部署流程：
1.安装docker 
2.通过docker拉取lightcdoc镜像 docker pull llight4d/lightdoc
3.运行lightdoc镜像并挂载本地的markdown文件夹（hostmddir） docker run  --name=lightdoc  -v hostmddir:/var/doc/  -p 8080:8080 -p 8000:8000 light4d/lightdoc
4.访问路径：localhost:8080
开发接口：
1.markdown文件读取

