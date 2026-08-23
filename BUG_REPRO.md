# BUG_REPRO: 审计/状态机并发读写数据竞争

## Bug 是什么
`ops_audit.go` 的 For/Since/Count 与 `ops_state.go` 的 Move/History/Last/Reset 丢失了互斥锁保护，并发写审计事件、推进状态机与读历史时发生 data race，历史与审计快照在并发读写下不稳定。

## 如何触发
多个 goroutine 同时调用 `OpsAudit.Add` / `OpsStateMachine.Move` 并读取 `For`/`Since`/`History`/`Last`/`Count`，用 `-race` 运行即可稳定复现。

## 真实错误信息
```
WARNING: DATA RACE
Write at 0x00c000100ca8 by goroutine 9:
  example.com/fire-door-inspection-service.(*OpsAudit).Add()
      .../backend/ops_audit.go:24 +0x2a4
Previous read at 0x00c000100ca8 by goroutine 15:
  example.com/fire-door-inspection-service.(*OpsAudit).For()
      .../backend/ops_audit.go:29 +0xcc
```
