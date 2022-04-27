package main

import (
	"fmt"
	"github.com/nicolachoquet06250/golang-tuto/arraysSlices"
	"github.com/nicolachoquet06250/golang-tuto/channels"
	"github.com/nicolachoquet06250/golang-tuto/constants"
	"github.com/nicolachoquet06250/golang-tuto/customErrors"
	_defer "github.com/nicolachoquet06250/golang-tuto/defer"
	"github.com/nicolachoquet06250/golang-tuto/errorHandling"
	"github.com/nicolachoquet06250/golang-tuto/functions"
	"github.com/nicolachoquet06250/golang-tuto/goroutines"
	"github.com/nicolachoquet06250/golang-tuto/helloWorld"
	"github.com/nicolachoquet06250/golang-tuto/ifElse"
	"github.com/nicolachoquet06250/golang-tuto/inheritence"
	"github.com/nicolachoquet06250/golang-tuto/interfaces"
	"github.com/nicolachoquet06250/golang-tuto/loops"
	"github.com/nicolachoquet06250/golang-tuto/maps"
	"github.com/nicolachoquet06250/golang-tuto/methods"
	"github.com/nicolachoquet06250/golang-tuto/mutexs"
	"github.com/nicolachoquet06250/golang-tuto/panicFunction"
	"github.com/nicolachoquet06250/golang-tuto/pointers"
	"github.com/nicolachoquet06250/golang-tuto/reflection"
	"github.com/nicolachoquet06250/golang-tuto/selects"
	"github.com/nicolachoquet06250/golang-tuto/strings"
	"github.com/nicolachoquet06250/golang-tuto/structs"
	"github.com/nicolachoquet06250/golang-tuto/structsClasses"
	"github.com/nicolachoquet06250/golang-tuto/switchCase"
	"github.com/nicolachoquet06250/golang-tuto/types"
	"github.com/nicolachoquet06250/golang-tuto/variables"
	"github.com/nicolachoquet06250/golang-tuto/variadicFunctions"
)

func main() {
	fmt.Println("\n--------- Hello World ---------")
	helloWorld.Main()

	fmt.Println("\n---------- Variables ----------")
	variables.Main()

	fmt.Println("\n------------ Types ------------")
	types.Main()

	fmt.Println("\n---------- Constants ----------")
	constants.Main()

	fmt.Println("\n---------- Functions ----------")
	functions.Main()

	fmt.Println("\n----------- If Else -----------")
	ifElse.Main()

	fmt.Println("\n------------ Loops ------------")
	loops.Main()

	fmt.Println("\n--------- Switch Case ---------")
	switchCase.Main()

	fmt.Println("\n------- Arrays & Slices -------")
	arraysSlices.Main()

	fmt.Println("\n------ Variadic Functions -----")
	variadicFunctions.Main()

	fmt.Println("\n------------- Maps ------------")
	maps.Main()

	fmt.Println("\n------------ Strings ----------")
	strings.Main()

	fmt.Println("\n----------- Pointers ----------")
	pointers.Main()

	fmt.Println("\n----------- Structs -----------")
	structs.Main()

	fmt.Println("\n----------- Methods -----------")
	methods.Main()

	fmt.Println("\n--------- Interfaces ----------")
	interfaces.Main()

	fmt.Println("\n--------- Go Routines ---------")
	goroutines.Main()

	fmt.Println("\n---------- Channels -----------")
	channels.Main()

	fmt.Println("\n----------- Select ------------")
	selects.Main()

	fmt.Println("\n----------- Mutexs ------------")
	mutexs.Main()

	fmt.Println("\n----- Structs et classes ------")
	structsClasses.Main()

	fmt.Println("\n--------- Inheritance ---------")
	inheritence.Main()

	fmt.Println("\n------------ Defer ------------")
	_defer.Main()

	fmt.Println("\n------- Error Handling --------")
	errorHandling.Main()

	fmt.Println("\n-------- Custom errors --------")
	customErrors.Main()

	fmt.Println("\n------- Panic function --------")
	panicFunction.Main()

	fmt.Println("\n--------- Reflection ----------")
	reflection.Main()
}
