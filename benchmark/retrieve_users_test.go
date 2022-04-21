package main

import (
	. "benchmark/blockchain"
	"testing"
)

func init() {
	// Call initUsers(n) to insert n users.
	//initUsers(10000)
}

func BenchmarkRetrieveUsers(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := Query(FuncRetrieveUserByEmail, []byte("admin@dl4csr.org"))
		if err != nil {
			b.FailNow()
		}
	}
}
