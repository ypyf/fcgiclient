package main

import (
	"io/ioutil"
	"log"

	"fcgiclient/fastcgi"
)

func main() {
	reqParams := "name=value"

	env := make(map[string]string)
	env["SCRIPT_FILENAME"] = "/home/yanfen/test.php"
	env["SERVER_SOFTWARE"] = "go / fcgiclient "
	env["REMOTE_ADDR"] = "127.0.0.1"
	env["QUERY_STRING"] = reqParams

	// fcgi, err := fcgiclient.New("tcp", "127.0.0.1:9000")
	// if err != nil {
	// 	log.Println("err:", err)
	// }

	fcgi, err := fastcgi.Dial("tcp", "127.0.0.1:9000")
	if err != nil {
		log.Fatalln(err)
	}

	resp, err := fcgi.Get(env)
	if err != nil {
		log.Fatalln(err)
	}

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("content:", string(content))
}
