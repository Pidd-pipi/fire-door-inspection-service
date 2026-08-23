package main

import (
	"context"
	"errors"
	"testing"
)

// R3: 错误链必须保留哨兵，errors.Is/opsCode 分类可用。

func TestWrapOpsPreservesCause(t *testing.T) {
	err := wrapOps("conflict", "store.put", ErrOpsConflict)
	if !errors.Is(err, ErrOpsConflict) {
		t.Fatalf("wrapOps 断链: %v", err)
	}
}

func TestOpsTransitionConflictKeepsSentinel(t *testing.T) {
	svc := newOpsService([]OpsRecord{{ID: "op-1", Subject: "door check", Owner: "alice", Priority: OpsPriorityNormal, Revision: 3, Labels: map[string]string{"site": "s1"}}})
	_, err := svc.Transition(context.Background(), "op-1", 99, OpsStatusActive, "tester")
	if !errors.Is(err, ErrOpsConflict) {
		t.Fatalf("版本冲突应保留 ErrOpsConflict: %v", err)
	}
	if got := opsCode(err); got != "conflict" {
		t.Fatalf("版本冲突分类错误: %q", got)
	}
}

func TestOpsTransitionMoveKeepsSentinel(t *testing.T) {
	svc := newOpsService([]OpsRecord{{ID: "op-1", Subject: "door check", Owner: "alice", Priority: OpsPriorityNormal, Revision: 3, Labels: map[string]string{"site": "s1"}}})
	_, err := svc.Transition(context.Background(), "op-1", 0, OpsStatus("bogus"), "tester")
	if !errors.Is(err, ErrOpsTransition) {
		t.Fatalf("非法迁移应保留 ErrOpsTransition: %v", err)
	}
	if got := opsCode(err); got != "transition" {
		t.Fatalf("非法迁移分类错误: %q", got)
	}
}

func TestOpsGetMissingPreservesNotFoundSentinel(t *testing.T) {
	svc := newOpsService(nil)
	if _, err := svc.Get(context.Background(), "nope"); !errors.Is(err, ErrOpsNotFound) {
		t.Fatalf("Get 不存在记录应保留 ErrOpsNotFound: %v", err)
	}
}

func TestOpsCreatePolicyErrorKeepsSentinel(t *testing.T) {
	svc := newOpsService(nil)
	record := OpsRecord{ID: "op-p", Subject: "door check", Owner: "bob", Priority: OpsPriorityNormal}
	if _, err := svc.Create(context.Background(), record); !errors.Is(err, ErrOpsPolicy) {
		t.Fatalf("策略拒绝应保留 ErrOpsPolicy: %v", err)
	}
}

func TestOpsCreateStoreErrorKeepsSentinel(t *testing.T) {
	svc := newOpsService(nil)
	rec := OpsRecord{ID: "op-9", Subject: "door check", Owner: "bob", Priority: OpsPriorityNormal, Labels: map[string]string{"site": "s1"}}
	if _, err := svc.Create(context.Background(), rec); err != nil {
		t.Fatal(err)
	}
	_, err := svc.Create(context.Background(), rec)
	if !errors.Is(err, ErrOpsConflict) {
		t.Fatalf("重复创建应保留 ErrOpsConflict: %v", err)
	}
	if got := opsCode(err); got == "internal" {
		t.Fatalf("重复创建被误判为 internal: %v", err)
	}
}
