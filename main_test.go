package gohttpclientv2

import (
	"net/url"
	"testing"
	"log"
)

func TestGet(t *testing.T) {
	s, err := Get("https://www.baidu.com").Exec().String()

	if err != nil {
		t.Fatal(err)
	}

	log.Println(s)
}

func TestPost(t *testing.T) {

	res := PostForm("http://jd.com/pageNotExists",url.Values{
		"a":[]string{"b"},
	}).Exec()

	str, err := res.String()



	if err != nil {
		t.Fatal("error:", err)
	}

	log.Println(str)
}
