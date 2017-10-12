package aliyun

import "testing"

func TestNewBaseRequest(t *testing.T) {
	req := NewBaseRequest("test")
	req.ToArgs()
	if req.Args.Get("Format") != JSONResponseFormat {
		t.Fatal("ToArgs")
	}
	cloneReq := req.Clone().(*BaseRequest)
	if cloneReq.ResponseFormat() != JSONResponseFormat {
		t.Fatal("cloneReq ToArgs")
	}

	req.SetArgs("TestKey", "TestVal")
	if req.Args.Get("TestKey") != "TestVal" {
		t.Fatal("SetArgs")
	}

	req.DelArgs("TestKey")
	if req.Args.Get("TestKey") != "" {
		t.Fatal("DelArgs")
	}
	req.Url = "http://"
	t.Log(req.String())
	req.Sign(NewCredentials("123456", "123456"))
	oldSignature := req.Signature
	t.Log("oldSignature", oldSignature)
	if oldSignature == "" {
		t.Fatal("Sign")
	}
	req.SetArgs("TestKey", "TestVal")
	req.Sign(NewCredentials("123456", "123456"))
	if oldSignature == req.Signature {
		t.Fatal("reSign")
	}
	httpReq, err := req.HttpRequest()
	if err != nil {
		t.Fatal("HttpRequest")
	}
	if httpReq.Method != req.Method {
		t.Fatal("HttpRequest Method")
	}
	t.Log("HttpRequest url", httpReq.URL)
}
