package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const helmRelease = `{
  "block": {
    "attributes": {
      "atomic": {
        "description": "If set, installation process purges chart on fail. The wait flag will be set automatically if atomic is used. Defaults to ` + "`" + `false` + "`" + `.",
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
      "cleanup_on_fail": {
        "description": "Allow deletion of new resources created in this upgrade when upgrade fails. Defaults to ` + "`" + `false` + "`" + `.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "create_namespace": {
        "description": "Create the namespace if it does not exist. Defaults to ` + "`" + `false` + "`" + `.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "dependency_update": {
        "description": "Run helm dependency update before installing the chart. Defaults to ` + "`" + `false` + "`" + `.",
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
      "disable_crd_hooks": {
        "description": "Prevent CRD hooks from, running, but run other hooks.  See helm install --no-crd-hook",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "disable_openapi_validation": {
        "description": "If set, the installation process will not validate rendered templates against the Kubernetes OpenAPI Schema. Defaults to ` + "`" + `false` + "`" + `.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "disable_webhooks": {
        "description": "Prevent hooks from running.Defaults to ` + "`" + `false` + "`" + `.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "force_update": {
        "description": "Force resource update through delete/recreate if needed. Defaults to ` + "`" + `false` + "`" + `.",
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
      "keyring": {
        "description": "Location of public keys used for verification. Used only if ` + "`" + `verify` + "`" + ` is true. Defaults to ` + "`" + `/.gnupg/pubring.gpg` + "`" + ` in the location set by ` + "`" + `home` + "`" + `.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "lint": {
        "description": "Run helm lint when planning. Defaults to ` + "`" + `false` + "`" + `.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "manifest": {
        "computed": true,
        "description": "The rendered manifest as JSON.",
        "description_kind": "plain",
        "type": "string"
      },
      "max_history": {
        "description": "Limit the maximum number of revisions saved per release. Use 0 for no limit. Defaults to 0 (no limit).",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "metadata": {
        "computed": true,
        "description": "Status of the deployed release.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "app_version": "string",
              "chart": "string",
              "first_deployed": "number",
              "last_deployed": "number",
              "name": "string",
              "namespace": "string",
              "notes": "string",
              "revision": "number",
              "values": "string",
              "version": "string"
            }
          ]
        ]
      },
      "name": {
        "description": "Release name. The length must not be longer than 53 characters.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "namespace": {
        "description": "Namespace to install the release into. Defaults to ` + "`" + `default` + "`" + `.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "pass_credentials": {
        "description": "Pass credentials to all domains. Defaults to ` + "`" + `false` + "`" + `.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "recreate_pods": {
        "description": "Perform pods restart during upgrade/rollback. Defaults to ` + "`" + `false` + "`" + `.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "render_subchart_notes": {
        "description": "If set, render subchart notes along with the parent. Defaults to ` + "`" + `true` + "`" + `.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "replace": {
        "description": "Re-use the given name, even if that name is already used. This is unsafe in production. Defaults to ` + "`" + `false` + "`" + `.",
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
        "description": "When upgrading, reset the values to the ones built into the chart. Defaults to ` + "`" + `false` + "`" + `.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "reuse_values": {
        "description": "When upgrading, reuse the last release's values and merge in any overrides. If 'reset_values' is specified, this is ignored. Defaults to ` + "`" + `false` + "`" + `.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "skip_crds": {
        "description": "If set, no CRDs will be installed. By default, CRDs are installed if not already present. Defaults to ` + "`" + `false` + "`" + `.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "status": {
        "computed": true,
        "description": "Status of the release.",
        "description_kind": "plain",
        "type": "string"
      },
      "timeout": {
        "description": "Time in seconds to wait for any individual kubernetes operation. Defaults to 300 seconds.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "upgrade_install": {
        "description": "If true, the provider will install the release at the specified version even if a release not controlled by the provider is present: this is equivalent to running 'helm upgrade --install' with the Helm CLI. WARNING: this may not be suitable for production use -- see the 'Upgrade Mode' note in the provider documentation. Defaults to ` + "`" + `false` + "`" + `.",
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
        "description": "Verify the package before installing it.Defaults to ` + "`" + `false` + "`" + `.",
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
        "description": "Will wait until all resources are in a ready state before marking the release as successful. Defaults to ` + "`" + `true` + "`" + `.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "wait_for_jobs": {
        "description": "If wait is enabled, will wait until all Jobs have been completed before marking the release as successful. Defaults to ` + "`" + `false` + "`" + `` + "`" + `.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      }
    },
    "block_types": {
      "postrender": {
        "block": {
          "attributes": {
            "args": {
              "description": "an argument to the post-renderer (can specify multiple)",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
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
          "description": "Custom list values to be merged with the values.",
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
      }
    },
    "description_kind": "plain"
  },
  "version": 1
}`

func HelmReleaseSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(helmRelease), &result)
	return &result
}
