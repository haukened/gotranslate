package gotranslate

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func Translate(text, srcLanguage, dstLanguage string) (string, error) {
	url := fmt.Sprintf("https://translate.googleapis.com/translate_a/single?client=gtx&sl=%s&tl=%s&dt=t&q=%s", srcLanguage, dstLanguage, url.QueryEscape(text))

	r, err := http.Get(url)

	if err != nil {
		return "", err
	}

	if r.Body != nil {
		defer r.Body.Close()
	}

	if r.StatusCode != 200 {
		return "", fmt.Errorf("HTTP response %d: %+v", r.StatusCode, http.StatusText(r.StatusCode))
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		return "", err
	}

	var jsonResult []interface{}
	err = json.Unmarshal(body, &jsonResult)
	if err != nil {
		return "", err
	}

	if len(jsonResult) > 0 {
		inner := jsonResult[0]
		for _, slice := range inner.([]interface{}) {
			for _, value := range slice.([]interface{}) {
				// find the first match
				return fmt.Sprintf("%v", value), nil
			}
		}
	}
	return "", errors.New("no translated text found")
}
