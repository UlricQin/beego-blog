package g

var (
	RootEmail    string
	RootName     string
	RootPass     string
	RootPortrait string
)

func initCfg() {
	RootName = Cfg.String("root_name")
	RootEmail = Cfg.String("root_email")
	RootPass = Cfg.String("root_pass")
	RootPortrait = Cfg.String("root_portrait")
}
