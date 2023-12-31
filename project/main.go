package project

import "fmt"

// Version is ruleset version
const Version string = "0.3.0"

// ReferenceLink returns the rule reference link
func ReferenceLink(name string) string {
	return fmt.Sprintf("https://github.com/KazanEpxeress/tflint-ruleset/blob/v%s/docs/%s.md", Version, name)
}
