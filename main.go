package main

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/novalagung/gubrak/v2"
)

const puncts = "?!.;,"

type Result struct {
	Text  string `json:"text"`
	Count int    `json:"count"`
}

type Body struct {
	Text string `json:"text"`
}

func main() {
	r := gin.Default()
	r.POST("/submit", process)
	r.Run(":3000")
}

func process(c *gin.Context) {

	body := Body{}
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	str := body.Text
	newStr := cleanString(str)

	loweredCaseStr := strings.ToLower(newStr)
	spl := strings.Fields(loweredCaseStr)

	grpby := gubrak.From(spl).GroupBy(func(each string) string {
		return each
	}).Result().(map[string][]string)

	res := []Result{}
	for i, v := range grpby {
		var each Result
		each.Text = grpby[i][0]
		each.Count = len(v)
		res = append(res, each)
	}

	// --- order desc
	orderDesc := gubrak.From(res).OrderBy(func(each Result) int {
		return each.Count
	}, false).Result()

	// --- top 10 result
	top10Results := []Result{}
	for i, v := range orderDesc.([]Result) {
		if i > 9 {
			break
		}
		top10Results = append(top10Results, v)
	}

	c.JSON(http.StatusOK, top10Results)
}

func isPunctuation(c string) bool {
	if strings.Contains(puncts, c) {
		return true
	}
	return false
}

func cleanString(input string) string {

	size := len(input)
	temp := ""
	var prevChar string

	for i := 0; i < size; i++ {
		str := string(input[i]) // convert to string for easier operation
		if (str == " " && prevChar != " ") || !isPunctuation(str) {
			temp += str
			prevChar = str
		} else if prevChar != " " && isPunctuation(str) {
			temp += " "
		}
	}
	return temp
}
