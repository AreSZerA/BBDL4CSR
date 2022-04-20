package main

import (
	. "benchmark/blockchain"
	"testing"
)

func init() {
	//initUsers(1000)
}

func BenchmarkRetrieveUsers(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := Query(FuncRetrieveUserByEmail, []byte("admin@dl4csr.org"))
		if err != nil {
			b.FailNow()
		}
	}
}
