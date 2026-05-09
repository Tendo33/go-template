# Go Template

一个偏实用、可直接起项目的 Go 全栈模板：Gin 后端、React + Vite 前端、基础工程化、模板维护脚本和 AI 协作文档都已经准备好，目标不是“再造一个 starter”，而是把项目最开始那段最容易重复搭建的部分先收拾好。

## 适合什么场景

- 想从一个干净、直接、不预装过多框架约束的 Go 仓库起步
- 希望后端、前端、CI、Docker、lint、测试和文档入口一开始就统一
- 希望模板在本地开发体验和生产运行方式之间有清晰边界
- 希望仓库对 AI 助手和人类协作者都足够友好

## 你会直接得到什么

- 后端基线：Go、Gin、`zap`、环境变量配置、健康检查接口、结构化日志
- 前端基线：React 19、TypeScript 6、Vite 8、Vitest
- 工程化基线：`golangci-lint`、GitHub Actions、Dockerfile、`.dockerignore`
- 模板维护入口：项目重命名、版本号同步
- 协作文档基线：`.trellis/spec/` 作为当前实现和工程约定的事实源

## Quick Start

### 1. 克隆仓库

```bash
git clone https://github.com/Tendo33/go-template.git
cd go-template
```

### 2. 安装依赖

```bash
go mod download
pnpm --prefix frontend install
```

### 3. 初始化环境变量

```bash
cp .env.example .env
cp frontend/.env.example frontend/.env.local
```

默认环境变量：

```env
APP_ENV=development
PORT=8080
LOG_LEVEL=INFO
SERVICE_NAME=go-template
VITE_API_BASE_URL=http://localhost:8080
```

补充说明：

- `APP_ENV=development|dev|local` 时使用更易读的彩色 console 日志
- 其它环境默认输出 JSON，方便容器和日志平台采集
- Gin 中间件会自动透传或生成 `X-Request-ID`
- 访问日志与业务日志会共享同一个 `request_id`

### 4. 启动后端

```bash
go run ./cmd/server
```

默认监听 `http://localhost:8080`，健康检查接口为 `GET /health`。

### 5. 启动前端

```bash
pnpm --prefix frontend dev
```

如果没有配置 `frontend/.env.local`，当前 `vite.config.ts` 也会把 `/health` 代理到本地 `8080` 后端。

### 6. 最小联调检查

```bash
curl http://localhost:8080/health
```

预期响应：

```json
{
  "status": "ok",
  "service": "go-template"
}
```

## 用模板起新仓库

建议在项目刚创建时就完成重命名和版本号同步，不要等文档、包名和模块引用散开后再回头统一改。

### 1. 创建新仓库

GitHub 上可以直接点 `Use this template`。如果你是在本地复制：

```bash
git clone https://github.com/Tendo33/go-template.git my-new-project
cd my-new-project
```

### 2. 先跑一次重命名 dry-run

```bash
sh ./scripts/rename_project.sh \
  --project-name my-new-project \
  --module-name github.com/acme/my-new-project \
  --frontend-package-name @acme/my-new-project-frontend \
  --dry-run
```

确认无误后正式执行：

```bash
sh ./scripts/rename_project.sh \
  --project-name my-new-project \
  --module-name github.com/acme/my-new-project \
  --frontend-package-name @acme/my-new-project-frontend
```

它会同步更新：

- `go.mod` 里的模块名
- 默认服务名
- `frontend/package.json` 的 `name`
- `README.md`、`.trellis/spec/`、前后端代码里的模板名称引用

### 3. 更新版本号

```bash
sh ./scripts/update_version.sh --version 0.2.0 --dry-run
sh ./scripts/update_version.sh --version 0.2.0
```

当前版本同步脚本会同时更新根目录 `VERSION` 和 `frontend/package.json`，并以 `VERSION` 作为模板版本的单一事实来源。

### 4. 初始化后立刻做一次完整校验

```bash
make ci-check
```

如果只是局部改动，则按 [.trellis/spec/shared/verification.md](.trellis/spec/shared/verification.md) 里的任务类型选择对应检查。

## 日常开发常用命令

