学习内容：「Go 框架三件套详解(Web/RPC/ORM)」
- 三件套(Gorm、Kitex、Hertz)介绍
- 三件套的使用
- 实战案例

**课程目标：**
- 将前面几节课所学到的知识应用到项目中
- 掌握 Hertz/Kitex/Gorm 的基本用法
- 通过学习实战案例，可以使用 Hertz/Kitex/Gorm 完成日常后端开发任务

# 三件套介绍

![三件套介绍](../assets/%E4%B8%89%E4%BB%B6%E5%A5%97%E4%BB%8B%E7%BB%8D.png)

# 三件套使用

## Gorm

### 基本使用

[Gorm指南-模型定义](https://gorm.cn/zh_CN/docs/models.html)

![Gorm 的基本使用](../assets/Gorm%E5%9F%BA%E6%9C%AC%E4%BD%BF%E7%94%A8.png)
- 结构体是 Gorm 的模型，对应数据库的一张表，字段是数据库里的每个字段
- TableName() 返回的是表明，是一个约定（默认）
- gorm.Open() 根据不同驱动支持不同数据库，通过第二个参数：&gorm.Config{} 传递自定义配置，这样就初始化了一个数据库的连接，初始化了 gorm 的对象 `db`

gorm的基本操作
- db.Create() 支持创造一条或多条数据，通过传递一个对象或一个切片
- db.First() 查询数据，需要传指针，只支持传递一条数据
- db.Model(&product).Update() 更新字段，db.Model(&product).Updates() 传递多条，通过传递 Map 可更新零值
- db.Delete() 删除数据

Gorm 的约定（默认）：
- Gorm 使用名为 ID 的字段 作为主键
- 使用结构体的 蛇形负数作为表名
- 字段名的蛇形作为列名
- 使用 CreatedAt、UpdatedAt 字段作为创建更新时间

**Gorm 支持的数据库**

GORM 目前支持 MySQL、SQLServer、PostgreSQL、SQLite。
```go
//连接 SQLServer 数据库为例
import (
"gorm.io/driver/sqlserver"
"gorm.io/gorm"
)
// github.com/denisenkom/go-mssqldb
dsn := "sqlserver://gorm:LoremIpsum86@localhost:9930?database=gorm"
db,err := gorm.Open(sqlserver.Open(dsn),&gorm.Config{})
```
GORM 通过驱动来连接数据库，如果需要连接其它类型的数据库，可以复用/自行开发驱动。

[什么是DSN-MySQL](https://github.com/go-sql-driver/mysql#dsn-data-source-name)
- dsn 就是包含一些连接数据库的信息

### Gorm创建数据

[Gorm指南-创建数据](https://gorm.cn/zh_CN/docs/create.html)

![Gorm创建数据](../assets/Gorm%E5%88%9B%E5%BB%BA%E6%95%B0%E6%8D%AE.png)
- db.Create() 创建一条数据时，因为 gorm 是链式调用，所以会返回一个 gorm 对象
- 创建多条数据通过切片

**更新数据**
- 使用 clause.OnConflict 处理数据冲突
- 图中处理方式为：不处理冲突，把数据加进去

**如何使用默认值**
- 通过 default 标签为字段定义默认值

### GORM 查询数据

[Gorm指南-查询数据](https://gorm.cn/zh_CN/docs/query.html)

![GORM 查询数据](../assets/GORM%E6%9F%A5%E8%AF%A2%E6%95%B0%E6%8D%AE.png)
- .RowsAffected 返回找到的记录数

First 的使用踩坑
- 使用 First 时，需要注意查询不到数据会返回 `ErrRecordNotFound`
- 使用 Find 查询多条数据，查询不到数据不会返回错误（通常使用）

使用结构体作为查询条件
- 当使用结构作为条件查询时，GORM 只会查询非零值字段。这意味着如果您的字段值为 0、"、false 或其他 零值，该字段不会被用于构建查询条件
- 使用Map 来构建查询条件

### GORM 更新数据

[Gorm指南-更新数据](https://gorm.cn/zh_CN/docs/update.html)

![GORM更新数据](../assets/Gorm%E6%9B%B4%E6%96%B0%E6%95%B0%E6%8D%AE.png)



