package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HitApiExternal(c *gin.Context) {

	url := "https://cat-fact.herokuapp.com/facts"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Cookie", "connect.sid=s%3AQ0uDpHx7jrfouqsIjc8ip4nMd9nM7mhN.eu3ABOzTnreBR%2FzUwktTI%2FyuZlPSEGQ3xzpDq4BzO90")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	var data interface{}

	json.Unmarshal(body, &data)

	c.JSON(200, map[string]interface{}{
		"DATA": data,
	})

}
