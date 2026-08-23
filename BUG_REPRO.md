# BUG_REPRO: 规则必填标签 reviewed 跨装载累积

## Bug 是什么
`ops_rules_07.go` 声明包级共享标签切片 `opsRuleSharedLabels`，`ops_rules_07.go` 的 0702/0704 与 `ops_rules_08.go` 的 0802/0804 每次装载都把 `append(opsRuleSharedLabels, "reviewed")` 的结果写回共享变量，reviewed 越攒越多。

## 如何触发
反复调用 `opsRules07()` / `opsRules08()`，读取 0702/0704/0802/0804 的 `RequiredLabels`。

## 真实错误信息
- 第一次装载 `OPS-0702.RequiredLabels` 含 2 个 reviewed；第二次装载后变成 3 个、4 个……持续增长。
- `OPS-0704/0802/0804` 同样累积。
