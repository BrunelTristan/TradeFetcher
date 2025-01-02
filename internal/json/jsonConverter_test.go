package json

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
	bitgetModel "tradeFetcher/model/bitget"
	"tradeFetcher/model/error"
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
		BooleanField:  true,
		IntegerField:  6,
		FloatingField: 0.354,
		StringField:   "wonder",
		ListObjectField: []JsonTesterObject{
			JsonTesterObject{
				IntegerField: 38,
				StringField:  "other",
			},
			JsonTesterObject{
				IntegerField: 796354,
				StringField:  "a",
			},
		},
	})

	assert.Nil(t, err)
	assert.Equal(t, "{\"b\":true,\"i\":6,\"f\":0.354,\"s\":\"wonder\",\"l\":[{\"i\":38,\"s\":\"other\"},{\"i\":796354,\"s\":\"a\"}]}", output)
}

func TestImportBitgetSpotGetFills(t *testing.T) {
	converter := NewJsonConverter[bitgetModel.ApiSpotGetFills]()

	output, err := converter.Import("{\"code\":\"00000\",\"msg\":\"success\",\"requestTime\":1735824191441,\"data\":[{\"userId\":\"123\",\"symbol\":\"XRPUSDT\",\"orderId\":\"659864654\",\"tradeId\":\"32640454\",\"orderType\":\"market\",\"side\":\"sell\",\"priceAvg\":\"2.3723\",\"size\":\"63.7035\",\"amount\":\"151.12381305\",\"feeDetail\":{\"deduction\":\"yes\",\"feeCoin\":\"BGB\",\"totalDeductionFee\":\"0\",\"totalFee\":\"0.0204152398581561\"},\"tradeScope\":\"taker\",\"cTime\":\"61468411\",\"uTime\":\"3568148\"},{\"userId\":\"123\",\"symbol\":\"XRPUSDT\",\"orderId\":\"98765432\",\"tradeId\":\"369871255\",\"orderType\":\"market\",\"side\":\"buy\",\"priceAvg\":\"2.3191\",\"size\":\"63.7035\",\"amount\":\"147.73478685\",\"feeDetail\":{\"deduction\":\"yes\",\"feeCoin\":\"BGB\",\"totalDeductionFee\":\"0\",\"totalFee\":\"0.0198668397175996\"},\"tradeScope\":\"taker\",\"cTime\":\"17569862352\",\"uTime\":\"1654983235\"}]}")

	assert.NotEmpty(t, output)
	assert.Nil(t, err)

	assert.Equal(t, "00000", output.Code)
	assert.Equal(t, "success", output.Message)
	assert.Len(t, output.Data, 2)

	if len(output.Data) == 2 {
		assert.Equal(t, "XRPUSDT", output.Data[0].Symbol)
		assert.Equal(t, "sell", output.Data[0].Side)
		assert.Equal(t, "2.3723", output.Data[0].Price)
		assert.Equal(t, "63.7035", output.Data[0].Size)
		assert.Equal(t, "0.0204152398581561", output.Data[0].FeeDetail.FeesValue)
		assert.Equal(t, "BGB", output.Data[0].FeeDetail.FeeToken)
		assert.Equal(t, "XRPUSDT", output.Data[1].Symbol)
		assert.Equal(t, "buy", output.Data[1].Side)
		assert.Equal(t, "2.3191", output.Data[1].Price)
		assert.Equal(t, "63.7035", output.Data[1].Size)
		assert.Equal(t, "0.0198668397175996", output.Data[1].FeeDetail.FeesValue)
		assert.Equal(t, "BGB", output.Data[1].FeeDetail.FeeToken)
	}
}
