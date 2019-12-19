package mws

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

	u "github.com/syncfuture/go/util"
)

func (x *apiBase) newClient(action string) *amazonMWSClient {
	if action == "" {
		panic("action cannot be empty")
	}

	r := new(amazonMWSClient)
	baseURL := ConfigProvider.GetString("BaseUrl")
	if baseURL == "" {
		panic("'BaseUrl' section cannot be empty in configs.json")
	}

	r.URL = baseURL + x.Module + "/" + x.Version
	r.MarketplaceID = ConfigProvider.GetString("MWS." + x.Seller + ".MarketplaceId")
	if r.MarketplaceID == "" {
		panic("'MarketplaceId' section cannot be empty in configs.json")
	}

	r.AccessSecret = ConfigProvider.GetString("MWS." + x.Seller + ".AccessSecret")
	if r.AccessSecret == "" {
		panic("'AccessSecret' section cannot be empty in configs.json")
	}
	r.Parameters = make(map[string]string)
	r.Parameters["Action"] = action
	r.Parameters["Version"] = x.Version
	r.Parameters["AWSAccessKeyId"] = ConfigProvider.GetString("MWS." + x.Seller + ".AccessKey")
	if r.Parameters["AWSAccessKeyId"] == "" {
		panic("'AccessKey' section cannot be empty in configs.json")
	}
	r.Parameters["SellerId"] = ConfigProvider.GetString("MWS." + x.Seller + ".SellerId")
	if r.Parameters["SellerId"] == "" {
		panic("'SellerId' section cannot be empty in configs.json")
	}
	r.Parameters["SignatureMethod"] = ConfigProvider.GetString("SignatureMethod")
	if r.Parameters["SignatureMethod"] == "" {
		panic("'SignatureMethod' section cannot be empty in configs.json")
	}
	r.Parameters["SignatureVersion"] = ConfigProvider.GetString("SignatureVersion")
	if r.Parameters["SignatureVersion"] == "" {
		panic("'SignatureVersion' section cannot be empty in configs.json")
	}

	return r
}

type amazonMWSClient struct {
	URL           string
	MarketplaceID string
	AccessSecret  string
	Parameters    map[string]string
}

func newMWSClient(seller, module, version, action string) *amazonMWSClient {
	if seller == "" {
		panic("seller cannot be empty")
	}
	if module == "" {
		panic("module cannot be empty")
	}
	if version == "" {
		panic("version cannot be empty")
	}
	if action == "" {
		panic("action cannot be empty")
	}

	r := new(amazonMWSClient)
	baseURL := ConfigProvider.GetString("MWS." + seller + ".BaseUrl")
	if baseURL == "" {
		panic("'BaseUrl' section cannot be empty in configs.json")
	}

	r.URL = baseURL + module + "/" + version
	r.MarketplaceID = ConfigProvider.GetString("MWS." + seller + ".MarketplaceId")
	if r.MarketplaceID == "" {
		panic("'MarketplaceId' section cannot be empty in configs.json")
	}

	r.AccessSecret = ConfigProvider.GetString("MWS." + seller + ".AccessSecret")
	if r.AccessSecret == "" {
		panic("'AccessSecret' section cannot be empty in configs.json")
	}
	r.Parameters = make(map[string]string)
	r.Parameters["Action"] = action
	r.Parameters["Version"] = version
	r.Parameters["AWSAccessKeyId"] = ConfigProvider.GetString("MWS." + seller + ".AccessKey")
	if r.Parameters["AWSAccessKeyId"] == "" {
		panic("'AccessKey' section cannot be empty in configs.json")
	}
	r.Parameters["SellerId"] = ConfigProvider.GetString("MWS." + seller + ".SellerId")
	if r.Parameters["SellerId"] == "" {
		panic("'SellerId' section cannot be empty in configs.json")
	}
	r.Parameters["SignatureMethod"] = ConfigProvider.GetString("SignatureMethod")
	if r.Parameters["SignatureMethod"] == "" {
		panic("'SignatureMethod' section cannot be empty in configs.json")
	}
	r.Parameters["SignatureVersion"] = ConfigProvider.GetString("SignatureVersion")
	if r.Parameters["SignatureVersion"] == "" {
		panic("'SignatureVersion' section cannot be empty in configs.json")
	}

	return r
}

func (x *amazonMWSClient) Get() (r string, err error) {
	return x.do(http.Get)
}

func (x *amazonMWSClient) do(do func(url string) (*http.Response, error)) (r string, err error) {
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

func (x *amazonMWSClient) generateURL() (r string, err error) {
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

func (x *amazonMWSClient) signAmazonUrl(origUrl *url.URL) (signedUrl string, err error) {
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

func (x *amazonMWSClient) setParameter(key, value string) {
	if key != "" && value != "" {
		x.Parameters[key] = value
	} else if value == "" {
		delete(x.Parameters, key)
	}
}
