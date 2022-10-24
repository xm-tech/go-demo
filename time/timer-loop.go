package main

import (
	"fmt"
	"time"
)

const GAME_FPS = 60

func main() {
	AddLoopTimer()
}

func AddLoopTimer() {
	fmt.Printf("AddLoopTimer ")
	runLoopTickerFlag := true
	maxTickerNum := 10
	go func() {
		//const FPS int32 = 60
		// //1000*1000/GAME_FPS =每个fps的微秒数
		//1000/GAME_FPS 每个fps的毫秒数
		sleepTime := time.Duration(1000/GAME_FPS) * time.Millisecond
		fmt.Printf("AddLoopTimer sleepTime %s \n", sleepTime)
		ticker := time.NewTicker(sleepTime)

		lastT := time.Now()
		var maxCostTimeMs int64 = 0
		var maxTickerTime time.Duration
		TickId := 0
		var tickerS time.Time
		for i := 0; i < maxTickerNum; i++ {
			tickerS = time.Now()
			TickId = i
			<-ticker.C
			checkTickerSleep(tickerS, sleepTime, TickId, &maxTickerTime)
			curT := time.Now()

			tickRun()
			checkTickFunRunCost(curT, lastT, sleepTime, TickId, maxCostTimeMs)
			lastT = curT
		}
		fmt.Println("pass AddLoopTimer ", maxTickerNum)
		runLoopTickerFlag = false
	}()
	fmt.Println("runLoopTickerFlag=", runLoopTickerFlag)
}

func checkTickerSleep(tickerS time.Time, sleepTime time.Duration, TickId int, maxTickerTime *time.Duration) {
	fmt.Printf("tickerS = %+v, sleepTime = %+v, TickId = %+v, maxTickerTime = %+v\n", tickerS, sleepTime, TickId, maxTickerTime)
}

func tickRun() {
	fmt.Println("tickRun")
}

func checkTickFunRunCost(curT time.Time, lastT time.Time, sleepTime time.Duration, TickId int, maxCostTimeMs int64) {
	fmt.Println("checkTickFunRunCost", curT, lastT, sleepTime, TickId, maxCostTimeMs)
}
