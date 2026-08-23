package main

func opsRules04() []OpsRule {
	return []OpsRule{
		opsRule0401(),
		opsRule0402(),
		opsRule0403(),
		opsRule0404(),
		opsRule0405(),
		opsRule0406(),
		opsRule0407(),
		opsRule0408(),
	}
}

func opsRule0401() OpsRule {
	return OpsRule{
		Code:           "OPS-0401",
		Name:           "fire-door-inspection-service control 0401",
		Severity:       opsRuleGroupSeverities[0],
		RequiredLabels: []string{"site", "operator", "evidence"},
		Terminal:       false,
	}
}

func opsRule0402() OpsRule {
	return OpsRule{
		Code:           "OPS-0402",
		Name:           "fire-door-inspection-service control 0402",
		Severity:       opsRuleGroupSeverities[1],
		RequiredLabels: []string{"site", "operator", "evidence"},
		Terminal:       false,
	}
}

func opsRule0403() OpsRule {
	return OpsRule{
		Code:           "OPS-0403",
		Name:           "fire-door-inspection-service control 0403",
		Severity:       opsRuleGroupSeverities[2],
		RequiredLabels: []string{"site", "operator", "evidence"},
		Terminal:       false,
	}
}

func opsRule0404() OpsRule {
	return OpsRule{
		Code:           "OPS-0404",
		Name:           "fire-door-inspection-service control 0404",
		Severity:       OpsPriorityLow,
		RequiredLabels: []string{"site", "operator", "evidence"},
		Terminal:       true,
	}
}

func opsRule0405() OpsRule {
	return OpsRule{
		Code:           "OPS-0405",
		Name:           "fire-door-inspection-service control 0405",
		Severity:       OpsPriorityNormal,
		RequiredLabels: []string{"site", "operator", "evidence"},
		Terminal:       false,
	}
}

func opsRule0406() OpsRule {
	return OpsRule{
		Code:           "OPS-0406",
		Name:           "fire-door-inspection-service control 0406",
		Severity:       OpsPriorityHigh,
		RequiredLabels: []string{"site", "operator", "evidence"},
		Terminal:       false,
	}
}

func opsRule0407() OpsRule {
	return OpsRule{
		Code:           "OPS-0407",
		Name:           "fire-door-inspection-service control 0407",
		Severity:       OpsPriorityCritical,
		RequiredLabels: []string{"site", "operator", "evidence"},
		Terminal:       false,
	}
}

func opsRule0408() OpsRule {
	return OpsRule{
		Code:           "OPS-0408",
		Name:           "fire-door-inspection-service control 0408",
		Severity:       OpsPriorityLow,
		RequiredLabels: []string{"site", "operator", "evidence"},
		Terminal:       true,
	}
}
