# BUG_REPRO: 错误链断裂导致错误分类全部失效

## Bug 是什么
`ops_errors.go` 的 `wrapOps` 用 `errors.New(cause.Error())` 构造 Cause 断开哨兵链；`ops_service.go` 的 Create/Get/Transition 多处用 `%v` 包装错误丢失哨兵。`errors.Is/As` 与 `opsCode` 的分类全部失效，所有错误都被判为 internal。

## 如何触发
创建重复 ID 记录、查询不存在的记录、对记录做非法状态迁移。

## 真实错误信息
- 重复创建：`opsCode(err)` 返回 `"internal"`（应为 `"conflict"`），`errors.Is(err, ErrOpsConflict)` 为 false。
- 查询不存在：`errors.Is(err, ErrOpsNotFound)` 为 false。
- 非法迁移：`errors.Is(err, ErrOpsTransition)` 为 false。
