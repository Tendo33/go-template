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
