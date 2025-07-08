package resource_test

import (
	"testing"

	tfjson "github.com/hashicorp/terraform-json"
	"github.com/lonegunmanb/terraform-helm-schema/v3/generated/resource"
	"github.com/stretchr/testify/assert"
)

func TestHelmReleaseSchema(t *testing.T) {
	defaultSchema := &tfjson.Schema{}
	s := resource.HelmReleaseSchema()
	assert.NotNil(t, s)
	assert.NotEqual(t, defaultSchema, s)
}
