// customError
// Written by J.F. Gratton <jean-francois@famillegratton.net>
// Original filename: debug/main.go
// Original timestamp: 2024/04/01 15:56

// ===========================================
// WARNING  WARNING  WARNING  WARNING  WARNING
// ===========================================
// THIS IS FOR DEBUGGING PURPOSES ONLY. ONCE READY TO PUBLISH, COMMENT
// ALL FILES IN THIS DIRECTORY
// ===========================================
// WARNING  WARNING  WARNING  WARNING  WARNING
// ===========================================

//package main

//
//import (
//	"fmt"
//	"github.com/jeanfrancoisgratton/customError"
//	"os"
//	"strings"
//)
//
//func main() {
//	if len(os.Args) == 1 {
//		fmt.Println("No tests are enacted, exiting.")
//		os.Exit(0)
//	}
//	switch strings.ToLower(os.Args[1]) {
//	case "fatal":
//		testFatal()
//	case "warning":
//		testWarning()
//	case "continuable":
//		testContinuable()
//	default:
//		fmt.Println("I do not understand that subcommand, exiting")
//		os.Exit(0)
//	}
//}
//func testFatal() {
//	var ce customError.CustomError
//	customError.ClearTerminal()
//
//	fmt.Println("Fatal with title and code")
//	ce = customError.CustomError{Fatality: customError.Fatal, Title: "FATAL", Message: "My Error message", Code: 222}
//	fmt.Println(ce.Error())
//
//	fmt.Println("Fatal with title, no code")
//	ce = customError.CustomError{Fatality: customError.Fatal, Title: "FATAL", Message: "My Error message"}
//	fmt.Println(ce.Error())
//
//	fmt.Println("Fatal no title, with code")
//	ce = customError.CustomError{Fatality: customError.Fatal, Code: 222, Message: "My Error message"}
//	fmt.Println(ce.Error())
//
//	fmt.Println("Fatal no title, no code")
//	ce = customError.CustomError{Fatality: customError.Fatal, Message: "My Error message"}
//	fmt.Println(ce.Error())
//}
//
//func testWarning() {
//	var ce customError.CustomError
//	customError.ClearTerminal()
//
//	fmt.Println("Warning with title and code")
//	ce = customError.CustomError{Fatality: customError.Warning, Title: "WARNING", Message: "My Error message", Code: 222}
//	fmt.Println(ce.Error())
//
//	fmt.Println("Warning with title, no code")
//	ce = customError.CustomError{Fatality: customError.Warning, Title: "WARNING", Message: "My Error message"}
//	fmt.Println(ce.Error())
//
//	fmt.Println("Warning no title with code")
//	ce = customError.CustomError{Fatality: customError.Warning, Message: "My Error message", Code: 222}
//	fmt.Println(ce.Error())
//
//	fmt.Println("Warning no title, no code")
//	ce = customError.CustomError{Fatality: customError.Warning, Message: "My Error message"}
//	fmt.Println(ce.Error())
//}
//
//func testContinuable() {
//	var ce customError.CustomError
//	customError.ClearTerminal()
//
//	fmt.Println("Continuable with title and code")
//	ce = customError.CustomError{Fatality: customError.Continuable, Title: "FATAL", Message: "My Error message", Code: 222}
//	fmt.Println(ce.Error())
//
//	fmt.Println("Continuable with title, no code")
//	ce = customError.CustomError{Fatality: customError.Continuable, Title: "FATAL", Message: "My Error message"}
//	fmt.Println(ce.Error())
//
//	fmt.Println("Continuable no title, with code")
//	ce = customError.CustomError{Fatality: customError.Continuable, Code: 222, Message: "My Error message"}
//	fmt.Println(ce.Error())
//
//	fmt.Println("Continuable no title, no code")
//	ce = customError.CustomError{Fatality: customError.Continuable, Message: "My Error message"}
//	fmt.Println(ce.Error())
//}
