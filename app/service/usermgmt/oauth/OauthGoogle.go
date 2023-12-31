package oauth

import (
	"bytes"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"hash"
	"idstar-idp/rest-api/app/config"
	"idstar-idp/rest-api/app/dto/request/login"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

type OauthGoogle struct {
	ProviderName string
	OauthPath    string
	clientId     string
	clientSecret string
	scope        string
	redirectUri  string
	authUri      string
	tokenUri     string
	userInfoUri  string
	timeout      int
	sha1         hash.Hash
}

type GoogleToken struct {
	AccessToken string
	ExpireIn    time.Time
	IdToken     string
}

func InitOauthGoogle() *OauthGoogle {
	return &OauthGoogle{
		ProviderName: "Google",
		OauthPath:    "google",
		clientId:     config.GetConfigValue("oauth.google.client.id"),
		clientSecret: config.GetConfigValue("oauth.google.client.secret"),
		scope:        "email profile",
		redirectUri:  config.GetConfigValue("oauth.google.uri.redirect"),
		authUri:      config.GetConfigValue("oauth.google.uri.auth"),
		tokenUri:     config.GetConfigValue("oauth.google.uri.token"),
		userInfoUri:  config.GetConfigValue("oauth.google.uri.userinfo"),
		timeout:      config.GetConfigIntValue("oauth.timeout.ms"),
		sha1:         sha1.New(),
	}
}

func (o *OauthGoogle) getState() string {
	reqTime := time.Now().Format(time.RFC3339Nano)
	o.sha1.Write([]byte(reqTime))
	return hex.EncodeToString(o.sha1.Sum(nil))
}

func (o *OauthGoogle) GetGoogleAuthAddress() string {
	values := url.Values{}
	values.Add("response_type", "code")
	values.Add("scope", o.scope)
	values.Add("client_id", o.clientId)
	values.Add("redirect_uri", o.redirectUri)
	values.Add("state", o.getState())

	return o.authUri + "?" + values.Encode()
}

func (o *OauthGoogle) GetGoogleOauthToken(code string) (*GoogleToken, error) {
	values := url.Values{}
	values.Add("grant_type", "authorization_code")
	values.Add("code", code)
	values.Add("client_id", o.clientId)
	values.Add("client_secret", o.clientSecret)
	values.Add("redirect_uri", o.redirectUri)

	req, err := http.NewRequest("POST", o.tokenUri, bytes.NewBufferString(values.Encode()))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	client := http.Client{
		Timeout: time.Millisecond * time.Duration(int64(o.timeout)),
	}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		return nil, errors.New("could not retrieve token")
	}

	var resBody bytes.Buffer
	_, err = io.Copy(&resBody, res.Body)
	if err != nil {
		return nil, err
	}
	var resMap map[string]interface{}
	if err := json.Unmarshal(resBody.Bytes(), &resMap); err != nil {
		return nil, err
	}

	expireIn, err := strconv.Atoi(resMap["access_token"].(string))
	if err != nil {
		expireIn = 3600
	}
	token := &GoogleToken{
		AccessToken: resMap["access_token"].(string),
		ExpireIn:    time.Now().Add(time.Second * time.Duration(int64(expireIn))),
		IdToken:     resMap["id_token"].(string),
	}
	return token, nil
}

func (o *OauthGoogle) GetGoogleUserInfo(accessToken string, idToken string) (*login.RegistrationLoginRequest, error) {
	values := url.Values{}
	values.Add("alt", "json")
	values.Add("access_token", accessToken)

	req, err := http.NewRequest("GET", o.userInfoUri+"?"+values.Encode(), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", idToken))
	client := http.Client{
		Timeout: time.Millisecond * time.Duration(int64(o.timeout)),
	}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		return nil, errors.New("could not retrieve user info")
	}

	var resBody bytes.Buffer
	_, err = io.Copy(&resBody, res.Body)
	if err != nil {
		return nil, err
	}
	var resMap map[string]interface{}
	if err := json.Unmarshal(resBody.Bytes(), &resMap); err != nil {
		return nil, err
	}

	userInfo := &login.RegistrationLoginRequest{
		Name:     resMap["name"].(string),
		Username: resMap["email"].(string),
	}
	if userInfo.Name == "" || userInfo.Username == "" {
		return nil, errors.New("could not retrieve user info")
	}
	return userInfo, nil
}
