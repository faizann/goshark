package main

import (
	"fmt"
	"goshark"
	"log"
)

func main() {

	file := "2.pcap"
	d := goshark.CreateDecoder()
	if err := d.DecodeStart(file); err != nil {
		log.Println("Decode start fail:", err)
		return
	}
	defer d.DecodeEnd()

	f, err := d.NextPacket()
	if err != nil {
		log.Println("Get packet fail:", err)
		return
	}

	fmt.Printf("%s", f)
}