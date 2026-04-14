# Go Template

一个面向开发者的 Go 全栈项目模板：Gin 后端、React + Vite 前端、基础工程化、模板维护脚本和 AI 协作文档都已经准备好，适合直接起新仓库，而不是从空目录重新拼装。

## 为什么用这个模板

这个仓库解决的不是“怎么再造一个 starter”，而是“怎么把项目第 1 天到第 30 天最容易重复手工做的事情先准备好”。

- 后端基线已经就位：Go、Gin、`golangci-lint`、结构化日志、环境变量配置、健康检查接口
- 前端 starter 已可直接运行：React 19 + TypeScript 6 + Vite 8 + Vitest
- 常见工程动作已统一：测试、构建、lint、开发启动、Docker、GitHub Actions
- 模板维护入口已预留：项目重命名、版本号同步
- `ai_docs/` 不是装饰目录，而是给 AI 助手和协作者共用的项目事实源

如果你想从一个尽量直接、尽量省心、又不会一上来就引入过多框架约束的 Go 全栈仓库起步，这个模板就是为这种场景准备的。默认开发环境按 macOS / Linux 组织，Windows 只保留兼容路径。

## Quick Start

### 1. 克隆仓库

```bash
git clone https://github.com/Tendo33/go-template.git
cd go-template
```

### 2. 下载 Go 依赖

```bash
go mod download
```

### 3. 安装前端依赖

```bash
pnpm --prefix frontend install
```

### 4. 初始化环境变量

Bash:

```bash
cp .env.example .env
cp frontend/.env.example frontend/.env.local
```

PowerShell:

```powershell
Copy-Item .env.example .env
Copy-Item frontend/.env.example frontend/.env.local
```

当前默认环境变量包括：

```env
APP_ENV=development
PORT=8080
LOG_LEVEL=INFO
SERVICE_NAME=go-template
VITE_API_BASE_URL=http://localhost:8080
```

### 5. 启动后端

```bash
go run ./cmd/server
```

后端默认监听 `http://localhost:8080`，健康检查接口为 `GET /health`。

### 6. 启动前端

```bash
pnpm --prefix frontend dev
```

前端默认运行在 Vite 开发服务器上。  
如果没有配置 `frontend/.env.local`，当前 `vite.config.ts` 也会把 `/health` 代理到本地 `8080` 后端。

### 7. 运行验证

```bash
go test ./...
go build ./...
go run github.com/golangci/golangci-lint/v2/cmd/golangci-lint@v2.8.0 run
pnpm --prefix frontend test --run
pnpm --prefix frontend build
```

完整验证说明见 [ai_docs/reference/verification.md](ai_docs/reference/verification.md)。

## Use This Template

建议在项目刚创建时就把模块名、服务名、包名和仓库名替换掉，不要等文档和文件变多了再统一修改。

### 1. 基于模板创建新仓库

GitHub 上可以直接点 `Use this template`。如果你是在本地复制：

```bash
git clone https://github.com/Tendo33/go-template.git my-new-project
cd my-new-project
```

### 2. 第一时间改项目名

Bash:

```bash
sh ./scripts/rename_project.sh \
  --project-name my-new-project \
  --module-name github.com/acme/my-new-project \
  --frontend-package-name @acme/my-new-project-frontend \
  --dry-run

sh ./scripts/rename_project.sh \
  --project-name my-new-project \
  --module-name github.com/acme/my-new-project \
  --frontend-package-name @acme/my-new-project-frontend
```

PowerShell:

```powershell
.\scripts\rename_project.ps1 -ProjectName my-new-project -ModuleName github.com/acme/my-new-project -FrontendPackageName @acme/my-new-project-frontend -DryRun
.\scripts\rename_project.ps1 -ProjectName my-new-project -ModuleName github.com/acme/my-new-project -FrontendPackageName @acme/my-new-project-frontend
```

脚本会同步更新：

- `go.mod` 中的模块名
- 默认服务名
- `frontend/package.json` 的 `name`
- `README.md`、`ai_docs/`、前后端代码里的模板名引用

### 3. 更新版本号

Bash:

```bash
sh ./scripts/update_version.sh --version 0.2.0 --dry-run
sh ./scripts/update_version.sh --version 0.2.0
```

PowerShell:

```powershell
.\scripts\update_version.ps1 -Version 0.2.0 -DryRun
.\scripts\update_version.ps1 -Version 0.2.0
```

当前版本同步脚本会同时更新根目录 `VERSION` 和 `frontend/package.json`，用 `VERSION` 作为模板版本的单一事实来源。

## Project Structure

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
├── ai_docs/                      # AI / 协作文档与工程规则
├── .github/workflows/            # CI workflow
├── docs/plans/                   # 设计与实施计划
├── Makefile
├── go.mod
└── README.md
```

更细的结构说明见 [ai_docs/reference/project-structure.md](ai_docs/reference/project-structure.md)。

## Backend

后端默认栈：

- Go 1.26
- Gin
- `slog`
- `golangci-lint`

当前后端已经内置这些可直接复用的入口：

- 配置：`internal/config`
- 日志：`internal/observability`
- HTTP 服务：`internal/httpserver`
- 健康检查服务：`internal/service`
- 响应模型：`internal/model`

最小示例：

```bash
go run ./cmd/server
curl http://localhost:8080/health
```

当前不会预装数据库、ORM、鉴权、Swagger 或 gRPC；这些都留给真实业务项目按需添加。

## Frontend

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

如果项目继续长大，建议按 `ai_docs/reference/project-structure.md` 和 `ai_docs/standards/frontend.md` 里的约定逐步扩展。

## Scripts

当前模板维护脚本位于 `scripts/`：

- `rename_project.sh`：默认的 macOS / Linux 项目重命名脚本
- `update_version.sh`：默认的 macOS / Linux 版本同步脚本
- `rename_project.ps1`：PowerShell 兼容版本的重命名脚本
- `update_version.ps1`：PowerShell 兼容版本的版本同步脚本

脚本现状与使用场景见 [ai_docs/current/scripts.md](ai_docs/current/scripts.md)。

## Verification

验证命令的唯一来源是 [ai_docs/reference/verification.md](ai_docs/reference/verification.md)。

常用本地验证：

```bash
go test ./...
go build ./...
go run github.com/golangci/golangci-lint/v2/cmd/golangci-lint@v2.8.0 run
pnpm --prefix frontend test --run
pnpm --prefix frontend build
```

默认开发环境按 macOS / Linux 组织，优先使用 `make`、`.sh` 脚本和 Bash 命令。  
Windows 下如果没有 `make`，请直接使用文档中的原始命令；`Makefile` 和 `.sh` 脚本都不是唯一入口。

## AI Docs

阅读入口：

1. [ai_docs/START_HERE.md](ai_docs/START_HERE.md)
2. [ai_docs/INDEX.md](ai_docs/INDEX.md)
3. [ai_docs/reference/verification.md](ai_docs/reference/verification.md)

`ai_docs/` 只维护当前真实实现、默认工程约束和共享参考，不把未来规划写成现状。

## Release

当前仓库只内置 CI，没有自动 release workflow。  
发版时建议：

1. 运行完整验证
2. 用 `scripts/update_version.sh` 或 `scripts/update_version.ps1` 更新版本号，并确认 `VERSION` 与 `frontend/package.json` 已同步
3. 手工创建并推送 tag
4. 按项目需要在 GitHub 上补 release 说明

当前发版现状见 [ai_docs/current/release.md](ai_docs/current/release.md)。

## License

当前仓库提供 MIT 许可证，见 [LICENSE](LICENSE)。
