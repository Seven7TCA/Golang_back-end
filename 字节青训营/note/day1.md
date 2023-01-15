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

[基础语法](../Data/go-by-example/example/)

note：
- switch 内置 break，可以使用任意类型，可以代替 if-else 语句
- 业务逻辑代码中函数通常返回多个值，第一个值是真正的返回结构，第二个值是错误信息
- Go 中也有指针
- 结构体可以定义方法
- [错误处理](../Data/go-by-example/example/14-error/main.go)
    - 无错误返回 `error = nil`
    - `errors.New("")` 编辑错误信息
- [fmt](../Data/go-by-example/example/16-fmt/main.go)
    - %v 打印任意类型变量，%+v、%#v 打印更详细信息  
    - %.2f 打印2位小数
- [JSON 处理](../Data/go-by-example/example/17-json/main.go)
    - 保证结构体每个字段首字母为大写（Go 中的公开字段），这个结构体就能用 `json.Marshal` 数列化
    - `json.Unmarshal` 反数列化到空变量里
    - 如果要输出小写风格，在结构体字段后面加`'json.“xx”'`
- [time](../Data/go-by-example/example/18-time/main.go)
    - `time.Now()` 获取当前时间
    - `time.Date()` 构造带时区的时间
    - `t.Sub(t2)` 对时间做减法得到一个时间段
    - `t.Format()`、`time.Parse()`  格式化
    - `time.Unix()` 获取时间戳
- [strconv](../Data/go-by-example/example/19-strconv/main.go)
    - 字符串和数字之间的转换
    - `strconv.ParseFloat`、`strconv.ParseInt`解析字符串，`strconv.ParseInt`有3个参数，1-字符串，2-进制（传0时自动解析），3-返回进度
    - `strconv.Atoi()` 快速转换成10进制
    - 输入不合法都会返回错误
- [env](../Data/go-by-example/example/20-env/main.go)
    - 获取进程信息
    - `os.Args` 获取进程在执行时的命令号参数
    - `os.Getenv()`、`os.Setenv()` 获取、写入环境变量
    - `exec.Command()` 快速启动子进程、获取其输入输出

# 