// Code generated by aliensboot. DO NOT EDIT.
// source: defaultmodule_interface.proto
package service

import "github.com/KylinHe/aliensboot-server/protocol"


//
func handleBenchmark(authID int64, gateID string, request *protocol.Benchmark, response *protocol.BenchmarkRet) {
    response.RespContent = request.RequestContent
    return
}

