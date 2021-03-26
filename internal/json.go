/*
--------------------------------------------------
 File Name: xxjson.go
 Author: hanxu
 AuthorSite: http://www.googx.top/
 GitSource: https://github.com/googx/linuxShell
 Created Time: 2020-6-6-下午5:27
---------------------说明--------------------------

---------------------------------------------------
*/

package internal

import (
	"bytes"
	"encoding/json"
)

func PrettyJson(jsonstr string) (string, error) {
	buf := bytes.Buffer{}
	if e := json.Indent(&buf, []byte(jsonstr), "", "\t"); e != nil {
		return "", e
	}
	return buf.String(), nil
}
