package aliyun

//Credentials API请求凭证
type Credentials struct {
	//Access Key Id
	AccessKeyID string
	//Access Key Secret
	AccessKeySecret string
}

func NewCredentials(accessKeyID, accessKeySecret string) *Credentials {
	return &Credentials{
		AccessKeyID:     accessKeyID,
		AccessKeySecret: accessKeySecret,
	}
}
