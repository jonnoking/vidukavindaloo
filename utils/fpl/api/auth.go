package api

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

var AuthCookies map[string]FPLCookie

//FPLCookie for persisting auth cookies
type FPLCookie struct {
	Name       string    `json:"name"`
	Value      string    `json:"value"`
	Path       string    `json:"path"`
	Domain     string    `json:"domain"`
	HTTPOnly   bool      `json:"http_only"`
	Secure     bool      `json:"secure"`
	MaxAge     int       `json:"max_age"`
	RawExpires string    `json:"raw_expires"`
	Acquired   time.Time `json:"acquired`
}

// BuildFPLRequest build a request object with current auth cookies populated
func (api *API) BuildFPLRequest(apiURL string, method string) (*http.Request, error) {

	var fplCookies map[string]FPLCookie

	r, _ := http.NewRequest(method, apiURL, nil)

	cookies, err := api.readCookieCache()
	isValid, vErr := api.validateCookies(cookies)

	if err != nil || vErr != nil || !isValid {
		cookies, rcErr := api.refreshCookies()
		api.cacheCookies(cookies)
		fplCookies = cookies
		if rcErr != nil {
			return nil, rcErr
		}
	}

	fplCookies = cookies

	for cookieName, cookie := range fplCookies {
		if cookieName != "elevate" {
			//log.Println(cookie.Name)
			r.AddCookie(&http.Cookie{
				Name:   cookie.Name,
				Value:  cookie.Value,
				Domain: cookie.Domain,
				Path:   cookie.Path,
			})
		}
	}

	r.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.11; rv:40.0) Gecko/20100101 Firefox/40.0'")

	return r, nil
}

// RefreshCookies get auth cooies from FPL
func (api *API) refreshCookies() (map[string]FPLCookie, error) {

	//cfg := config.New()

	fplCookies := make(map[string]FPLCookie)

	loginURL := "https://users.premierleague.com/accounts/login/"

	data := url.Values{}
	data.Set("password", api.Config.Login.Password)
	data.Set("login", api.Config.Login.User)
	data.Set("redirect_uri", api.Config.Login.RedirectURI)
	data.Set("app", api.Config.Login.App)

	u, _ := url.ParseRequestURI(loginURL)
	log.Println("URL: ", data.Encode())

	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			// required to ensure that the cookies are accessible
			return http.ErrUseLastResponse
		}}

	r, _ := http.NewRequest(http.MethodPost, u.String(), strings.NewReader(data.Encode()))
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))
	// required otherwise cache proxy will intercept request
	r.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.11; rv:40.0) Gecko/20100101 Firefox/40.0'")

	resp, respErr := client.Do(r)
	check(respErr)
	defer resp.Body.Close()

	log.Println("Number of cookies: ", len(resp.Cookies()))

	if resp.StatusCode >= 400 {
		e := errors.New("FPL login failed with status " + strconv.Itoa(resp.StatusCode))
		return nil, e
	}

	now := time.Now().UTC()

	for _, cookie := range resp.Cookies() {
		cc := FPLCookie{
			Name:       cookie.Name,
			Value:      cookie.Value,
			Domain:     cookie.Domain,
			HTTPOnly:   cookie.HttpOnly,
			MaxAge:     cookie.MaxAge,
			Path:       cookie.Path,
			RawExpires: cookie.RawExpires,
			Secure:     cookie.Secure,
			Acquired:   now,
		}

		fplCookies[cookie.Name] = cc
	}
	return fplCookies, nil
}

// CacheCookies saves FPL auth cookies to file
// should read from global vars
func (api *API) cacheCookies(cookies map[string]FPLCookie) error {

	file, fErr := json.MarshalIndent(cookies, "", "")
	if fErr != nil {
		return nil
	}

	sErr := ioutil.WriteFile("./fpl-auth-cache.json", file, 0644)
	if sErr != nil {
		return nil
	}
	return nil
}

// ReadCookieCache reads the cookies from file and into memory
func (api *API) readCookieCache() (map[string]FPLCookie, error) {
	c := map[string]FPLCookie{}
	cookies, err := ioutil.ReadFile("./fpl-auth-cache.json")
	if err != nil {
		return nil, err
	}

	json.Unmarshal(cookies, &c)

	return c, nil
}

// ValidateCookies verifies that pl_profile or sessionid expiry date has not passed
func (api *API) validateCookies(cookies map[string]FPLCookie) (bool, error) {

	now := time.Now()

	sessionIDCookie := cookies["sessionid"]
	plProfileCookie := cookies["pl_profile"]

	var sidDuration time.Duration = -8 * time.Hour

	if sessionIDCookie.Acquired.Add(sidDuration).After(now.Add(-1000)) {
		return false, nil
	}

	if plProfileCookie.Acquired.Add(sidDuration).After(now.Add(-1000)) {
		return false, nil
	}

	return true, nil
}
