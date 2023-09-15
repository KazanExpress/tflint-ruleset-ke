package rules

import (
	"strings"

	"github.com/KazanExpress/tflint-ruleset/project"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/terraform-linters/tflint-ruleset-terraform/terraform"
	"golang.org/x/mod/semver"
)

// TerraformModuleVersionTypeRule checks whether terraform module version is not very old.
type TerraformModuleVersionTypeRule struct {
	tflint.DefaultRule
}

// NewTerraformModuleVersionTypeRule returns a new rule
func NewTerraformModuleVersionTypeRule() *TerraformModuleVersionTypeRule {
	return &TerraformModuleVersionTypeRule{}
}

type terraformModuleVersionConfig struct {
	LastValidVersion string `hclext:"last_valid_version,optional"`
}

// Name returns the rule name
func (r *TerraformModuleVersionTypeRule) Name() string {
	return "terraform_module_version"
}

// Enabled returns whether the rule is enabled by default
func (r *TerraformModuleVersionTypeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *TerraformModuleVersionTypeRule) Severity() tflint.Severity {
	return tflint.WARNING
}

// Link returns the rule reference link
func (r *TerraformModuleVersionTypeRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks whether terraform module version is not very old.
func (r *TerraformModuleVersionTypeRule) Check(rr tflint.Runner) error {

	runner := rr.(*terraform.Runner)

	path, err := runner.GetModulePath()
	if err != nil {
		return err
	}
	if !path.IsRoot() {
		// This rule does not evaluate child modules.
		return nil
	}

	config := terraformModuleVersionConfig{}
	config.LastValidVersion = "0.0.0"
	if err := runner.DecodeRuleConfig(r.Name(), &config); err != nil {
		return err
	}

	calls, diags := runner.GetModuleCalls()
	if diags.HasErrors() {
		return diags
	}

	for _, call := range calls {
		if strings.HasPrefix(call.Source, "git::") {
			// find string after ref=
			ref := strings.Split(call.Source, "ref=")[1]
			if semver.IsValid(ref) {
				if semver.Compare(ref, config.LastValidVersion) < 0 {
					runner.EmitIssue(
						r,
						"Module version is too old. Please update to the latest version",
						call.SourceAttr.Range,
					)
				}
			}

		}
	}
	return nil
}
