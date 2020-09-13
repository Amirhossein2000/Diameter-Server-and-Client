package main

import (
	"context"
	"github.com/fiorix/go-diameter/v4/diam"
	"github.com/fiorix/go-diameter/v4/diam/avp"
	"github.com/fiorix/go-diameter/v4/diam/datatype"
	"github.com/fiorix/go-diameter/v4/diam/sm"
	"log"
	"net"
	"time"
)

func newDiameterClient() (*sm.Client, chan *diam.Message, readyFunc, closeFunc) {
	cfg := &sm.Settings{
		OriginHost:       datatype.DiameterIdentity("lab"),
		OriginRealm:      datatype.DiameterIdentity("parspooyesh.com"),
		VendorID:         13,
		ProductName:      "go-diameter",
		OriginStateID:    datatype.Unsigned32(time.Now().Unix()),
		FirmwareRevision: 1,
		HostIPAddresses: []datatype.Address{
			datatype.Address(net.ParseIP("127.0.0.1")),
		},
	}

	mux := sm.New(cfg)
	ctx, cancel := context.WithCancel(context.Background())
	closeHandler := func() {
		cancel()
	}
	go printErrors(ctx, mux.ErrorReports())

	cli := &sm.Client{
		Dict:               DefaultDict,
		Handler:            mux,
		MaxRetransmits:     3,
		RetransmitInterval: time.Second,
		EnableWatchdog:     true,
		WatchdogInterval:   50 * time.Second,
		AuthApplicationID: []*diam.AVP{
			// Advertise support for credit control application
			diam.NewAVP(avp.AuthApplicationID, avp.Mbit, 0, datatype.Unsigned32(4)), // RFC 4006
		},
	}

	respChannel := make(chan *diam.Message, 10)
	mux.Handle("CCA", handleCCA(ctx, respChannel))
	ready := handShake(mux, time.Second*10)

	return cli, respChannel, ready, closeHandler
}

func dialConnections(cli *sm.Client, parallelCount int) ([]diam.Conn, func()) {
	var err error
	diamConnections := make([]diam.Conn, parallelCount)

	for i := range diamConnections {
		diamConnections[i], err = cli.DialTimeout("127.0.0.1:3868", time.Second*10)
		if err != nil {
			log.Println("Dial Network err: %s", err.Error())
		}
	}

	return diamConnections, getClose(diamConnections)
}
