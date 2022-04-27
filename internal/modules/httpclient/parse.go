package httpclient

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// OriginalCategory 原始字符串类别
type OriginalCategory int8

const (
	Original OriginalCategory = iota // 原始形式
	EDGE                             // EDGE形式
)

// LineCategory 当前行类别
type LineCategory int8

const (
	Url    LineCategory = iota // url部分
	Header                     // 头信息部分
	Body                       // body部分
)

// ParseRequest 解析原始内容
func ParseRequest(data string) (*http.Request, error) {
	method := ``
	url := ``
	body := ``
	headers := make(map[string]string)
	data = SubstitutionVariables(data)
	reader := bufio.NewReader(strings.NewReader(data))
	originalCategory := Original
	lineCategory := Url
	index := 0
	for {
		index++
		line, err := reader.ReadString('\n')
		if err != nil || err == io.EOF {
			break
		}
		if index == 1 && line[0] == ':' {
			originalCategory = EDGE
		}
		switch originalCategory {
		case Original:
			if line == `` {
				lineCategory = Body
				continue
			}
			if index != 1 {
				lineCategory = Header
			}
			switch lineCategory {
			case Url:
				words := strings.Split(line, ` `)
				method = words[0]
				url = words[1]
			case Header:
				words := strings.Split(line, `:`)
				// 从EDGE复制出来带一个空格，这里去掉
				headers[words[0]] = words[1][1:]
			case Body:
				body += line + "\n"
			}
		case EDGE:
			if line == `` {
				lineCategory = Body
				continue
			}
			words := strings.Split(line, `:`)
			if len(words) == 2 {
				lineCategory = Header
			}
			switch lineCategory {
			case Url:
				switch words[1] {
				case `authority`:
					// 从EDGE复制出来带一个空格，这里去掉
					url = words[2][1:]
				case `method`:
					// 从EDGE复制出来带一个空格，这里去掉
					method = words[2][1:]
				case `path`:
					// 从EDGE复制出来带一个空格，这里去掉
					url = url + words[2][1:]
				case `scheme`:
					// 从EDGE复制出来带一个空格，这里去掉
					url = fmt.Sprintf(`%v://%v`, words[2][1:], url)
				}
			case Header:
				// 从EDGE复制出来带一个空格，这里去掉
				headers[words[0]] = words[1][1:]
			case Body:
				body += line + "\n"
			}
		}
	}
	myRequest, err := http.NewRequest(method, url, strings.NewReader(body))
	if err != nil {
		return nil, err
	}
	for k, v := range headers {
		myRequest.Header.Set(k, v)
	}
	return myRequest, nil
}

// SubstitutionVariables 替换变量
func SubstitutionVariables(data string) string {
	result := strings.ReplaceAll(data, `{now}`, strconv.FormatInt(time.Now().Unix(), 10))
	result = strings.ReplaceAll(result, `{now_mill}`, strconv.FormatInt(time.Now().UnixNano()/1e6, 10))
	result = strings.ReplaceAll(result, `{now_nano}`, strconv.FormatInt(time.Now().UnixNano(), 10))
	result = strings.ReplaceAll(result, "\r\n", "\n")
	return result
}
