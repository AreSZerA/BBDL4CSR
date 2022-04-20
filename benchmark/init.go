package main

import (
	. "benchmark/blockchain"
	"github.com/google/uuid"
	"log"
	"strconv"
)

func initUsers(n int) {
	resp, err := Query(FuncRetrieveUsers, []byte("count"))
	if err != nil {
		log.Panicln("Failed to initialize:", err.Error())
	}
	count, _ := strconv.Atoi(string(resp.Payload))
	if count > n {
		log.Panicln("Number of users out of range:", count, ">", 10)
	}
	for i := 0; i < n-count; i++ {
		id, _ := uuid.NewUUID()
		_, err = Execute(FuncCreateUser, []byte(id.String()), []byte(id.String()), []byte("12345678901234567890123456789012"))
		if err != nil {
			log.Panicln("Failed to initialize:", err.Error())
		}
	}
}
