# textractor.gui

一个网页正文提取的可视化工具，核心是textractor，采用aardio制作。核心项目地址： https://github.com/gloomyzerg/textractor

#### dll编译说明

1. 根据不同的配置，修改build.bat中go，gcc涉及的环境变量
2. 生成的dll文件用upx减小体积


#### 参考

go最小化编译

https://golangforall.com/en/post/go-minimum-docker-image.html

C.free编译报错解决

https://stackoverflow.com/q/68899200

cgo传递char *字符串的方法

http://blog.codeg.cn/post/blog/2016-04-20-golang-cgo

golang 生成 shared object 供其他语言使用

https://segmentfault.com/a/1190000018707206