# Go Fullstack Template Implementation Plan

> **For Claude:** REQUIRED SUB-SKILL: Use superpowers:executing-plans to implement this plan task-by-task.

**Goal:** 构建一个可复用的 Go 全栈模板仓库，默认提供 `Gin` 后端、`React + Vite` 前端、基础工程化能力、维护脚本与 AI 协作文档。

**Architecture:** 后端采用 `cmd + internal` 的常见 Go 布局，前端保留独立 `frontend/` 工作区，根目录统一承载 CI、脚本、文档和模板元信息。首版只提供“可直接起项目”的基础设施，不预置数据库、MQ、认证系统等重依赖。

**Tech Stack:** Go、Gin、React、TypeScript、Vite、Vitest、golangci-lint、air、Makefile、GitHub Actions、Docker

---

### Task 1: 写入模板计划并建立目录骨架

**Files:**
- Create: `docs/plans/2026-04-14-go-fullstack-template.md`
- Create: `cmd/server/main.go`
- Create: `internal/config/config.go`
- Create: `internal/observability/logger.go`
- Create: `internal/httpserver/router.go`
- Create: `internal/httpserver/server.go`
- Create: `internal/service/health_service.go`
- Create: `internal/model/health.go`
- Create: `internal/httpserver/router_test.go`

**Step 1: Write the failing test**

在 `internal/httpserver/router_test.go` 中编写 `GET /health` 返回 `200` 与 JSON body 的测试。

**Step 2: Run test to verify it fails**

Run: `go test ./...`
Expected: 因路由和服务尚未实现而失败。

**Step 3: Write minimal implementation**

实现 `config`、`logger`、`health service`、`router` 和 `main`，让 `/health` 返回最小健康检查响应。

**Step 4: Run test to verify it passes**

Run: `go test ./...`
Expected: 后端测试通过。

**Step 5: Commit**

```bash
git add .
git commit -m "feat: add go backend template skeleton"
```

### Task 2: 补齐后端工程化基础

**Files:**
- Create: `Makefile`
- Create: `.gitignore`
- Create: `.golangci.yml`
- Create: `.air.toml`
- Create: `Dockerfile`
- Create: `.env.example`
- Create: `.github/workflows/ci.yml`

**Step 1: Write the failing test**

没有直接的单元测试，先用命令验证法约束开发体验。

**Step 2: Run test to verify it fails**

Run: `make test`
Expected: `Makefile` 尚不存在，命令失败。

**Step 3: Write minimal implementation**

补齐 lint、test、build、dev 所需配置，保证模板具备基础工程化能力。

**Step 4: Run test to verify it passes**

Run:
- `make test`
- `make build`

Expected: 均成功。

**Step 5: Commit**

```bash
git add .
git commit -m "chore: add backend tooling and ci"
```

### Task 3: 接入前端 starter

**Files:**
- Create: `frontend/package.json`
- Create: `frontend/tsconfig.json`
- Create: `frontend/tsconfig.node.json`
- Create: `frontend/vite.config.ts`
- Create: `frontend/index.html`
- Create: `frontend/src/main.tsx`
- Create: `frontend/src/app/App.tsx`
- Create: `frontend/src/lib/api.ts`
- Create: `frontend/src/styles/globals.css`
- Create: `frontend/src/test/setup.ts`
- Create: `frontend/src/app/App.test.tsx`

**Step 1: Write the failing test**

在 `frontend/src/app/App.test.tsx` 中编写页面渲染与健康状态展示测试。

**Step 2: Run test to verify it fails**

Run: `pnpm --prefix frontend test --run`
Expected: 前端文件尚未创建而失败。

**Step 3: Write minimal implementation**

实现一个最小 React/Vite 页面，调用后端 `/health` 并展示响应。

**Step 4: Run test to verify it passes**

Run:
- `pnpm --prefix frontend install`
- `pnpm --prefix frontend test --run`
- `pnpm --prefix frontend build`

Expected: 前端测试与构建成功。

**Step 5: Commit**

```bash
git add .
git commit -m "feat: add frontend starter"
```

### Task 4: 文档与 AI 协作入口

**Files:**
- Modify: `README.md`
- Create: `ai_docs/START_HERE.md`
- Create: `ai_docs/INDEX.md`
- Create: `ai_docs/current/backend.md`
- Create: `ai_docs/current/frontend.md`
- Create: `ai_docs/current/scripts.md`
- Create: `ai_docs/reference/project-structure.md`
- Create: `ai_docs/reference/verification.md`

**Step 1: Write the failing test**

使用文档可执行性验证：按 README 执行命令前先确认缺少说明。

**Step 2: Run test to verify it fails**

Run: 人工检查 README
Expected: 当前 README 无法指导完整初始化与验证。

**Step 3: Write minimal implementation**

补齐 README 与最小版 `ai_docs`，明确项目结构、启动方式与验证命令。

**Step 4: Run test to verify it passes**

Run:
- `go test ./...`
- `pnpm --prefix frontend test --run`
- `pnpm --prefix frontend build`

Expected: 文档中的关键命令全部成立。

**Step 5: Commit**

```bash
git add .
git commit -m "docs: add template usage and ai docs"
```
