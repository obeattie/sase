package main

import (
	"io/ioutil"
	"os"

	log "github.com/cihub/seelog"

	"github.com/obeattie/sase-parser/parser"
)

func main() {
	defer log.Flush()

	bytes, _ := ioutil.ReadAll(os.Stdin)
	if _, err := parser.Parse(string(bytes)); err != nil {
		log.Error(err.Error())
	}
}
