package data_test

import (
	"testing"

	tfjson "github.com/hashicorp/terraform-json"
	"github.com/lonegunmanb/terraform-helm-schema/v3/generated/data"
	"github.com/stretchr/testify/assert"
)

func TestHelmTemplateSchema(t *testing.T) {
	defaultSchema := &tfjson.Schema{}
	s := data.HelmTemplateSchema()
	assert.NotNil(t, s)
	assert.NotEqual(t, defaultSchema, s)
}
