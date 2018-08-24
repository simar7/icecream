package main

import "fmt"

func (lbc LiftBridgeClient) Sub() {
	defer lbc.wg.Done()
	fmt.Println("real Sub() called")
}

func sub(lbcIface EventStreamClient) {
	lbcIface.Sub()
}
