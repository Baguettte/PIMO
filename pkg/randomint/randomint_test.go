package randomint

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"makeit.imfr.cgi.com/makeit2/scm/lino/pimo/pkg/model"
)

func TestMaskingShouldReplaceSensitiveValueByRandomNumber(t *testing.T) {
	min := 7
	max := 77
	ageMask := NewMask(min, max, 0)
	config := model.NewMaskConfiguration().
		WithEntry("age", ageMask)

	maskingEngine := model.MaskingEngineFactory(config)
	data := model.Dictionary{"age": 83}
	result, err := maskingEngine.Mask(data)
	assert.Equal(t, nil, err, "error should be nil")
	resultmap := result.(map[string]model.Entry)

	assert.NotEqual(t, data, result, "Should be masked")
	assert.True(t, resultmap["age"].(int) >= min, "Should be more than min")
	assert.True(t, resultmap["age"].(int) <= max, "Should be less than max")
}

func TestNewMaskFromConfigShouldCreateAMask(t *testing.T) {
	maskingConfig := model.Masking{Mask: model.MaskType{RandomInt: model.RandIntType{Min: 18, Max: 25}}}
	mask, present, err := NewMaskFromConfig(maskingConfig, 0)
	intMask := mask.(MaskEngine)
	assert.Equal(t, 18, intMask.min, "should be equal")
	assert.Equal(t, 25, intMask.max, "datemax should be te same")
	assert.True(t, present, "should be true")
	assert.Nil(t, err, "error should be nil")
}

func TestNewMaskFromConfigShouldNotCreateAMaskFromAnEmptyConfig(t *testing.T) {
	maskingConfig := model.Masking{Mask: model.MaskType{}}
	mask, present, err := NewMaskFromConfig(maskingConfig, 0)
	assert.Nil(t, mask, "should be nil")
	assert.False(t, present, "should be false")
	assert.Nil(t, err, "error should be nil")
}