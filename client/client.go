package client

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/rolex01/hearthstone_api_client/utils"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
	"io/ioutil"
	"net/http"
	"strings"
)

var (
	client *Client
)

type Client struct {
	client                 *http.Client
	cfg                    clientcredentials.Config
	authorizedCfg          oauth2.Config
	oauth                  OAuth
	oauthURL               string
	apiURL                 string
	dynamicNamespace       string
	profileNamespace       string
	staticNamespace        string
	staticClassicNamespace string
	region                 Region
	locale                 Locale
}

type OAuth struct {
	ClientID     string
	ClientSecret string
	Token        *oauth2.Token
}

type Region int

const (
	_ Region = iota
	US
	EU
	KR
	TW
	CN
)

func (region Region) String() string {
	var r = []string{"", "us", "eu", "kr", "tw", "cn"}

	return r[region]
}

func (region Region) Int() int {
	return int(region)
}

type Locale string

const (
	DeDE = Locale("de_DE")
	EnUS = Locale("en_US")
	EsES = Locale("es_ES")
	EsMX = Locale("es_MX")
	FrFR = Locale("fr_FR")
	ItIT = Locale("it_IT")
	JaJP = Locale("ja_JP")
	KoKR = Locale("ko_KR")
	PlPL = Locale("pl_PL")
	PtBR = Locale("pt_BR")
	RuRU = Locale("ru_RU")
	ThTH = Locale("th_TH")
	ZhCN = Locale("zh_CN")
	ZhTW = Locale("zh_TW")
)

func (locale Locale) String() string {
	return string(locale)
}

func NewClient() *Client {
	clientID := utils.GetEnvString("HSAPI_ID", "")
	clientSecret := utils.GetEnvString("HSAPI_SECRET", "")
	region := Region(utils.GetEnvInt("HSAPI_REGION", EU.Int()))
	locale := Locale(utils.GetEnvString("HSAPI_LOCALE", RuRU.String()))

	var c = Client{
		oauth: OAuth{
			ClientID:     clientID,
			ClientSecret: clientSecret,
		},
		locale: locale,
	}

	c.cfg = clientcredentials.Config{
		ClientID:     c.oauth.ClientID,
		ClientSecret: c.oauth.ClientSecret,
	}

	c.SetRegion(region)

	return &c
}

func (c *Client) SetRegion(region Region) {
	c.region = region

	switch region {
	case CN:
		c.oauthURL = "https://www.battlenet.com.cn"
		c.apiURL = "https://gateway.battlenet.com.cn"
		c.dynamicNamespace = "dynamic-zh"
		c.profileNamespace = "profile-zh"
		c.staticNamespace = "static-zh"
		c.staticClassicNamespace = "static-classic-zh"
	default:
		c.oauthURL = fmt.Sprintf("https://%s.battle.net", region)
		c.apiURL = fmt.Sprintf("https://%s.api.blizzard.com", region)
		c.dynamicNamespace = fmt.Sprintf("dynamic-%s", region)
		c.profileNamespace = fmt.Sprintf("profile-%s", region)
		c.staticNamespace = fmt.Sprintf("static-%s", region)
		c.staticClassicNamespace = fmt.Sprintf("static-classic-%s", region)
	}

	c.cfg.TokenURL = c.oauthURL + "/oauth/token"
	c.client = c.cfg.Client(context.Background())
}

func (c *Client) getURLBody(url, namespace string) ([]byte, error) {
	var (
		req  *http.Request
		res  *http.Response
		body []byte
		err  error
	)

	req, err = http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/json")

	if namespace != "" {
		req.Header.Set("Battlenet-Namespace", namespace)
	}

	res, err = c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	switch res.StatusCode {
	case http.StatusNotFound:
		return nil, errors.New(res.Status)
	}

	return body, nil
}

func (c *Client) AccessTokenRequest() (err error) {
	var (
		req *http.Request
		res *http.Response
		b   []byte
	)

	req, err = http.NewRequest("POST", c.oauthURL+"/oauth/token",
		strings.NewReader("grant_type=client_credentials"),
	)
	if err != nil {
		return
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	res, err = c.client.Do(req)
	if err != nil {
		return
	}
	defer res.Body.Close()

	b, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}

	err = json.Unmarshal(b, &c.oauth.Token)

	return
}
