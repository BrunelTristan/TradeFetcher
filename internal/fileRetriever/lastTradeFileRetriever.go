package fileRetriever

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"tradeFetcher/internal/processUnit"
)

type LastTradeFileRetriever struct {
	filePath string
}

func NewLastTradeFileRetriever(fPath string) processUnit.ILastProceedRetriever {
	return &LastTradeFileRetriever{
		filePath: fPath,
	}
}

func (r *LastTradeFileRetriever) GetLastProceedTimestamp() int64 {
	file, err := os.Open(r.filePath)
	if err != nil {
		return 0
	}
	defer file.Close()

	line := r.getLastLine(file)
	timestamp, err := strconv.ParseInt(strings.Split(line, ";")[0], 10, 64)
	if err != nil {
		return 0
	}

	return timestamp
}

func (r *LastTradeFileRetriever) getLastLine(file *os.File) string {
	stat, _ := file.Stat()
	filesize := stat.Size()
	if filesize == 0 {
		return ""
	}

	line := ""
	char := make([]byte, 1)

	for cursor := int64(-1); cursor >= -filesize; cursor-- {
		_, _ = file.Seek(cursor, io.SeekEnd)
		_, _ = file.Read(char)

		if cursor != -1 && (char[0] == 10 || char[0] == 13) {
			break
		}

		line = fmt.Sprintf("%s%s", string(char), line)
	}

	return line
}
