# Python 开发者的 Go 快速入门教程

> 适合谁：已经会 Python，但第一次系统学习 Go，并希望尽快进入后端和工程实战的人。
>
> 最后核实日期：2026-04-14  
> 版本说明：本文中“当前 Go 官方下载版本”以 [go.dev/dl](https://go.dev/dl/) 为准。写作时官方 Featured downloads 显示为 `Go 1.26.1`。

## 先说结论

如果你是 Python 开发者，学习 Go 最容易踩坑的地方不是语法，而是思维方式。

- Python 的核心体验是解释执行、动态类型、运行时灵活、生态丰富。
- Go 的核心体验是编译执行、静态类型、显式错误处理、标准库强、部署简单。

所以不要把 Go 理解成“语法更啰嗦但更快的 Python”。更准确的说法是：

- Python 擅长快速表达想法。
- Go 擅长把服务写得简单、稳定、易部署、易维护。

如果你带着这个前提看下面的内容，很多设计选择都会自然很多。

---

## 1. Python 和 Go 的核心区别

先看一张迁移视角下最重要的对照表。

| 主题 | Python | Go | 你需要切换的思维 |
| --- | --- | --- | --- |
| 执行方式 | 解释执行 | 编译执行 | Go 更强调构建产物和编译期检查 |
| 类型系统 | 动态类型 | 静态类型 | 很多错误会提前在编译期暴露 |
| 包管理 | `pip` / `uv` / `poetry` | `go mod` + `go` toolchain | Go 官方工具链覆盖更多默认工作流 |
| 环境隔离 | 常见是 `.venv` | 通常不按项目创建语言运行时隔离 | Go 更像“装一个 toolchain，项目按 module 管依赖” |
| 错误处理 | `raise` / `try` / `except` | `error` 返回值 + `if err != nil` | 错误处理更显式，也更靠近调用现场 |
| OOP | 类、继承、鸭子类型 | `struct`、method、interface、组合 | 少想继承，多想组合和接口 |
| 并发 | `threading` / `multiprocessing` / `asyncio` | goroutine / channel / `context` | Go 并发是语言一等公民，但不是越多越好 |
| 部署 | 常带解释器和依赖环境 | 常见是单个可执行文件 | Go 在服务部署上通常更省心 |
| 测试 | `pytest` 很强大 | 标准库 `testing` 足够强 | Go 默认鼓励轻量、直接、表驱动测试 |

### 一句话理解 Go 的气质

Go 很少追求“最优雅的语言表达力”，它更追求：

- 团队里不同水平的人都能读懂
- 工具链统一
- 编译快
- 部署简单
- 线上服务行为可预测

这也是为什么很多 Python 开发者初学 Go 时会觉得“有点笨”，但写几周服务之后又会觉得“很省心”。

---

## 2. 从 `uv` 视角理解 Go 工具链

如果你已经熟悉 `uv`，那你会很关心一个问题：

> Go 里谁相当于 `uv`？

答案是：**没有一个一模一样的单点工具**。  
Go 的做法不是把所有能力都收进一个第三方工具，而是把核心能力直接做进官方 `go` 命令。

### 2.1 先理解 Go 里“环境”和 Python 很不一样

Python 项目通常是：

1. 安装 Python 解释器
2. 创建 `.venv`
3. 安装依赖
4. 运行脚本或应用

Go 项目通常是：

1. 安装 Go toolchain
2. 创建或进入一个 module
3. 让 `go` 解析和下载依赖
4. 直接 `go run`、`go test`、`go build`

默认情况下，Go 不要求你为每个项目创建一个独立运行时环境。依赖版本写在 `go.mod` 和 `go.sum` 里，由官方工具链管理。

### 2.2 `uv` 和 Go 常用命令对照

| Python / uv 心智模型 | Go 中更接近的东西 | 说明 |
| --- | --- | --- |
| `uv python install` | 安装 Go toolchain | Go 通常直接安装一个官方版本 |
| `uv venv` | 通常没有完全等价物 | Go 很少按项目创建解释器隔离环境 |
| `uv init` | `go mod init` | 初始化一个 module |
| `uv add requests` | `go get module/path` | 向 `go.mod` 添加依赖 |
| `uv sync` | `go mod tidy` / `go mod download` | 同步和整理依赖 |
| `uv lock` | `go.sum` 自动维护 | Go 没有完全一样的独立 lock 命令 |
| `uv run app.py` | `go run .` 或 `go run ./cmd/server` | 直接编译并运行入口 |
| `uv tool install ruff` | `go install module/path@version` | 安装 Go CLI 工具 |
| `uv pip list` | `go list -m all` | 查看当前模块依赖 |

### 2.3 你最该先会的 Go 命令

```bash
go version
go env
go mod init example.com/hello
go mod tidy
go run .
go test ./...
go build ./...
go install github.com/golangci/golangci-lint/v2/cmd/golangci-lint@latest
```

它们分别可以这样理解：

- `go version`：看当前 toolchain 版本
- `go env`：看 Go 的环境信息
- `go mod init`：创建当前项目的 module
- `go mod tidy`：整理依赖，把真正用到的写入 `go.mod` / `go.sum`
- `go run`：编译后直接运行
- `go test`：跑测试
- `go build`：生成可执行文件或检查可构建性
- `go install module@version`：安装某个 Go 命令行工具

### 2.4 `GOROOT`、`GOPATH`、`GOMODCACHE` 到底是什么

这是 Python 开发者第一次学 Go 时很容易困惑的点。

- `GOROOT`：Go 安装目录，通常不用手动改
- `GOPATH`：老时代很重要，现在主要用来放缓存、安装的工具和 module cache
- `GOMODCACHE`：Go 依赖缓存目录

最重要的实际结论只有两个：

1. 大多数现代 Go 项目里，你几乎只需要关心 `go.mod`
2. 大多数情况下，你不需要手工折腾 `GOPATH`

可以先用下面命令看看本机情况：

```bash
go env GOROOT GOPATH GOMODCACHE
```

---

## 3. 安装 Go 和开始第一个项目

### 3.1 安装

请优先按 Go 官方下载页安装：

- 下载页：[https://go.dev/dl/](https://go.dev/dl/)
- 安装说明：[https://go.dev/doc/install](https://go.dev/doc/install)

安装后先确认：

```bash
go version
```

如果你看到类似下面输出，就说明 toolchain 已经可用：

```bash
go version go1.26.1 windows/amd64
```

你的具体 patch 版本可能更新，只要是当前稳定版即可。

### 3.2 创建第一个 Go 项目

```bash
mkdir hello-go
cd hello-go
go mod init example.com/hello-go
```

新建 `main.go`：

```go
package main

import "fmt"

func main() {
	fmt.Println("Hello, Go")
}
```

运行：

```bash
go run .
```

### 3.3 这里发生了什么

- `package main`：这个包会生成可执行程序
- `func main()`：程序入口
- `go mod init`：创建 `go.mod`
- `go run .`：把当前目录作为入口构建并执行

对于 Python 开发者来说，可以这样类比：

- `module` 更像“项目级依赖边界”
- `package` 更像“代码组织单元”
- `main.go` 更像“应用入口脚本”，但它不是解释执行，而是先编译再运行

### 3.4 Go 的导出规则非常简单

Go 没有 `public/private` 关键字。  
规则只有一个：

- 首字母大写：导出
- 首字母小写：包内私有

例如：

```go
type User struct{}      // 可导出
type userRepo struct{}  // 包内可见

func NewUser() User {}  // 可导出
func loadUser() User {} // 包内可见
```

这和 Python 完全不同。Python 更依赖约定，Go 把这件事做成了语言层级的显式规则。

---

## 4. 语法：只学 Python 开发者最需要的部分

你不需要一开始就把整本语法手册背下来。先掌握这些最常用的结构。

### 4.1 变量、类型和短变量声明

Python：

```python
name = "tudou"
count = 3
```

Go：

```go
var name string = "tudou"
count := 3
```

要点：

- `:=` 只能在函数内部用
- Go 会做类型推断，但它仍然是静态类型
- 类型不匹配多数会在编译时报错，不会拖到运行时

### 4.2 数组、slice、map

这是最容易“看着像 Python，实际不是”的地方。

#### 数组

Go 数组长度是类型的一部分，实际业务里不如 slice 常用。

```go
var a [3]int = [3]int{1, 2, 3}
```

#### slice

Python 的 `list` 更接近 Go 的 `slice`，但不完全一样。

```go
nums := []int{1, 2, 3}
nums = append(nums, 4)
fmt.Println(nums[1:3])
```

记住三件事：

- `slice` 是对底层数组的视图
- `append` 可能返回新的底层存储，所以要接住返回值
- `nil slice` 和空 slice 在某些场景里要区分

#### map

Python：

```python
user = {"name": "tudou", "age": 18}
```

Go：

```go
user := map[string]any{
	"name": "tudou",
	"age":  18,
}
```

读取 map 时要习惯“第二返回值”：

```go
name, ok := user["name"]
if !ok {
	// key 不存在
}
```

这和 Python 的 `dict.get()` 心智不同。Go 倾向显式说明“到底有没有这个值”。

### 4.3 `struct`：先把它理解成“有字段的数据结构”

Python 常用 `class` 或 `dataclass`：

```python
from dataclasses import dataclass

@dataclass
class User:
    name: str
    age: int
```

Go 更常见是：

```go
type User struct {
	Name string
	Age  int
}
```

### 4.4 方法

Go 没有类，但可以给类型绑定方法。

```go
type User struct {
	Name string
}

func (u User) Greet() string {
	return "hello, " + u.Name
}
```

如果你把 Go 强行理解成“没有 class 的 OOP 残缺版”，会很难受。  
更好的理解是：

- 数据结构：`struct`
- 行为：method
- 复用方式：组合优先

### 4.5 指针

Python 开发者一看到指针常常会紧张，但 Go 的指针比 C 简单很多。

你主要只需要知道：

- 用 `*T` 表示“指向 T 的指针”
- 用 `&x` 取地址
- Go 没有指针运算

典型场景是“方法要修改原对象”：

```go
type Counter struct {
	Value int
}

func (c *Counter) Inc() {
	c.Value++
}
```

### 4.6 接口

这是 Go 最优雅、也最容易被 Python 开发者误解的地方。

Python 经常依赖鸭子类型：

```python
def write_data(writer, data):
    writer.write(data)
```

Go 会显式写接口：

```go
type Writer interface {
	Write([]byte) (int, error)
}
```

但 Go 的关键点是：

> 类型不需要显式声明“我实现了某接口”，只要方法集匹配就算实现了。

这意味着 Go 既比 Java 轻，也比 Python 更可静态检查。

另一个非常重要的实践是：

> 接口通常定义在使用方，不定义在实现方。

这个习惯会让你的设计更小、更灵活。

### 4.7 多返回值

Python 会返回 tuple：

```python
def divide(a, b):
    if b == 0:
        raise ValueError("division by zero")
    return a / b
```

Go 更常见：

```go
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return a / b, nil
}
```

调用：

```go
result, err := divide(10, 2)
if err != nil {
	return err
}
fmt.Println(result)
```

这就是 Go 里错误处理的基本节奏。

---

## 5. 错误处理：从异常思维切到显式返回

如果你只记住一件事，那就是：

> Go 不把“普通业务错误”当成异常机制来处理。

### 5.1 Python 的习惯

```python
def load_config(path: str) -> dict:
    with open(path) as f:
        return json.load(f)
```

错误由异常向上冒泡。

### 5.2 Go 的习惯

```go
func LoadConfig(path string) ([]byte, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read config %s: %w", path, err)
	}
	return data, nil
}
```

要点：

- 错误是返回值，不是控制流的“隐藏通道”
- 出错时要加上下文
- 优先返回原始错误链，而不是只打印日志

### 5.3 你应该掌握的错误处理 API

```go
fmt.Errorf("load user %d: %w", id, err)
errors.Is(err, os.ErrNotExist)
errors.As(err, &targetErr)
```

### 5.4 Go 错误处理最佳实践

- 不要在底层到处 `log.Println` 再继续返回错误
- 日志记录通常留在边界层做
- 包裹错误时补充上下文，但不要把无意义字符串一层层堆满
- 不要把每个错误都设计成复杂的自定义类型

---

## 6. 面向对象：少想继承，多想组合

Python 项目里，很多人会自然写出继承层级。  
Go 社区的默认答案通常是：**组合优于继承**。

### 6.1 一个常见误区

Python 开发者初学 Go，常会下意识寻找：

- 抽象基类
- service 基类
- repository 基类
- controller 基类

Go 一般不鼓励这样。

### 6.2 Go 更推荐的方式

```go
type Clock interface {
	Now() time.Time
}

type Service struct {
	clock Clock
}
```

重点是：

- 依赖什么，就注入什么
- 只为“调用方真正需要的能力”定义接口
- 不为对称和层次感而设计抽象

### 6.3 一句经验

如果你发现自己在 Go 里写出一堆 `BaseXxx`、`AbstractXxx`、`CommonXxx`，大概率已经开始把 Go 写成 Java 或 Python 了。

---

## 7. 并发：Go 的强项，但不是让你无脑并发

### 7.1 从 Python 迁移时怎么理解

| Python | Go |
| --- | --- |
| `threading` | goroutine |
| `queue.Queue` | channel |
| `asyncio` 的 task/cancel | goroutine + `context.Context` |
| `multiprocessing` | Go 通常先尝试 goroutine，CPU 密集再具体分析 |

### 7.2 第一个 goroutine

```go
go func() {
	fmt.Println("run in goroutine")
}()
```

这会启动一个并发任务。  
但注意：主协程退出后，程序就结束了，所以演示时通常要配合同步手段。

### 7.3 用 `sync.WaitGroup`

```go
var wg sync.WaitGroup

for i := 0; i < 3; i++ {
	wg.Add(1)
	go func(n int) {
		defer wg.Done()
		fmt.Println(n)
	}(i)
}

wg.Wait()
```

### 7.4 channel 基础

```go
ch := make(chan string)

go func() {
	ch <- "done"
}()

msg := <-ch
fmt.Println(msg)
```

什么时候适合用 channel：

- goroutine 之间要传递结果
- 你确实想表达“通信和同步”

什么时候不一定适合用 channel：

- 只是想保护共享数据
- 只是想等一组任务完成

这种时候 `sync.Mutex`、`sync.WaitGroup` 可能更简单。

### 7.5 `context.Context` 很重要

在 Go 后端里，`context.Context` 基本是必学项。

它主要解决：

- 超时
- 取消
- 请求级边界传播

比如数据库、HTTP 请求、RPC 调用，都常常会传 `context.Context`。

```go
ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
defer cancel()
```

### 7.6 并发最佳实践

- 先写对，再并发
- 不要为了“看起来更快”就到处开 goroutine
- 有共享状态时，先把数据边界想清楚
- 跑测试时学会使用 race detector：

```bash
go test -race ./...
```

---

## 8. Web 开发：Go 里怎么选标准库和框架

这是 Python 开发者最常问的问题之一。

> Go 里有 Django / FastAPI 那样的“默认答案”吗？

没有完全等价的单一答案。

Go 的常见路线是：

1. 标准库 `net/http`
2. 在此基础上加轻量路由器或中间件
3. 只在必要时上更完整的 Web 框架

### 8.1 我的默认建议

如果你是第一次学 Go：

- 学习路线首选：`net/http` + `chi`
- 团队快速落地：`Gin`
- 想要更简洁均衡：`Echo`
- 想要更接近 Express 风格：`Fiber`

### 8.2 框架选择建议

| 方案 | 适合谁 | 优点 | 注意点 |
| --- | --- | --- | --- |
| `net/http` + `chi` | 想学 Go 原生风格的人 | 贴近标准库、长期维护舒服、依赖少 | 需要自己多做一点组装 |
| `Gin` | 团队项目、资料优先 | 生态成熟、示例多、上手快 | 有自己的一套风格，抽象略重于原生 |
| `Echo` | 想要 API 简洁的人 | 易读、功能均衡 | 社区体量略小于 Gin |
| `Fiber` | 喜欢 Express 风格、在意性能的人 | 写法顺手、性能卖点强 | 与标准库生态思路差异更大 |

### 8.3 什么时候不用框架

如果你的需求只是：

- 一个内部服务
- 少量接口
- 健康检查
- 简单 JSON API

那么直接用 `net/http` 完全没问题。

先看一个最小 HTTP 服务：

```go
package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]string{
			"status": "ok",
		})
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
```

这已经能支撑很多真实内部工具和小服务的起点。

### 8.4 如果你正在用这个仓库

这个仓库当前模板基线是 Gin。  
也就是说：

- 学 Go 本身时，我建议你理解 `net/http` 和 `chi`
- 但如果你是要在当前仓库基础上直接开发，那就遵循仓库现有 Gin 结构，不要为了“更原生”再推翻模板

---

## 9. 数据库访问：不要一上来就沉迷 ORM

Python 生态里，Django ORM、SQLAlchemy 非常强，所以很多人学 Go 时也会先找“最好用 ORM”。

我的建议是：

- 先理解 `database/sql`
- 然后优先考虑 `sqlc`
- 真有需要再上 `GORM`

### 9.1 为什么先学 `database/sql`

因为它是 Go 标准库的一部分，能帮助你理解：

- 连接池
- 查询生命周期
- `context`
- 扫描结果
- 显式错误处理

最小示例：

```go
func GetUserName(ctx context.Context, db *sql.DB, id int64) (string, error) {
	var name string
	err := db.QueryRowContext(ctx, "SELECT name FROM users WHERE id = ?", id).Scan(&name)
	if err != nil {
		return "", fmt.Errorf("query user %d: %w", id, err)
	}
	return name, nil
}
```

### 9.2 为什么推荐 `sqlc`

`sqlc` 的思路是：

- SQL 你自己写
- 它帮你生成类型安全的 Go 代码

这很适合已经有 SQL 基础、又不想失去查询可控性的后端开发者。

### 9.3 什么时候用 `GORM`

适合这些情况：

- 你要快速搭一个 CRUD 后台
- 团队对 ORM 很熟
- 业务查询不复杂

但要注意：

- ORM 抽象会隐藏 SQL 成本
- 复杂查询和性能调优时，经常还是要回到 SQL

### 9.4 默认建议

如果你是 Python 后端转 Go：

1. 先学 `database/sql`
2. 项目里优先评估 `sqlc`
3. 不要把 ORM 当成默认起点

---

## 10. 测试：Go 的默认测试体验是什么样

Python 开发者常问：

> Go 里是不是一定要找 `pytest` 替代品？

通常不需要。  
Go 标准库的 `testing` 已经足够覆盖大多数单测需求。

### 10.1 最小测试

文件名：`math_test.go`

```go
package mathx

import "testing"

func Add(a, b int) int {
	return a + b
}

func TestAdd(t *testing.T) {
	got := Add(1, 2)
	if got != 3 {
		t.Fatalf("want 3, got %d", got)
	}
}
```

运行：

```bash
go test ./...
```

### 10.2 Go 很常见的表驱动测试

```go
func TestAdd(t *testing.T) {
	tests := []struct {
		name string
		a    int
		b    int
		want int
	}{
		{name: "positive", a: 1, b: 2, want: 3},
		{name: "zero", a: 0, b: 0, want: 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Add(tt.a, tt.b)
			if got != tt.want {
				t.Fatalf("want %d, got %d", tt.want, got)
			}
		})
	}
}
```

它的好处是：

- 新增用例很容易
- 结构统一
- 很适合边界条件覆盖

### 10.3 测试之外，你还应该熟悉这些命令

```bash
go test ./...
go test -race ./...
go test -cover ./...
go test -bench=. ./...
go fmt ./...
go vet ./...
```

如果项目用了 `golangci-lint`，也建议把它加入日常工作流。

### 10.4 和 Python 测试思维的差异

- Go 默认更轻，少魔法
- monkeypatch 风格不常见
- 更推荐把可替换依赖通过接口和构造函数注入
- 你会更频繁地为了测试而收紧接口边界

---

## 11. Go 项目结构：不要一上来就分层过深

这是 Python 开发者迁移 Go 时的高频误区。

很多人会想当然地写出：

```text
app/
services/
repositories/
utils/
helpers/
common/
base/
```

但 Go 社区通常不太喜欢这种“看起来很整齐、实际上语义模糊”的目录风格。

### 11.1 更推荐的思路

小项目先从扁平开始。

例如一个中小型服务，可以这样起步：

```text
myapp/
├── cmd/
│   └── server/
│       └── main.go
├── internal/
│   ├── user/
│   ├── order/
│   ├── platform/
│   └── httpserver/
├── go.mod
└── go.sum
```

### 11.2 `cmd/` 和 `internal/` 怎么理解

- `cmd/`：程序入口
- `internal/`：项目内部实现，不对外暴露

对于服务型项目，这是非常常见、也非常稳妥的组织方式。

### 11.3 项目结构最佳实践

- 从业务边界拆目录，而不是从“技术层名词”拆目录
- 少建 `utils`、`common`、`helpers`
- 如果一个抽象只用一次，就先不要抽
- 先让目录反映真实业务，再谈优雅架构

---

## 12. 你真的需要记住的 Go 最佳实践

这一节我直接给结论。

### 12.1 工程和代码层面

- 优先标准库，再引入第三方库
- package 保持小而清晰
- 接口定义在消费方
- 组合优先于继承式思维
- 配置优先环境变量
- 错误向上返回，日志留在边界层
- `context.Context` 只传请求边界相关数据，不塞业务参数

### 12.2 Web 和数据库

- 小服务可以从 `net/http` 开始
- 需要成熟工程脚手架时可选 `Gin`
- 数据访问优先 `database/sql` + `sqlc`
- 不要因为“写得少”就默认上重 ORM

### 12.3 测试和质量

- 写 table-driven tests
- 跑 `go test ./...`
- 并发代码跑 `go test -race ./...`
- 先清晰，再优化
- 性能问题要用 benchmark 或 pprof 证明，不要靠感觉

### 12.4 代码风格

- 不要执着链式写法
- 不要执着“把所有重复都抽掉”
- 允许一点点直白重复，换来更清晰的局部可读性

这点对 Python 开发者尤其重要。  
Go 社区对“过度抽象”的容忍度通常比 Python 社区更低。

---

## 13. 给 Python 开发者的常见误区清单

### 误区 1：总想找 Go 版 Django / FastAPI 的唯一标准答案

Go 没有那么强的“单框架中心化”文化。  
很多事情标准库就够。

### 误区 2：总想造通用抽象

Go 更奖励直接、局部清晰的代码，而不是“提前为未来做框架”。

### 误区 3：把接口当作 Java 的接口来写

Go 接口应该小、具体、由使用方定义，而不是一开始就全局声明一个巨大接口。

### 误区 4：把 goroutine 当成性能按钮

并发不是免费午餐。  
共享状态、取消、超时、排队、背压，都会变复杂。

### 误区 5：遇到错误先打日志再继续返回

这样容易重复打日志，还会让错误边界混乱。

### 误区 6：项目一开始就堆满框架和脚手架

Go 的优势之一就是可以从很小的核心开始长起来。  
你不需要在第一天就把所有中间件、ORM、代码生成器、插件系统都接满。

---

## 14. 我给你的学习路线建议

如果你想快速上手，我建议按这个顺序学。

### 第 1 周：工具链和基础语法

- 安装 Go
- 学会 `go mod init`、`go run`、`go test`
- 理解 package、module、导出规则
- 学变量、slice、map、struct、method

### 第 2 周：错误处理、接口、测试

- 把 `if err != nil` 用顺手
- 学会 `errors.Is` / `errors.As`
- 理解接口为什么由消费方定义
- 自己写 5 到 10 个 table-driven tests

### 第 3 周：Web 服务和数据库

- 用 `net/http` 写一个 `/health` 接口
- 再试一次 `chi` 或 `Gin`
- 用 `database/sql` 做一个查询
- 理解 `context.Context`

### 第 4 周：并发和工程实践

- 写 goroutine + `WaitGroup`
- 理解 channel 的边界
- 跑 `go test -race`
- 学基本项目结构和日志/配置组织

---

## 15. 一个建议的技术栈起点

如果你现在就要开始一个真实 Go 后端项目，而你又是 Python 背景，我建议优先这样选：

### 保守、稳妥、长期维护优先

- Web：`net/http` + `chi`
- 数据库：`database/sql` + `sqlc`
- 测试：标准库 `testing`
- 配置：环境变量
- 日志：`zap`
- 质量：`go fmt`、`go vet`、`golangci-lint`

### 团队想更快起步

- Web：`Gin`
- 数据库：`sqlc` 或 `GORM`
- 测试：标准库 `testing`，必要时加 `testify`

### 不建议的起手式

- 一开始就重 ORM
- 一开始就复杂 DDD 分层
- 一开始就服务框架大礼包
- 还没理解标准库就把所有事情交给第三方库

---

## 16. 一个最小实战练习清单

学完本文后，你应该自己动手做下面这些练习。

1. 安装 Go，并确认 `go version`
2. 新建一个 module，并运行 `Hello, Go`
3. 把一个 Python 函数改写成 Go 版本
4. 写一个返回 `(value, error)` 的函数
5. 写一个 `/health` HTTP 接口
6. 写一个 `database/sql` 查询函数
7. 写一个 table-driven test
8. 写一个 goroutine + `WaitGroup` 示例
9. 运行 `go test ./...`
10. 运行 `go test -race ./...`

如果这些你都能独立完成，说明你已经过了“会看不会写”的阶段。

---

## 17. 推荐阅读顺序

如果你准备继续深入，建议按这个顺序读官方资料：

1. [Tutorial: Get started with Go](https://go.dev/doc/tutorial/getting-started)
2. [Tutorial: Create a Go module](https://go.dev/doc/tutorial/create-module)
3. [Effective Go](https://go.dev/doc/effective_go)
4. [Go Code Review Comments](https://go.dev/wiki/CodeReviewComments)
5. [`database/sql` package docs](https://pkg.go.dev/database/sql)
6. [Go workspaces tutorial](https://go.dev/doc/tutorial/workspaces)

如果你想继续理解 Python 工具链对照，可以看：

- [uv 官方文档](https://docs.astral.sh/uv/)
- [uv 安装文档](https://docs.astral.sh/uv/getting-started/installation/)

---

## 18. 最后给你的迁移建议

你最开始的目标，不应该是“把 Go 学全”。  
更现实的目标是：

- 能读懂 Go 项目结构
- 能写出一个小 HTTP 服务
- 能处理错误
- 能写测试
- 能不把 Go 写成 Python

一旦过了这个阶段，后面的提升主要就是工程经验，而不是语法记忆。

如果你愿意，我建议你的下一步不是继续读更多概念，而是立刻做一个很小的项目，比如：

- 待办事项 API
- 文件上传服务
- 一个带 PostgreSQL 的用户列表接口

Go 这门语言非常适合“边写边学”。  
你真正形成手感，通常不是在读教程的时候，而是在你写完第一个可以部署的小服务之后。

---

## 19. 最小实战项目：从零写一个用户服务

如果你已经看完前面的概念，我最建议你马上做一个小项目。  
不要上来就做复杂电商、权限系统或者微服务。先做一个非常小、但闭环完整的 HTTP 服务。

### 19.1 目标

做一个最小用户服务，只有 3 个接口：

- `GET /health`
- `GET /users`
- `POST /users`

第一版先不接数据库，直接用内存存数据。  
这个项目的价值不在于“能不能上线”，而在于你会完整走一遍：

- 建 module
- 写 `struct`
- 写 handler
- 写 JSON 编解码
- 写错误处理
- 写测试

### 19.2 为什么先用标准库

这个练习我建议你先用标准库 `net/http` 做，而不是先上 Gin。

原因很简单：

- 你会更清楚请求和响应到底是什么
- 你会理解 Go Web 服务的基础接口
- 后面再切 Gin、chi、Echo 都更容易

### 19.3 推荐目录

先别搞复杂分层，够用就好：

```text
mini-user-api/
├── main.go
├── user.go
├── store.go
├── handler.go
└── handler_test.go
```

### 19.4 第一步：初始化项目

```bash
mkdir mini-user-api
cd mini-user-api
go mod init example.com/mini-user-api
```

### 19.5 第二步：定义用户模型

`user.go`

```go
package main

type User struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}
```

这一步对应 Python 里你可能会写的 `dataclass` 或 Pydantic model。  
在 Go 里，先把结构体写清楚，再让 JSON 标签显式声明序列化字段。

### 19.6 第三步：写一个内存存储

`store.go`

```go
package main

import "sync"

type UserStore struct {
	mu     sync.Mutex
	nextID int64
	users  []User
}

func NewUserStore() *UserStore {
	return &UserStore{
		nextID: 1,
		users:  []User{},
	}
}

func (s *UserStore) List() []User {
	s.mu.Lock()
	defer s.mu.Unlock()

	out := make([]User, len(s.users))
	copy(out, s.users)
	return out
}

func (s *UserStore) Create(name string) User {
	s.mu.Lock()
	defer s.mu.Unlock()

	user := User{
		ID:   s.nextID,
		Name: name,
	}
	s.nextID++
	s.users = append(s.users, user)
	return user
}
```

这里顺手能学到几件事：

- `struct` 不只是“数据”，也可以承载行为
- 有共享状态时，可以先用 `sync.Mutex`
- 返回 slice 时最好拷贝一份，避免外部直接改内部数据

### 19.7 第四步：写 HTTP handler

`handler.go`

```go
package main

import (
	"encoding/json"
	"net/http"
	"strings"
)

type Handler struct {
	store *UserStore
}

func NewHandler(store *UserStore) Handler {
	return Handler{store: store}
}

func (h Handler) Register(mux *http.ServeMux) {
	mux.HandleFunc("/health", h.handleHealth)
	mux.HandleFunc("/users", h.handleUsers)
}

func (h Handler) handleHealth(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	writeJSON(w, http.StatusOK, map[string]string{
		"status": "ok",
	})
}

func (h Handler) handleUsers(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		writeJSON(w, http.StatusOK, h.store.List())
	case http.MethodPost:
		var req struct {
			Name string `json:"name"`
		}

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "invalid json", http.StatusBadRequest)
			return
		}

		req.Name = strings.TrimSpace(req.Name)
		if req.Name == "" {
			http.Error(w, "name is required", http.StatusBadRequest)
			return
		}

		user := h.store.Create(req.Name)
		writeJSON(w, http.StatusCreated, user)
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}
```

如果你来自 FastAPI，这段会让你觉得有点“手工”。  
这是正常的。Go 的常见哲学就是：少一点魔法，多一点显式控制。

### 19.8 第五步：启动服务

`main.go`

```go
package main

import (
	"log"
	"net/http"
)

func main() {
	store := NewUserStore()
	handler := NewHandler(store)

	mux := http.NewServeMux()
	handler.Register(mux)

	log.Println("listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
```

运行：

```bash
go run .
```

测试接口：

```bash
curl http://localhost:8080/health
curl http://localhost:8080/users
curl -X POST http://localhost:8080/users \
  -H "Content-Type: application/json" \
  -d '{"name":"tudou"}'
```

### 19.9 第六步：写测试

`handler_test.go`

```go
package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealth(t *testing.T) {
	store := NewUserStore()
	handler := NewHandler(store)
	mux := http.NewServeMux()
	handler.Register(mux)

	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	rec := httptest.NewRecorder()

	mux.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("want %d, got %d", http.StatusOK, rec.Code)
	}
}

func TestCreateUser(t *testing.T) {
	store := NewUserStore()
	handler := NewHandler(store)
	mux := http.NewServeMux()
	handler.Register(mux)

	body := bytes.NewBufferString(`{"name":"tudou"}`)
	req := httptest.NewRequest(http.MethodPost, "/users", body)
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	mux.ServeHTTP(rec, req)

	if rec.Code != http.StatusCreated {
		t.Fatalf("want %d, got %d", http.StatusCreated, rec.Code)
	}
}
```

运行：

```bash
go test ./...
```

### 19.10 这个小项目你练到了什么

做完这一轮，你其实已经掌握了 Go 后端最关键的一批基础：

- `struct`
- 方法和构造函数风格
- slice 和内存数据
- `net/http`
- JSON 编解码
- 错误处理
- 测试
- 一点点并发安全

如果你第一次学 Go，我非常建议你先把这个项目独立手敲一遍，再继续上数据库和框架。

---

## 20. 进阶实战：把这个思路落到当前仓库

如果你不是单独练习，而是直接基于当前仓库开发，那最佳路径不是重写一套 `net/http`，而是顺着仓库现有模板往下扩。

### 20.1 当前仓库后端真实基线

当前仓库后端已经有这些事实：

- 入口在 `cmd/server/main.go`
- HTTP 路由由 `internal/httpserver/router.go` 组装
- 当前唯一接口是 `GET /health`
- 健康检查服务位于 `internal/service/health_service.go`
- 响应模型位于 `internal/model/health.go`
- 当前框架是 Gin，不是 `net/http` 原生 mux

也就是说，这个仓库已经替你做完了“最小服务骨架”。

### 20.2 最适合你的第一个练习

不要一上来改一堆架构。  
最好的练习是：**在当前模板上新增一个用户列表接口**。

建议目标：

- 保留现有 `GET /health`
- 新增 `GET /users`
- 返回内存中的固定用户列表

这是个非常好的第一步，因为它只新增一个读接口，不涉及数据库、鉴权和复杂输入校验。

### 20.3 你可以按这个顺序扩展

#### 第一步：新增响应模型

你可以增加一个用户模型，例如：

```go
type UserResponse struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}
```

按当前仓库风格，它更适合放在 `internal/model/`。

#### 第二步：新增用户服务

你可以新增一个简单服务，比如：

```go
type UserService struct{}

func NewUserService() UserService {
	return UserService{}
}

func (s UserService) List() []model.UserResponse {
	return []model.UserResponse{
		{ID: 1, Name: "tudou"},
		{ID: 2, Name: "python-user"},
	}
}
```

这一步的重点是体会当前仓库的模式：

- 业务逻辑在 `internal/service/`
- 路由层只负责 HTTP 和 JSON
- model 层负责输出结构

#### 第三步：在路由里注册接口

当前仓库的 `internal/httpserver/router.go` 已经有健康检查路由。  
你可以仿照同样方式，新增：

```go
router.GET("/users", func(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, userService.List())
})
```

#### 第四步：补测试

最自然的测试点有两个：

- `GET /users` 返回 `200`
- 返回体是 JSON 数组，且至少包含预期字段

如果你还不熟悉测试，可以先只做“状态码 + 响应包含关键字段”的断言。

### 20.4 为什么我建议你先做“只读列表接口”

因为这一步刚好能练到：

- struct
- service 分层
- 路由注册
- JSON 输出
- HTTP 测试

但又不会一下子把复杂度推到：

- 请求校验
- 数据持久化
- 并发写入
- 事务
- 数据库驱动

### 20.5 第二个练习再做什么

当 `GET /users` 完成后，你的下一步建议是二选一：

- 增加 `POST /users`，练输入校验和错误处理
- 接入数据库，练 `database/sql` 和 `context.Context`

如果你是第一次接触 Go，我建议优先顺序是：

1. `GET /users`
2. `POST /users`
3. `database/sql`
4. 再考虑 ORM 或代码生成器

### 20.6 这部分实战的真正目的

不是“把模板改复杂”，而是让你学会一个很重要的迁移能力：

> 看到一个 Go 仓库后，能沿着它现有结构做增量开发，而不是先推翻再重构。

这对 Python 开发者尤其重要，因为很多人一开始会本能地想先搭一套自己熟悉的抽象层。  
而在 Go 项目里，更推荐的路线通常是：先沿用已有结构，等真实重复出现，再抽象。

---

## 参考资料

以下资料已在写作时核实，且优先使用官方或项目主站：

- Go 下载页：[https://go.dev/dl/](https://go.dev/dl/)
- Go release 历史：[https://go.dev/doc/devel/release](https://go.dev/doc/devel/release)
- Go 安装说明：[https://go.dev/doc/install](https://go.dev/doc/install)
- Go 入门教程：[https://go.dev/doc/tutorial/getting-started](https://go.dev/doc/tutorial/getting-started)
- Go module 教程：[https://go.dev/doc/tutorial/create-module](https://go.dev/doc/tutorial/create-module)
- Go workspace 教程：[https://go.dev/doc/tutorial/workspaces](https://go.dev/doc/tutorial/workspaces)
- Effective Go：[https://go.dev/doc/effective_go](https://go.dev/doc/effective_go)
- Go Code Review Comments：[https://go.dev/wiki/CodeReviewComments](https://go.dev/wiki/CodeReviewComments)
- `database/sql` 文档：[https://pkg.go.dev/database/sql](https://pkg.go.dev/database/sql)
- uv 官方文档：[https://docs.astral.sh/uv/](https://docs.astral.sh/uv/)
- Gin：[https://github.com/gin-gonic/gin](https://github.com/gin-gonic/gin)
- chi：[https://github.com/go-chi/chi](https://github.com/go-chi/chi)
- Echo：[https://github.com/labstack/echo](https://github.com/labstack/echo)
- Fiber：[https://github.com/gofiber/fiber](https://github.com/gofiber/fiber)
- sqlc：[https://github.com/sqlc-dev/sqlc](https://github.com/sqlc-dev/sqlc)
- GORM：[https://github.com/go-gorm/gorm](https://github.com/go-gorm/gorm)
