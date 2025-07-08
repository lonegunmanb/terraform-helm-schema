package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const helmRelease = `{
  "block": {
    "attributes": {
      "atomic": {
        "computed": true,
        "description": "If set, installation process purges chart on fail. The wait flag will be set automatically if atomic is used",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "chart": {
        "description": "Chart name to be installed. A path may be used",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "cleanup_on_fail": {
        "computed": true,
        "description": "Allow deletion of new resources created in this upgrade when upgrade fails",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "create_namespace": {
        "computed": true,
        "description": "Create the namespace if it does not exist",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "dependency_update": {
        "computed": true,
        "description": "Run helm dependency update before installing the chart",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "description": {
        "computed": true,
        "description": "Add a custom description",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "devel": {
        "description": "Use chart development versions, too. Equivalent to version '\u003e0.0.0-0'. If 'version' is set, this is ignored",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "disable_crd_hooks": {
        "computed": true,
        "description": "Prevent CRD hooks from running, but run other hooks. See helm install --no-crd-hook",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "disable_openapi_validation": {
        "computed": true,
        "description": "If set, the installation process will not validate rendered templates against the Kubernetes OpenAPI Schema",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "disable_webhooks": {
        "computed": true,
        "description": "Prevent hooks from running",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "force_update": {
        "computed": true,
        "description": "Force resource update through delete/recreate if needed.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "keyring": {
        "description": "Location of public keys used for verification, Used only if 'verify is true'",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "lint": {
        "computed": true,
        "description": "Run helm lint when planning",
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
        "computed": true,
        "description": "Limit the maximum number of revisions saved per release. Use 0 for no limit",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "metadata": {
        "computed": true,
        "description": "Status of the deployed release.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "app_version": {
              "computed": true,
              "description": "The version number of the application being deployed",
              "description_kind": "plain",
              "type": "string"
            },
            "chart": {
              "computed": true,
              "description": "The name of the chart",
              "description_kind": "plain",
              "type": "string"
            },
            "first_deployed": {
              "computed": true,
              "description": "FirstDeployed is an int64 which represents timestamp when the release was first deployed.",
              "description_kind": "plain",
              "type": "number"
            },
            "last_deployed": {
              "computed": true,
              "description": "LastDeployed is an int64 which represents timestamp when the release was last deployed.",
              "description_kind": "plain",
              "type": "number"
            },
            "name": {
              "computed": true,
              "description": "Name is the name of the release",
              "description_kind": "plain",
              "type": "string"
            },
            "namespace": {
              "computed": true,
              "description": "Namespace is the kubernetes namespace of the release",
              "description_kind": "plain",
              "type": "string"
            },
            "notes": {
              "computed": true,
              "description": "Notes is the description of the deployed release, rendered from templates.",
              "description_kind": "plain",
              "type": "string"
            },
            "revision": {
              "computed": true,
              "description": "Version is an int32 which represents the version of the release",
              "description_kind": "plain",
              "type": "number"
            },
            "values": {
              "computed": true,
              "description": "Set of extra values. added to the chart. The sensitive data is cloaked. JSON encoded.",
              "description_kind": "plain",
              "type": "string"
            },
            "version": {
              "computed": true,
              "description": "A SemVer 2 conformant version string of the chart",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "name": {
        "description": "Release name. The length must not be longer than 53 characters",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "namespace": {
        "computed": true,
        "description": "Namespace to install the release into",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "pass_credentials": {
        "computed": true,
        "description": "Pass credentials to all domains",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "postrender": {
        "description": "Postrender command config",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "args": {
              "description": "An argument to the post-renderer (can specify multiple)",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "binary_path": {
              "description": "The common binary path",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "recreate_pods": {
        "computed": true,
        "description": "Perform pods restart during upgrade/rollback",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "render_subchart_notes": {
        "computed": true,
        "description": "If set, render subchart notes along with the parent",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "replace": {
        "computed": true,
        "description": "Re-use the given name, even if that name is already used. This is unsafe in production",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "repository": {
        "description": "Repository where to locate the requested chart. If it is a URL, the chart is installed without installing the repository",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "repository_ca_file": {
        "description": "The Repositories CA file",
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
        "computed": true,
        "description": "When upgrading, reset the values to the ones built into the chart",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "reuse_values": {
        "computed": true,
        "description": "When upgrading, reuse the last release's values and merge in any overrides. If 'reset_values' is specified, this is ignored",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "set": {
        "description": "Custom values to be merged with the values",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "type": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "set_list": {
        "description": "Custom sensitive values to be merged with the values",
        "description_kind": "plain",
        "nested_type": {
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
          "nesting_mode": "list"
        },
        "optional": true
      },
      "set_sensitive": {
        "description": "Custom sensitive values to be merged with the values",
        "description_kind": "plain",
        "nested_type": {
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
          "nesting_mode": "list"
        },
        "optional": true
      },
      "set_wo": {
        "description": "Custom values to be merged with the values",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "name": {
              "description_kind": "plain",
              "required": true,
              "type": "string",
              "write_only": true
            },
            "type": {
              "description_kind": "plain",
              "optional": true,
              "type": "string",
              "write_only": true
            },
            "value": {
              "description_kind": "plain",
              "required": true,
              "type": "string",
              "write_only": true
            }
          },
          "nesting_mode": "list"
        },
        "optional": true,
        "write_only": true
      },
      "set_wo_revision": {
        "description": "The current revision of the write-only \"set_wo\" attribute. Incrementing this integer value will cause Terraform to update the write-only value.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "skip_crds": {
        "computed": true,
        "description": "If set, no CRDs will be installed. By default, CRDs are installed if not already present",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "status": {
        "computed": true,
        "description": "Status of the release",
        "description_kind": "plain",
        "type": "string"
      },
      "timeout": {
        "computed": true,
        "description": "Time in seconds to wait for any individual kubernetes operation",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "values": {
        "description": "List of values in raw YAML format to pass to helm",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "verify": {
        "computed": true,
        "description": "Verify the package before installing it.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "version": {
        "computed": true,
        "description": "Specify the exact chart version to install. If this is not specified, the latest version is installed",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "wait": {
        "computed": true,
        "description": "Will wait until all resources are in a ready state before marking the release as successful.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "wait_for_jobs": {
        "computed": true,
        "description": "If wait is enabled, will wait until all Jobs have been completed before marking the release as successful.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      }
    },
    "description": "Schema to define attributes that are available in the resource",
    "description_kind": "plain"
  },
  "version": 2
}`

func HelmReleaseSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(helmRelease), &result)
	return &result
}
