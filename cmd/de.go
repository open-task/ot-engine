package main

import (
	"flag"
	"fmt"
	"github.com/xyths/ot-engine/collect"
)

const version string = "0.1.0"

// 实际中应该用更好的变量名
var (
	h bool
	v bool

	n string
	a string
	f int
	t int

	e string
)

func init() {
	flag.BoolVar(&h, "h", false, "this help")
	flag.BoolVar(&v, "v", false, "show version and exit")

	flag.StringVar(&n, "n", "r", "`network`: r rinkeby, m mainnet")
	flag.StringVar(&a, "a", "", "smart contract `address`")
	flag.IntVar(&f, "f", 1, "block `from`")
	flag.IntVar(&t, "t", 1, "block `to`")
	flag.StringVar(&e, "e", "", "event `type`: p publish, a accept, r reject, c comfirm")

	// 改变默认的 Usage
	flag.Usage = usage
}

func main() {
	flag.Parse()

	if h {
		flag.Usage()
		return
	}

	if v {
		fmt.Println("de:", version)
		return
	}

	download(n, a, f, t, e)
}

func usage() {
	fmt.Fprintf(flag.CommandLine.Output(), `de: download event log from Ethereum Network, version: %s
Usage: de [-hv] [-n network] [-a address] [-f from] [-t to] [-e type]

Options:
`, version)
	flag.PrintDefaults()
}

func download(network string, address string, from int, to int, t string) {
	if network == "rinkeby" || network == "r" {
		// rinkeby
		fmt.Println("network is rinkeby.")
		server := "https://rinkeby.infura.io/v3/e17969db9bc94e75a474b3d3c5257a75"
		collect.Collect(server, address, from,to, t)
	}
}
