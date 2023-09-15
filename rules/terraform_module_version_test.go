package rules

import (
	"testing"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_TerraformUnusedRequiredProvidersRule(t *testing.T) {
	rule := NewTerraformModuleVersionTypeRule()

	cases := []struct {
		Name     string
		Content  string
		Config   string
		Expected helper.Issues
	}{
		{
			Name:     "empty",
			Content:  "",
			Expected: helper.Issues{},
		},
		{
			Name: "default config",
			Content: `
				terraform {
					source = "git::git@github.com:KazanExpress/ke-infra-modules.git//modules/org/group?ref=v0.400.0"
				}
			`,
			Expected: helper.Issues{},
		},
		{
			Name: "module version is less than last_valid_version",
			Content: `
				module "group" {
					source = "git::git@github.com:KazanExpress/ke-infra-modules.git//modules/org/group?ref=v0.100.0"
				}
			`,
			Config: `
rule "terraform_module_version" {
    enabled = true
    last_valid_version = "v0.400.0"
}
            `,
			Expected: helper.Issues{{
				Rule:    rule,
				Message: "Module version is too old. Please update to the latest version",
				Range: hcl.Range{
					Filename: "module.tf",
					Start: hcl.Pos{
						Line:   3,
						Column: 6,
					},
					End: hcl.Pos{
						Line:   3,
						Column: 102,
					},
				},
			}},
		},
		{
			Name: "module version equals last_valid_version",
			Content: `
				module "group" {
					source = "git::git@github.com:KazanExpress/ke-infra-modules.git//modules/org/group?ref=v0.400.0"
				}
			`,
			Config: `
rule "terraform_module_version" {
    enabled = true
    last_valid_version = "v0.400.0"
}
            `,
			Expected: helper.Issues{},
		},
		{
			Name: "module version is not semver",
			Content: `
				module "group" {
					source = "git::git@github.com:KazanExpress/ke-infra-modules.git//modules/org/group?ref=master"
				}
			`,
			Config: `
rule "terraform_module_version" {
    enabled = true
    last_valid_version = "v0.400.0"
}
            `,
			Expected: helper.Issues{},
		},
	}

	for _, tc := range cases {
		t.Run(tc.Name, func(t *testing.T) {
			runner := testRunner(t, map[string]string{"module.tf": tc.Content, ".tflint.hcl": tc.Config})

			if err := rule.Check(runner); err != nil {
				t.Fatalf("Unexpected error occurred: %s", err)
			}

			helper.AssertIssues(t, tc.Expected, runner.Runner.(*helper.Runner).Issues)
		})
	}
}
