package main

import (
	"fmt"
	"sync"
	"tencentsms/qcloudsms"
	"time"
)

var sesscache map[string]string
var sesslock sync.Mutex

func main() {
	//获取时间戳
	timestamp := time.Now().Unix()
	fmt.Println(timestamp)

	//格式化为字符串,tm为Time类型
	tm := time.Unix(timestamp, 0)
	fmt.Println(tm.Format("2006-01-02 03:04:05 PM"))
	fmt.Println(tm.Format("02/01/2006 15:04:05 PM"))

	//从字符串转为时间戳，第一个参数是格式，第二个是要转换的时间字符串
	tm2, _ := time.Parse("01/02/2006", "02/08/2015")
	fmt.Println(tm2.Unix())

	sendsms := qcloudsms.NewQcloudSms()
	sendsms.Doinit(qcloudsms.Appid, qcloudsms.Appkey, qcloudsms.Sign, qcloudsms.Nationcode)
	var params []string
	randseed := qcloudsms.GenRandom(1000, 9999)
	fmt.Println("randseed: ", string(randseed))
	params = append(params, fmt.Sprintf("%d", randseed))
	params = append(params, fmt.Sprintf("%d", 5))
	fmt.Println("params: ", params)
	// sendsms.SendSms(355535, tencentsms.SmsSign, tencentsms.Nationcode, "18128867300", params)
	// 18819669333
	// sendsms.SendCheckCode("18128867300", fmt.Sprintf("%d", randseed))
	// sendsms.SendNoitceSms("18128867300", "许", "2019-06-17", "18819669333")
	// sendsms.SendAdminSms("18128867300")

	sesscache = make(map[string]string)
	sesslock.Lock()
	sesscache["asd1"] = "asdfsdf1"
	sesscache["asd2"] = "asdfsdf2"
	sesscache["asd3"] = "asdfsdf3"
	sesscache["asd4"] = "asdfsdf4"
	sesscache["asd5"] = "asdfsdf5"
	sesscache["asd5"] = "asdfsdf6"
	sesslock.Unlock()
	// return hashValue
	CleanSess()

	//时间戳
	t := time.Now()
	fmt.Println(t.Weekday())
}

func CleanSess() {
	sesslock.Lock()
	for k, v := range sesscache {
		fmt.Println("k: ", k, " v: ", v)
	}
	sesslock.Unlock()
}
