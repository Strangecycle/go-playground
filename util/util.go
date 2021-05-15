package util

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

// GenCaptcha create a validation code
func GenCaptcha() string {
	numeric := [10]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	width := 6
	r := len(numeric)
	rand.Seed(time.Now().Unix())

	var sb strings.Builder
	for i := 0; i < width; i++ {
		fmt.Fprintf(&sb, "%d", numeric[rand.Intn(r)])
	}
	return sb.String()
}

// GetParentDir 获取当前目录的父级目录
func GetParentDir() string {
	// 当前目录
	wd, _ := os.Getwd()
	idx := strings.LastIndex(wd, "/")
	return wd[:idx]
}
