# Engineering Standards

## When to read

- 开始任何实现或文档任务前先读这里。
- 需要确认最小改动原则和验证原则时回到这里。

## Rules and defaults

- 先确认目标、作用范围和风险点
- 改动前先核对真实实现，不靠记忆写文档或规则
- 做最小、可审查的修改
- 行为、结构、脚本、公共导出或入口文档变了就更新文档
- 不在根入口文件中重复维护 `ai_docs/` 的完整正文

## Quality bar

- backend 代码保持简洁、可读、直接
- frontend 改动遵守当前 starter 和设计系统约束
- 错误处理显式，不走静默失败路径
- 文档描述当前实现，不把未来规划写成当前状态
- 不把模板仓库做成过度框架化的脚手架
- 默认优先保证 macOS / Linux 开发体验，同时保留 Windows 的最小兼容路径

## Shared references

- 验证命令见 [verification.md](../reference/verification.md)
- 项目结构见 [project-structure.md](../reference/project-structure.md)
- 具体任务路径请从 `ai_docs/INDEX.md` 进入
