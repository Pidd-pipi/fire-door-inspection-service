# BUG_REPRO: 规则组严重度被共享 severity 表串改

## Bug 是什么
`ops_rules_03.go` 声明包级共享切片 `opsRuleGroupSeverities`（值为 03 组模式旋转 2 位后的序列），`ops_rules_03.go` 的 0301/0302/0303 与 `ops_rules_04.go` 的 0401/0402/0403 都按偏移读取这张表，表内容与两组各自定义都不符，严重度跨组串改。

## 如何触发
调用 `opsRules03()` / `opsRules04()` 读取前三条规则的 `Severity`。

## 真实错误信息
- `OPS-0301` 为 high（应为 low）、`OPS-0302` 为 critical（应为 normal）、`OPS-0303` 为 low（应为 high）。
- `OPS-0401` 为 high（应为 normal）、`OPS-0402` 为 critical（应为 high）、`OPS-0403` 为 low（应为 critical）。
