# BUG_REPRO: 健康检查与页面入口跨层状态错位

## Bug 是什么
`health/health.go` 的 Handler 不校验请求方法且响应缺少 `service` 身份字段；`web/web.go` 的 Handler 不再服务 `/app.js` 静态资源；`main.go` 的 `newMux` 把页面处理器注册在 `/static/` 前缀下，根路径与 `/app.js` 全部 404，页面打不开但健康检查仍返回 ok。

## 如何触发
`GET /healthz`、`POST /healthz`、`GET /`、`GET /app.js`。

## 真实错误信息
- `GET /healthz` 返回 `{"status":"ok"}`（缺 `service` 身份）。
- `POST /healthz` 返回 200（应为 405）。
- `GET /` 与 `GET /app.js` 返回 404（应 200）。
