package main

import (
	"fmt"
	"encoding/json"
	"strconv"
	"math/rand"
	"time"
)
type Product struct {
	Name string
	ProductID int64
	Number int
	Price float64 //json:"price"
	IsOnSale bool //json:"is_on_sale,omitempty"
}

func main() {
	//data: ={"feedbackId":"57471","suggest":null,"targetArr":[{"questionNo":"1","questionVal":"1"},
	//{"questionNo":"2","questionVal":"1"},{"questionNo":"3","questionVal":"1"},
	//{"questionNo":"4","questionVal":"1"},{"questionNo":"5","questionVal":"1"},
	//{"questionNo":"6","questionVal":"1"}],"anonymity":0}
	rand.Seed(time.Now().UnixNano())
	randNum:=rand.Intn(3)
	fmt.Println(randNum)
	questionSlice:=make([]string,0)
	for i := 1; i <= 6; i++ {
		o:=&Obj{}
		o.QuestionNo = strconv.Itoa(i)
		o.QuestionVal = "2"

		data1,_:=json.Marshal(o)
		questionSlice=append(questionSlice, string(data1))
	}
	fmt.Println(questionSlice)
}
type Obj struct {
	QuestionNo string
	QuestionVal string
}
