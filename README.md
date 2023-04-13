## 编译成可执行文件
```shell
go build main.go
```

## 编译成deb文件
1. 把生成的可执行文件移到go-nmap/usr/local/bin下并改名为go-nmap
2. 回到最上层go-nmap目录的父目录
3. 安装依赖
```shell
sudo apt-get install dpkg-dev
```
4. 打包
```shell
dpkg-deb --build go-nmap
```
5. 安装卸载
```shell
sudo dpkg -i go-nmap.deb
sudo dpkg -r go-nmap
```
