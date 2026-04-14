# Current Frontend

## When to read

- 要改 frontend 代码前先读这里。
- 需要了解当前 frontend starter 的真实范围和可复用入口时读这里。

## Current truth

当前 `frontend/` 是一个很轻的单页 starter，不是完整业务前端。

已经落地的能力：

- React + TypeScript + Vite 前端基础运行链路
- 一个最小主页入口
- 一个调用后端健康检查的 API helper
- Vitest + Testing Library 测试基线
- Vite 开发环境下的 `/health` 本地代理

## Current UI scope

当前页面主要用于证明前端 starter 可工作，而不是提供完整产品视觉系统。页面包括：

- 一个标题区域
- 一段模板说明文案
- 一个后端状态展示区域

## Current frontend structure

- `frontend/src/app`：当前页面入口与测试
- `frontend/src/lib`：当前 API helper
- `frontend/src/styles`：全局样式
- `frontend/src/test`：测试初始化

## What is not here yet

- 当前还没有复杂路由
- 当前还没有状态管理方案
- 当前还没有组件库或设计系统
- 当前还没有多页面结构

## Shared references

- frontend 约束见 [frontend.md](../standards/frontend.md)
- 设计基线见 [design-system.md](../standards/design-system.md)
- 项目结构见 [project-structure.md](../reference/project-structure.md)
- 验证命令见 [verification.md](../reference/verification.md)
