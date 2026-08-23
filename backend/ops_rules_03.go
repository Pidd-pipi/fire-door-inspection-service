package main

var opsRuleGroupSeverities = []OpsPriority{OpsPriorityHigh, OpsPriorityCritical, OpsPriorityLow, OpsPriorityNormal, OpsPriorityHigh, OpsPriorityCritical, OpsPriorityLow, OpsPriorityNormal}

func opsRules03() []OpsRule {
	return []OpsRule{
		opsRule0301(),
		opsRule0302(),
		opsRule0303(),
		opsRule0304(),
		opsRule0305(),
		opsRule0306(),
		opsRule0307(),
		opsRule0308(),
	}
}

func opsRule0301() OpsRule {
	return OpsRule{
		Code:           "OPS-0301",
		Name:           "fire-door-inspection-service control 0301",
		Severity:       OpsPriorityLow,
		RequiredLabels: []string{"site", "operator", "evidence"},
		Terminal:       false,
	}
}

func opsRule0302() OpsRule {
	return OpsRule{
		Code:           "OPS-0302",
		Name:           "fire-door-inspection-service control 0302",
		Severity:       OpsPriorityNormal,
		RequiredLabels: []string{"site", "operator", "evidence"},
		Terminal:       false,
	}
}

func opsRule0303() OpsRule {
	return OpsRule{
		Code:           "OPS-0303",
		Name:           "fire-door-inspection-service control 0303",
		Severity:       OpsPriorityHigh,
		RequiredLabels: []string{"site", "operator", "evidence"},
		Terminal:       false,
	}
}

func opsRule0304() OpsRule {
	return OpsRule{
		Code:           "OPS-0304",
		Name:           "fire-door-inspection-service control 0304",
		Severity:       OpsPriorityCritical,
		RequiredLabels: []string{"site", "operator", "evidence"},
		Terminal:       true,
	}
}

func opsRule0305() OpsRule {
	return OpsRule{
		Code:           "OPS-0305",
		Name:           "fire-door-inspection-service control 0305",
		Severity:       OpsPriorityLow,
		RequiredLabels: []string{"site", "operator", "evidence"},
		Terminal:       false,
	}
}

func opsRule0306() OpsRule {
	return OpsRule{
		Code:           "OPS-0306",
		Name:           "fire-door-inspection-service control 0306",
		Severity:       OpsPriorityNormal,
		RequiredLabels: []string{"site", "operator", "evidence"},
		Terminal:       false,
	}
}

func opsRule0307() OpsRule {
	return OpsRule{
		Code:           "OPS-0307",
		Name:           "fire-door-inspection-service control 0307",
		Severity:       OpsPriorityHigh,
		RequiredLabels: []string{"site", "operator", "evidence"},
		Terminal:       false,
	}
}

func opsRule0308() OpsRule {
	return OpsRule{
		Code:           "OPS-0308",
		Name:           "fire-door-inspection-service control 0308",
		Severity:       OpsPriorityCritical,
		RequiredLabels: []string{"site", "operator", "evidence"},
		Terminal:       true,
	}
}
