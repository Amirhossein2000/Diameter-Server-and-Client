package main

import (
	"github.com/fiorix/go-diameter/v4/diam"
	"log"
	"net"
)

func handleCCR(c diam.Conn, message *diam.Message) {
	clientCCR, err := ParseCCR(message)
	if err != nil {
		log.Println("error", err)
		return
	}
	if clientCCR == nil {
		log.Println("nil CCR")
	}

	// check some of broken struct conditions
	if !isValid(clientCCR.ServiceInformation.PSInformation.TGPPPPDPAddress) ||
		!isValid(clientCCR.ServiceInformation.PSInformation.GGSNAddress) ||
		!isValid(clientCCR.ServiceInformation.PSInformation.SGSNAddress) {

		// we have the race condition right now but we enable print for more details
		//data, err := json.Marshal(clientCCR)
		//if err != nil {
		//	log.Println("broken CCR Found but could not marshal")
		//	return
		//}
		//log.Printf("broken CCR Found len: %d data: %s\n", message.Header.MessageLength, data)
	}
}

func ParseCCR(m *diam.Message) (*CCR, error) {
	var serverCCR CCR
	err := m.Unmarshal(&serverCCR)
	if err != nil {
		return nil, err
	}

	return &serverCCR, nil
}

func isValid(ip *net.IP) bool {
	if ip == nil {
		// this is ok
		return true
	}

	dereferenceIP := *ip
	for i := range dereferenceIP {
		if dereferenceIP[i] == 0 {
			return false
		}
	}

	return true
}
