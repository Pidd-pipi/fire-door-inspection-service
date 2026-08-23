package main

import (
	"context"
	"testing"
	"time"
)

// R2: 取消/超时必须向下游传播。

func TestOpsDelayHonorsCancellation(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	start := time.Now()
	if err := opsDelay(ctx, 5*time.Second); err == nil {
		t.Fatal("已取消 ctx 的 opsDelay 应返回错误")
	}
	if time.Since(start) > 500*time.Millisecond {
		t.Fatalf("opsDelay 未及时响应取消: %v", time.Since(start))
	}
}

func TestOpsContextKeepsParentDeadline(t *testing.T) {
	parent, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	child, childCancel := opsContext(parent, 10*time.Second)
	defer childCancel()
	parentDeadline, _ := parent.Deadline()
	childDeadline, ok := child.Deadline()
	if !ok {
		t.Fatal("opsContext 结果应保留 deadline")
	}
	if childDeadline.After(parentDeadline) {
		t.Fatalf("opsContext 丢父 deadline: parent=%v child=%v", parentDeadline, childDeadline)
	}
}

func TestOpsStoreGetHonorsCancelledContext(t *testing.T) {
	s := newOpsStore(nil)
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	if _, err := s.Get(ctx, "missing"); err != context.Canceled {
		t.Fatalf("取消 ctx 的 Get 应返回 context.Canceled，实际 %v", err)
	}
}

func TestOpsStorePutHonorsCancelledContext(t *testing.T) {
	s := newOpsStore(nil)
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	rec := OpsRecord{ID: "op-1", Subject: "door check", Owner: "alice", Priority: OpsPriorityNormal}
	if err := s.Put(ctx, rec); err != context.Canceled {
		t.Fatalf("取消 ctx 的 Put 应返回 context.Canceled，实际 %v", err)
	}
}

func TestOpsStoreUpdateHonorsCancelledContext(t *testing.T) {
	s := newOpsStore(nil)
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	rec := OpsRecord{ID: "op-1", Subject: "door check", Owner: "alice", Priority: OpsPriorityNormal}
	if err := s.Update(ctx, rec, 0); err != context.Canceled {
		t.Fatalf("取消 ctx 的 Update 应返回 context.Canceled，实际 %v", err)
	}
}
