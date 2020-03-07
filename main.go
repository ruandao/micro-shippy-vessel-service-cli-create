package main

import (
	"context"
	"fmt"
	microclient "github.com/micro/go-micro/client"
	"github.com/micro/go-micro/config/cmd"
	"github.com/ruandao/micro-shippy-vessel-service-ser/lib"
	go_micro_srv_vessel "github.com/ruandao/micro-shippy-vessel-service-ser/proto/vessel"
	"log"
	"time"
)

func main() {
	cmd.Init()
	now := time.Now().Unix()
	id := fmt.Sprintf("createAt_%v", now)
	client := go_micro_srv_vessel.NewVesselServiceClient(lib.CONST_SER_NAME_VESSEL, microclient.DefaultClient)
	resp, err := client.Create(context.Background(), &go_micro_srv_vessel.Vessel{
		Id:                   id,
		Capacity:             1000,
		MaxWeight:            1000,
		Name:                 id,
		Available:            true,
		OwnerId:              "ljy",
	})

	if err != nil {
		log.Fatalf("Create Vessel err: %v", err)
	}

	log.Printf("vessel: %v", resp.Vessels)
}
