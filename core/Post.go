package core

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/syncfuture/go/u"
)

func (x *MWSClient) Post() (r string, err error) {
	return x.doPost(http.PostForm)
}

func (x *MWSClient) doPost(do func(url string, body url.Values) (*http.Response, error)) (r string, err error) {
	uuu, _ := url.Parse(x.URL)
	values := url.Values{}
	for k, v := range x.Parameters {
		values.Set(k, v)
	}
	values.Set("Timestamp", time.Now().UTC().Format(time.RFC3339))

	hash, err := x.getSignature(uuu, values.Encode())
	if u.LogError(err) {
		return "", err
	}
	values.Set("Signature", hash)

	resp, err := do(uuu.String(), values)
	if u.LogError(err) {
		return "", err
	} else if resp != nil {
		defer resp.Body.Close()
	} else {
		return
	}

	bytes, err := ioutil.ReadAll(resp.Body)
	if u.LogError(err) {
		return "", err
	}

	return string(bytes), nil
}

func (x *MWSClient) getSignature(origUrl *url.URL, encodeParams string) (signedUrl string, err error) {
	params := strings.Replace(encodeParams, ",", "%2C", -1)
	params = strings.Replace(params, ":", "%3A", -1)

	toSign := fmt.Sprintf("POST\n%s\n%s\n%s", origUrl.Host, origUrl.Path, params)

	hasher := hmac.New(sha256.New, []byte(x.AccessSecret))
	_, err = hasher.Write([]byte(toSign))
	if err != nil {
		return "", err
	}
	hash := base64.StdEncoding.EncodeToString(hasher.Sum(nil))
	// hash = url.QueryEscape(hash)

	return hash, err
}
