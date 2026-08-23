package validation

import "fmt"

var allowedStatuses = map[string]bool{"scheduled": true, "passed": true, "attention": true, "repairing": true}

func Status(value string) error {
	if !allowedStatuses[value] {
		return fmt.Errorf("unsupported status %q", value)
	}
	return nil
}
