package util

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

const (
	speedOfReadBytePerSecond = 15 // 15 bytes per second
	timeSecondsPerPicture    = 20

	prefixOfPicture = "https://victor-bucket.oss-cn-shanghai.aliyuncs.com"
)

func CharacterLen(content string) int {
	return utf8.RuneCountInString(content)
}

func ReadDurationAnalyze(content string) string {

	t := readDurationAnalyze(content) / 60
	if t == 0 {
		t = 1
	}

	return fmt.Sprintf("约 %d 分钟", t)
}

func readDurationAnalyze(v string) int {

	t1 := algoReadCharacters(len(v))

	t2 := algoReadPictures(strings.Count(v, prefixOfPicture))

	return t1 + t2
}

func algoReadCharacters(length int) int {
	return length / speedOfReadBytePerSecond
}

func algoReadPictures(amount int) int {
	return amount * timeSecondsPerPicture
}
