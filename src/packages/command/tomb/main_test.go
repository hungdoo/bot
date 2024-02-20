package tomb

import (
	"os"
	"path"
	"runtime"
	"testing"

	"github.com/hungdoo/bot/src/packages/dotenv"
	"github.com/hungdoo/bot/src/packages/log"
)

func init() {
	_, filename, _, _ := runtime.Caller(0)
	dir := path.Join(path.Dir(filename), "../../../..")
	err := os.Chdir(dir)
	if err != nil {
		panic(err)
	}
	dotenv.InitEnv()
}

func TestMain(m *testing.M) {
	log.GeneralLogger.Println("Init tests!")
	runTests := m.Run()
	log.GeneralLogger.Println("Done tests!")
	os.Exit(runTests)
}
