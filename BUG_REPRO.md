# BUG_REPRO: context 取消/超时不向下游传播

## Bug 是什么
`ops_clock.go` 的 `opsContext` 用 `context.Background()` 派生超时上下文，丢掉父 ctx 的 deadline；`opsDelay` 不再监听 `ctx.Done()`；`ops_store.go` 的 Get/Put/Update 丢失 ctx 取消检查。调用方取消请求后，后台存储与延迟仍继续执行。

## 如何触发
构造已取消的 context 调用 `opsDelay`、`OpsStore.Get/Put/Update`，或让父 context 带更早的 deadline 再经 `opsContext` 派生。

## 真实错误信息
- `opsDelay(ctx, 5s)` 在 ctx 已取消时不返回 `context.Canceled`，等待完整 5 秒。
- `OpsStore.Get` 用已取消 ctx 查询不存在的记录返回 `ErrOpsNotFound` 而非 `context.Canceled`。
- `opsContext(parent, 10s)` 在父 deadline 为 2s 时仍返回 10s 的 deadline。
