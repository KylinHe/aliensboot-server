package main

import (
	"flag"
	"fmt"
	"strconv"
	"test/protobuf_client/base"
	"test/protobuf_client/object"
	"time"
)

var g_Players map[int32]*object.Player

var (
	gameserver    string
	accountserver string
	preaccount    string
	password      string
	accountnum    int
	synctime      int
	exist         bool
	secretkey     string
)

//stresstest -accountserver "127.0.0.1:18811" -gameserver "127.0.0.1:18812" -preaccount "test_" -password "11111111" -accountnum 3000 -synctime 5 -exist=true
func main() {
	//形参
	//gameserver := flag.String("gameserver", "", "gameserver address")
	//accountserver := flag.String("accountserver", "", "accountserver address")
	//preaccount := flag.String("preaccount", "", "pre account")
	//password := flag.String("password", "", "password")
	//accountnum := flag.Int("accountnum", "", "account number")
	//synctime := flag.Int("synctime", "", "synctime=5s")
	//exist := flag.Bool("exist", "", "account is exist?")
	flag.StringVar(&gameserver, "gameserver", "", "gameserver address")
	flag.StringVar(&accountserver, "accountserver", "", "accountserver address")
	flag.StringVar(&preaccount, "preaccount", "", "pre account")
	flag.StringVar(&password, "password", "", "password")
	flag.StringVar(&secretkey, "secretkey", "", "cipher key")
	flag.IntVar(&accountnum, "accountnum", 3000, "account number")
	flag.IntVar(&synctime, "synctime", 5, "synctime=5s")
	flag.BoolVar(&exist, "exist", false, "account is exist?")

	flag.Parse()

	if gameserver == "" || accountserver == "" || preaccount == "" || password == "" {
		println("Please input correct params => stresstest -h")
		println("stresstest -accountserver \"127.0.0.1:18811\" -gameserver \"127.0.0.1:18812\" -preaccount \"test\" -password \"11111111\" -secretkey \"abcd\" -accountnum 3000 -synctime 5 -exist=true")
		return
	}

	println("secript key: " + secretkey)
	g_Players = make(map[int32]*object.Player)

	fmt.Printf("##Create %v connections\n", accountnum)
	time.Sleep(3 * time.Second)

	if exist {
		fmt.Println("##Login account")
	} else {
		fmt.Println("##Register account")
	}

	for i := 1; i <= accountnum; i++ {
		p := &object.Player{}
		p.Username = preaccount + "_" + strconv.Itoa(i)
		p.Password = password
		//p.Init(int32(i), "127.0.0.1:3568")
		p.Init(int32(i), accountserver, gameserver, int64(synctime), secretkey)
		g_Players[p.GetIdx()] = p
	}

	//time.Sleep(3 * time.Second)
	//for _, p := range g_Players {
	//	if exist {
	//		p.AcceptOp(base.OP_LOGIN)
	//	} else {
	//		p.AcceptOp(base.OP_REGISTER)
	//	}
	//	time.Sleep(10 * time.Millisecond)
	//}

	fmt.Println("##Sync data")

	for _, p := range g_Players {
		p.AcceptOp(base.OP_SYNC)
	}
	time.Sleep(time.Hour)
	//for {
	//
	//	//time.Sleep(time.Duration(synctime) * time.Second)
	//}
}
