package processUnit

import (
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"os"
	"testing"
	"tradeFetcher/internal/generatedMocks"
	"tradeFetcher/model/trading"
)

func TestNewTradeFileSaver(t *testing.T) {
	fileSaver := NewTradeFileSaver(nil, "")

	assert.NotNil(t, fileSaver)
}

func TestProcessTradesFileSaverOnEmptySliceWithUnexistingFile(t *testing.T) {
	mockController := gomock.NewController(t)

	fileName := "/nop/never/ever.tkt"
	_ = os.Remove(fileName)

	tradeFormatterMock := generatedMocks.NewMockITradeFormatter(mockController)

	tradeFormatterMock.
		EXPECT().
		Format(gomock.Any()).
		Times(0)

	fileSaver := NewTradeFileSaver(tradeFormatterMock, fileName)

	assert.NotNil(t, fileSaver)

	trades := []*trading.Trade{}

	err := fileSaver.ProcessTrades(trades)
	assert.Nil(t, err)

	_, err = os.Stat(fileName)

	assert.NotNil(t, err)
	if err != nil {
		assert.True(t, os.IsNotExist(err))
	}
}

func TestProcessTradesFileSaverOnEmptySliceWithExistingFile(t *testing.T) {
	mockController := gomock.NewController(t)

	fileName := "/tmp/nop/but/whynot.perhaps"
	err := os.MkdirAll("/tmp/nop/but", 0770)
	assert.Nil(t, err)

	_, err = os.Create(fileName)
	assert.Nil(t, err)

	initFileInfo, err := os.Stat(fileName)
	assert.Nil(t, err)

	tradeFormatterMock := generatedMocks.NewMockITradeFormatter(mockController)

	tradeFormatterMock.
		EXPECT().
		Format(gomock.Any()).
		Times(0)

	fileSaver := NewTradeFileSaver(tradeFormatterMock, fileName)

	trades := []*trading.Trade{}

	err = fileSaver.ProcessTrades(trades)
	assert.Nil(t, err)

	newFileInfo, err := os.Stat(fileName)

	assert.Nil(t, err)
	assert.Equal(t, initFileInfo, newFileInfo)
}

func TestProcessTradesFileSaverWithValuesButUnexistingFile(t *testing.T) {
	mockController := gomock.NewController(t)

	content := "A trade but which one ?"
	fileName := "/tmp/test/shouldnot.exist"
	_ = os.Remove(fileName)

	tradeFormatterMock := generatedMocks.NewMockITradeFormatter(mockController)

	tradeFormatterMock.
		EXPECT().
		Format(gomock.Any()).
		Return(content).
		Times(2)

	fileSaver := NewTradeFileSaver(tradeFormatterMock, fileName)

	assert.NotNil(t, fileSaver)

	trades := []*trading.Trade{
		&trading.Trade{},
		&trading.Trade{},
	}

	err := fileSaver.ProcessTrades(trades)
	assert.Nil(t, err)

	newFileInfo, err := os.Stat(fileName)

	assert.Nil(t, err)
	assert.NotNil(t, newFileInfo)

	if newFileInfo != nil {
		assert.Equal(t, int64(2*(len(content)+1)), newFileInfo.Size())
	}
}

func TestProcessTradesFileSaverWithValuesButNotRealFile(t *testing.T) {
	mockController := gomock.NewController(t)

	fileName := "/tmp/test/folder/"

	tradeFormatterMock := generatedMocks.NewMockITradeFormatter(mockController)

	tradeFormatterMock.
		EXPECT().
		Format(gomock.Any()).
		Return("").
		Times(0)

	fileSaver := NewTradeFileSaver(tradeFormatterMock, fileName)

	assert.NotNil(t, fileSaver)

	trades := []*trading.Trade{
		&trading.Trade{},
		&trading.Trade{},
	}

	err := fileSaver.ProcessTrades(trades)
	assert.NotNil(t, err)
}

func TestProcessTradesFileSaverWithValuesButNoRightInFolder(t *testing.T) {
	mockController := gomock.NewController(t)

	fileName := "/dev/null/notEnoughRight.lst"

	tradeFormatterMock := generatedMocks.NewMockITradeFormatter(mockController)

	tradeFormatterMock.
		EXPECT().
		Format(gomock.Any()).
		Return("").
		Times(0)

	fileSaver := NewTradeFileSaver(tradeFormatterMock, fileName)

	assert.NotNil(t, fileSaver)

	trades := []*trading.Trade{
		&trading.Trade{},
		&trading.Trade{},
	}

	err := fileSaver.ProcessTrades(trades)
	assert.NotNil(t, err)
}

func TestProcessTradesFileSaverWithValuesOnExistingFile(t *testing.T) {
	mockController := gomock.NewController(t)

	previousContent := "Some trade\nWith other trades\n and so on\n et cetera\n"
	content := "A trade but which one ?"
	fileName := "/tmp/test/saveFile.pwd"
	err := os.MkdirAll("/tmp/test", 0770)
	assert.Nil(t, err)

	file, err := os.Create(fileName)
	assert.Nil(t, err)

	_, err = file.Write([]byte(previousContent))
	assert.Nil(t, err)
	err = file.Close()
	assert.Nil(t, err)

	tradeFormatterMock := generatedMocks.NewMockITradeFormatter(mockController)

	tradeFormatterMock.
		EXPECT().
		Format(gomock.Any()).
		Return(content).
		Times(1)

	fileSaver := NewTradeFileSaver(tradeFormatterMock, fileName)

	assert.NotNil(t, fileSaver)

	trades := []*trading.Trade{
		&trading.Trade{},
	}

	err = fileSaver.ProcessTrades(trades)
	assert.Nil(t, err)

	newFileInfo, err := os.Stat(fileName)

	assert.Nil(t, err)
	assert.NotNil(t, newFileInfo)

	if newFileInfo != nil {
		assert.Equal(t, int64(len(previousContent)+len(content)+1), newFileInfo.Size())
	}
}
