package generated

import (
	tfjson "github.com/hashicorp/terraform-json"
	"github.com/lonegunmanb/terraform-helm-schema/v2/generated/data"
	"github.com/lonegunmanb/terraform-helm-schema/v2/generated/resource"
)

var Resources map[string]*tfjson.Schema
var DataSources map[string]*tfjson.Schema

func init() {
	resources := make(map[string]*tfjson.Schema, 0)
	dataSources := make(map[string]*tfjson.Schema, 0)  
	resources["helm_release"] = resource.HelmReleaseSchema()  
	dataSources["helm_template"] = data.HelmTemplateSchema()  
	Resources = resources
	DataSources = dataSources
}