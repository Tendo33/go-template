# Naming And Paths Reference

## Purpose

本文件集中维护命名、导入路径、文档路径和根入口文件规则。其他文档需要引用这些规则时，只链接这里。

## Package and import rules

- Go 代码优先从模块路径导入，不从相对路径拼装导入
- 当前模块路径定义在 `go.mod`
- `internal/` 下的包只供仓库内部使用，不把它当作外部 SDK
- 文档和示例只引用当前真实存在的包和文件

## Path naming rules

- Go 文件和目录使用小写命名，按现有目录结构扩展
- 前端文件与目录遵循当前 starter 约定
- `ai_docs/` 使用 `current/`、`standards/`、`reference/`
- 禁止从 `frontend/dist/`、`node_modules/`、`.vite/` 等生成目录导入或引用源码

## AI docs layering rules

- `current/` 只写当前真实实现
- `standards/` 只写约束、默认值和工作准则
- `reference/` 只写共享事实

## Root entrypoint policy

- 根入口文件为 `AGENTS.md` 和 `CLAUDE.md`
- 根入口文件里的链接必须指向真实存在的文档
