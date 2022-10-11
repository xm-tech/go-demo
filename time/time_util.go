package main

import (
	"flag"
	"fmt"
	"strconv"
	"time"
)

const timeLayout = "2006-01-02 15:04:05"

type TimeConvert struct {
	op  string
	val string
}

func (tc *TimeConvert) u2t() {
	timestamp, err := strconv.Atoi(tc.val)
	if err != nil {
		fmt.Println("u2t fail, error val type,val :", tc.val)
		panic(err)
	}
	timeStr := Format(int64(timestamp))
	fmt.Println(timeStr)
}

func (tc *TimeConvert) t2u() {
	utime := ToUnix(tc.val)
	fmt.Println(utime)
}

func main() {
	flag.Parse()
	args := flag.Args()

	op := args[0]
	val := args[1]

	tc := &TimeConvert{op: op, val: val}

	switch tc.op {
	case "u2t":
		tc.u2t()
	case "t2u":
		tc.t2u()
	}
}

func Format(timestamp int64) string {
	return time.Unix(timestamp, 0).Format(timeLayout)
}

func ToUnix(timestr string) int64 {
	timeObj, _ := time.Parse(timeLayout, timestr)
	return timeObj.Unix()
}
