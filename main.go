package main

import (
	"io/ioutil"
	"os"

	log "github.com/cihub/seelog"

	"github.com/obeattie/sase-parser/query"
)

func main() {
	defer log.Flush()

	bytes, _ := ioutil.ReadAll(os.Stdin)
	query, err := query.Parse(string(bytes))
	if err != nil {
		panic(err)
	}
	panic(query.Query())
}
