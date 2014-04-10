package g

var (
	RootEmail    string
	RootName     string
	RootPass     string
	RootPortrait string
	BlogTitle string
	BlogResume string
	BlogLogo string
	QiniuAccessKey string
	QiniuSecretKey string
	UseQiniu bool
)

func initCfg() {
	RootName = Cfg.String("root_name")
	RootEmail = Cfg.String("root_email")
	RootPass = Cfg.String("root_pass")
	RootPortrait = Cfg.String("root_portrait")
	BlogTitle = Cfg.String("blog_title")
	BlogResume = Cfg.String("blog_resume")
	BlogLogo = Cfg.String("blog_logo")
	QiniuAccessKey = Cfg.String("qiniu_access_key")
	QiniuSecretKey = Cfg.String("qiniu_secret_key")
	UseQiniu = QiniuAccessKey != "" && QiniuSecretKey != ""
}
