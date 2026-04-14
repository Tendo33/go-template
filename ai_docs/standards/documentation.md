# Documentation Standards

## When to read

- 修改 `README.md`、`ai_docs/`、根入口文件或脚本说明前先读这里。

## Layering rules

- `current/` 只写当前真实实现
- `standards/` 只写工程约束和默认值
- `reference/` 只写共享事实

## Writing rules

- 先写人和 agent 能直接执行的信息，再写背景解释
- 把“当前实现”和“推荐扩展”分开
- 不用 `[待定]` 占位冒充规范
- 示例只引用当前仓库真实存在的文件、命令和导出
- 如果脚本、目录或入口文件已经删除，文档中也要同步移除，不保留假入口

## Linking rules

- 命令统一链接到 [verification.md](../reference/verification.md)
- 目录结构统一链接到 [project-structure.md](../reference/project-structure.md)
- 路径与命名统一链接到 [naming-and-paths.md](../reference/naming-and-paths.md)

## README rules

- README 负责人类入口，不承担完整 AI 协作契约
- README 中的 AI Docs 部分只说明入口和阅读路径
- 详细规则留在 `ai_docs/`

## Root entrypoint rules

- `AGENTS.md` 和 `CLAUDE.md` 都应保持轻量
- 两者都必须只引用当前存在的 `ai_docs/` 文档
