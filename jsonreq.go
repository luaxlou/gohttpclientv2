package gohttpclientv2

import (
	"bytes"
	"encoding/json"
)

func (c *GoHttpClient) Header(k, v string) *GoHttpClient {

	c.req.Header.Set(k, v)

	return c

}


//Start with post a json object body
func PostBody(url string, reqObj interface{}) *GoHttpClient {

	bf := bytes.NewBuffer([]byte{})
	jsonEncoder := json.NewEncoder(bf)
	jsonEncoder.SetEscapeHTML(false)
	err := jsonEncoder.Encode(reqObj)

	c := Raw(url, bf.Bytes())

	if err != nil {
		return c
	}

	c.Header("Content-Type", "application/json")

	return c
}
