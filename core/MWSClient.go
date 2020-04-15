package core

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"time"

	mwsconfig "github.com/syncfuture/mws/config"

	"github.com/syncfuture/go/u"
)

type (
	APIBase struct {
		Config  *mwsconfig.MWSConfig
		Module  string
		Version string
	}
	QueryBase struct {
	}
)

func (x *APIBase) NewClient(action string) (r *MWSClient, err error) {
	if action == "" {
		panic("action cannot be empty")
	}

	r = new(MWSClient)
	r.URL = x.Config.BaseURL + x.Module + "/" + x.Version
	r.MarketplaceID = x.Config.MarketplaceID
	r.AccessSecret = x.Config.AccessSecret
	r.Parameters = make(map[string]string)
	r.Parameters["Action"] = action
	r.Parameters["Version"] = x.Version
	r.Parameters["AWSAccessKeyId"] = x.Config.AccessKey
	r.Parameters["SellerId"] = x.Config.SellerID
	r.Parameters["SignatureMethod"] = x.Config.SignatureMethod
	r.Parameters["SignatureVersion"] = x.Config.SignatureVersion
	if x.Config.AuthToken != "" {
		r.Parameters["MWSAuthToken"] = x.Config.AuthToken
	}

	return
}

type MWSClient struct {
	URL           string
	MarketplaceID string
	AccessSecret  string
	Parameters    map[string]string
}

func (x *MWSClient) Get() (r string, err error) {
	return x.do(http.Get)
}

func (x *MWSClient) do(do func(url string) (*http.Response, error)) (r string, err error) {
	url, err := x.generateURL()
	if u.LogError(err) {
		return "", err
	}
	resp, err := do(url)
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

func (x *MWSClient) generateURL() (r string, err error) {
	uuu, _ := url.Parse(x.URL)

	values := url.Values{}

	for k, v := range x.Parameters {
		values.Set(k, v)
	}

	values.Set("Timestamp", time.Now().UTC().Format(time.RFC3339))

	params := values.Encode()
	uuu.RawQuery = params

	r, err = x.signAmazonUrl(uuu)

	return r, err
}

func (x *MWSClient) signAmazonUrl(origUrl *url.URL) (signedUrl string, err error) {
	escapeUrl := strings.Replace(origUrl.RawQuery, ",", "%2C", -1)
	escapeUrl = strings.Replace(escapeUrl, ":", "%3A", -1)

	params := strings.Split(escapeUrl, "&")
	sort.Strings(params)
	sortedParams := strings.Join(params, "&")

	toSign := fmt.Sprintf("GET\n%s\n%s\n%s", origUrl.Host, origUrl.Path, sortedParams)

	hasher := hmac.New(sha256.New, []byte(x.AccessSecret))
	_, err = hasher.Write([]byte(toSign))
	if err != nil {
		return "", err
	}

	hash := base64.StdEncoding.EncodeToString(hasher.Sum(nil))

	hash = url.QueryEscape(hash)

	newParams := fmt.Sprintf("%s&Signature=%s", sortedParams, hash)

	origUrl.RawQuery = newParams

	return origUrl.String(), nil
}

func (x *MWSClient) SetParameter(key, value string) {
	if key != "" && value != "" {
		x.Parameters[key] = value
	} else if value == "" {
		delete(x.Parameters, key)
	}
}
