# Current Release

## When to read

- 准备发版、更新版本号或调整 CI / tag 流程时先读这里。

## Current truth

当前仓库只有 `.github/workflows/ci.yml`，没有单独的 release workflow。
当前发布仍然是手工流程：

- 更新版本号
- 确认根目录 `VERSION` 与 `frontend/package.json` 已同步
- 运行本地验证
- 创建并推送 tag
- 视项目需要在 GitHub 上补 release 说明

当前 CI 会校验四类门禁：

- backend 测试、构建与 lint
- frontend 安装、测试与构建
- 文档链接检查
- Docker 镜像构建

默认开发环境按 macOS / Linux 组织，因此版本更新脚本当前只提供 `.sh` 入口。

## Release notes

- 当前仓库还没有独立的 release notes 生成脚本
- 如果后续新增自动 release 或生成脚本，需要同步更新本文件和 `current/scripts.md`

## Shared references

- 当前脚本现状见 [scripts.md](scripts.md)
- 验证命令见 [verification.md](../reference/verification.md)
