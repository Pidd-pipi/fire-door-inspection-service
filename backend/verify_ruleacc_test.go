package main

import "testing"

// R10: 规则必填标签必须独立，reviewed 不随装载次数累积（diagnosis 红灯复现）。

func reviewedCount(labels []string) int {
	n := 0
	for _, l := range labels {
		if l == "reviewed" {
			n++
		}
	}
	return n
}

func ruleByCode(rules []OpsRule, code string) OpsRule {
	for _, r := range rules {
		if r.Code == code {
			return r
		}
	}
	return OpsRule{}
}

func TestRule0702ReviewedDoesNotAccumulate(t *testing.T) {
	first := ruleByCode(opsRules07(), "OPS-0702")
	if got := reviewedCount(first.RequiredLabels); got != 1 {
		t.Fatalf("首次装载 OPS-0702 reviewed 数量应为 1，实际 %d", got)
	}
	second := ruleByCode(opsRules07(), "OPS-0702")
	if got := reviewedCount(second.RequiredLabels); got != 1 {
		t.Fatalf("二次装载 OPS-0702 reviewed 数量仍应为 1，实际 %d", got)
	}
}

func TestRule0704ReviewedDoesNotAccumulate(t *testing.T) {
	first := ruleByCode(opsRules07(), "OPS-0704")
	if got := reviewedCount(first.RequiredLabels); got != 1 {
		t.Fatalf("首次装载 OPS-0704 reviewed 数量应为 1，实际 %d", got)
	}
	second := ruleByCode(opsRules07(), "OPS-0704")
	if got := reviewedCount(second.RequiredLabels); got != 1 {
		t.Fatalf("二次装载 OPS-0704 reviewed 数量仍应为 1，实际 %d", got)
	}
}

func TestRule0802ReviewedDoesNotAccumulate(t *testing.T) {
	first := ruleByCode(opsRules08(), "OPS-0802")
	if got := reviewedCount(first.RequiredLabels); got != 1 {
		t.Fatalf("首次装载 OPS-0802 reviewed 数量应为 1，实际 %d", got)
	}
	second := ruleByCode(opsRules08(), "OPS-0802")
	if got := reviewedCount(second.RequiredLabels); got != 1 {
		t.Fatalf("二次装载 OPS-0802 reviewed 数量仍应为 1，实际 %d", got)
	}
}

func TestRule0804ReviewedDoesNotAccumulate(t *testing.T) {
	first := ruleByCode(opsRules08(), "OPS-0804")
	if got := reviewedCount(first.RequiredLabels); got != 1 {
		t.Fatalf("首次装载 OPS-0804 reviewed 数量应为 1，实际 %d", got)
	}
	second := ruleByCode(opsRules08(), "OPS-0804")
	if got := reviewedCount(second.RequiredLabels); got != 1 {
		t.Fatalf("二次装载 OPS-0804 reviewed 数量仍应为 1，实际 %d", got)
	}
}
