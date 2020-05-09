package test

import (
	"os"
	"fmt"
	"io/ioutil"
	"encoding/xml"
	"log"
)

// DemoFile demo file
func DemoFile() {
	//fileOps()
	// demoXML()
	demoLog()

}

func fileOps() {
	dirpath := "./files"
	filepath := "./files/test.txt"
	err:= os.MkdirAll(dirpath, os.ModeDir)
	if err != nil {
		fmt.Println("folder created error", err)
		return
	}
	os.Chmod(dirpath, 0777)

	file, err := os.Create(filepath)
	if err != nil {
		fmt.Println("file created error", err)
		return
	}
	fmt.Println("file created success", file.Name())

	fw, err:= os.OpenFile(filepath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	defer fw.Close()
	if err != nil {
		fmt.Println("failed to open file", err)
		return
	}
	filesize,errmsg:=fw.Write([]byte("hello world!\n hello China!\n"))
	if errmsg != nil {
		fmt.Println("failed to write file", errmsg)
		return
	}
	fmt.Println("file write success, file size:", filesize)
	fw.WriteString("hello usa")

	fr, err:= os.Open(filepath) // 只读
	if err != nil {
		fmt.Println("failed to fetch file", err)
		return
	}
	fileInfo, err:= fr.Stat()
	if err != nil {
		fmt.Println("failed to fetch fileinfo", err)
		return
	}
	fmt.Println(fileInfo.Name(), fileInfo.Size(), fileInfo.ModTime(), fileInfo.Mode(), fileInfo.IsDir())

	b:= make([]byte, fileInfo.Size())
	fr.Read(b)
	fmt.Println("file content:", string(b))

	br1,err:=ioutil.ReadAll(fr)
	if err != nil {
		fmt.Println("1: failed to read file", err)
		return
	}
	fmt.Println("read 1: ", string(br1)) // read nothing since already read

	ioutil.WriteFile(filepath, []byte("新写的数据"), 0666) // ! clean data before

	br2,err:=ioutil.ReadFile(filepath)
	if err != nil {
		fmt.Println("2: failed to read file", err)
		return
	}
	fmt.Println("read 2: ", string(br2))

}

// Employees : 
type Employees struct {
	XMLName xml.Name	`xml:"employees"`
	Version string		`xml:"version,attr"`
	Emps []Employee		`xml:"employee"`
}

// Employee : 
type Employee struct {
	XMLName xml.Name	`xml:"employee"`
	Eid int				`xml:"eid,attr"`
	Name string			`xml:"name"`
	Age int				`xml:"age"`
}

func demoXML() {
	fmt.Println("==== start demo xml")
	emps:=new(Employees)
	b,_:=ioutil.ReadFile("./files/test.xml")
	xml.Unmarshal(b, emps)
	fmt.Println(emps)

	emp := Employee{Eid:12, Name:"Tom", Age:30}
	xb, _ := xml.MarshalIndent(emp, "", "	")
	xb = append([]byte(xml.Header), xb...)
	ioutil.WriteFile("./files/employee.xml", xb, 0666)
	fmt.Println("==== end demo xml")
}

func demoLog() {
	log.Println("开始打印日志。。。")
	// log.Panicln("打印panic日志信息")
	// log.Fatal("打印fatal信息，立即终止程序")

	f,_ := os.OpenFile("./files/go.log", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	logger := log.New(f, "[Info]", log.Ltime)
	logger.Println("print log info here")
}