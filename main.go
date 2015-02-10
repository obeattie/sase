package main

import (
	"fmt"
	"io/ioutil"
	"os"

	log "github.com/cihub/seelog"

	"github.com/obeattie/sase/query"
)

func main() {
	defer log.Flush()

	bytes, _ := ioutil.ReadAll(os.Stdin)
	query, err := query.Parse(string(bytes))
	if err != nil {
		fmt.Printf("Error parsing: %s\n", err.Error())
		os.Exit(1)
	}
	fmt.Println(query.QueryText())
}
