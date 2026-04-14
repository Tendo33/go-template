# Frontend Design System

## When to read

- 写任何 UI 代码前先读这里。
- 如果要改版、重构 UI 或扩展视觉语言，先改这里，再改代码。

## Design goal

当前前端不是业务产品，而是模板仓库的演示壳层。设计目标是：

- 让新仓库一 clone 下来就有一个干净、可运行的前端入口
- 展示最基础的 API 调用、状态展示和页面骨架
- 保持克制，不抢未来真实产品设计的空间

## Current UI scope

当前只实现了一个单页 starter，位于 `frontend/src/app/App.tsx`。页面由以下部分组成：

- 标题
- 模板说明文案
- 后端状态展示

## Visual direction

- 中性
- 清爽
- 工具型
- 不做过度品牌化

关键词：

- neutral
- minimal
- technical
- calm

## Styling rules

- 当前没有正式设计系统或组件库
- 当前样式只服务于 starter 演示，不把临时样式写成通用设计规范
- 如果后续引入真实设计系统，应先更新本文件，再扩展组件实现

## Motion and responsive rules

- 当前只用轻量布局和状态切换
- 小屏优先
- 优先保证可读性和可运行，而不是做复杂动效

## Accessibility rules

- 所有可交互元素要有可见 focus 样式
- 颜色不是唯一状态表达方式
- 文案和布局在小屏下仍应可读

## Shared references

- 当前 frontend 现状见 [frontend.md](../current/frontend.md)
- 其它 frontend 约束见 [frontend.md](frontend.md)
