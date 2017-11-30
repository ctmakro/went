# went
quietsocks server/client in Go

## 其实就是翻墙工具

之前是用Node.js搞的（见<https://github.com/ctmakro/quietsocks>），体积和消耗都大，处理速度也有限。换成go会好很多。

## 编译

从golang.org下载安装适合你的系统的 Go 1.92（需要翻墙）

在Windows上 请cmd运行build.....bat
在OS X/Linux 上请bash运行build.....sh

将会生成三个文件，分别是windows, linux和OS X下的可执行文件。

在同个系统下，客户端和服务端使用同一个可执行文件，通过命令行开关区别模式。

## 运行

- 服务端

  对外开放8338端口，确保本机8228端口未占用，然后

  ```bash
  $ ./went_linux64
  ```

- 客户端

  确保本机8118端口未占用，然后

  ```bash
  >\ went_win64.exe --connect [服务端IP]
  ```
  然后将浏览器SOCKS5代理指向本机8118端口即可。
  
## 备注

- 练习作品，代码一如既往的简单
- 并没有加密，只是作了按位取反以避免关键词过滤。
- 凡是运营商不认识的TCP协议（既不是HTTP也不是TLS）都可能被严重限速。使用kcptun可以得到极大改善，以后可能并入本项目。
