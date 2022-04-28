package httpclient

import (
	"bufio"
	"fmt"
	"github.com/zmhuanf/gocron/internal/modules/logger"
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
		line = strings.ReplaceAll(line, "\n", ``)
		if index == 1 && strings.HasPrefix(line, `:`) {
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
				words := strings.Split(line, `: `)
				headers[words[0]] = strings.Join(words[1:], `: `)
			case Body:
				body += line + "\n"
			}
		case EDGE:
			if line == `` {
				lineCategory = Body
				continue
			}
			if lineCategory == Url && !strings.HasPrefix(line, `:`) {
				lineCategory = Header
			}
			words := strings.Split(line, `: `)
			switch lineCategory {
			case Url:
				switch words[0] {
				case `:authority`:
					url = strings.Join(words[1:], `: `)
				case `:method`:
					method = strings.Join(words[1:], `: `)
				case `:path`:
					url = url + strings.Join(words[1:], `: `)
				case `:scheme`:
					url = fmt.Sprintf(`%v://%v`, strings.Join(words[1:], `: `), url)
				}
			case Header:
				headers[words[0]] = strings.Join(words[1:], `: `)
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
	logger.Debug(`======解析request======`)
	logger.Debug(`url = `, url)
	logger.Debug(`method = `, method)
	logger.Debug(fmt.Sprintf(`header = %+v`, headers))
	logger.Debug(fmt.Sprintf(`request header = %+v`, myRequest.Header))
	logger.Debug(`======      ======`)
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
