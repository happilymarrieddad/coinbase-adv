package client

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"sync"
	"sync/atomic"
	"time"
)

const (
	CoinbaseAdvV3endpoint = "https://api.coinbase.com/api/v3"
	CoinbaseAdvV2endpoint = "https://api.coinbase.com/v2"
	DefaultAPIRateLimit   = 100 // (ms) throttled below
)

type Client struct {

	// API Key based auth
	apiKey  string
	apiSKey string

	// OAuth2 based auth
	accessToken           string
	accessTokenExpiration time.Time
	refreshToken          string

	httpClient     *http.Client
	sessionHeaders map[string]string

	lastRequest int64
	reqLock     sync.Mutex

	rateLimit atomic.Int64
}

type Credentials struct {
	AccessToken string // Bearer token, either a 'test' development key, or one obtained via OAuth2

	ApiKey  string
	ApiSKey string
}

//
//// AuthExpiredError returned when token needs to be refreshed
//type AuthExpiredError struct{}
//
//func (e *AuthExpiredError) Error() string {
//	return fmt.Sprint("Authentication token expired")
//}

func NewClient(creds *Credentials) (cb CoinbaseClient) {
	c := &Client{
		httpClient: &http.Client{Timeout: time.Second * 10},
	}
	c.rateLimit.Store(DefaultAPIRateLimit)
	c.sessionHeaders = make(map[string]string)
	if creds != nil {
		c.accessToken = creds.AccessToken
		c.apiKey = creds.ApiKey
		c.apiSKey = creds.ApiSKey
	}
	if len(c.accessToken) > 0 {
		c.sessionHeaders["Authorization"] = "Bearer " + c.accessToken
	}
	return c
}

func (c *Client) HttpClient() *http.Client {
	return c.httpClient
}

func (c *Client) AddSessionHeader(k string, v string) {
	c.sessionHeaders[k] = v
}

// Rate limit the client since Coinbase only allows API calls / second.
// Default is 100ms - DefaultAPIRateLimit
func (c *Client) SetRateLimit(ms int64) {
	c.rateLimit.Store(ms)
}

func (c *Client) checkThrottle() {
	limit := c.rateLimit.Load()

	if limit > 0 {
		c.reqLock.Lock()
		defer c.reqLock.Unlock()

		tm := time.Now().UnixMilli()
		df := tm - c.lastRequest
		if df < limit {
			waitTm := limit - df
			time.Sleep(time.Millisecond * time.Duration(waitTm))
		}
		c.lastRequest = time.Now().UnixMilli()
	}
}

// GetAndDecode retrieves from the endpoint and unmarshals resulting json into
// the provided destination interface, which must be a pointer.
func (c *Client) GetAndDecode(ctx context.Context, URL url.URL, dest interface{}, headers *map[string]string, urlValues *map[string]string) error {
	//if time.Now().After(c.AccessTokenExpiration) {
	//	return &AuthExpiredError{}
	//}

	v := url.Values{}
	if urlValues != nil {
		for key, val := range *urlValues {
			v.Add(key, val)
		}
	}
	URL.RawQuery = v.Encode()

	if req, err := http.NewRequestWithContext(ctx, http.MethodGet, URL.String(), nil); err != nil {
		return err
	} else if req == nil {
		return fmt.Errorf("unable to create request")
	} else {
		if len(c.sessionHeaders) > 0 {
			for key, val := range c.sessionHeaders {
				req.Header.Add(key, val)
			}
		}
		if headers != nil {
			for key, val := range *headers {
				req.Header.Add(key, val)
			}
		}
		c.CheckAuthentication(req, nil)
		return c.DoAndDecode(req, dest)
	}
}

// PostAndDecode retrieves from the endpoint and unmarshals resulting json into
// the provided destination interface, which must be a pointer.
func (c *Client) PostAndDecode(ctx context.Context, URL url.URL, dest interface{}, headers *map[string]string, urlValues *map[string]string, payload []byte) error {
	//if c.AccessToken != "" {
	//	if time.Now().After(c.AccessTokenExpiration) {
	//		return &AuthExpiredError{}
	//	}
	//}

	v := url.Values{}
	if urlValues != nil {
		for key, val := range *urlValues {
			v.Set(key, val)
		}
	}
	URL.RawQuery = v.Encode()
	uStr := URL.String()
	if req, err := http.NewRequestWithContext(ctx, http.MethodPost, uStr, bytes.NewReader(payload)); err != nil {
		return err
	} else if req == nil {
		return fmt.Errorf("unable to create request")
	} else {
		if len(c.sessionHeaders) > 0 {
			for key, val := range c.sessionHeaders {
				req.Header.Add(key, val)
			}
		}
		if headers != nil {
			for key, val := range *headers {
				req.Header.Add(key, val)
			}
		}
		c.CheckAuthentication(req, payload)
		return c.DoAndDecode(req, dest)
	}
}

// DoAndDecode provides useful abstractions around common errors and decoding
// issues. Ideally unmarshals into `dest`. On error, it'll use the Webull `ErrorBody` model.
// Last fallback is a plain interface.
func (c *Client) DoAndDecode(req *http.Request, dest interface{}) (err error) {

	c.checkThrottle()

	// st := time.Now()

	req.Header.Add("Content-Type", "application/json")
	res, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}

	defer func() {
		_ = res.Body.Close()
	}()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("got read error on body: %s", err.Error())
	}

	//et := time.Now()
	//df := et.Sub(st)
	//fmt.Printf("[%d], [%s] [%s] \n", df.Milliseconds(), req.URL.Path, req.URL.RawQuery)

	if res.StatusCode/100 != 2 {
		return fmt.Errorf(string(body))
	}
	if err = json.Unmarshal(body, &dest); err != nil {
		return fmt.Errorf(string(body))
	}
	return err
}

func (c *Client) apiKeyAuth(req *http.Request, reqPath string, body []byte) {

	// The CB-ACCESS-SIGN header is generated by creating a sha256 HMAC object
	// using the API secret key on the string timestamp + method + requestPath + body.

	ts := strconv.FormatInt(time.Now().UTC().Unix(), 10)
	message := ts + req.Method + reqPath + string(body)
	h := hmac.New(sha256.New, []byte(c.apiSKey))
	h.Write([]byte(message))
	signature := hex.EncodeToString(h.Sum(nil))

	req.Header.Set("CB-ACCESS-KEY", c.apiKey)
	req.Header.Set("CB-ACCESS-TIMESTAMP", ts)
	req.Header.Set("CB-ACCESS-SIGN", signature)
}

func (c *Client) addInt32Param(params map[string]string, k string, i int32) {
	vStr := strconv.FormatInt(int64(i), 10)
	params[k] = vStr
}

func (c *Client) addStringParam(params map[string]string, k string, v string) {
	if len(v) > 0 {
		params[k] = v
	}
}

func (c *Client) CheckAuthentication(req *http.Request, body []byte) {

	if len(c.accessToken) > 0 {
		// ok, using Oauth2, it will be
		// added as Auth header  Bearer
	}

	apiSig, ok := req.Header["CB-ACCESS-SIGN"]
	if ok && len(apiSig) > 0 {
		return
	}

	reqPath := req.URL.Path
	// API Key auth.
	c.apiKeyAuth(req, reqPath, body)
}

func (c *Client) IsTokenValid(_ int64) bool {
	// TODO -- check expiry and renew
	return true
}
