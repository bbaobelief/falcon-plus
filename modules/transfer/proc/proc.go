// Copyright 2017 Xiaomi, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package proc

import (
	"log"

	nproc "github.com/toolkits/proc"
)

// trace
var (
	RecvDataTrace = nproc.NewDataTrace("RecvDataTrace", 3)
)

// filter
var (
	RecvDataFilter = nproc.NewDataFilter("RecvDataFilter", 5)
)

// 统计指标的整体数据
var (
	// 计数统计,正确计数,错误计数, ...
	RecvCnt       = nproc.NewSCounterQps("RecvCnt")
	RpcRecvCnt    = nproc.NewSCounterQps("RpcRecvCnt")
	HttpRecvCnt   = nproc.NewSCounterQps("HttpRecvCnt")
	SocketRecvCnt = nproc.NewSCounterQps("SocketRecvCnt")

	SendToKafkaCnt    = nproc.NewSCounterQps("SendToKafkaCnt")
	SendToInfluxDBCnt = nproc.NewSCounterQps("SendToInfluxDBCnt")
	SendToJudgeCnt    = nproc.NewSCounterQps("SendToJudgeCnt")
	SendToTsdbCnt     = nproc.NewSCounterQps("SendToTsdbCnt")
	SendToGraphCnt    = nproc.NewSCounterQps("SendToGraphCnt")
	SendToPolyCnt     = nproc.NewSCounterQps("SendToPolyCnt")

	SendToKafkaDropCnt = nproc.NewSCounterQps("SendToKafkaDropCnt")
	SendToJudgeDropCnt = nproc.NewSCounterQps("SendToJudgeDropCnt")
	SendToPolyDropCnt  = nproc.NewSCounterQps("SendToPolyDropCnt")
	SendToTsdbDropCnt  = nproc.NewSCounterQps("SendToTsdbDropCnt")
	SendToGraphDropCnt = nproc.NewSCounterQps("SendToGraphDropCnt")

	SendToKafkaFailCnt = nproc.NewSCounterQps("SendToKafkaFailCnt")
	SendToJudgeFailCnt = nproc.NewSCounterQps("SendToJudgeFailCnt")
	SendToTsdbFailCnt  = nproc.NewSCounterQps("SendToTsdbFailCnt")
	SendToGraphFailCnt = nproc.NewSCounterQps("SendToGraphFailCnt")
	SendToPolyFailCnt  = nproc.NewSCounterQps("SendToPolyFailCnt")

	// 发送缓存大小
	KafkaQueuesCnt    = nproc.NewSCounterBase("KafkaSendCacheCnt")
	InfluxDBQueuesCnt = nproc.NewSCounterBase("InfluxDBQueuesCnt")
	JudgeQueuesCnt    = nproc.NewSCounterBase("JudgeSendCacheCnt")
	TsdbQueuesCnt     = nproc.NewSCounterBase("TsdbSendCacheCnt")
	GraphQueuesCnt    = nproc.NewSCounterBase("GraphSendCacheCnt")
	PolyQueuesCnt     = nproc.NewSCounterBase("PolySendCacheCnt")

	// http请求次数
	HistoryRequestCnt = nproc.NewSCounterQps("HistoryRequestCnt")
	InfoRequestCnt    = nproc.NewSCounterQps("InfoRequestCnt")
	LastRequestCnt    = nproc.NewSCounterQps("LastRequestCnt")
	LastRawRequestCnt = nproc.NewSCounterQps("LastRawRequestCnt")

	// http回执的监控数据条数
	HistoryResponseCounterCnt = nproc.NewSCounterQps("HistoryResponseCounterCnt")
	HistoryResponseItemCnt    = nproc.NewSCounterQps("HistoryResponseItemCnt")
	LastRequestItemCnt        = nproc.NewSCounterQps("LastRequestItemCnt")
	LastRawRequestItemCnt     = nproc.NewSCounterQps("LastRawRequestItemCnt")
)

func Start() {
	log.Println("proc.Start, ok")
}

func GetAll() []interface{} {
	ret := make([]interface{}, 0)

	// recv cnt
	ret = append(ret, RecvCnt.Get())
	ret = append(ret, RpcRecvCnt.Get())
	ret = append(ret, HttpRecvCnt.Get())
	ret = append(ret, SocketRecvCnt.Get())

	// send cnt
	ret = append(ret, SendToKafkaCnt.Get())
	ret = append(ret, SendToJudgeCnt.Get())
	ret = append(ret, SendToTsdbCnt.Get())
	ret = append(ret, SendToGraphCnt.Get())
	ret = append(ret, SendToPolyCnt.Get())

	// drop cnt
	ret = append(ret, SendToKafkaDropCnt.Get())
	ret = append(ret, SendToJudgeDropCnt.Get())
	ret = append(ret, SendToTsdbDropCnt.Get())
	ret = append(ret, SendToGraphDropCnt.Get())

	// send fail cnt
	ret = append(ret, SendToKafkaFailCnt.Get())
	ret = append(ret, SendToJudgeFailCnt.Get())
	ret = append(ret, SendToTsdbFailCnt.Get())
	ret = append(ret, SendToGraphFailCnt.Get())
	ret = append(ret, SendToPolyFailCnt.Get())

	// cache cnt
	ret = append(ret, KafkaQueuesCnt.Get())
	ret = append(ret, JudgeQueuesCnt.Get())
	ret = append(ret, TsdbQueuesCnt.Get())
	ret = append(ret, GraphQueuesCnt.Get())
	ret = append(ret, PolyQueuesCnt.Get())

	// http request
	ret = append(ret, HistoryRequestCnt.Get())
	ret = append(ret, InfoRequestCnt.Get())
	ret = append(ret, LastRequestCnt.Get())
	ret = append(ret, LastRawRequestCnt.Get())

	// http response
	ret = append(ret, HistoryResponseCounterCnt.Get())
	ret = append(ret, HistoryResponseItemCnt.Get())
	ret = append(ret, LastRequestItemCnt.Get())
	ret = append(ret, LastRawRequestItemCnt.Get())

	return ret
}
