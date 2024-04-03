package log

import (
	"bytes"
	"encoding/json"
)

func StdoutFormtJson(j ...interface{}) {
	for _, item := range j {
		log, _ := json.Marshal(item)

		var str bytes.Buffer
		_ = json.Indent(&str, log, "", "    ")
		println(str.String())
	}
}
