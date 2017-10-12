package aliyun

//Api请求凭证
type Credentials struct {
	//Access Key Id
	AccessKeyId string
	//Access Key Secret
	AccessKeySecret string
}

func NewCredentials(accessKeyId, accessKeySecret string) *Credentials {
	return &Credentials{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
	}
}
