学习内容：Go 语言基础 - 基础语法
- Go 语言基础语法
- Go 语言的实战案例

[toc]

[手册-1](https://juejin.cn/post/7188225875211452476)

# Go 语言基础语法

## Go 语言简介

什么是 Go 语言

1. 高性能、高并发
- 内定高并发的支持，不需要寻找高度性能优化的第三方性能库
2. 语法简单、学习曲线平缓
3. 丰富的标准库
- 标准库有很高的稳定性和兼容性保障
4. 完善的工具链
- 支持单元测试、性能测试、代码覆盖率等
5. 静态链接
- 编译结构默认都是静态链接的，只需要拷贝编译之后唯一的可执行文件
6. 快速编译
7. 跨平台
- 交叉编译
8. 垃圾回收
- 无需考虑内存分配释放

## Go 入门

开发环境 - 配置集成开发环境
- VS code or Goland
- 基于云的开发环境

[基础语法](https://github.com/wangkechun/go-by-example/tree/master/example)

note：
- switch 内置 break，可以使用任意类型，可以代替 if-else 语句
- 业务逻辑代码中函数通常返回多个值，第一个值是真正的返回结构，第二个值是错误信息
- Go 中也有指针
- 结构体可以定义方法
- 错误处理
    - 无错误返回 `error = nil`
    - `errors.New("")` 编辑错误信息
```go
package main

import (
	"errors"
	"fmt"
)

type user struct {
	name     string
	password string
}

func findUser(users []user, name string) (v *user, err error) {
	for _, u := range users {
		if u.name == name {
			return &u, nil
		}
	}
	return nil, errors.New("not found")
}

func main() {
	u, err := findUser([]user{{"wang", "1024"}}, "wang")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(u.name) // wang

	if u, err := findUser([]user{{"wang", "1024"}}, "li"); err != nil {
		fmt.Println(err) // not found
		return
	} else {
		fmt.Println(u.name)
	}
}
```
- fmt
    - %v 打印任意类型变量，%+v、%#v 打印更详细信息  
    - %.2f 打印2位小数
```go
package main

import "fmt"

type point struct {
	x, y int
}

func main() {
	s := "hello"
	n := 123
	p := point{1, 2}
	fmt.Println(s, n) // hello 123
	fmt.Println(p)    // {1 2}

	fmt.Printf("s=%v\n", s)  // s=hello
	fmt.Printf("n=%v\n", n)  // n=123
	fmt.Printf("p=%v\n", p)  // p={1 2}
	fmt.Printf("p=%+v\n", p) // p={x:1 y:2}
	fmt.Printf("p=%#v\n", p) // p=main.point{x:1, y:2}

	f := 3.141592653
	fmt.Println(f)          // 3.141592653
	fmt.Printf("%.2f\n", f) // 3.14
}
```
- JSON 处理
    - 保证结构体每个字段首字母为大写（Go 中的公开字段），这个结构体就能用 `json.Marshal` 序列化
    - `json.Unmarshal` 反数列化到空变量里
    - 如果要输出小写风格，在结构体字段后面加`'json.“xx”'`
```go
package main

import (
	"encoding/json"
	"fmt"
)

type userInfo struct {
	Name  string
	Age   int `json:"age"`
	Hobby []string
}

func main() {
	a := userInfo{Name: "wang", Age: 18, Hobby: []string{"Golang", "TypeScript"}}
	buf, err := json.Marshal(a)
	if err != nil {
		panic(err)
	}
	fmt.Println(buf)         // [123 34 78 97...]
	fmt.Println(string(buf)) // {"Name":"wang","age":18,"Hobby":["Golang","TypeScript"]}

	buf, err = json.MarshalIndent(a, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(buf))

	var b userInfo
	err = json.Unmarshal(buf, &b)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", b) // main.userInfo{Name:"wang", Age:18, Hobby:[]string{"Golang", "TypeScript"}}
}
```
- time
    - `time.Now()` 获取当前时间
    - `time.Date()` 构造带时区的时间
    - `t.Sub(t2)` 对时间做减法得到一个时间段
    - `t.Format()`、`time.Parse()`  格式化
    - `time.Unix()` 获取时间戳
```go
package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now) // 2022-03-27 18:04:59.433297 +0800 CST m=+0.000087933
	t := time.Date(2022, 3, 27, 1, 25, 36, 0, time.UTC)
	t2 := time.Date(2022, 3, 27, 2, 30, 36, 0, time.UTC)
	fmt.Println(t)                                                  // 2022-03-27 01:25:36 +0000 UTC
	fmt.Println(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute()) // 2022 March 27 1 25
	fmt.Println(t.Format("2006-01-02 15:04:05"))                    // 2022-03-27 01:25:36
	diff := t2.Sub(t)
	fmt.Println(diff)                           // 1h5m0s
	fmt.Println(diff.Minutes(), diff.Seconds()) // 65 3900
	t3, err := time.Parse("2006-01-02 15:04:05", "2022-03-27 01:25:36")
	if err != nil {
		panic(err)
	}
	fmt.Println(t3 == t)    // true
	fmt.Println(now.Unix()) // 1648738080
}
```
- strconv
    - 字符串和数字之间的转换
    - `strconv.ParseFloat`、`strconv.ParseInt`解析字符串，`strconv.ParseInt`有3个参数，1-字符串，2-进制（传0时自动解析），3-返回进度
    - `strconv.Atoi()` 快速转换成10进制
    - 输入不合法都会返回错误
```go
package main

import (
	"fmt"
	"strconv"
)

func main() {
	f, _ := strconv.ParseFloat("1.234", 64)
	fmt.Println(f) // 1.234

	n, _ := strconv.ParseInt("111", 10, 64)
	fmt.Println(n) // 111

	n, _ = strconv.ParseInt("0x1000", 0, 64)
	fmt.Println(n) // 4096

	n2, _ := strconv.Atoi("123")
	fmt.Println(n2) // 123

	n2, err := strconv.Atoi("AAA")
	fmt.Println(n2, err) // 0 strconv.Atoi: parsing "AAA": invalid syntax
}
```
- env
    - 获取进程信息
    - `os.Args` 获取进程在执行时的命令号参数
    - `os.Getenv()`、`os.Setenv()` 获取、写入环境变量
    - `exec.Command()` 快速启动子进程、获取其输入输出
```go
package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	// go run example/20-env/main.go a b c d
	fmt.Println(os.Args)           // [/var/folders/8p/n34xxfnx38dg8bv_x8l62t_m0000gn/T/go-build3406981276/b001/exe/main a b c d]
	fmt.Println(os.Getenv("PATH")) // /usr/local/go/bin...
	fmt.Println(os.Setenv("AA", "BB"))

	buf, err := exec.Command("grep", "127.0.0.1", "/etc/hosts").CombinedOutput()
	if err != nil {
		panic(err)
	}
	fmt.Println(string(buf)) // 127.0.0.1       localhost
}
```

# Go 语言的实战案例

- 学习代码生成
- 一定要在本地运行一下，理解更透彻

## [猜谜游戏](https://github.com/wangkechun/go-by-example/tree/master/guessing-game)

生成随机数：(见 v2 版本)
- 需要用 `math/rand` 包
- `rand.Intn(maxNum)` 生成随机数，范围是 0~`maxNum`
	- 需要设置随机数种子，否则每次都会生成相同的随机数类(见 v1 版本)
	- 用时间戳来初始化随机数种子，加上一行` rand.Seed(time.Now().UnixNano)`

读取用户输入：(见 v3 版本)
- 每个程序执行时都会打开几个文件，包括 Stdin、Stdout 等
- `os.Stdin` 得到
- `bufio.NewReader` 将 Stdin 转化为只读的流
- `reader.ReadString('\n')` 读取一行
- `input = strings.Trim(input, "\r\n")` 去掉Tab、换行符
- `strconv.Atoi(input)` 转换为数字

实现判断逻辑（见 v4 版本）、实现游戏循环（见 v5 版本）

## [在线词典介绍](https://github.com/wangkechun/go-by-example/tree/master/simpledict)

学习如何用 Go 语言发送 http 请求、解析 json 、如何使用代码生成提高开发效率

抓包
- 在网站中右键点“检查”（以[彩云翻译app](https://fanyi.caiyunapp.com/)为例）
- 点翻译，在 Network 的 name 中找到 `dict`（需要是Request Method 为 POST 的），在Payload 和 Preview 中可以看到相关信息
-  右键 dict 选择 copy as curl（有 bash 选 bash）得到 curl

生成代码
- [代码生成网址](https://curlconverter.com/go/)，粘贴复制的 curl 请求
- copy Golang 代码，把转义导致的编译错误的代码删掉即可

生成代码解读（代码见 v1 版本）
```go
var data = strings.NewReader(`{"trans_type":"en2zh","source":"good"}`)
req, err := http.NewRequest("POST", "https://api.interpreter.caiyunai.com/v1/dict", data)
```
- 创建请求，3个参数：POST、curl、data
- data 是一个流，将输入的字符串转换成流，只需要占用很少的内存去创建请求
- `req.Header.Set()` 设置请求头
- `resp, err := client.Do(req)` 发起请求
	- resp 是个流，为了防止资源泄露，需要 `defer resp.Body.Close()` 手动关闭，defer 会在函数结束后从下往上触发
- `bodyText, err := ioutil.ReadAll(resp.Body)` 将流读成 byte 数组

生成 request body（v2 版本）
- 将 json 序列化，需要构造结构体，与 json 的结构是对应的，`json.Marshal()` 序列化为一个 byte 数组
- `json.Marshal()` 返回的是 byte 数组，所以换成 `bytes.NewReader()`

解析（v3 版本）
- [代码生成](https://oktools.net/json2go)，将 网站 Priview 中的 json 复制上去 转换-嵌套
```go
var dictResponse DictResponse
	err = json.Unmarshal(bodyText, &dictResponse)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%#v\n", dictResponse)
```
- 用 `json.Unmarshal(bodyText, &dictResponse)` 反序列化到变量中
- %#v 详细打印结构体

打印结果（v4 版本）
```go
if resp.StatusCode != 200 {
	log.Fatal("bad StatusCode:", resp.StatusCode, "body", string(bodyText))
}
```
- 防御性编程，检测返回的 response 是正确的

## [SOCKS5代理](https://github.com/wangkechun/go-by-example/tree/master/proxy)

- TCP echo server 源代码（v1 版本）
	- 用来测试 server 正确性，发送啥回复啥
- auth（v2 阶段）
	- 实现认证阶段（auth 函数）
- 请求阶段 （v3 版本）
- relay 阶段（v4 版本）
	- 用 `context.WithCancel(context.Background())` 创建 ctx
	- `<-ctx.Done()` 等待ctx 执行完成（是任何一个func（）中 cancel（）函数调用的时机）
```go
ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		_, _ = io.Copy(dest, reader)
		cancel()
	}()
	go func() {
		_, _ = io.Copy(conn, dest)
		cancel()
	}()

	<-ctx.Done()
```