### 后端

```bash
go test ./...
go build ./...
go run github.com/golangci/golangci-lint/v2/cmd/golangci-lint@v2.11.4 run
```

### 前端

```bash
pnpm --prefix frontend test --run
pnpm --prefix frontend build
```

### Makefile 快捷入口

```bash
make fullstack
make docs-check
make docker-build
make ci-check
```

### 容器构建

```bash
docker build -t go-template:local .
```

验证命令的唯一详细来源是 [.trellis/spec/shared/verification.md](.trellis/spec/shared/verification.md)。

## 后端默认约定

后端固定基线：

- Go 1.26
- Gin
- `zap`
- `golangci-lint`

日志与请求链路默认约定：

- 开发环境使用彩色 console encoder，优先可读性
- 非开发环境使用 JSON，优先结构化采集
- HTTP 请求默认透传或生成 `X-Request-ID`
- 请求级 logger 通过 `context.Context` 传递
- handler/service 统一通过 `observability.FromContext(...)` 读取请求级 logger
- `trace_id` / `span_id` 已预留 context 字段位，后续可在不改业务调用方式的前提下接入 OTel
- 服务启动前会执行配置校验，当前最小覆盖 `PORT`、`SERVICE_NAME`、`LOG_LEVEL`

当前已经内置的后端入口：

- 配置：`internal/config`
- 日志：`internal/observability`
- HTTP 服务：`internal/httpserver`
- 健康检查服务：`internal/service`
- 响应模型：`internal/model`

当前不会预装数据库、ORM、鉴权、Swagger 或 gRPC；这些留给真实业务项目按需添加。

## 前端默认约定

前端固定基线：

- `pnpm`
- React 19
- TypeScript 6
- Vite 8
- Vitest + Testing Library

当前 starter 刻意保持为一个很小的单页应用，只负责：

- 提供一个可运行的前端入口
- 演示如何调用后端 `/health`
- 提供一套最小测试与构建配置

## 项目结构

```text
go-template/
├── cmd/server/                   # Go 服务入口
├── internal/
│   ├── config/                   # 环境变量配置
│   ├── httpserver/               # Gin 路由、middleware、server 组装
│   ├── model/                    # HTTP 响应结构
│   ├── observability/            # 结构化日志
│   └── service/                  # 当前业务服务
├── frontend/                     # React + Vite starter
│   ├── src/app/                  # 当前页面入口与测试
│   ├── src/lib/                  # API helper
│   ├── src/styles/               # 全局样式
│   └── src/test/                 # 前端测试初始化
├── scripts/                      # 模板维护脚本
├── .trellis/spec/                # AI / 协作文档与工程规则
├── .github/workflows/            # CI workflow
├── docs/plans/                   # 设计与实施计划
├── Makefile
├── go.mod
└── README.md
```

更细的结构说明见 [.trellis/spec/shared/index.md](.trellis/spec/shared/index.md)。

## 脚本与文档入口

模板维护脚本位于 `scripts/`：

- `rename_project.sh`
- `update_version.sh`

当前仓库默认只提供 `.sh` 模板维护脚本。

协作文档入口：

1. [.trellis/spec/README.md](.trellis/spec/README.md)
2. [.trellis/spec/shared/index.md](.trellis/spec/shared/index.md)
3. [.trellis/spec/shared/verification.md](.trellis/spec/shared/verification.md)

相关说明：

- 脚本现状见 [.trellis/spec/shared/index.md](.trellis/spec/shared/index.md)
- 后端现状见 [.trellis/spec/backend/index.md](.trellis/spec/backend/index.md)
- 发版与验证见 [.trellis/spec/shared/verification.md](.trellis/spec/shared/verification.md)

## Release

当前仓库只内置 CI，没有自动 release workflow。发版时建议：

1. 运行完整验证
2. 用 `scripts/update_version.sh` 更新版本号，并确认 `VERSION` 与 `frontend/package.json` 已同步
3. 手工创建并推送 tag
4. 按项目需要在 GitHub 上补 release 说明

## License

当前仓库提供 MIT 许可证，见 [LICENSE](LICENSE)。
