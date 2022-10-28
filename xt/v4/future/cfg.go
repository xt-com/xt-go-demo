package future

import "os"

var (
	BaseUrl string
)

func init() {

	env := os.Getenv("APIENV")
	switch env {
	case "pro":
		BaseUrl = "https://fapi.xt.com"
	case "dev":
		BaseUrl = "http://fapi.xt.com"
	case "test":
		BaseUrl = "http://fapi.xt.com"
	default:
		BaseUrl = "http://fapi.xt.com"
	}

}
