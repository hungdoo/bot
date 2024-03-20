package test

import (
	"context"
	"os"
	"path"
	"runtime"
	"testing"

	"github.com/hungdoo/bot/src/packages/db"
	"github.com/hungdoo/bot/src/packages/dotenv"
	"github.com/hungdoo/bot/src/packages/log"
)

func init() {
	_, filename, _, _ := runtime.Caller(0)
	dir := path.Join(path.Dir(filename), "../..")
	err := os.Chdir(dir)
	if err != nil {
		panic(err)
	}
	dotenv.InitEnv()
}

func TestMain(m *testing.M) {
	_db := db.GetDb()
	log.GeneralLogger.Printf("Test db init. IsTestEnv: %v\n", _db.IsTestEnv)

	defer func() {
		if _db.IsTestEnv {
			log.GeneralLogger.Println("Clean test db")
			_db.GetCollection("commands").Drop(context.Background())
		}
		_db.Close()
	}()

	log.GeneralLogger.Println("Init tests!")
	m.Run()
	// runTests := m.Run()
	log.GeneralLogger.Println("Done tests!")
	// os.Exit(runTests)
}
