package json

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
	"tradeFetcher/internal/error"
)

type JsonTesterObject struct {
	BooleanField    bool               `json:"b"`
	IntegerField    int                `json:"i"`
	FloatingField   float64            `json:"f"`
	StringField     string             `json:"s"`
	ListObjectField []JsonTesterObject `json:"l"`
}

func TestNewJsonConverter(t *testing.T) {
	object := NewJsonConverter[JsonTesterObject]()

	assert.NotNil(t, object)
}

func TestImportWithBadJsonString(t *testing.T) {
	converter := NewJsonConverter[JsonTesterObject]()

	output, err := converter.Import("not a json")

	assert.Nil(t, output)
	assert.NotNil(t, err)

	assert.True(t, errors.As(err, new(*error.JsonError)))
	assert.Equal(t, "Error on json operation : Bad json format", err.Error())
}

/* Not manage by default encoding/json unmarshal
func TestImportWithMandatoryFieldNotPresent(t *testing.T) {
	converter := NewJsonConverter[JsonTesterObject]()

	output, err := converter.Import("{\"b\":true,\"f\":0.256}")

	assert.Nil(t, output)
	assert.NotNil(t, err)

	assert.True(t, errors.As(err, new(*error.JsonError)))
	assert.Equal(t, "Error on json operation : Missing mandatory field", err.Error())
}*/

func TestImportWithBadFieldType(t *testing.T) {
	converter := NewJsonConverter[JsonTesterObject]()

	output, err := converter.Import("{\"b\":true,\"f\":0.256,\"i\":\"nor an integer\"}")

	assert.Nil(t, output)
	assert.NotNil(t, err)

	assert.True(t, errors.As(err, new(*error.JsonError)))
	assert.Equal(t, "Error on json operation : Bad field format", err.Error())
}

func TestImportWithBadObjectType(t *testing.T) {
	converter := NewJsonConverter[error.RestApiError]()

	output, err := converter.Import("{\"b\":true,\"f\":0.256,\"i\":\"nor an integer\"}")

	assert.Empty(t, output)
	assert.Nil(t, err)
}

func TestImportWithGoodData(t *testing.T) {
	converter := NewJsonConverter[JsonTesterObject]()

	output, err := converter.Import("{\"b\":true,\"f\":0.256,\"l\":[{},{},{}]}")

	assert.NotEmpty(t, output)
	assert.Nil(t, err)

	assert.True(t, output.BooleanField)
	assert.Equal(t, 0.256, output.FloatingField)
	assert.Len(t, output.ListObjectField, 3)
}

func TestExportNil(t *testing.T) {
	// TODO
}

func TestExportEmptyObject(t *testing.T) {
	// TODO
}

func TestExportBadTypeObject(t *testing.T) {
	// TODO
}

func TestExportRealObject(t *testing.T) {
	// TODO
}
