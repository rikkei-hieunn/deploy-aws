package locale

import (
	"golang.org/x/text/language"
)

var languageMatcher language.Matcher

func initLanguage(conf *Config) {
	tags := make([]language.Tag, 0)
	tags = append(tags, language.MustParse(conf.Default))
	for _, s := range conf.Availables {
		tags = append(tags, language.MustParse(s))
	}
	languageMatcher = language.NewMatcher(tags)
}

// ExtractLocale 指定されたAcceptLanguage(優先順位)から、使用出来る言語を返却する
func ExtractLocale(acceptLanguage string) string {
	tag, _ := language.MatchStrings(languageMatcher, acceptLanguage)
	return tag.String()
}
