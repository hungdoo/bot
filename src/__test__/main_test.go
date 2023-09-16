package test

import (
	"os"
	"path"
	"runtime"
	"testing"

	u "github.com/hungdoo/bot/src/packages/utils"
)

func init() {
	_, filename, _, _ := runtime.Caller(0)
	dir := path.Join(path.Dir(filename), "../..")
	err := os.Chdir(dir)
	if err != nil {
		panic(err)
	}
	u.InitEnv()
}

func TestMain(m *testing.M) {
	u.GeneralLogger.Println("Init tests!")
	runTests := m.Run()
	u.GeneralLogger.Println("Done tests!")
	os.Exit(runTests)
}
