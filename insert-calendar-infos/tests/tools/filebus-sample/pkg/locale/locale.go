package locale

//Config 言語設定
type Config struct {
	Default        string   `yaml:"default"`
	Availables     []string `yaml:"availables"`
	MessageFileDir string   `yaml:"message_file_dir"`
}

//Init 初期化
func Init(conf *Config) {
	onceInit.Do(func() {
		initLanguage(conf)
		initBundle(conf)
	})
}
