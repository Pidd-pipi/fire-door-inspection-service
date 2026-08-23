# BUG_REPRO: 跨层状态机错位（reopened 状态口径不一致）

## Bug 是什么
校验层 `validation/validation.go` 允许新增的 `reopened` 中间态，但 `store/store.go` 的 `UpdateStatus` 对 `reopened` 直接返回 `ErrNotFound`；`httpapi/handler.go` 丢掉了状态校验（任意状态都能改成功），且放宽了路由方法约束（GET 也能改状态）。

## 如何触发
`POST /api/v1/inspections/fd-101/status` 提交 `{"status":"reopened"}`；提交 `{"status":"bogus"}`；用 GET 访问状态接口。

## 真实错误信息
- reopened 请求返回 404（应为 200）。
- bogus 请求返回 200（应为 400）。
- GET 请求返回 400/200（应为 405）。
