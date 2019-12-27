package main

import (
	"flag"
	"log"
	"os"
	"strings"
)

func main() {
	var dirname string
	flag.StringVar(&dirname,"d","","")
	flag.Parse()
	if dirname == "" {
		log.Fatal("dir path error")
	}
	err := os.Mkdir(dirname,os.ModePerm)
	if err != nil {
		log.Fatal("mkdir error :", err)
	}
	file_name := strings.Split(dirname,"_")[1]
	file, err := os.Create(dirname + string( os.PathSeparator) + file_name + ".go")
	if err != nil {
		log.Fatal("create file error :", err)
	}
	_, err = file.WriteString("package " + "_" + dirname[1:] + "\n\n")
	file.Close()
	if err != nil {
		log.Fatal("write file error :", err)
	}

	file , err = os.Create(dirname + string( os.PathSeparator) + file_name + "_test.go")
	if err != nil {
		log.Fatal("create file error :", err)
	}
	_, err = file.WriteString("package " + "_" + dirname[1:] + "\n\n" + "import \"testing\"\n\n\nfunc Test" +  strings.Title(file_name) + "(t *testing.T) {\n\n}")
	file.Close()
	if err != nil {
		log.Fatal("write file error :", err)
	}

}
