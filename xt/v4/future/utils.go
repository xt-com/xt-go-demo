package future

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/url"
	"sort"
	"strconv"
	"strings"
)

const (
	XT_VALIDATE_CONTENTTYPE_URLENCODE = "application/x-www-form-urlencoded"
	XT_VALIDATE_CONTENTTYPE_JSON      = "application/json;charset=UTF-8"
)

// ------------------------------
type Body struct {
	Result int64 `json:"result"`
}

var RepBody *Body

// ------------------------------

type Auth struct {
	urlencoded bool
	signed     SignedFutureHttpAPI
	path       string
	method     string
}

func NewAuth(signed SignedFutureHttpAPI, path, method string) *Auth {

	return &Auth{
		signed: signed,
		path:   path,
		method: method,
	}
}

// 生成签名
func createSigned(xy, secret string) string {
	keys := []byte(secret)
	h := hmac.New(sha256.New, keys)
	h.Write([]byte(xy))

	return hex.EncodeToString(h.Sum(nil))
}

// 判断是否进行urlencoded编码
func (a *Auth) SetUrlencode(value bool) {
	a.urlencoded = value
}

func (a *Auth) testGetServertime() {
	RepBody = &Body{}
	publicFutureHttpAPI := PublicFutureHttpAPI{}
	rep := publicFutureHttpAPI.GetServerTime()
	json.Unmarshal([]byte(rep.Data), RepBody)
}

// 生成请求头
func (a *Auth) createHeader() url.Values {

	// TODO ********************
	a.testGetServertime()
	// nt := time.Now().UnixMilli()
	nt := RepBody.Result
	// TODO ********************

	u := url.Values{}
	u.Set("xt-validate-appkey", a.signed.Accesskey)
	value := strconv.FormatInt(nt, 10)
	u.Set("xt-validate-timestamp", value)

	return u
}

// Change encode
func (a Auth) escape(data map[string]interface{}) (tmp string, err error) {

	u := make([]string, 0)
	keys := make([]string, 0, len(data))
	for k := range data {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		switch i := data[k].(type) {
		case string:
			u = append(u, fmt.Sprintf("%s=%s", k, i))
		case int64:
			value := strconv.FormatInt(i, 10)
			u = append(u, fmt.Sprintf("%s=%s", k, value))
		default:
			bt, err := json.Marshal(i)
			if err != nil {
				return "", err
			}
			u = append(u, fmt.Sprintf("%s=%s", k, string(bt)))
		}
	}
	tmp = strings.Join(u, "&")
	return
}

// 构造请求需要的请求头和参数
func (a Auth) createPayload(data map[string]interface{}) (headers map[string]string, err error) {
	var tmp, decode, X, Y string

	// 构造X
	header := a.createHeader()
	X = header.Encode()
	decode = XT_VALIDATE_CONTENTTYPE_JSON

	if a.urlencoded {
		tmp, err = a.escape(data)
		if err != nil {
			return
		}
		decode = XT_VALIDATE_CONTENTTYPE_URLENCODE
	}

	if len(data) <= 0 {
		Y = fmt.Sprintf("#%s", a.path)
	} else {
		bt, err := json.Marshal(data)
		if err != nil {
			return nil, err
		}

		param := string(bt)
		if tmp != "" {
			Y = fmt.Sprintf("#%s#%s", a.path, tmp)
		} else {
			Y = fmt.Sprintf("#%s#%s", a.path, param)
		}
	}

	fmt.Println("XY>>>>", X+Y)

	signature := createSigned(X+Y, a.signed.Secretkey)
	header.Set("xt-validate-signature", signature)
	header.Set("Content-Type", decode)

	headers = make(map[string]string)
	for k, v := range header {
		headers[k] = v[0]
	}
	return
}
