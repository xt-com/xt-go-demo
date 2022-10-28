package spot

import "os"

var (
	BaseUrl string
)

func init() {

	env := os.Getenv("APIENV")
	switch env {
	case "pro":
		BaseUrl = "https://sapi.xt.com"
	case "dev":
		BaseUrl = "https://sapi.xt.com"
	case "test":
		BaseUrl = "https://sapi.xt.com"
	default:
		BaseUrl = "https://sapi.xt.com"
	}

}
