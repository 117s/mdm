package utils

import (
	"github.com/117s/mdm/internal/global"
	"path"
	"path/filepath"
	"runtime"
)

func PreTest() {
	_, fp, _, _ := runtime.Caller(0)
	dir := filepath.Dir(fp)
	testConfig := "config_test.yml"
	testConfigPath := path.Join(dir, testConfig)
	global.Init(&testConfigPath)
}
