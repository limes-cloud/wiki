package main

import (
	"encoding/json"
	"fmt"
	"regexp"
)

func matchTarget(target []string, input string) bool {
	for _, pattern := range target {
		// 将*替换为匹配任意多字符的正则表达式
		pattern = "^" + pattern + "$"
		pattern = regexp.MustCompile("/\\*").ReplaceAllString(pattern, "/.*")

		// 编译正则表达式
		re := regexp.MustCompile(pattern)

		// 检查输入是否匹配正则表达式
		if re.MatchString(input) {
			return true
		}
	}
	return false
}

type AuditFailReason struct {
	Headline string `json:"headline"`
	Image    string `json:"image"`
	Abstract string `json:"abstract"`
	Content  string `json:"content"`
}

func main() {
	object := AuditFailReason{}
	err := json.Unmarshal([]byte("{这是一段话"), &object)
	fmt.Println(err)
	target := []string{"/hello/*", "/welcome/lihua"}

	input1 := "/hello/lihua"
	input2 := "/welcome/john"

	fmt.Printf("Input: %s, Match: %v\n", input1, matchTarget(target, input1))
	fmt.Printf("Input: %s, Match: %v\n", input2, matchTarget(target, input2))
}
