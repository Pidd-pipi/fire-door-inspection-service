# BUG_REPRO: HTTP 服务生命周期 defer 漏释放

## Bug 是什么
`runtime.go` 的 `serveHTTP` 把 `errCh` 改为无缓冲通道，服务经信号关停后监听 goroutine 向 errCh 发送永远阻塞（goroutine 泄露）；`recoveryMiddleware` 删除 recover defer，处理器 panic 直接逃逸；`requestIDMiddleware` 与 `opsEnterpriseMiddleware` 用 defer 在响应头提交后才写 `X-Request-ID`/`X-Operations-Latency-Ms`，头被丢弃。

## 如何触发
反复启动/关停服务（SIGTERM）；访问任意接口检查响应头；处理器内 panic。

## 真实错误信息
- 关停 3 轮后 `runtime.NumGoroutine()` 持续增长（每轮 +2）。
- 真实响应缺少 `X-Request-ID` 与 `X-Operations-Latency-Ms` 头。
- 处理器 panic 时请求直接中断，无 500 兜底。
