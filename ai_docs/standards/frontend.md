# Frontend Standards

## When to read

- frontend 任务开始前必读。
- 改 UI 前先读这里，再读设计系统。

## Fixed stack

- `pnpm`
- React 19
- TypeScript (`strict`)
- Vite 8
- Vitest + Testing Library + jsdom

除非用户明确要求，否则不要改这套栈。

## Working rules

- 当前 starter 还是单页壳层，不要把未来结构写成当前实现
- 当前共享能力优先沿用 `app/lib/styles/test`
- 当第二个以上稳定通用工具出现时，再继续扩展 `frontend/src/lib`
- 默认先保证可运行、可测试、可读，再考虑视觉包装

## TypeScript and UI rules

- 保持 `strict`
- 默认不使用 `any`
- 组件和页面改动要考虑响应式与可访问性
- 如果改视觉或交互基线，先更新设计系统文档

## Testing rules

- 组件测试用 Vitest + Testing Library
- 新增页面行为时同步补测试
- 当前 starter 测试很轻，优先证明行为而不是堆叠快照

## Shared references

- 当前 frontend 现状见 [frontend.md](../current/frontend.md)
- 项目结构见 [project-structure.md](../reference/project-structure.md)
- 验证命令见 [verification.md](../reference/verification.md)
- 设计系统基线见 [design-system.md](design-system.md)
