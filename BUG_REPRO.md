# BUG_REPRO: 规则必填标签被共享底层数组串改

## Bug 是什么
`ops_rules_01.go`/`ops_rules_02.go` 的规则标签改为共享包级 `opsRuleLabelsShared[:3]`（带多余容量），偶数规则原地 `append("reviewed")` 写入共享底层数组，装载时 enrich 又原地追加 `verified`，把其它规则（0102/0104/0108/0202/0204/0208）的 `reviewed` 覆盖成 `verified`，RequiredLabels 被串改。

## 如何触发
调用 `opsRules01()` / `opsRules02()` 后读取偶数规则的 `RequiredLabels`。

## 真实错误信息
- `OPS-0102.RequiredLabels` 实际为 `[site operator evidence verified verified]`，应为 `[site operator evidence reviewed verified]`。
- `OPS-0104/0108/0202/0204/0208` 同样被串改。
