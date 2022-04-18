package utils

import (
	"github.com/117s/mdm/internal/initialize"
	"github.com/rs/xid"
	"path"
	"path/filepath"
	"runtime"
)

func PreTest() {
	_, fp, _, _ := runtime.Caller(0)
	dir := filepath.Dir(fp)
	testConfig := "config_test.yml"
	testConfigPath := path.Join(dir, testConfig)
	initialize.Init(&testConfigPath)
}

func NewID() string {
	guid := xid.New()
	return guid.String()
}
