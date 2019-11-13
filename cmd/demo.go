/*
   Copyright © 2019 M.Watermann, 10247 Berlin, Germany
                   All rights reserved
               EMail : <support@mwat.de>
*/

package main

//lint:file-ignore ST1017 – I prefer Yoda conditions

import (
	"io"
	"log"
	"net/http"
	"os"

	"github.com/mwat56/apachelogger"
)

// `myHandler()` is a dummy for demonstration purposes.
func myHandler(aWriter http.ResponseWriter, aRequest *http.Request) {
	_, _ = io.WriteString(aWriter, "Hello world!")
} // myHandler()

func main() {
	// the filename should be taken from the commandline or a config file:
	logfile := "/dev/stderr"

	pageHandler := http.NewServeMux()
	pageHandler.HandleFunc("/", myHandler)

	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: apachelogger.Wrap(pageHandler, logfile),
	}
	apachelogger.SetErrLog(&server)

	log.Fatalf("%s: %v", os.Args[0], server.ListenAndServe())
} // main()

/* _EoF_ */
