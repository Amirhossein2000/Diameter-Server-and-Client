package main

import (
	"bytes"
	"flag"
	"github.com/fiorix/go-diameter/v4/diam"
	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
	"log"
	"net"
	"time"
)

var messageChan chan []byte

var (
	pcapFilePath *string
)

func init() {
	pcapFilePath = flag.String("f", "", "")
	flag.Parse()

	messageChan = make(chan []byte, 1000)
}

func main() {
	cli, _, ready, closeHandler := newDiameterClient()

	diamConnections, closeConnections := dialConnections(cli, 10)
	closeServerConn := ready()
	defer closeConnections()
	defer closeServerConn()
	defer closeHandler()

	pcapFileHandle, err := pcap.OpenOffline(*pcapFilePath)
	if err != nil {
		panic(err)
	}
	defer pcapFileHandle.Close()

	packetSource := gopacket.NewPacketSource(pcapFileHandle, pcapFileHandle.LinkType())
	packetSource.Lazy = true

	go assemble(packetSource)

	for i := range diamConnections {
		go sendRoutine(diamConnections[i])
	}

	ticker := time.NewTicker(time.Second * 10)
	for {
		<-ticker.C
		if len(packetSource.Packets()) == 0 {
			return
		}
	}

	<-ticker.C
}

func sendRoutine(conn diam.Conn) {
	counter := 0
	data := []byte{}

	for message := range messageChan {
		counter++
		data = append(data, message...)

		if counter == 1 {
			_, err := conn.Write(data)
			if err != nil {
				log.Println("write: err:", err)
				return
			}
			counter = 0
			data = []byte{}

			time.Sleep(time.Millisecond * 10)
		}
	}
}

func assemble(packetSource *gopacket.PacketSource) {
	//for {
	//	packet, _ := packetSource.NextPacket()
	//	if packet == nil {
	//		continue
	//	}
	//	messageChan <- packet.ApplicationLayer().Payload()
	//}
	//

	buffer := bytes.NewBuffer([]byte{})

	for {
		packet, _ := packetSource.NextPacket()
		if packet == nil {
			continue
		}
		buffer.Write(packet.ApplicationLayer().Payload())

		m, err := diam.ReadMessage(buffer, DefaultDict)
		//if err != nil && err != io.EOF &&
		//	strings.Contains(err.Error(), "Failed to decode AVP") &&
		//	strings.Contains(err.Error(), "Could not find preloaded") {
		//
		//	log.Println(err, buffer.Len())
		//	buffer.Reset()
		//}

		if err != nil {
			log.Println(err, ",buff len -->", buffer.Len())
			buffer.Reset()
			continue
		}

		if m != nil {
			ccr := CCR{}
			err := m.Unmarshal(&ccr)

			if err != nil {
				panic(err)
			}

			//if ccr.ServiceInformation == nil{
			//	b,_:=json.Marshal(m)
			//	fmt.Println(string(b))
			//	panic(nil)
			//}

			if ccr.ServiceInformation != nil && ccr.ServiceInformation.PSInformation.TGPPPPDPAddress != nil &&
				ccr.ServiceInformation.PSInformation.TGPPPPDPAddress.Equal(net.IP{0, 0, 0, 0}) {

				panic("fuck")
			}

			data, err := m.Serialize()
			if err != nil {
				panic(err)
			}
			messageChan <- data
		}

	}
}
