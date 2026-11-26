# go-template

Go + Gin 后端加 React/Vite 前端的全栈脚手架模板，同样会派生到新项目。`Makefile` 是唯一命令入口。

## 结构与边界

- `cmd/server/`：进程入口。
- `internal/`：业务代码，保持内部边界，不要为图方便把内部包暴露成公开 API。
- 保留 request ID 透传、请求级 zap 日志、env 配置的既有写法。
- `frontend/`：React 19 + TypeScript + Vite + Vitest，包管理用 pnpm。
- `VERSION`、`Dockerfile`、`.dockerignore`、`Makefile` 与模板脚本要互相对得上：
  `scripts/rename_project.sh` 派生时改项目名，`scripts/update_version.sh` 与 `VERSION` 联动。

## 命令

```bash
make dev
make test
make lint
make fmt
make build
make frontend-test
make frontend-build
make fullstack     # test + build + lint + frontend-test + frontend-build
make docker-build
make ci-check      # fullstack + docker-build，提交前跑这个
```

## 改模板时的规矩

- 改项目名、版本或构建产物路径，要同时检查 `scripts/` 下两个脚本与 `Makefile`。
- 不引入前端 UI 框架，保持模板轻量。
