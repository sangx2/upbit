package model

import (
	"encoding/json"
	"io"
)

// ResponseError 에러 발생시에 반환되는 정보
type ResponseError struct {
	Detail struct {
		Name    string `json:"name"`
		Message string `json:"message"`
	} `json:"error"`
}

func ResponseErrorFromJSON(r io.Reader) *ResponseError {
	var re *ResponseError

	json.NewDecoder(r).Decode(&re)

	return re
}

// Response1 입금 주소 생성 요청 시에 반환되는 정보
type Response1 struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func Response1FromJSON(r io.Reader) *Response1 {
	var r1 *Response1

	json.NewDecoder(r).Decode(&r1)

	return r1
}
