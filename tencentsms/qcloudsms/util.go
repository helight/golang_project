package qcloudsms

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

var msgsigtpl = "appkey=%s&random=%d&time=%d&mobile=%s"

func GenRandom(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}

// "sig" 字段根据公式 sha256（appkey=$appkey&random=$random&time=$time&mobile=$mobile）
/*
string strMobile = "13788888888"; //tel 的 mobile 字段的内容
string strAppKey = "5f03a35d00ee52a21327ab048186a2c4"; //sdkappid 对应的 appkey，需要业务方高度保密
string strRand = "7226249334"; //URL 中的 random 字段的值
string strTime = "1457336869"; //UNIX 时间戳
string sig = sha256(appkey=5f03a35d00ee52a21327ab048186a2c4&random=7226249334&time=1457336869&mobile=13788888888)
		= ecab4881ee80ad3d76bb1da68387428ca752eb885e52621a3129dcf4d9bc4fd4;
		hashlib.sha256(utf8(raw_text)).hexdigest()
*/
func GenSignature(appkey, mobile string, randseed int, timestamp int64) string {
	hash := sha256.New()
	msgsigtxt := fmt.Sprintf(msgsigtpl, appkey, randseed, timestamp, mobile)
	fmt.Println("msgsigtxt: ", msgsigtxt)
	hash.Write([]byte(msgsigtxt))
	hashInBytes := hash.Sum(nil)
	hashValue := hex.EncodeToString(hashInBytes)
	fmt.Println("msgsig: ", hashValue)
	return hashValue
}

// GET/POST
func HttpRequst(method, url, data string) (res []byte, err error) {
	client := &http.Client{}
	ret := []byte("")
	req, err := http.NewRequest(method, url, strings.NewReader(data))
	if err == nil {
		req.Header.Set("Content-Type", "pplication/json")

		resp, _ := client.Do(req)
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err == nil {
			fmt.Println(string(body))
			ret = body
		}
	}
	return ret, err
}
