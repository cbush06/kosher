package main

import (
	b64 "encoding/base64"
	"io/ioutil"
	"log"
	"os"
	"text/template"
)

type axeCore struct {
	Source string
}

const sourceTemplate = `package js

import b64 "encoding/base64"

var AxeCore, _ = b64.StdEncoding.DecodeString(` + "`" + `{{.Source}}` + "`" + `)`

func main() {
	var (
		axeSource []byte
		err       error
		destFile  *os.File
	)

	if axeSource, err = ioutil.ReadFile("libs/axe.min.js"); err != nil {
		log.Fatalf("Error encountered opening [libs/axe.min.js]: %s", err)
	}

	if destFile, err = os.OpenFile("resources/js/axecore.go", os.O_WRONLY|os.O_CREATE, 0755); err != nil {
		log.Fatalf("Error encountered opening file [resources/js/axecore.go] for writing: %s", err)
	}

	templ, _ := template.New("axeSourceTemplate").Parse(sourceTemplate)
	if err = templ.Execute(destFile, axeCore{b64.StdEncoding.EncodeToString(axeSource)}); err != nil {
		log.Fatalf("Error encountered generating templated content: %s", err)
	}

	destFile.Close()
}
