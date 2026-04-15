# Current Architecture

## When to read

- 想快速理解仓库现在已经落地了什么时先读这里。
- 想判断某项描述属于“当前实现”还是“后续扩展方向”时先读这里。

## Current truth

这个仓库当前是一个面向 AI 协作的 Go 全栈模板，而不是已经内建完整业务系统的脚手架。

它已经提供：

- Gin 后端基础设施：配置、日志、HTTP 服务装配与健康检查接口
- React + Vite 前端 starter
- 一组模板维护脚本
- Dockerfile 与 `.dockerignore` 组成的基础容器化入口
- 根目录 AI 入口文件：`AGENTS.md` 和 `CLAUDE.md`
- 一套精简后的 `ai_docs/` 文档系统

它当前没有内建：

- 数据库 schema、迁移或持久化层
- 完整的 service / repository / domain 分层
- 鉴权、Swagger、gRPC 或消息队列
- 多页面前端应用或复杂前端状态架构
- 自动 release workflow

## Subsystems

- backend 的真实实现见 [backend.md](backend.md)
- frontend 的真实实现见 [frontend.md](frontend.md)
- scripts 的真实实现见 [scripts.md](scripts.md)
- release 的真实实现见 [release.md](release.md)

## Shared references

- 项目结构见 [project-structure.md](../reference/project-structure.md)
- 命名与路径规则见 [naming-and-paths.md](../reference/naming-and-paths.md)
- 验证命令见 [verification.md](../reference/verification.md)
