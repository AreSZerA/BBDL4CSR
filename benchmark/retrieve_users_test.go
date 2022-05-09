package main

import (
	. "benchmark/blockchain"
	"testing"
)

func init() {
	// Call initUsers(n) to insert n users.
	//initUsers(10000)
}

const email = "admin@dl4csr.org"

func BenchmarkRetrieveByCompositeKey(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := Query(FuncRetrieveUserByEmail, []byte(email))
		if err != nil {
			b.FailNow()
		}
	}
}

const query1 = `{"selector":{"user_email":"admin@dl4csr.org"}}`

func BenchmarkRetrieveByQuery(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := Query(FuncRetrieveUsersByQuery, []byte(query1))
		if err != nil {
			b.FailNow()
		}
	}
}

const query2 = `{"selector":{"user_email":{"$regex":".*?dl4csr.*?"}}}`

func BenchmarkRetrieveByQueryRegexp(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := Query(FuncRetrieveUsersByQuery, []byte(query2))
		if err != nil {
			b.FailNow()
		}
	}
}
