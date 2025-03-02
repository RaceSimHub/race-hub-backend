package utils

import "strings"

type Exception struct {
	Key     string `json:"key,omitempty"`
	Message string `json:"message,omitempty"`
}

func (e Exception) Make(m string) Exception {
	var arr = strings.Split(m, " | ")

	e.Key = arr[0]

	for index, value := range arr {
		if index == 0 {
			continue
		}
		e.Message += value + " | "
	}

	if len(e.Message) > 3 {
		e.Message = e.Message[:len(e.Message)-3]
	}

	return e
}

type List struct {
	Pagination Pagination `json:"pagination"`
	Data       any        `json:"data"`
}

type Pagination struct {
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
	Total  int `json:"total"`
}

type Id struct {
	Id int `json:"id"`
}
