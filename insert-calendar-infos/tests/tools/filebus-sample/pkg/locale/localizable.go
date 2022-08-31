package locale

import (
	"github.com/rs/zerolog/log"
	"path/filepath"
	"sync"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/pkg/errors"
	"golang.org/x/text/language"
)

//JST 東京タイムゾーン
var JST = time.FixedZone("Asia/Tokyo", 9*60*60)

// Localizable ローカライズ可能のもの
type Localizable interface {
	//Localize ローカライズした文字列を返す
	Localize(l Localizer) string
}

// LocalizeKey ローカライズ用のメッセージIDを持つ
type LocalizeKey interface {
	LocalizeKey() string
}

// Localizer ローカライザー
type Localizer interface {
	Localize(messageID string, params map[string]interface{}) (string, error)
	AcceptLanguages() []string
	Location() *time.Location
}

type localizer struct {
	lo          *i18n.Localizer
	acceptLangs []string
	location    *time.Location
}

var bundle *i18n.Bundle

var onceInit sync.Once

func initBundle(conf *Config) {
	bundle = i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)

	files, err := list(conf.MessageFileDir)
	if err != nil {
		return
	}

	for _, f := range files {
		if _, err := bundle.LoadMessageFile(f); err != nil {
			log.Error().Msg("failed to load message file")
		}
	}
}

func list(dirPath string) ([]string, error) {
	return filepath.Glob(dirPath + "/*.toml")
}

// New Accept-Languageを渡し、ローカライザーを生成する
//     引数が空の場合はbundleで指定されたデフォルト言語となる
func New(acceptLangs ...string) Localizer {
	lo := i18n.NewLocalizer(bundle, acceptLangs...)
	return &localizer{lo: lo, acceptLangs: acceptLangs, location: JST}
}

func (l *localizer) AcceptLanguages() []string {
	return l.acceptLangs
}

func (l *localizer) Location() *time.Location {
	return l.location
}

func (l *localizer) Localize(messageID string, params map[string]interface{}) (string, error) {
	if len(messageID) == 0 {
		return "", errors.New("localization failed, Message ID is missing")
	}
	var data map[string]interface{}
	if len(params) > 0 {
		data = make(map[string]interface{}, len(params))
		for k, v := range params {
			localized := v
			switch value := v.(type) {
			case LocalizeKey:
				id := value.LocalizeKey()

				v, err := l.lo.Localize(&i18n.LocalizeConfig{MessageID: id})
				if err == nil {
					localized = v
				}
			case Localizable:
				localized = value.Localize(l)
			}
			data[k] = localized
		}
	}

	localized, err := l.lo.Localize(&i18n.LocalizeConfig{
		MessageID:    messageID,
		TemplateData: data,
	})
	if err != nil {
		return "", err
	}
	return localized, nil
}
