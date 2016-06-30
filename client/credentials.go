package client

//Api请求凭证
type Credentials struct {
	//Access Key Id
	AccessKeyId     string
	//Access Key Secret
	AccessKeySecret string
}

//func (c *Credentials)SetAccessKeyId(accessKeyId   string) {
//	c.AccessKeyId = accessKeyId
//}
//
//func (c *Credentials)SetAccessKeySecret(accessKeySecret  string) {
//	c.AccessKeySecret = accessKeySecret
//}

//var DefaultCredentials *Credentials
//
//func InitDefaultCredentials(accessKeyId, accessKeySecret string) {
//	DefaultCredentials = NewCredentials(accessKeyId, accessKeySecret)
//}

func NewCredentials(accessKeyId, accessKeySecret string) *Credentials {
	return &Credentials{
		AccessKeyId:accessKeyId,
		AccessKeySecret:accessKeySecret,
	}
}

