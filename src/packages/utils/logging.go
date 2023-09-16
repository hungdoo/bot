package utils

import (
	"log"
	"os"
)

// GeneralLogger exported
var GeneralLogger *log.Logger

// ErrorLogger exported
var ErrorLogger *log.Logger

func init() {
	// // Log to file
	// absPath, err := filepath.Abs("../../log")
	// if err != nil {
	// 	fmt.Println("Error reading given path:", err)
	// }
	// generalLog, err := os.OpenFile(absPath+"/general-log.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	// if err != nil {
	// 	fmt.Println("Error opening file:", err)
	// 	os.Exit(1)
	// }
	generalLog := os.Stdout
	GeneralLogger = log.New(generalLog, "General:\t", log.Ldate|log.Lshortfile)
	ErrorLogger = log.New(generalLog, "Error:\t", log.Ldate|log.Lshortfile)
}
