package json

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
	"tradeFetcher/internal/error"
)

type JsonTesterObject struct {
	BooleanField    bool               `json:"b,omitempty"`
	IntegerField    int                `json:"i"`
	FloatingField   float64            `json:"f,omitempty"`
	StringField     string             `json:"s"`
	ListObjectField []JsonTesterObject `json:"l,omitempty"`
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
	converter := NewJsonConverter[JsonTesterObject]()

	output, err := converter.Export(nil)

	assert.Nil(t, err)
	assert.Equal(t, "", output)
}

func TestExportEmptyObject(t *testing.T) {
	converter := NewJsonConverter[JsonTesterObject]()

	output, err := converter.Export(&JsonTesterObject{})

	assert.Nil(t, err)
	assert.Equal(t, "{\"i\":0,\"s\":\"\"}", output)
}

/* Manage by compiler : you can't pass an object with the bad type
func TestExportBadTypeObject(t *testing.T) {
	converter := NewJsonConverter[JsonTesterObject]()

	output, err := converter.Export(&error.RestApiError{})

	assert.NotNil(t, err)
	assert.Equal(t, "", output)
	
	assert.True(t, errors.As(err, new(*error.JsonError)))
	assert.Equal(t, "Error on json operation : Bad type object", err.Error())
}*/

func TestExportRealObject(t *testing.T) {
	converter := NewJsonConverter[JsonTesterObject]()

	output, err := converter.Export(&JsonTesterObject{
		BooleanField : true,
		IntegerField : 6,
		FloatingField : 0.354,
		StringField : "wonder",
		ListObjectField : []JsonTesterObject{
			JsonTesterObject{
				IntegerField: 38,
				StringField: "other",
			},
			JsonTesterObject{
				IntegerField: 796354,
				StringField: "a",
			},
		},
	})

	assert.Nil(t, err)
	assert.Equal(t, "{\"b\":true,\"i\":6,\"f\":0.354,\"s\":\"wonder\",\"l\":[{\"i\":38,\"s\":\"other\"},{\"i\":796354,\"s\":\"a\"}]}", output)
}
