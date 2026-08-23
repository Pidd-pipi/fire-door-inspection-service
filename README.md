# 防火门巡检服务

这是一个标准库实现的防火门巡检台账 HTTP 服务。服务默认监听 `8080`，也可通过 `PORT` 调整。

接口：`GET /healthz`；`GET /api/v1/inspections`；`POST /api/v1/inspections/{id}/status`，请求体为 `{"status":"passed|attention|repairing|scheduled"}`。根路径提供内置台账页面。启动命令为 `cd backend && go run .`。

## Directory Tree and API

```text
backend/       Go 模块、HTTP 服务、领域包和内嵌 web 静态资源
database/      数据库说明
output/        验证记录
README.md      项目说明
prompt.txt     任务提示
runtime_smoke.json  运行冒烟配置
```

健康检查：`GET /healthz`。API：`GET /api/v1/inspections`、`POST /api/v1/inspections/{id}/status`；页面入口：`GET /`。

## Verification

在项目目录执行 `gofmt -w $(rg --files -g '*.go')`、`go build ./...`、`go test ./...`，结果全部通过；HTTP 定向测试通过。

真实服务使用端口 `18101` 启动后验证：`GET /healthz` 返回 `{"service":"fire-door-inspection","status":"ok"}`；`GET /api/v1/inspections` 返回 2 条记录；`POST /api/v1/inspections/fd-101/status` 将状态更新为 `attention` 并返回 200；非法状态返回 400；未知记录返回 404；`/` 返回 822 字节，`/app.js` 返回 981 字节。验证后服务已关闭。

## Engineering Notes

防火门巡检流程代码按领域模型、校验、状态转换、并发安全存储、审计事件和 HTTP 生命周期分层。请求会保留请求标识并经过恢复与超时保护；状态写入使用版本校验，错误通过可识别的领域错误返回。

除现有接口回归测试外，项目还保留可复用的分页、过滤、策略、工作流和运行健康能力，便于后续扩展而不把业务规则堆积到处理器中。
