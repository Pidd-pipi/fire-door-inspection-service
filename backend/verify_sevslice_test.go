package main

import (
	"testing"
)

// R9: 每组规则的严重度必须独立，不能被其它组串改。

func sevOf(t *testing.T, rules []OpsRule, code string) OpsPriority {
	t.Helper()
	for _, r := range rules {
		if r.Code == code {
			return r.Severity
		}
	}
	t.Fatalf("规则 %s 不存在", code)
	return ""
}

func TestRule0301SeverityMatches(t *testing.T) {
	if got := sevOf(t, opsRules03(), "OPS-0301"); got != OpsPriorityLow {
		t.Fatalf("OPS-0301 严重度应为 low，实际 %s", got)
	}
}

func TestRule0302SeverityMatches(t *testing.T) {
	if got := sevOf(t, opsRules03(), "OPS-0302"); got != OpsPriorityNormal {
		t.Fatalf("OPS-0302 严重度应为 normal，实际 %s", got)
	}
}

func TestRule0303SeverityMatches(t *testing.T) {
	if got := sevOf(t, opsRules03(), "OPS-0303"); got != OpsPriorityHigh {
		t.Fatalf("OPS-0303 严重度应为 high，实际 %s", got)
	}
}

func TestRule0401SeverityMatches(t *testing.T) {
	if got := sevOf(t, opsRules04(), "OPS-0401"); got != OpsPriorityNormal {
		t.Fatalf("OPS-0401 严重度应为 normal，实际 %s", got)
	}
}

func TestRule0403SeverityMatches(t *testing.T) {
	if got := sevOf(t, opsRules04(), "OPS-0403"); got != OpsPriorityCritical {
		t.Fatalf("OPS-0403 严重度应为 critical，实际 %s", got)
	}
}

func TestRule0402SeverityMatches(t *testing.T) {
	if got := sevOf(t, opsRules04(), "OPS-0402"); got != OpsPriorityHigh {
		t.Fatalf("OPS-0402 严重度应为 high，实际 %s", got)
	}
}
