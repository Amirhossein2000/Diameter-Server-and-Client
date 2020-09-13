package main

import (
	"github.com/fiorix/go-diameter/v4/diam"
	"github.com/fiorix/go-diameter/v4/diam/datatype"
	"github.com/fiorix/go-diameter/v4/diam/sm"
	"log"
	"strings"
)

func main() {
	settings := &sm.Settings{
		OriginHost:       datatype.DiameterIdentity("cgw1"),
		OriginRealm:      datatype.DiameterIdentity("parspooyesh.com"),
		VendorID:         51670, // Parspooyesh
		ProductName:      "IBSng",
		FirmwareRevision: 3,
	}

	mux := sm.New(settings)

	handlerCCR := diam.HandlerFunc(handleCCR)
	mux.Handle("CCR", handlerCCR)

	// log error reports
	go func() {
		for ec := range mux.ErrorReports() {
			if !strings.Contains(ec.Error.Error(), "unhandled message for 'CCA'") {
				log.Printf("Diameter Server Error: %s\n", ec)
			}
		}
	}()

	if err := diam.ListenAndServe("127.0.0.1:3868", mux, DefaultDict); err != nil {
		panic(err)
	}
}
