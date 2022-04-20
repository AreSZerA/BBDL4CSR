# DL4CSR Benchmark

### Modified File

`${GOPATH}/pkg/mod/github.com/hyperledger/fabric-sdk-go@v1.0.0/internal/github.com/hyperledger/fabric/core/operations/system.go`
has been modified. In line 225, modify `go s.statsd.SendLoop(context.Background(), s.sendTicker.C, network, address)`
to `go s.statsd.SendLoop(s.sendTicker.C, network, address)`
. [Source code](https://github.com/hyperledger/fabric-sdk-go/commit/14047c6d88f0e995f09d55817bfbf735e245547a)
has been modified on GitHub but the latest release does not contain this modification.

### Benchmark Results

Benchmark environment

> goos: linux
>
> goarch: amd64
>
> cpu: Intel(R) Core(TM) i7-8550U CPU @ 1.80GHz

#### Search 1 user in 1000 users

|    #    |  ops  |  ns / op  |
|:-------:|:-----:|:---------:|
|    1    |  194  |  6388050  |
|    2    |  195  |  5884506  |
|    3    |  205  |  6701470  |
|    4    |  199  |  6049829  |
|    5    |  195  |  5992238  |
|    6    |  198  |  6086278  |
|    7    |  199  |  6006521  |
|    8    |  193  |  5979019  |
|    9    |  200  |  6345257  |
|   10    |  190  |  6011029  |
| average | 196.8 | 5543767.6 |