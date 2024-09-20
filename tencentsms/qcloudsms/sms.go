package qcloudsms

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"
)

/*
{
    "ext": "",
    "extend": "",
    "params": [
        "验证码",
        "1234",
        "4"
    ],
    "sig": "ecab4881ee80ad3d76bb1da68387428ca752eb885e52621a3129dcf4d9bc4fd4",
    "sign": "腾讯云",
    "tel": {
        "mobile": "13788888888",
        "nationcode": "86"
    },
    "time": 1457336869,
    "tpl_id": 19
}
*/

type SMSReq struct {
	EXT    string   `json:"ext"`
	Extend string   `json:"extend"`
	Params []string `json:"params"`
	Sig    string   `json:"sig"`
	Sign   string   `json:"sign"`
	Tel    Tel      `json:"tel"`
	Time   int64    `json:"time"`
	TplID  int      `json:"tpl_id"`
}

type Tel struct {
	Mobile     string `json:"mobile"`
	Nationcode string `json:"nationcode"`
}

/*
{
    "result": 0, // 错误码，0表示成功(计费依据)，非0表示失败，参考 错误码
    "errmsg": "OK", // 错误消息，result 非0时的具体错误信息
    "ext": "",  // 用户的 session 内容，腾讯 server 回包中会原样返回
    "fee": 1,  // 短信计费的条数，计费规则请参考 国内短信内容长度计算规则 或 国际短信内容长度计算规则
    "sid": "xxxxxxx"
}
*/
type SMSRet struct {
	Result int64  `json:"result"`
	Errmsg string `json:"errmsg"`
	EXT    string `json:"ext"`
	Fee    int64  `json:"fee"`
	Sid    string `json:"sid"`
}

var Appid = 1400177009                          // SDK AppID是1400开头// 短信应用SDK AppID
var Appkey = "872abce909b9f13fce8d8133b23fbe73" // 短信应用SDK AppKey
var Sign = "黑光技术"                               // 签名
var Nationcode = "86"

var smsurl = "https://yun.tim.qq.com/v5/tlssmssvr/sendsms?sdkappid=%d&random=%d"

var codeTpl = 355535   // 您的验证码为:{1}，仅用于本次预约申请，请与{2}分钟内完成操作。如非本人操作，请忽略本短信。
var noticeTpl = 355634 // {1} 先生/女生，您好！您预约的参观活动申请（参观时间：{2}）审核通过。联系电话：{3}
var adminTpl = 355633  // 管理员，您好！有单位预约的参观活动申请，请登录后台进行审批。

// POST https://yun.tim.qq.com/v5/tlssmssvr/sendsms?sdkappid=xxxxx&random=xxxx

type QcloudSms struct {
	initok bool
	appid  int
	appkey string
	sign   string
	ncode  string
}

func NewQcloudSms() *QcloudSms {
	return &QcloudSms{initok: false}
}

func (t *QcloudSms) Doinit(appid int, appkey, sign, ncode string) {
	t.appid = appid
	t.appkey = appkey
	t.sign = sign
	t.ncode = ncode
	t.initok = true
}

func (t *QcloudSms) SendSms(tplId int, mobile string, params []string) (ret int) {
	msg := &SMSReq{}

	randseed := GenRandom(100000, 999999)
	timestamp := time.Now().Unix()
	msg.Sig = GenSignature(t.appkey, mobile, randseed, timestamp)
	msg.Params = params
	msg.TplID = tplId
	msg.Sign = t.sign
	msg.Time = timestamp
	msg.Tel.Mobile = mobile
	msg.Tel.Nationcode = t.ncode
	strmsg, _ := json.Marshal(msg)
	fmt.Println("req: ", string(strmsg))

	url := fmt.Sprintf(smsurl, t.appid, randseed)
	retdata, _ := HttpRequst("POST", url, string(strmsg))
	fmt.Println("create group chat: ", string(retdata))
	jsondata := json.NewDecoder(bytes.NewReader(retdata))
	var smsRet SMSRet
	jsondata.Decode(&smsRet)
	fmt.Println("smsRet: ", smsRet)

	return 0
}

func (t *QcloudSms) SendCheckCode(mobile string, code string) (ret int) {
	var params []string
	params = append(params, code)
	params = append(params, "5")

	return t.SendSms(codeTpl, mobile, params)
}

func (t *QcloudSms) SendNoitceSms(mobile string, name, dtime, tel_num string) (ret int) {
	var params []string
	params = append(params, name)
	params = append(params, dtime)
	params = append(params, tel_num)

	return t.SendSms(noticeTpl, mobile, params)
}

func (t *QcloudSms) SendAdminSms(mobile string) (ret int) {
	var params []string
	params = append(params, "name")

	return t.SendSms(adminTpl, mobile, params)
}
