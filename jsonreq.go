package gohttpclientv2

import (
	"bytes"
	"encoding/json"
	"log"
)

func (c *GoHttpClient) Header(k, v string) *GoHttpClient {

	c.req.Header.Set(k, v)

	return c

}

//Start with post a json object body
func (c *GoHttpClient) PostBody(url string, reqObj interface{}) *GoHttpClient {

	bf := bytes.NewBuffer([]byte{})
	jsonEncoder := json.NewEncoder(bf)
	jsonEncoder.SetEscapeHTML(false)
	err := jsonEncoder.Encode(reqObj)

	if err != nil {
		log.Println(err.Error())
		return c
	}
	c.Raw(url, bf.Bytes())
	c.Header("Content-Type", "application/json")


	return  c
}
