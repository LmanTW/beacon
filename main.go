package main

import (
        "fmt"
        "os"

        "tinygo.org/x/bluetooth"
)

var adapter = bluetooth.DefaultAdapter
var advertisement = adapter.DefaultAdvertisement()

var hwid = []byte{0x00, 0x00, 0x00, 0x00, 0x00}
var rssi = []byte{0xC5}

func main() {
    err := adapter.Enable()
    if err != nil {
        fmt.Println("Failed enable the adapter:", err.Error())

        os.Exit(1)
    }

    serviceData := []byte{0x02}
    serviceData = append(serviceData, hwid...)
    serviceData = append(serviceData, rssi...)

    uuid, err2 := bluetooth.ParseUUID("FE6F")
    if err2 != nil {
        fmt.Println("Failed to parse UUID:", err2)

        os.Exit(1)
    }

    err = advertisement.Configure(bluetooth.AdvertisementOptions{
        LocalName: "Line Beacon",

        ServiceUUIDs: []bluetooth.UUID{uuid},
        ServiceData: []bluetooth.ServiceDataElement{{
            UUID: uuid,
            Data: serviceData,
        }},

        Interval: 250,
    })
    if err != nil {
        println("Failed to configure the advertisement:", err.Error())

        os.Exit(1)
    }

    err = advertisement.Start()
    if err != nil {
        println("Failed to start advertising:", err.Error())

        os.Exit(1)
    }

    fmt.Println("Beacon advertising")

    select {}
}
