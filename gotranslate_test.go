package gotranslate

import (
	"strings"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestTranslate(t *testing.T) {
	test_phrase := "this is only a test"
	known_translation := "esto es solo una prueba"
	source_lang := "en"
	target_lang := "es"
	result, err := Translate(test_phrase, source_lang, target_lang)
	result = strings.ToLower(result)
	if err != nil {
		t.Errorf("translate function produced error: %+v", err)
	}
	assert.Equal(t, result, known_translation)
}
