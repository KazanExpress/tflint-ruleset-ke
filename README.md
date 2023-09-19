# KazanExpress TFLint Ruleset
[![Build Status](https://github.com/KazanExpress/tflint-ruleset/workflows/build/badge.svg?branch=main)](https://github.com/KazanExpress/tflint-ruleset/actions)

This is tflint ruleset used in KazanExpress team for better terraform code quality.

## Requirements

- TFLint v0.42+
- Go v1.21

## Installation

You can install the plugin with `tflint --init`. Declare a config in `.tflint.hcl` as follows:

```hcl
plugin "ke" {
  enabled = true

  version = "0.3.1"
  source  = "github.com/KazanExpress/tflint-ruleset"

  signing_key = <<EOF
-----BEGIN PGP PUBLIC KEY BLOCK-----

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
/cTNgpPBAgoRDBYUiZaMqg==
=EPIV
-----END PGP PUBLIC KEY BLOCK-----

EOF
}
```

## Rules

|Name|Description|Severity|Enabled|Link|
| --- | --- | --- | --- | --- |
|terraform_module_version| Checks whether terraform module version is not very old|WARNING|✔||

## Building the plugin

Clone the repository locally and run the following command:

```
$ make
```

You can easily install the built plugin with the following:

```
$ make install
```

You can run the built plugin like the following:

```
$ cat << EOS > .tflint.hcl
plugin "template" {
  enabled = true
}
EOS
$ tflint
```
