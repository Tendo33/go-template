# Current Scripts

## When to read

- 想了解仓库维护脚本和模板初始化能力时先读这里。
- 想确认脚本当前行为时读这里。

## Current truth

仓库当前提供这些维护脚本：

- `rename_project.sh`：默认的 macOS / Linux 项目重命名脚本
- `update_version.sh`：默认的 macOS / Linux 版本同步脚本

仓库当前还提供这些常用 `Makefile` 校验入口：

- `make fullstack`
- `make docs-check`
- `make docker-build`
- `make ci-check`

## When to use which script

### Rename project and template references

- 新建仓库后第一时间改模板名称时优先用 `rename_project.sh`
- 它会更新 `go.mod`、默认服务名、前端包名以及文档和源码中的模板名称引用
- 当前维护脚本只提供 `.sh` 入口，默认面向 macOS / Linux 或兼容的 shell 环境
- 两者都使用参数式调用，不提供交互式界面
- Unix 环境默认用 `sh ./scripts/rename_project.sh ...` 调用，不依赖文件可执行位

### Update version

- 准备发版或同步版本号时优先用 `update_version.sh`
- 当前脚本会同步更新根目录 `VERSION` 与 `frontend/package.json`
- `VERSION` 是模板版本的单一事实来源，前端包版本跟随它同步
- Unix 环境默认用 `sh ./scripts/update_version.sh ...` 调用

### Verification shortcuts

- 想跑代码级完整验证时优先用 `make fullstack`
- 改了文档时额外跑 `make docs-check`
- 改了 Dockerfile 或镜像上下文时额外跑 `make docker-build`
- 想本地尽量贴近 CI 时直接跑 `make ci-check`

## Verification

- 脚本、README 和 `ai_docs` 的相关验证命令见 [verification.md](../reference/verification.md)

## Shared references

- 当前发布流程见 [release.md](release.md)
- 验证命令见 [verification.md](../reference/verification.md)
