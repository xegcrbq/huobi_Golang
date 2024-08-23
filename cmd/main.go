package main

import (
	"github.com/xegcrbq/huobi_golang/cmd/accountclientexample"
	"github.com/xegcrbq/huobi_golang/cmd/accountwebsocketclientexample"
	"github.com/xegcrbq/huobi_golang/cmd/algoorderclientexample"
	"github.com/xegcrbq/huobi_golang/cmd/commonclientexample"
	"github.com/xegcrbq/huobi_golang/cmd/crossmarginclientexample"
	"github.com/xegcrbq/huobi_golang/cmd/etfclientexample"
	"github.com/xegcrbq/huobi_golang/cmd/isolatedmarginclientexample"
	"github.com/xegcrbq/huobi_golang/cmd/marketclientexample"
	"github.com/xegcrbq/huobi_golang/cmd/marketwebsocketclientexample"
	"github.com/xegcrbq/huobi_golang/cmd/orderclientexample"
	"github.com/xegcrbq/huobi_golang/cmd/orderwebsocketclientexample"
	"github.com/xegcrbq/huobi_golang/cmd/stablecoinclientexample"
	"github.com/xegcrbq/huobi_golang/cmd/subuserclientexample"
	"github.com/xegcrbq/huobi_golang/cmd/walletclientexample"
	"github.com/xegcrbq/huobi_golang/logging/perflogger"
)

func main() {
	runAll()
}

// Run all examples
func runAll() {
	commonclientexample.RunAllExamples()
	accountclientexample.RunAllExamples()
	orderclientexample.RunAllExamples()
	algoorderclientexample.RunAllExamples()
	marketclientexample.RunAllExamples()
	isolatedmarginclientexample.RunAllExamples()
	crossmarginclientexample.RunAllExamples()
	walletclientexample.RunAllExamples()
	subuserclientexample.RunAllExamples()
	stablecoinclientexample.RunAllExamples()
	etfclientexample.RunAllExamples()
	marketwebsocketclientexample.RunAllExamples()
	accountwebsocketclientexample.RunAllExamples()
	orderwebsocketclientexample.RunAllExamples()
}

// Run performance test
func runPerfTest() {
	perflogger.Enable(true)

	commonclientexample.RunAllExamples()
	accountclientexample.RunAllExamples()
	orderclientexample.RunAllExamples()
	marketclientexample.RunAllExamples()
	isolatedmarginclientexample.RunAllExamples()
	crossmarginclientexample.RunAllExamples()
	walletclientexample.RunAllExamples()
	etfclientexample.RunAllExamples()
}
