package main

import (
	"examples/basic"
	"examples/begin"
	"examples/complexn"
	"examples/concurrency"
	"examples/fun"
	"examples/oop"
	"examples/paral"
	phonebook2 "examples/phoneBook2"
	"examples/phoneBook3"
	phonebook4 "examples/phoneBook4"
	phonebook5 "examples/phoneBook5"
	"examples/unix"
	"examples/webservices"
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
	case "timeouts":
		webservices.MtimoutS()
	case "timeoutc":
		webservices.MtimeoutC()
	case "withdeadline":
		webservices.MwithDeadline()
	case "getentries":
		webservices.MgetEntries()
	case "fileserver":
		webservices.MfileServer()
	case "wwwclient":
		webservices.MwwwClient()
	case "simpleclient":
		webservices.MsimpleClient()
	case "prometheus":
		webservices.Mprometheus()
	case "samplepro":
		webservices.MsamplePro()
	case "metrics":
		webservices.Mmetrics()
	case "wwwserver":
		webservices.MwebServices()
	case "freeman":
		concurrency.Mmain()
	case "semaphore":
		paral.Msemaphore()
	case "keyval":
		paral.MkeyVal()
	case "usecontext":
		paral.MuseContext()
	case "goclosurecorrect":
		paral.MgoclosureCorrect()
	case "goclosure":
		paral.MgoClosure()
	case "monitor":
		paral.Mmonitor()
	case "atomic":
		paral.Matomic()
	case "rwmutex":
		paral.MrwMutex()
	case "forgetmutex":
		paral.MforgetMutex()
	case "mmutex":
		paral.Mmutex()
	case "wpools":
		paral.MwPools()
	case "defineorder":
		paral.MdefineOrder()
	case "nilchannel":
		paral.MnilChannel()
	case "bufchannel":
		paral.MbufChannel()
	case "closenil":
		paral.McloseNil()
	case "timeout1":
		paral.Mtimeout1()
	case "timeout2":
		paral.Mtimeout2()
	case "select":
		paral.Mselect()
	case "chrace":
		paral.MchRace()
	case "channelfunc":
		paral.MchannelFunc()
	case "readclosech":
		paral.MreadCloseCh()
	case "channels":
		paral.Mchannels()
	case "randomfiles":
		paral.MrandomFiles()
	case "adddone":
		paral.MaddDone()
	case "vargoroutines":
		paral.Mvargoroutines()
	case "multiple":
		paral.Multiple()
	case "mcreate":
		paral.Mcreate()
	case "maxprocs":
		paral.MaxProcs()
	case "phonebook":
		begin.PhoneBook()
	case "phonebook2":
		phonebook2.PhoneBook()
	case "phonebook3":
		phoneBook3.PhoneBook()
	case "phonebook4":
		phonebook4.Phonebook()
	case "phonebook5":
		phonebook5.MPhoneBook()
	case "jsonviper":
		unix.MjsonViper()
	case "useviper":
		unix.MuseViper()
	case "yaml":
		unix.Myaml()
	case "json2xml":
		unix.Mjson2xml()
	case "xml":
		unix.Mxml()
	case "prettyprint":
		unix.MprettyPrint()
	case "jsonstreams":
		unix.MjsonStreams()
	case "tagsjson":
		unix.MtagsJSON()
	case "encodedecode":
		unix.MencodeDecode()
	case "writefile":
		unix.MwriteFile()
	case "readsize":
		unix.MreadSize()
	case "devrandom":
		unix.MdevRandom()
	case "bychar":
		unix.ByCharacter()
	case "byword":
		unix.ByWord()
	case "byline":
		unix.ByLine()
	case "iointerface":
		unix.MioInterface()
	case "signals":
		unix.Msignals()
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
