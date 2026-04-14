# Project Structure Reference

## Purpose

本文件集中说明当前仓库目录结构。其他文档需要引用目录、位置或扩展点时，只链接这里。

## Current repository shape

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
│   ├── src/app/                  # 当前页面入口和测试
│   ├── src/lib/                  # API helper
│   ├── src/styles/               # 全局样式
│   └── src/test/                 # 前端测试初始化
├── scripts/                      # 模板维护脚本
├── ai_docs/                      # Shared AI and engineering docs
│   ├── current/                  # Current implementation facts
│   ├── standards/                # Working rules and defaults
│   └── reference/                # Shared commands and structure references
├── .github/workflows/            # CI workflow
├── docs/plans/                   # 设计与实施计划
├── AGENTS.md                     # Cross-tool root entrypoint
├── CLAUDE.md                     # Claude root entrypoint
├── README.md
├── Makefile
└── go.mod
```

## Current frontend starter

- 当前页面入口在 `frontend/src/app/App.tsx`
- 当前前端通用工具位于 `frontend/src/lib/api.ts`
- 当前只提供一个很小的页面壳层和一个健康检查请求示例

## Recommended future expansion

如果项目继续长大，优先沿着现有目录扩展，而不是平铺新增顶层目录：

- 在 `frontend/src/features/*` 下按业务拆分前端模块
- 在 `frontend/src/lib` 下沉淀稳定的前端通用工具
- backend 再按真实需要引入 `repository`、`domain` 或其它边界

## Usage rule

- 描述当前目录现状时，只引用 `Current repository shape`
- 描述预留扩展方向时，只引用 `Recommended future expansion`
