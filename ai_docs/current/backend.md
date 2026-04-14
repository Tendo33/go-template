# Current Backend

## When to read

- 要改 backend 代码前先读这里。
- 需要了解当前 backend 真实实现、公共入口和配置基线时读这里。

## Current truth

当前 backend 是一个轻量模板服务，服务入口位于 `cmd/server/`，核心实现位于 `internal/`。

已经落地的模块：

- `config`：环境变量配置加载
- `observability`：基于 `slog` 的结构化日志初始化
- `httpserver`：Gin 路由、middleware 和 `http.Server` 组装
- `service`：当前只有健康检查服务
- `model`：健康检查响应结构

## Settings and configuration

- 当前配置定义在 `internal/config`
- 基线字段只有 `APP_ENV`、`PORT`、`LOG_LEVEL`、`SERVICE_NAME`
- 默认服务名是 `go-template`
- 当前没有引入配置文件解析或多层配置来源

## HTTP and public behavior

- 当前服务入口为 `cmd/server/main.go`
- 当前只暴露一个 HTTP 接口：`GET /health`
- 成功时返回 `200` 和 JSON：

```json
{
  "status": "ok",
  "service": "go-template"
}
```

## What is not here yet

- 当前还没有数据库或 repository 层
- 当前还没有鉴权、Swagger 或 gRPC
- 当前还没有按真实业务领域拆开的复杂 service 结构

这些都属于后续项目扩展方向，具体做法以 standards 文档和当前代码现状为准。

## Shared references

- backend 约束见 [backend.md](../standards/backend.md)
- 项目结构见 [project-structure.md](../reference/project-structure.md)
- 验证命令见 [verification.md](../reference/verification.md)
