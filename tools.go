package archeage

import (
	"io"
	"net/url"
	"strings"
)

var ServerNameMap = map[string]string{
	"안탈론":  "ANTHALON",
	"크라켄":  "KRAKEN",
	"노아르타": "NOARTA",
	"키리오스": "KYRIOS",
	"누이":   "NUI",
	"곤":    "GON",
	"키프로사": "KYPROSA",
}

func form(m map[string]string) io.Reader {
	data := url.Values{}
	for k, v := range m {
		data.Set(k, v)
	}
	return strings.NewReader(data.Encode())
}
