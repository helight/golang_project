package main

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"time"

	"github.com/tealeg/xlsx"
)

func test1() {
	var file *xlsx.File
	var sheet *xlsx.Sheet
	var row *xlsx.Row
	var cell *xlsx.Cell
	var err error

	file = xlsx.NewFile()
	sheet, _ = file.AddSheet("Sheet1")
	row = sheet.AddRow()
	cell = row.AddCell()
	cell.Value = "000101"
	cell = row.AddCell()
	cell.Value = "中文"
	err = file.Save("MyXLSXFile.xlsx")
	if err != nil {
		fmt.Printf(err.Error())
	}
}

func test2() {
	var file *xlsx.File
	var sheet *xlsx.Sheet
	var row *xlsx.Row
	var cell *xlsx.Cell
	var err error

	file, _ = xlsx.OpenFile("MyXLSXFile.xlsx")
	sheet = file.Sheet["Sheet1"]
	row = sheet.AddRow()
	cell = row.AddCell()
	cell.Value = "000101"
	cell = row.AddCell()
	cell.Value = "中文1"
	err = file.Save("MyXLSXFile1.xlsx")
	if err != nil {
		fmt.Printf(err.Error())
	}
}

func SubstrByByte(str string, length int) string {
	bs1 := []byte(str)
	fmt.Println("bs len: ", len(bs1))

	bs := []byte(str)[len(bs1)-length:]
	bl := 0
	for i := len(bs) - 1; i >= 0; i-- {
		switch {
		case bs[i] >= 0 && bs[i] <= 127:
			return string(bs[:i+1])
		case bs[i] >= 128 && bs[i] <= 191:
			bl++
		case bs[i] >= 192 && bs[i] <= 253:
			cl := 0
			switch {
			case bs[i]&252 == 252:
				cl = 6
			case bs[i]&248 == 248:
				cl = 5
			case bs[i]&240 == 240:
				cl = 4
			case bs[i]&224 == 224:
				cl = 3
			default:
				cl = 2
			}
			if bl+1 == cl {
				return string(bs[:i+cl])
			}
			return string(bs[:i])
		}
	}
	return ""
}

func test3() {
	conn, err := tls.Dial("tcp", "www.helight.info:443", nil)
	if err != nil {
		fmt.Print("Server doesn't support SSL certificate err: " + err.Error())
		return
	}

	err = conn.VerifyHostname("www.helight.info")
	if err != nil {
		fmt.Print("Hostname doesn't match with certificate: " + err.Error())
		return
	}
	expiry := conn.ConnectionState().PeerCertificates[0].NotAfter
	fmt.Printf("Issuer: %s\n证书到期时间 Expiry: %v\n", conn.ConnectionState().PeerCertificates[0].Issuer, expiry.Format(time.RFC850))
}

func checkSSL(beforeTime time.Duration) {
	client := &http.Client{
		Transport: &http.Transport{
			// 注意如果证书已过期，那么只有在关闭证书校验的情况下链接才能建立成功
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
		// 10s 超时后认为服务挂了
		Timeout: 10 * time.Second,
	}
	resp, err := client.Get("https://www.helight.info")
	if err != nil {
		fmt.Println("Server err: " + err.Error())
		return
	}
	defer func() { _ = resp.Body.Close() }()

	expiry := resp.TLS.PeerCertificates[0].NotAfter
	fmt.Printf("Issuer: %s\n证书到期时间 Expiry: %v\n", resp.TLS.PeerCertificates[0].Issuer, expiry.Format("2006-01-02 15:04:05"))
	// 遍历所有证书
	for _, cert := range resp.TLS.PeerCertificates {
		fmt.Println("certificate: " + cert.NotAfter.Local().Format("2006-01-02 15:04:05"))

		// 检测证书是否已经过期
		if !cert.NotAfter.After(time.Now()) {
			fmt.Printf("Website [www.helight.info] certificate has expired: %s \n", cert.NotAfter.Local().Format("2006-01-02 15:04:05"))
			return
		}
		fmt.Printf("cert.NotAfter.Sub(time.Now()): %f \r\n", cert.NotAfter.Sub(time.Now()))
		// 检测证书距离当前时间 是否小于 beforeTime
		// 例如 beforeTime = 7d，那么在证书过期前 6d 开始就发出警告
		if cert.NotAfter.Sub(time.Now()) < beforeTime {
			fmt.Printf("Website [www.helight.info] certificate will expire, remaining time: %fh \n", cert.NotAfter.Sub(time.Now()).Hours())
			return
		}
	}
}

func main() {
	// test1()
	// test2()
	// test3()
	checkSSL(800000 * time.Second)
	str := "斯顿发送到发送阿瑟大发送到发送分阿"
	strret := SubstrByByte(str, 36)
	fmt.Println("str: ", str)
	fmt.Println("strret: ", strret)

	// 格式化
	t := "2006-01-02 15:04:05"
	// 日期字符串
	str = "2017-09-13 00:00:00"
	loc, _ := time.LoadLocation("Local")
	time1, _ := time.ParseInLocation(t, str, loc)

	fmt.Println(time1)

	tt := time.Now()
	fmt.Println(tt.Weekday().String())
	fmt.Println(int(tt.Weekday()))

}
