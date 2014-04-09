package g

var (
	RootEmail    string
	RootName     string
	RootPass     string
	RootPortrait string
	BlogTitle string
	BlogResume string
)

func initCfg() {
	RootName = Cfg.String("root_name")
	RootEmail = Cfg.String("root_email")
	RootPass = Cfg.String("root_pass")
	RootPortrait = Cfg.String("root_portrait")
	BlogTitle = Cfg.String("blog_title")
	BlogResume = Cfg.String("blog_resume")
}
