# Template Hardening Implementation Plan

> **For Claude:** REQUIRED SUB-SKILL: Use superpowers:executing-plans to implement this plan task-by-task.

**Goal:** 把当前仓库补成更适合公开复用的 Go 模板，解决校验不稳定、缺少许可证和版本同步不完整这 3 个问题。

**Architecture:** 保持现有前后端骨架不变，只收敛模板基础设施。通过固定 `golangci-lint` 版本提升可复现性，通过根目录 `VERSION` 文件让模板版本有单一事实来源，再同步 README、AI 文档和维护脚本。

**Tech Stack:** Go、PowerShell、POSIX shell、GitHub Actions、Markdown

---

### Task 1: 固定 lint 版本

**Files:**
- Modify: `Makefile`
- Modify: `.github/workflows/ci.yml`
- Modify: `README.md`
- Modify: `.trellis/spec/shared/verification.md`

**Step 1: 先确认当前校验命令仍在使用 `@latest`**

Run: `Select-String -Path README.md,Makefile,.github/workflows/ci.yml,.trellis/spec/shared/verification.md -Pattern '@latest'`
Expected: 能看到 `golangci-lint@latest`

**Step 2: 写最小实现**

- 把 `golangci-lint` 版本固定到单一值
- 同步本地、CI 和文档里的命令

**Step 3: 运行验证**

Run:
- `go run github.com/golangci/golangci-lint/v2/cmd/golangci-lint@v2.8.0 run`
- `go build ./...`

Expected: lint 与 build 都能通过

### Task 2: 补齐公开模板元信息

**Files:**
- Create: `LICENSE`
- Modify: `README.md`

**Step 1: 先确认 `LICENSE` 缺失**

Run: `Test-Path LICENSE`
Expected: `False`

**Step 2: 写最小实现**

- 添加 MIT 许可证
- 在 README 中把许可证现状改成已提供

**Step 3: 运行验证**

Run: `Test-Path LICENSE`
Expected: `True`

### Task 3: 扩展版本同步脚本

**Files:**
- Create: `VERSION`
- Modify: `scripts/update_version.sh`
- Modify: `scripts/update_version.ps1`
- Modify: `README.md`
- Modify: `.trellis/spec/shared/index.md`
- Modify: `.trellis/spec/shared/verification.md`

**Step 1: 先确认当前 dry-run 只会更新前端版本**

Run: `.\scripts\update_version.ps1 -Version 0.2.0 -DryRun`
Expected: 输出里只有 `frontend/package.json`

**Step 2: 写最小实现**

- 新增根目录 `VERSION`
- 两个更新脚本同时同步 `VERSION` 和 `frontend/package.json`
- 文档同步描述新的版本单一事实来源

**Step 3: 运行验证**

Run:
- `.\scripts\update_version.ps1 -Version 0.2.0 -DryRun`
- `.\scripts\update_version.ps1 -Version 0.1.0 -DryRun`

Expected:
- 目标版本不同的时候，会同时显示两个文件的变更
- 目标版本相同的时候，提示版本已经是当前值
