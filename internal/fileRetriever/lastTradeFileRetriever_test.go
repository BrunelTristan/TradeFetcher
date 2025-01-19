package fileRetriever

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestNewLastTradeFileRetriever(t *testing.T) {
	fileReader := NewLastTradeFileRetriever("")

	assert.NotNil(t, fileReader)
}

func TestGetLastProceedTimestampWithoutAnyFile(t *testing.T) {
	fileName := "/nop/it/doesnt/exist/ever.forget"
	_ = os.Remove(fileName)

	fileReader := NewLastTradeFileRetriever(fileName)

	assert.NotNil(t, fileReader)

	timestamp := fileReader.GetLastProceedTimestamp()

	assert.Equal(t, int64(0), timestamp)
}

func TestGetLastProceedTimestampOnEmptyFile(t *testing.T) {
	fileName := "/tmp/test/empty.file"
	err := os.MkdirAll("/tmp/test", 0770)
	assert.Nil(t, err)

	_, err = os.Create(fileName)
	assert.Nil(t, err)

	fileReader := NewLastTradeFileRetriever(fileName)

	timestamp := fileReader.GetLastProceedTimestamp()

	assert.Equal(t, int64(0), timestamp)
}

func TestGetLastProceedTimestampBadContentFile(t *testing.T) {
	fileName := "/tmp/test/another.file"
	previousContent := "something that could not be understood"
	err := os.MkdirAll("/tmp/test", 0770)
	assert.Nil(t, err)

	file, err := os.Create(fileName)
	assert.Nil(t, err)

	_, err = file.Write([]byte(previousContent))
	assert.Nil(t, err)
	err = file.Close()
	assert.Nil(t, err)

	fileReader := NewLastTradeFileRetriever(fileName)

	timestamp := fileReader.GetLastProceedTimestamp()

	assert.Equal(t, int64(0), timestamp)
}

func TestGetLastProceedTimestampOnlyTimestamp(t *testing.T) {
	fileName := "/tmp/test/number.file"
	previousContent := "1234568"
	err := os.MkdirAll("/tmp/test", 0770)
	assert.Nil(t, err)

	file, err := os.Create(fileName)
	assert.Nil(t, err)

	_, err = file.Write([]byte(previousContent))
	assert.Nil(t, err)
	err = file.Close()
	assert.Nil(t, err)

	fileReader := NewLastTradeFileRetriever(fileName)

	timestamp := fileReader.GetLastProceedTimestamp()

	assert.Equal(t, int64(1234568), timestamp)
}

func TestGetLastProceedTimestampOnlyOneTrade(t *testing.T) {
	fileName := "/tmp/test/trade.file"
	previousContent := "1736660716;C;L;AAVEUSDT;0.10000000;289.25000000;0.01735500"
	err := os.MkdirAll("/tmp/test", 0770)
	assert.Nil(t, err)

	file, err := os.Create(fileName)
	assert.Nil(t, err)

	_, err = file.Write([]byte(previousContent))
	assert.Nil(t, err)
	err = file.Close()
	assert.Nil(t, err)

	fileReader := NewLastTradeFileRetriever(fileName)

	timestamp := fileReader.GetLastProceedTimestamp()

	assert.Equal(t, int64(1736660716), timestamp)
}

func TestGetLastProceedTimestampMultipleTimestamp(t *testing.T) {
	fileName := "/tmp/test/numbers.file"
	previousContent := "1234568\n4699754\n032468"
	err := os.MkdirAll("/tmp/test", 0770)
	assert.Nil(t, err)

	file, err := os.Create(fileName)
	assert.Nil(t, err)

	_, err = file.Write([]byte(previousContent))
	assert.Nil(t, err)
	err = file.Close()
	assert.Nil(t, err)

	fileReader := NewLastTradeFileRetriever(fileName)

	timestamp := fileReader.GetLastProceedTimestamp()

	assert.Equal(t, int64(32468), timestamp)
}

func TestGetLastProceedTimestampMultipleTrades(t *testing.T) {
	fileName := "/tmp/test/trades.file"
	previousContent := "1736565011;O;S;AAVEUSDT;0.10000000;281.12000000;0.01686720\n1736571909;O;L;AAVEUSDT;0.10000000;281.89000000;0.01691340\n1736572574;C;L;AAVEUSDT;0.10000000;281.90000000;0.01691400\n1736574010;C;S;AAVEUSDT;0.10000000;283.19000000;0.01699140\n1736580614;O;S;AAVEUSDT;0.10000000;281.15000000;0.01686900\n1736582412;F;;AAVEUSDT;0.00000000;0.00000000;-0.00281387\n"
	err := os.MkdirAll("/tmp/test", 0770)
	assert.Nil(t, err)

	file, err := os.Create(fileName)
	assert.Nil(t, err)

	_, err = file.Write([]byte(previousContent))
	assert.Nil(t, err)
	err = file.Close()
	assert.Nil(t, err)

	fileReader := NewLastTradeFileRetriever(fileName)

	timestamp := fileReader.GetLastProceedTimestamp()

	assert.Equal(t, int64(1736582412), timestamp)
}
