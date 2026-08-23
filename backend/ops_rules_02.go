package main

func opsRules02() []OpsRule {
	rules := []OpsRule{
		opsRule0201(),
		opsRule0202(),
		opsRule0203(),
		opsRule0204(),
		opsRule0205(),
		opsRule0206(),
		opsRule0207(),
		opsRule0208(),
	}
	for i := range rules {
		rules[i].RequiredLabels = append(rules[i].RequiredLabels, "verified")
	}
	return rules
}

func opsRule0201() OpsRule {
	labels := opsRuleLabelsShared[:3]
	return OpsRule{
		Code:           "OPS-0201",
		Name:           "fire-door-inspection-service control 0201",
		Severity:       OpsPriorityCritical,
		RequiredLabels: labels,
		Terminal:       false,
	}
}

func opsRule0202() OpsRule {
	labels := opsRuleLabelsShared[:3]
	labels = append(labels, "reviewed")
	return OpsRule{
		Code:           "OPS-0202",
		Name:           "fire-door-inspection-service control 0202",
		Severity:       OpsPriorityLow,
		RequiredLabels: labels,
		Terminal:       false,
	}
}

func opsRule0203() OpsRule {
	labels := opsRuleLabelsShared[:3]
	return OpsRule{
		Code:           "OPS-0203",
		Name:           "fire-door-inspection-service control 0203",
		Severity:       OpsPriorityNormal,
		RequiredLabels: labels,
		Terminal:       false,
	}
}

func opsRule0204() OpsRule {
	labels := opsRuleLabelsShared[:3]
	labels = append(labels, "reviewed")
	return OpsRule{
		Code:           "OPS-0204",
		Name:           "fire-door-inspection-service control 0204",
		Severity:       OpsPriorityHigh,
		RequiredLabels: labels,
		Terminal:       true,
	}
}

func opsRule0205() OpsRule {
	labels := []string{"site", "operator", "evidence"}
	return OpsRule{
		Code:           "OPS-0205",
		Name:           "fire-door-inspection-service control 0205",
		Severity:       OpsPriorityCritical,
		RequiredLabels: labels,
		Terminal:       false,
	}
}

func opsRule0206() OpsRule {
	labels := []string{"site", "operator", "evidence"}
	labels = append(labels, "reviewed")
	return OpsRule{
		Code:           "OPS-0206",
		Name:           "fire-door-inspection-service control 0206",
		Severity:       OpsPriorityLow,
		RequiredLabels: labels,
		Terminal:       false,
	}
}

func opsRule0207() OpsRule {
	labels := opsRuleLabelsShared[:3]
	return OpsRule{
		Code:           "OPS-0207",
		Name:           "fire-door-inspection-service control 0207",
		Severity:       OpsPriorityNormal,
		RequiredLabels: labels,
		Terminal:       false,
	}
}

func opsRule0208() OpsRule {
	labels := opsRuleLabelsShared[:3]
	labels = append(labels, "reviewed")
	return OpsRule{
		Code:           "OPS-0208",
		Name:           "fire-door-inspection-service control 0208",
		Severity:       OpsPriorityHigh,
		RequiredLabels: labels,
		Terminal:       true,
	}
}
