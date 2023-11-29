package main

import "github.com/go-resty/resty/v2"

/*
	- Trên các trình duyệt sẽ có cơ chế lưu lại cookie (như trong local storage...) để các request sau không cần phải đăng nhập
	- cookieJar cũng hoạt động tương tự như vậy, nó cũng lưu lại cookie vào một thành phần gọi là cookie jar để các request có thể sử dụng lại
	- Thường sử dụng khi ta cần gọi api của bên thứ 3 mà không muốn phải set cookie nhiều
	- Trong Golang:
		+ Sử dụng thư viện: net/http/cookiejar
		+ Một số thư viện nổi tiếng để call API như resty cũng sử dụng thư viện này. Sử dụng với resty thì ta không cần quan tâm việc set cookie jar vì đã hỗ trợ sẵn
*/

func main() {
	client := resty.New()

	client.R().SetFormData(map[string]string{
		"email": "xxx",
		"pass":  "xxx",
	}).Post("api để login")

	client.R().Get("api để get dữ liệu") // nếu không call api login thì api này sẽ bị lỗi xác thực
}
