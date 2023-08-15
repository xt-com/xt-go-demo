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

/**
 * @description:
 * @param {SignedFutureHttpAPI} signed
 * @param {*} path
 * @param {string} method
 * @return {*}
 */
func NewAuth(signed SignedFutureHttpAPI, path, method string) *Auth {

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

/**
 * @description:
 * @param {*} rep
 * @return {*}
 */
func (a *Auth) testGetServertime() {

	RepBody = &Body{}
	publicFutureHttpAPI := PublicFutureHttpAPI{}
	rep := publicFutureHttpAPI.GetServerTime()
	json.Unmarshal([]byte(rep.Data), RepBody)

}

// Generating request headers
/**
 * @description:
 * @param {*} xt
 * @param {*} value
 * @return {*}
 */
func (a *Auth) createHeader() url.Values {

	a.testGetServertime()
	nt := RepBody.Result

	u := url.Values{}
	u.Set("validate-appkey", a.signed.Accesskey)
	value := strconv.FormatInt(nt, 10)
	u.Set("validate-timestamp", value)

	return u
}

// Change encode
/**
 * @description:
 * @param {*} u
 * @return {*}
 */
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

	signature := createSigned(X+Y, a.signed.Secretkey)
	header.Set("validate-signature", signature)
	header.Set("Content-Type", decode)

	headers = make(map[string]string)
	for k, v := range header {
		headers[k] = v[0]
	}

	return
}
