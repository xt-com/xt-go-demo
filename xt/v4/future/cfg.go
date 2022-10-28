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
		BaseUrl = "https://fapi.xt.com"
	case "test":
		BaseUrl = "https://fapi.xt.com"
	default:
		BaseUrl = "https://fapi.xt.com"
	}

}
