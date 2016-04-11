package g

import (
	"github.com/qiniu/api.v6/io"
	"github.com/qiniu/api.v6/rs"
)

func UploadFile(localFile string, destName string) (addr string, err error) {
	policy := new(rs.PutPolicy)
	policy.Scope = QiniuScope
	uptoken := policy.Token(nil)

	var ret io.PutRet
	var extra = new(io.PutExtra)
	err = io.PutFile(nil, &ret, uptoken, destName, localFile, extra)
	if err != nil {
		return
	}
	addr = "http://" + QiniuScope + ".qiniudn.com/" + destName
	return
}
