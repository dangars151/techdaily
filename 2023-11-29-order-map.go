package main

import (
	"bytes"
	"encoding/json"
	"sort"

	"github.com/gin-gonic/gin"
)

/*
	- Trong Golang khi sử dụng json.Unmarshal để parse dữ liệu kiểu map thì thứ tự các phần tử trong map sẽ không được đảm bảo
	- Có thể google map với cụm từ: "golang encoding/json: no way to preserve the order of map keys" sẽ có khá nhiều bài nói về vấn đề này
	- Điều này là không mong muốn trong một số TH ta muốn giữ nguyên thứ tự các phần tử được truyền lên
	- Ý tưởng: Sử dụng hàm bytes.Index để xác định vị trí của các key trong dữ liệu bytes truyền lên
*/

func main() {
	r := gin.Default()
	r.POST("ping", func(c *gin.Context) {
		var req Request
		c.Bind(&req)

		c.JSON(200, gin.H{
			"responseWithNoOrder": ParseWithoutOrder(req),
			"responseWithOrder":   ParseWithOrder(req),
		})
	})
	r.Run()
}

type Request struct {
	User json.RawMessage `json:"user"`
}

type Response struct {
	Name  string
	Value interface{}
}

func ParseWithoutOrder(req Request) []Response {
	data := make(map[string]interface{}, 0)
	json.Unmarshal(req.User, &data)

	response := make([]Response, 0)
	for k, v := range data {
		response = append(response, Response{
			Name:  k,
			Value: v,
		})
	}

	return response
}

func ParseWithOrder(req Request) []Response {
	tmpData := make(map[string]interface{}, 0)
	json.Unmarshal(req.User, &tmpData)

	orders := make([]string, 0)
	for key := range tmpData {
		orders = append(orders, key) // khởi tạo mảng orders, ban đầu mảng orders sẽ không theo thứ tự
	}

	index := make(map[string]int)
	for key := range tmpData {
		esc, _ := json.Marshal(key)
		index[key] = bytes.Index(req.User, esc) // xác định vị trí của key trong dữ liệu bytes truyền lên
	}

	sort.Slice(orders, func(i, j int) bool {
		return index[orders[i]] < index[orders[j]] // sắp xếp mảng orders theo vị trí key
	})

	response := make([]Response, 0)
	for i := 0; i < len(tmpData); i++ {
		response = append(response, Response{
			Name:  orders[i],
			Value: tmpData[orders[i]],
		})
	}

	return response
}
