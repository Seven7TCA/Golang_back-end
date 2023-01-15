学习内容：Go 语言基础 - 基础语法
- Go 语言基础语法
- Go 语言的实战案例

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
    - 保证结构体每个字段首字母为大写（Go 中的公开字段），这个结构体就能用 `json.Marshal` 数列化
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

# 