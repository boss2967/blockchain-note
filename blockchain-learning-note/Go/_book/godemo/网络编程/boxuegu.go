package main

import (
	"net/http"
	"net/url"
	"fmt"
	"os"
	"strings"
	"io"
	"strconv"
	"encoding/json"
	"github.com/json-iterator/go"
	"math/rand"
	"time"
)

const (
	loginUrl        = "http://ntlias-stu.boxuegu.com/user/login"
	getQuestionsUrl = "http://ntlias-stu.boxuegu.com/feedback/queryTodayFeedback?r_=212600302202344.88"
	postAnswersUrl  = "http://ntlias-stu.boxuegu.com/feedback/save"
	loginName       = "A180700717"
	password        = "A180700717"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	response, err := http.PostForm(loginUrl, url.Values{"loginName": {loginName}, "password": {password}})
	if err != nil {
		fmt.Println("postform err:", err)
		return
	}
	defer response.Body.Close()
	buf := make([]byte, 4096)
	var resultStr string
	for {
		n, err := response.Body.Read(buf)
		resultStr += string(buf[:n])
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("response.Body.Read err:", err)
			return
		}
	}
	fmt.Println(resultStr)
	val := []byte(resultStr)
	isSuccess := jsoniter.Get(val, "success").ToString()
	if isSuccess == "false" {
		fmt.Println("---------------------[登录失败]")
		return
	}
	cookie := response.Cookies()[0].Value
	fmt.Println("---------------------[登录成功] cookie=", cookie)
	getTable(cookie)
}
func getTable(cookie string) {
	var r http.Request
	r.ParseForm()
	bodystr := strings.TrimSpace(r.Form.Encode())
	request, err := http.NewRequest("GET", getQuestionsUrl, strings.NewReader(bodystr))
	if err != nil {
		fmt.Println("http newrequest err:", err)
		return
	}
	request.Header.Set("Cookie", "SESSION="+cookie)
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		fmt.Println("http get err:", err)
		return
	}
	defer response.Body.Close()
	buf := make([]byte, 4096)
	fp, err := os.Create("result.html")
	if err != nil {
		fmt.Println("os.Create err:", err)
		return
	}
	var resultStr string
	for {
		n, err := response.Body.Read(buf)
		resultStr += string(buf[:n])
		fp.Write(buf[:n])
		if err == io.EOF {
			break
		}
		if err != nil {
			return
		}
	}
	fmt.Println(resultStr)
	val := []byte(resultStr)
	isSuccess := jsoniter.Get(val, "success").ToString()
	if isSuccess == "false" {
		fmt.Println("---------------------[获取feedbackId和count失败]")
		return
	}
	count := jsoniter.Get(val, "resultObject", "list", 0, "questionCount").ToString()
	feedbackId := jsoniter.Get(val, "resultObject", "list", 0, "feedbackId").ToString()
	fmt.Printf("---------------------[获取feedbackId和count成功] feedbackId=%s,count=%s\n", feedbackId, count)
	num, _ := strconv.Atoi(count)
	postYourAnswer(cookie, num, feedbackId)
}
func postYourAnswer(cookie string, count int, feedbackId string) {
	questionSlice := make([]string, 0)
	for i := 1; i <= count; i++ {
		o := &Obj{}
		o.QuestionNo = strconv.Itoa(i)
		randNum := rand.Intn(3)
		if randNum == 0 {
			o.QuestionVal = "1"
		} else {
			o.QuestionVal = "2"
		}

		data1, _ := json.Marshal(o)
		questionSlice = append(questionSlice, string(data1))
	}
	fmt.Println("questionSlice",questionSlice)
	postValue := url.Values{"feedbackId": {feedbackId}, "suggest": {"null"}, "targetArr": questionSlice,"anonymity":{"1"}}
	postString := postValue.Encode()
	fmt.Println(postValue)
	fmt.Println(postString)
	req, err := http.NewRequest("POST", postAnswersUrl, strings.NewReader(postString))
	if err != nil {
		fmt.Println("NewRequest err:", err)
		return
	}
	//* Accept: application/json, text/plain, *
	//* Accept-Encoding: gzip, deflate
	//* Accept-Language: en-US,en;q=0.9,zh-CN;q=0.8,zh-TW;q=0.7,zh;q=0.6
	//* Connection: keep-alive
	//* Content-Length: 500
	//* Content-Type: application/x-www-form-urlencoded
	//* Cookie: SESSION=9b7b5b55-039c-4a75-a72e-1b5d1e3c4318
	//* Host: ntlias-stu.boxuegu.com
	//* Origin: http://ntlias-stu.boxuegu.com
	//* Referer: http://ntlias-stu.boxuegu.com/
	//* User-Agent: Mozilla/5.0 (iPhone; CPU iPhone OS 11_0 like Mac OS X) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/11.0 Mobile/15A372 Safari/604.1
	req.Header.Set("Cookie", "SESSION="+cookie)
	req.Header.Set("Content-Length", "500")
	req.Header.Set("Host", "ntlias-stu.boxuegu.com")
	req.Header.Set("Accept", "application/json, text/plain, *")
	req.Header.Set("Content-Type","application/x-www-form-urlencoded")
	req.Header.Set("Origin","http://ntlias-stu.boxuegu.com")
	req.Header.Set("Referer","http://ntlias-stu.boxuegu.com")
	req.Header.Set("User-Agent","Mozilla/5.0 (iPhone; CPU iPhone OS 11_0 like Mac OS X) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/11.0 Mobile/15A372 Safari/604.1")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("client.do err:", err)
		return
	}
	var resultStr string
	buf := make([]byte, 4096)
	for {
		n, err := resp.Body.Read(buf)
		resultStr += string(buf[:n])
		if err == io.EOF {
			break
		}
		if err != nil {
			return
		}
	}
	val := []byte(resultStr)
	isSuccess := jsoniter.Get(val, "success").ToString()
	fmt.Println(resultStr)
	if isSuccess == "false" {
		fmt.Println("---------------------[反馈提交失败]")
		return
	} else {
		fmt.Println("---------------------[反馈提交成功]")
	}

}

type Obj struct {
	QuestionNo  string
	QuestionVal string
}
