package main

import (
	"reflect"
	"testing"
)

// R5: 规则标签必须独立，不能被其它规则串改。

func equalStrings(a, b []string) bool { return reflect.DeepEqual(a, b) }

func findRule(t *testing.T, rules []OpsRule, code string) OpsRule {
	t.Helper()
	for _, r := range rules {
		if r.Code == code {
			return r
		}
	}
	t.Fatalf("规则 %s 不存在", code)
	return OpsRule{}
}

func TestRule0102LabelsNotPolluted(t *testing.T) {
	r := findRule(t, opsRules01(), "OPS-0102")
	want := []string{"site", "operator", "evidence", "reviewed", "verified"}
	if !equalStrings(r.RequiredLabels, want) {
		t.Fatalf("OPS-0102 RequiredLabels 被污染: %v", r.RequiredLabels)
	}
}

func TestRule0104LabelsNotPolluted(t *testing.T) {
	r := findRule(t, opsRules01(), "OPS-0104")
	want := []string{"site", "operator", "evidence", "reviewed", "verified"}
	if !equalStrings(r.RequiredLabels, want) {
		t.Fatalf("OPS-0104 RequiredLabels 被污染: %v", r.RequiredLabels)
	}
}

func TestRule0202LabelsNotPolluted(t *testing.T) {
	r := findRule(t, opsRules02(), "OPS-0202")
	want := []string{"site", "operator", "evidence", "reviewed", "verified"}
	if !equalStrings(r.RequiredLabels, want) {
		t.Fatalf("OPS-0202 RequiredLabels 被污染: %v", r.RequiredLabels)
	}
}

func TestRule0108LabelsNotPolluted(t *testing.T) {
	r := findRule(t, opsRules01(), "OPS-0108")
	want := []string{"site", "operator", "evidence", "reviewed", "verified"}
	if !equalStrings(r.RequiredLabels, want) {
		t.Fatalf("OPS-0108 RequiredLabels 被污染: %v", r.RequiredLabels)
	}
}

func TestRule0204LabelsNotPolluted(t *testing.T) {
	r := findRule(t, opsRules02(), "OPS-0204")
	want := []string{"site", "operator", "evidence", "reviewed", "verified"}
	if !equalStrings(r.RequiredLabels, want) {
		t.Fatalf("OPS-0204 RequiredLabels 被污染: %v", r.RequiredLabels)
	}
}

func TestRule0208LabelsNotPolluted(t *testing.T) {
	r := findRule(t, opsRules02(), "OPS-0208")
	want := []string{"site", "operator", "evidence", "reviewed", "verified"}
	if !equalStrings(r.RequiredLabels, want) {
		t.Fatalf("OPS-0208 RequiredLabels 被污染: %v", r.RequiredLabels)
	}
}

func TestRule0101LabelsClean(t *testing.T) {
	r := findRule(t, opsRules01(), "OPS-0101")
	want := []string{"site", "operator", "evidence", "verified"}
	if !equalStrings(r.RequiredLabels, want) {
		t.Fatalf("OPS-0101 RequiredLabels 异常: %v", r.RequiredLabels)
	}
}
