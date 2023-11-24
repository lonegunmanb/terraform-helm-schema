package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const helmTemplate = `{
  "block": {
    "attributes": {
      "api_versions": {
        "description": "Kubernetes api versions used for Capabilities.APIVersions",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "atomic": {
        "description": "If set, installation process purges chart on fail. The wait flag will be set automatically if atomic is used",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "chart": {
        "description": "Chart name to be installed. A path may be used.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "crds": {
        "computed": true,
        "description": "List of rendered CRDs from the chart.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "create_namespace": {
        "description": "Create the namespace if it does not exist",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "dependency_update": {
        "description": "Run helm dependency update before installing the chart",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "description": {
        "description": "Add a custom description",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "devel": {
        "description": "Use chart development versions, too. Equivalent to version '\u003e0.0.0-0'. If ` + "`" + `version` + "`" + ` is set, this is ignored",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "disable_openapi_validation": {
        "description": "If set, the installation process will not validate rendered templates against the Kubernetes OpenAPI Schema",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "disable_webhooks": {
        "description": "Prevent hooks from running.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "include_crds": {
        "description": "Include CRDs in the templated output",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "is_upgrade": {
        "description": "Set .Release.IsUpgrade instead of .Release.IsInstall",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "keyring": {
        "description": "Location of public keys used for verification. Used only if ` + "`" + `verify` + "`" + ` is true",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "kube_version": {
        "description": "Kubernetes version used for Capabilities.KubeVersion",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "manifest": {
        "computed": true,
        "description": "Concatenated rendered chart templates. This corresponds to the output of the ` + "`" + `helm template` + "`" + ` command.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "manifests": {
        "computed": true,
        "description": "Map of rendered chart templates indexed by the template name.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "name": {
        "description": "Release name.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "namespace": {
        "description": "Namespace to install the release into.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "notes": {
        "computed": true,
        "description": "Rendered notes if the chart contains a ` + "`" + `NOTES.txt` + "`" + `.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "pass_credentials": {
        "description": "Pass credentials to all domains",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "render_subchart_notes": {
        "description": "If set, render subchart notes along with the parent",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "replace": {
        "description": "Re-use the given name, even if that name is already used. This is unsafe in production",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "repository": {
        "description": "Repository where to locate the requested chart. If is a URL the chart is installed without installing the repository.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "repository_ca_file": {
        "description": "The Repositories CA File",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "repository_cert_file": {
        "description": "The repositories cert file",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "repository_key_file": {
        "description": "The repositories cert key file",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "repository_password": {
        "description": "Password for HTTP basic authentication",
        "description_kind": "plain",
        "optional": true,
        "sensitive": true,
        "type": "string"
      },
      "repository_username": {
        "description": "Username for HTTP basic authentication",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "reset_values": {
        "description": "When upgrading, reset the values to the ones built into the chart",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "reuse_values": {
        "description": "When upgrading, reuse the last release's values and merge in any overrides. If 'reset_values' is specified, this is ignored",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "show_only": {
        "description": "Only show manifests rendered from the given templates",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "skip_crds": {
        "description": "If set, no CRDs will be installed. By default, CRDs are installed if not already present",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "skip_tests": {
        "description": "If set, tests will not be rendered. By default, tests are rendered",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "timeout": {
        "description": "Time in seconds to wait for any individual kubernetes operation.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "validate": {
        "description": "Validate your manifests against the Kubernetes cluster you are currently pointing at. This is the same validation performed on an install",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "values": {
        "description": "List of values in raw yaml format to pass to helm.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "verify": {
        "description": "Verify the package before installing it.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "version": {
        "computed": true,
        "description": "Specify the exact chart version to install. If this is not specified, the latest version is installed.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "wait": {
        "description": "Will wait until all resources are in a ready state before marking the release as successful.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      }
    },
    "block_types": {
      "postrender": {
        "block": {
          "attributes": {
            "binary_path": {
              "description": "The command binary path.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description": "Postrender command configuration.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "set": {
        "block": {
          "attributes": {
            "name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "type": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description": "Custom values to be merged with the values.",
          "description_kind": "plain"
        },
        "nesting_mode": "set"
      },
      "set_list": {
        "block": {
          "attributes": {
            "name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "value": {
              "description_kind": "plain",
              "required": true,
              "type": [
                "list",
                "string"
              ]
            }
          },
          "description": "Custom sensitive values to be merged with the values.",
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "set_sensitive": {
        "block": {
          "attributes": {
            "name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "type": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "description_kind": "plain",
              "required": true,
              "sensitive": true,
              "type": "string"
            }
          },
          "description": "Custom sensitive values to be merged with the values.",
          "description_kind": "plain"
        },
        "nesting_mode": "set"
      },
      "set_string": {
        "block": {
          "attributes": {
            "name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "value": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "deprecated": true,
          "description": "Custom string values to be merged with the values.",
          "description_kind": "plain"
        },
        "nesting_mode": "set"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func HelmTemplateSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(helmTemplate), &result)
	return &result
}
