# Verification Reference

## Purpose

本文档是仓库验证命令的唯一详细事实源。其他 AI 文档和根入口文件只链接这里，不重复抄写完整命令。

## Backend

```bash
go test ./...
go build ./...
go run github.com/golangci/golangci-lint/v2/cmd/golangci-lint@v2.8.0 run
```

## Frontend

```bash
pnpm --prefix frontend test --run
pnpm --prefix frontend build
```

## Full stack

```bash
go test ./...
go build ./...
go run github.com/golangci/golangci-lint/v2/cmd/golangci-lint@v2.8.0 run
pnpm --prefix frontend test --run
pnpm --prefix frontend build
```

`Full stack` 表示本地完整验证，包含前端测试与构建。

## CI gate

GitHub Actions 当前以 `.github/workflows/ci.yml` 为准，等价门禁检查为：

```bash
go test ./...
go build ./...
go run github.com/golangci/golangci-lint/v2/cmd/golangci-lint@v2.8.0 run
pnpm --prefix frontend install --frozen-lockfile
pnpm --prefix frontend test --run
pnpm --prefix frontend build
```

## Docs and links

默认环境可用 Bash 时，优先用：

```bash
node - <<'JS'
const fs = require("fs");
const path = require("path");

const docs = [
  ...walk(path.join(process.cwd(), "ai_docs")),
  path.join(process.cwd(), "README.md"),
  path.join(process.cwd(), "AGENTS.md"),
  path.join(process.cwd(), "CLAUDE.md"),
];

const pattern = /\[[^\]]+\]\(([^)#]+)\)/g;
const missing = [];

for (const doc of docs) {
  const text = fs.readFileSync(doc, "utf8");
  for (const match of text.matchAll(pattern)) {
    const rel = match[1];
    if (rel.includes("://") || rel.startsWith("#")) {
      continue;
    }

    const target = path.resolve(path.dirname(doc), rel);
    if (!fs.existsSync(target)) {
      missing.push(`${doc}: ${rel}`);
    }
  }
}

if (missing.length > 0) {
  console.error(missing.join("\n"));
  process.exit(1);
}

function walk(dir) {
  const entries = fs.readdirSync(dir, { withFileTypes: true });
  const files = [];

  for (const entry of entries) {
    const fullPath = path.join(dir, entry.name);
    if (entry.isDirectory()) {
      files.push(...walk(fullPath));
      continue;
    }

    if (entry.isFile() && fullPath.endsWith(".md")) {
      files.push(fullPath);
    }
  }

  return files;
}
JS
```

PowerShell 等价命令：

```powershell
$docs = @()
$docs += (Get-ChildItem -Path ai_docs -Recurse -Filter *.md).FullName
$docs += (Resolve-Path README.md).Path
$docs += (Resolve-Path AGENTS.md).Path
$docs += (Resolve-Path CLAUDE.md).Path
$pattern = '\[[^\]]+\]\(([^)#]+)\)'
$missing = @()

foreach ($doc in $docs) {
    $text = Get-Content -Raw -LiteralPath $doc
    foreach ($match in [regex]::Matches($text, $pattern)) {
        $rel = $match.Groups[1].Value
        if ($rel -match '://' -or $rel.StartsWith('#')) {
            continue
        }

        $baseDir = Split-Path -Parent $doc
        $target = [System.IO.Path]::GetFullPath((Join-Path $baseDir $rel))
        if (-not (Test-Path -LiteralPath $target)) {
            $missing += ('{0}: {1}' -f $doc, $rel)
        }
    }
}

if ($missing.Count -gt 0) {
    $missing
    exit 1
}
```

## Usage rule

- backend-only 任务跑 `Backend`
- frontend-only 任务跑 `Frontend`
- 同时改前后端、脚本或模板文档时跑 `Full stack`
- 改 `ai_docs/`、`README.md`、`AGENTS.md` 或 `CLAUDE.md` 时，先跑对应任务的本地验证，再额外跑 `Docs and links`
- 判断自动化门禁时看 `CI gate`，不要把它和 `Full stack` 视为同一组检查
