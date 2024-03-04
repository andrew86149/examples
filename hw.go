package main

import (
	"examples/basic"
	"examples/begin"
	"examples/complexn"
	"examples/fun"
	"examples/oop"
	phonebook2 "examples/phoneBook2"
	"examples/phoneBook3"
	phonebook4 "examples/phoneBook4"
	"fmt"
	"os"
	"strconv"
)

func hw() {
	arguments := os.Args
	ln := len(arguments)
	fmt.Println("len(args) =", ln)
	if ln == 1 {
		fmt.Println("No arguments. Default message: Hello world!")
		return
	}
	argument := arguments[1]
	switch argument {
	case "phonebook":
		begin.PhoneBook()
	case "phonebook2":
		phonebook2.PhoneBook()
	case "phonebook3":
		phoneBook3.PhoneBook()
	case "phonebook4":
		phonebook4.Phonebook()
	case "mpostgo":
		fun.MpostGo()
	case "getschema":
		fun.Mmain()
	case "mdefer":
		fun.Mdefer()
	case "mvariadic":
		fun.Mvariadic()
	case "returnfunction":
		fun.MreturnFunction()
	case "msorting":
		fun.Msorting()
	case "namedreturn":
		fun.MnamedReturn()
	case "mfunctions":
		fun.Mfunctions()
	case "mobj":
		oop.Mobj()
	case "msortcsv":
		oop.Msortcsv()
	case "msortshapes":
		oop.MsortShapes()
	case "mshape2d":
		oop.Mshape2D()
	case "merrorint":
		oop.MerrorInt()
	case "mmapempty":
		oop.MmapEmpty()
	case "massertions":
		oop.Massertions()
	case "mtypeswitch":
		oop.MtypeSwitch()
	case "mempty":
		oop.Mempty()
	case "msort":
		oop.Msort()
	case "methods":
		oop.Methods()
	case "setvalues":
		oop.SetValues()
	case "reflection":
		oop.Reflection()
	case "csvdata":
		complexn.CsvData()
	case "fieldsre":
		complexn.FieldsRE()
	case "intre":
		complexn.IntRE()
	case "namesurre":
		complexn.NameSurRE()
	case "slicestructures":
		complexn.SliceStructures()
	case "structures":
		complexn.Structures()
	case "formaps":
		complexn.ForMaps()
	case "nilmap":
		complexn.NilMap()
	case "cryptorand":
		basic.CryptoRand()
	case "genpass":
		basic.GenPass()
	case "randomnumbers":
		basic.RandomNumbers()
	case "pointers":
		basic.Pointers()
	case "sortslice":
		basic.SortSlice()
	case "copyslice":
		basic.CopySlice()
	case "slicearrays":
		basic.SliceArrays()
	case "byteslices":
		basic.ByteSlices()
	case "deleteslice":
		basic.DeleteSlice()
	case "partslice":
		basic.PartSlice()
	case "caplen":
		basic.CapLen()
	case "goslices":
		basic.GoSlices()
	case "constants":
		basic.Constants()
	case "converttimes":
		basic.ConvertTimes()
	case "dates":
		basic.Dates()
	case "usestring":
		basic.UseString()
	case "municode":
		basic.Municode()
	case "intstring":
		basic.IntString()
	case "mtext":
		basic.Mtext()
	case "numbers":
		basic.Numbers()
	case "merror":
		basic.Merror()
	case "generics":
		begin.Generics()
	case "customloglinenumber":
		begin.CustomLogLineNumber()
	case "customlog":
		begin.CustomLog()
	case "logs":
		begin.Logs()
	case "systemlog":
		begin.SystemLog()
		// journalctl -xe
	case "which":
		begin.Which()
	case "goroutines":
		begin.GoRoutines()
	case "process":
		begin.Process()
	case "cla":
		begin.Cla()
	case "input":
		begin.Input()
	case "forloops":
		begin.Forloops()
	case "variables":
		begin.Variables()
	case "fib":
		if ln == 3 {
			n, err := strconv.Atoi(arguments[2])
			if err == nil {
				fmt.Printf("fib(%d)=%d\n", n, fib(n))
			}
		} else {
			fmt.Println("fib(10) = ", fib(10))
		}
	default:
		fmt.Println("selected default...")
	}
}
func fib(n int) int {
	if n < 2 {
		return 1
	} else {
		return fib(n-1) + fib(n-2)
	}
}
