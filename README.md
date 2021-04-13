# wechat-dat2picture
将微信目录中的.dat后缀的图片转换成jpeg、png格式

# 编译

#### for linux
```
make linux
```
#### for windows
```
make windows
```
# 使用
```
./wechat-dat2picture --help
Usage of ./wechat-dat2picture:
  -datPath string
    	待转换的dat文件目录 (default "./dat")
  -picPath string
    	转换后的存放目录 (default "./pic")

示例：
./wechat-dat2picture -datPath=./dat -picPath=./pic
```