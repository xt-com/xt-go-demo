package spot

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
	"time"
)

const (
	XT_VALIDATE_ALGORITHMS            = "HmacSHA256"
	XT_VALIDATE_RECVWINDOW            = "5000"
	XT_VALIDATE_CONTENTTYPE_URLENCODE = "application/x-www-form-urlencoded"
	XT_VALIDATE_CONTENTTYPE_JSON      = "application/json;charset=UTF-8"
)

type Auth struct {
	urlencoded bool
	signed     SignedHttpAPI
	path       string
	method     string
}

/**
 * @description:
 * @param {SignedHttpAPI} signed
 * @param {*} path
 * @param {string} method
 * @return {*}
 */
func NewAuth(signed SignedHttpAPI, path, method string) *Auth {

	return &Auth{
		signed: signed,
		path:   path,
		method: method,
	}
}

// To generate the signature
/**
 * @description:
 * @param {*} nil
 * @return {*}
 */
func createSigned(xy, secret string) string {
	keys := []byte(secret)
	h := hmac.New(sha256.New, keys)
	h.Write([]byte(xy))

	return hex.EncodeToString(h.Sum(nil))
}

// urlencode encoding is determined
/**
 * @description:
 * @param {bool} value
 * @return {*}
 */
func (a *Auth) SetUrlencode(value bool) {
	a.urlencoded = value
}

// Generating request headers
/**
 * @description:
 * @param {*} xt
 * @param {*} value
 * @return {*}
 */
func (a *Auth) createHeader() url.Values {
	u := url.Values{}
	u.Set("validate-algorithms", XT_VALIDATE_ALGORITHMS)
	u.Set("validate-appkey", a.signed.Accesskey)
	u.Set("validate-recvwindow", XT_VALIDATE_RECVWINDOW)
	nt := time.Now().UnixMilli()
	value := strconv.FormatInt(nt, 10)
	u.Set("validate-timestamp", value)

	return u
}

// The headers and parameters needed to construct the request
/**
 * @description:
 * @param {*} map
 * @return {*}
 */
func (a Auth) createPayload(data map[string]interface{}) (headers map[string]string, err error) {
	var tmp, decode, X, Y string

	// 构造X
	header := a.createHeader()
	X = header.Encode()
	decode = XT_VALIDATE_CONTENTTYPE_JSON

	if a.urlencoded {
		u := url.Values{}
		for k, v := range data {
			switch i := v.(type) {
			case string:
				u.Set(k, i)
			case int64:
				value := strconv.FormatInt(i, 10)
				u.Set(k, value)
			default:
				bt, err := json.Marshal(i)
				if err != nil {
					return nil, err
				}
				u.Set(k, string(bt))
			}
		}
		tmp = u.Encode()
		decode = XT_VALIDATE_CONTENTTYPE_URLENCODE
	}

	if len(data) <= 0 {
		Y = fmt.Sprintf("#%s#%s", a.method, a.path)
	} else {
		bt, err := json.Marshal(data)
		if err != nil {
			return nil, err
		}

		param := string(bt)
		if tmp != "" {
			Y = fmt.Sprintf("#%s#%s#%s", a.method, a.path, tmp)
		} else {
			Y = fmt.Sprintf("#%s#%s#%s", a.method, a.path, param)
		}
	}

	signature := createSigned(X+Y, a.signed.Secretkey)
	header.Set("validate-signature", signature)
	header.Set("Content-Type", decode)

	headers = make(map[string]string)
	for k, v := range header {
		headers[k] = v[0]
	}

	return
}
