# Backend Standards

## When to read

- backend 任务开始前必读。
- 新增配置、服务边界、HTTP 行为或公共导入面时回到这里。

## Default stack

- Go 1.26
- Gin
- `zap`
- `golangci-lint`
- Go 自带测试工具链

日志与请求链路默认约定：

- 开发环境（`development`、`dev`、`local`）优先可读性，使用彩色 console encoder
- 非开发环境优先机器采集，使用 JSON 结构化日志
- HTTP 请求默认透传或生成 `X-Request-ID`
- 请求链路日志通过 `context.Context` 传递，不把请求态 logger 放进全局变量或 service struct
- handler/service 读取请求级 logger 时统一使用 `observability.FromContext(...)`
- `trace_id` / `span_id` 可作为后续 OpenTelemetry 接入的 context 字段位，模板阶段不直接引入 OTel SDK

只有在项目确实需要时，再引入数据库、ORM、Swagger、gRPC 或消息队列。

## Design rules

- handler 保持轻薄
- 业务流程尽量落在 `service`
- 优先放在 `internal/` 下，不急着做外部复用包
- 任何抽象如果不能让代码更清晰，就不要引入
- 避免为了分层而分层，先用直接、可扫描的实现

## Public API and import rules

- 对外可见行为只引用当前真实存在的 HTTP 接口或环境变量
- 文档和示例不引用不存在的模块
- 不把 `internal/` 当成稳定外部 API

## Configuration and safety

- 配置统一从环境变量读取
- 服务启动前必须显式校验关键配置，至少覆盖端口、服务名和日志级别这类基础项
- secrets 只从环境变量或受控配置源读取
- 日志不打印 token、密码或敏感个人数据
- 外部输入必须显式校验

## Testing rules

- 新行为补测试
- bugfix 至少补一个回归测试
- 纯逻辑优先单元测试
- 真正涉及进程、网络或脚本行为时再补集成测试

## Shared references

- 当前 backend 现状见 [backend.md](../current/backend.md)
- 命名与路径见 [naming-and-paths.md](../reference/naming-and-paths.md)
- 验证命令见 [verification.md](../reference/verification.md)
