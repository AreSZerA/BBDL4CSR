# DL4CSR Benchmark

### Modified File

`${GOPATH}/pkg/mod/github.com/hyperledger/fabric-sdk-go@v1.0.0/internal/github.com/hyperledger/fabric/core/operations/system.go`
has been modified. In line 225, modify `go s.statsd.SendLoop(context.Background(), s.sendTicker.C, network, address)`
to `go s.statsd.SendLoop(s.sendTicker.C, network, address)`
. [Source code](https://github.com/hyperledger/fabric-sdk-go/commit/14047c6d88f0e995f09d55817bfbf735e245547a)
has been modified on GitHub but the latest release does not contain this modification.
