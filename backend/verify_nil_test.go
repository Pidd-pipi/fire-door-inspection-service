package main

import (
	"context"
	"os"
	"testing"

	"example.com/fire-door-inspection-service/config"
)

// R4: 零值/nil 路径。

func TestOpsNormalizeUnlabeledRecordNoPanic(t *testing.T) {
	rec := normalizeOpsRecord(OpsRecord{ID: "op-x", Subject: "  door  check  ", Owner: "alice"})
	if rec.Labels == nil {
		t.Fatal("normalize 后 Labels 仍为 nil")
	}
	if rec.LabelValue("site") != "" {
		t.Fatalf("normalize 不应伪造 site 标签，实际 %q", rec.LabelValue("site"))
	}
}

func TestMissingSiteLabelCreateRejected(t *testing.T) {
	svc := newOpsService(nil)
	record := OpsRecord{ID: "op-y", Subject: "door check", Owner: "bob", Priority: OpsPriorityNormal}
	if _, err := svc.Create(context.Background(), record); err == nil {
		t.Fatal("缺少 site 标签的记录不应被放行")
	}
}

func TestOpsCloneReturnsIndependentLabels(t *testing.T) {
	original := OpsRecord{ID: "op-w", Subject: "door check", Owner: "dave", Priority: OpsPriorityNormal,
		Labels: map[string]string{"site": "b1"}}
	cloned := original.Clone()
	cloned.Labels["extra"] = "polluted"
	if original.LabelValue("extra") != "" {
		t.Fatalf("Clone 与原始记录共享 Labels: %+v", original.Labels)
	}
}

func TestPortRejectsOutOfRange(t *testing.T) {
	old, had := os.LookupEnv("PORT")
	os.Setenv("PORT", "99999")
	defer func() {
		if had {
			os.Setenv("PORT", old)
		} else {
			os.Unsetenv("PORT")
		}
	}()
	cfg := config.Load()
	if cfg.Port != 8080 {
		t.Fatalf("越界端口应回退默认 8080，实际 %d", cfg.Port)
	}
}

func TestAddressUsesConfiguredPort(t *testing.T) {
	cfg := config.Config{Port: 9090}
	if got := cfg.Address(); got != ":9090" {
		t.Fatalf("Address 应使用配置端口 9090，实际 %s", got)
	}
}

func TestPortFallsBackToDefault(t *testing.T) {
	old, had := os.LookupEnv("PORT")
	os.Setenv("PORT", "0")
	defer func() {
		if had {
			os.Setenv("PORT", old)
		} else {
			os.Unsetenv("PORT")
		}
	}()
	cfg := config.Load()
	if cfg.Port != 8080 {
		t.Fatalf("零值端口应回退默认 8080，实际 %d", cfg.Port)
	}
}
