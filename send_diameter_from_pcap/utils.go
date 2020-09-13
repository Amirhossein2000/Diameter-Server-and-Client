package main

import (
	"context"
	"github.com/fiorix/go-diameter/v4/diam"
	"github.com/fiorix/go-diameter/v4/diam/sm"
	"log"
	"time"
)

type closeFunc func()
type readyFunc func() closeFunc

func handShake(mux *sm.StateMachine, timeout time.Duration) readyFunc {
	ready := make(chan struct{})
	hsc := make(chan diam.Conn, 1)

	go func() {
		close(ready)
		c := <-(*mux).HandshakeNotify()
		hsc <- c
	}()

	return func() closeFunc {
		var h diam.Conn
		<-ready

		select {
		case <-time.After(timeout):
			log.Fatalf("diameter handshake timeout mux: %+v", *mux)
		case h = <-hsc:
			log.Println("diameter handshake successful")
		}

		return func() {
			h.Close()
			h.(diam.CloseNotifier).CloseNotify()
		}
	}
}

func printErrors(ctx context.Context, errChan <-chan *diam.ErrorReport) {
	for {
		select {
		case <-ctx.Done():
			return
		case err := <-errChan:
			log.Println("CCA Handler err:", err.Error)
		}
	}
}

func handleCCA(ctx context.Context, respChannel chan<- *diam.Message) diam.HandlerFunc {
	return func(conn diam.Conn, m *diam.Message) {
		select {
		case <-ctx.Done():
			conn.Close()
			return
		default:
			//respChannel <- m
		}
	}
}

func getClose(diamConnections []diam.Conn) func() {
	return func() {
		for i := 0; i < len(diamConnections); i++ {
			go func(conn diam.Conn) {
				conn.Close()
				conn.(diam.CloseNotifier).CloseNotify()
			}(diamConnections[i])
		}
		<-time.After(time.Second * 2)
	}
}
