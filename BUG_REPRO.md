# BUG_REPRO: 零值路径 nil map 写入 panic 与配置端口失效

## Bug 是什么
`ops_model.go` 的 `normalizeOpsRecord` 对 nil `Labels` 无保护地写默认 `site` 标签，触发 nil map 写入 panic，且给记录伪造站点标签绕过策略校验；`OpsRecord.Clone` 丢失深拷贝共享 Labels；`config.Load` 放行 0/越界端口且 `Address` 无视配置端口固定返回 `:8080`。

## 如何触发
创建不带 `Labels` 的巡检记录；克隆记录后修改标签；设置 `PORT=0` 或 `PORT=99999` 启动服务。

## 真实错误信息
```
panic: assignment to entry in nil map [recovered]
	panic: assignment to entry in nil map

goroutine 35 [running]:
example.com/fire-door-inspection-service.normalizeOpsRecord(...)
	.../backend/ops_model.go:112 +0xf4
```
另外：`PORT=9090` 时服务仍监听 8080；`PORT=99999` 时被直接采用。
