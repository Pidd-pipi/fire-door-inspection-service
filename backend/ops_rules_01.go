package main

// opsRuleLabelsShared holds the common read-only base labels. It must keep
// cap == len so that every `append` on a slice taken from it allocates a fresh
// backing array; adding spare capacity here caused each rule's `append(...,
// "reviewed"/"verified")` to mutate this shared array and cross-contaminate
// other rules' RequiredLabels.
var opsRuleLabelsShared = []string{"site", "operator", "evidence"}

func opsRules01() []OpsRule {
	rules := []OpsRule{
		opsRule0101(),
		opsRule0102(),
		opsRule0103(),
		opsRule0104(),
		opsRule0105(),
		opsRule0106(),
		opsRule0107(),
		opsRule0108(),
	}
	for i := range rules {
		rules[i].RequiredLabels = append(rules[i].RequiredLabels, "verified")
	}
	return rules
}

func opsRule0101() OpsRule {
	labels := opsRuleLabelsShared[:3]
	return OpsRule{
		Code:           "OPS-0101",
		Name:           "fire-door-inspection-service control 0101",
		Severity:       OpsPriorityHigh,
		RequiredLabels: labels,
		Terminal:       false,
	}
}

func opsRule0102() OpsRule {
	labels := opsRuleLabelsShared[:3]
	labels = append(labels, "reviewed")
	return OpsRule{
		Code:           "OPS-0102",
		Name:           "fire-door-inspection-service control 0102",
		Severity:       OpsPriorityCritical,
		RequiredLabels: labels,
		Terminal:       false,
	}
}

func opsRule0103() OpsRule {
	labels := opsRuleLabelsShared[:3]
	return OpsRule{
		Code:           "OPS-0103",
		Name:           "fire-door-inspection-service control 0103",
		Severity:       OpsPriorityLow,
		RequiredLabels: labels,
		Terminal:       false,
	}
}

func opsRule0104() OpsRule {
	labels := opsRuleLabelsShared[:3]
	labels = append(labels, "reviewed")
	return OpsRule{
		Code:           "OPS-0104",
		Name:           "fire-door-inspection-service control 0104",
		Severity:       OpsPriorityNormal,
		RequiredLabels: labels,
		Terminal:       true,
	}
}

func opsRule0105() OpsRule {
	labels := []string{"site", "operator", "evidence"}
	return OpsRule{
		Code:           "OPS-0105",
		Name:           "fire-door-inspection-service control 0105",
		Severity:       OpsPriorityHigh,
		RequiredLabels: labels,
		Terminal:       false,
	}
}

func opsRule0106() OpsRule {
	labels := []string{"site", "operator", "evidence"}
	labels = append(labels, "reviewed")
	return OpsRule{
		Code:           "OPS-0106",
		Name:           "fire-door-inspection-service control 0106",
		Severity:       OpsPriorityCritical,
		RequiredLabels: labels,
		Terminal:       false,
	}
}

func opsRule0107() OpsRule {
	labels := opsRuleLabelsShared[:3]
	return OpsRule{
		Code:           "OPS-0107",
		Name:           "fire-door-inspection-service control 0107",
		Severity:       OpsPriorityLow,
		RequiredLabels: labels,
		Terminal:       false,
	}
}

func opsRule0108() OpsRule {
	labels := opsRuleLabelsShared[:3]
	labels = append(labels, "reviewed")
	return OpsRule{
		Code:           "OPS-0108",
		Name:           "fire-door-inspection-service control 0108",
		Severity:       OpsPriorityNormal,
		RequiredLabels: labels,
		Terminal:       true,
	}
}
