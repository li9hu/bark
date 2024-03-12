package bark

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"

	"github.com/li9hu/rehttp"
)

type Sounds struct {
	S滴嘟滴嘟   string
	S敲钟     string
	S鸟叫     string
	S急促     string
	S叮叮     string
	S谷故估谷故估 string
}

var Sound = &Sounds{"alarm", "bell", "birdsong", "electronic", "glass", "horn"}

type Bark struct {
	Key string
	Url string
}

type Result struct {
	Code      int    `json:"code"`
	Message   string `json:"message"`
	Timestamp int    `json:"timestamp"`
}

var bark *Bark

func SetUp(key, url string) {
	b := &Bark{key, url}
	bark = b
}

// Run 请求Bark 返回消息唯一ID
func Run(group, title, data, sound string) (string, error) {
	if bark == nil {
		return "", fmt.Errorf("please use bark.SetUp first")
	}

	// 构造 bark 请求
	url := bark.Url
	head := map[string]string{
		"Content-Type": "application/json;charset=utf-8",
	}
	body := map[string]string{
		"title":      title,
		"body":       data,
		"device_key": bark.Key,
		"group":      group,
		"sound":      sound,
	}
	b, err := json.Marshal(body)
	if err != nil {
		return "", fmt.Errorf("bark request body to json: %v", err)
	}

	// 生成唯一ID
	hasher := md5.New()
	hasher.Write(b)
	hashBytes := hasher.Sum(nil)
	// 将哈希值转换为固定长度的字符串ID
	id := hex.EncodeToString(hashBytes)

	// 尝试发送 Bark
	ret := rehttp.Post(url, head, b)
	if ret.Err != nil {
		return "", fmt.Errorf("bark post: %v", err)
	}
	barkResult := &Result{}
	err = json.Unmarshal([]byte(ret.ResponseBody), barkResult)
	if err != nil {
		return "", fmt.Errorf("bark result to json: %v", err)
	}

	if barkResult.Code != 200 {
		return "", fmt.Errorf("bark status != 200: %s", barkResult.Message)
	}
	return id, nil
}
