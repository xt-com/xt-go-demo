package v4

import (
	"encoding/json"
	"fmt"
	"sdk/response"
	"strconv"
	"time"

	"github.com/parnurzeal/gorequest"
	r "github.com/solos/requests"
)

const (
	TIMEOUT = 10
)

type RequestPerpare struct {
	request *r.Request
}

func NewRequestPerpare() *RequestPerpare {
	return &RequestPerpare{
		request: &r.Request{
			Args: r.M{
				"timeout": TIMEOUT,
			},
		},
	}
}

/**
 * @description:
 * @param {*} k
 * @param {*} queryVal
 * @return {*}
 */
func (rp *RequestPerpare) queryStruct(s *gorequest.SuperAgent, content interface{}) (err error) {
	if marshalContent, err := json.Marshal(content); err != nil {
		return err
	} else {
		var val map[string]interface{}
		if err := json.Unmarshal(marshalContent, &val); err != nil {
			s.Errors = append(s.Errors, err)
		} else {
			for k, v := range val {
				var queryVal string
				switch t := v.(type) {
				case string:
					queryVal = t
				case float64:
					queryVal = strconv.FormatFloat(t, 'f', -1, 64)
				case time.Time:
					queryVal = t.Format(time.RFC3339)
				default:
					j, err := json.Marshal(v)
					if err != nil {
						continue
					}
					queryVal = string(j)
				}
				s.QueryData.Add(k, queryVal)
			}
		}
	}

	return
}

// Make the request QueryString
/**
 * @description:
 * @param {*} body
 * @param {*} Success
 * @param {*} url
 * @param {*} true
 * @return {*}
 */
func (rp *RequestPerpare) RequesParam(
	method, url string,
	headers map[string]string,
	data map[string]interface{},
) *response.APIBody {

	request := gorequest.New()

	switch method {
	case "GET":
		request = request.Get(url)
	case "POST":
		request = request.Post(url)
	default:
		request = request.Get(url)
	}
	err := rp.queryStruct(request, data)
	if err != nil {
		return response.APIResponse(err.Error(), "Failed", url, false)
	}

	for k, v := range headers {
		request.AppendHeader(k, v)
	}

	_, body, errs := request.End()
	if errs != nil {
		return response.APIResponse(errs[0].Error(), "Failed", url, false)
	}

	return response.APIResponse(body, "Success", url, true)
}

// Make a request to JsonBody
/**
 * @description:
 * @param {*} body
 * @param {*} Success
 * @param {*} url
 * @param {*} true
 * @return {*}
 */
func (rp *RequestPerpare) RequesJson(
	method, url string,
	headers map[string]string,
	data map[string]interface{},
) *response.APIBody {

	request := gorequest.New()

	switch method {
	case "POST":
		request = request.Post(url).Send(data)
	case "DELETE":
		request = request.Delete(url).Send(data)
	default:
		request = request.Post(url).Send(data)
	}

	for k, v := range headers {
		request.AppendHeader(k, v)
	}

	_, body, err := request.End()
	if err != nil {
		fmt.Println(err[0].Error())
		return response.APIResponse(body, "Failed", url, false)
	}

	return response.APIResponse(body, "Success", url, true)

}
