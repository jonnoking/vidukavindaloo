package main

import (
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

/*
	Get Client
	Refresh Cookies
	Get Cookies from Cache
	Check Cache Expiry
	Save Cookies to Cache

*/

func main() {

	//loginFPL()

	// cookies, err := RefreshCookies()
	// if err != nil {
	// 	log.Println(err)
	// }
	// CacheCookies(cookies)

	// cookies, _ := ReadCookieCache()
	// log.Println(cookies["pl_profile"].Value)
	// log.Println(cookies["pl_profile"].RawExpires)

	// isValid, _ := ValidateCookies(cookies)
	// log.Println(isValid)

	GetMyTeam()

}

type FPLCookie struct {
	Name       string    `json:"name"`
	Value      string    `json:"value"`
	Path       string    `json:"path"`
	Domain     string    `json:"domain"`
	HttpOnly   bool      `json:"http_only"`
	Secure     bool      `json:"secure"`
	MaxAge     int       `json:"max_age"`
	RawExpires string    `json:"raw_expires"`
	Acquired   time.Time `json:"acquired`
}

func GetMyTeam() {
	apiURL := "https://fantasy.premierleague.com/api/my-team/1759299/"

	client := &http.Client{}

	r, _ := BuildFPLRequest(apiURL, "GET")

	resp, respErr := client.Do(r)
	check(respErr)

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Panic("\nStatus Code: ", resp.StatusCode)
		log.Println(resp.Status)
		return
	}

	SaveBodyToFile(resp.Body, "./fpl-myteam.json")
}

// SaveBodyToFile save response body to file
func SaveBodyToFile(body io.ReadCloser, filename string) error {
	byteValue, readErr := ioutil.ReadAll(body)
	if readErr != nil {
		return readErr
	}

	fErr := ioutil.WriteFile(filename, byteValue, 0644)
	if fErr != nil {
		return fErr
	}
	return nil
}

// BuildFPLRequest build a request object with current auth cookies populated
func BuildFPLRequest(apiURL string, method string) (*http.Request, error) {

	var fplCookies map[string]FPLCookie

	r, _ := http.NewRequest(method, apiURL, nil)

	cookies, err := ReadCookieCache()
	isValid, vErr := ValidateCookies(cookies)

	if err != nil || vErr != nil || !isValid {
		cookies, rcErr := RefreshCookies()
		fplCookies = cookies
		if rcErr != nil {
			return nil, rcErr
		}
	}

	fplCookies = cookies

	for cookieName, cookie := range fplCookies {
		if cookieName != "elevate" {
			log.Println(cookie.Name)
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
func RefreshCookies() (map[string]FPLCookie, error) {

	fplCookies := make(map[string]FPLCookie)

	loginURL := "https://users.premierleague.com/accounts/login/"

	data := url.Values{}
	data.Set("password", "")
	data.Set("login", "jonno.king@gmail.com")
	data.Set("redirect_uri", "https://fantasy.premierleague.com/")
	data.Set("app", "plfpl-web")

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
			HttpOnly:   cookie.HttpOnly,
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
func CacheCookies(cookies map[string]FPLCookie) error {

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

// ReadCookies reads the cookies from file and into memory
func ReadCookieCache() (map[string]FPLCookie, error) {
	c := map[string]FPLCookie{}
	cookies, err := ioutil.ReadFile("./fpl-auth-cache.json")
	if err != nil {
		return nil, err
	}

	json.Unmarshal(cookies, &c)

	return c, nil
}

// ValidateCookies verifies that pl_profile or sessionid expiry date has not passed
func ValidateCookies(cookies map[string]FPLCookie) (bool, error) {

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

func loginFPL() {

	loginURL := "https://users.premierleague.com/accounts/login/"

	data := url.Values{}
	data.Set("password", "")
	data.Set("login", "jonno.king@gmail.com")
	data.Set("redirect_uri", "https://fantasy.premierleague.com/")
	data.Set("app", "plfpl-web")

	u, _ := url.ParseRequestURI(loginURL)

	log.Printf(u.String())
	log.Println("URL: ", data.Encode())

	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			// required to ensure that the cookies are accessible
			return http.ErrUseLastResponse
		}}

	r, _ := http.NewRequest(http.MethodPost, u.String(), strings.NewReader(data.Encode()))
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))
	// required otherwise proxy will intercept request
	r.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.11; rv:40.0) Gecko/20100101 Firefox/40.0'")

	resp, respErr := client.Do(r)
	check(respErr)
	defer resp.Body.Close()

	byteValue, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	fErr := ioutil.WriteFile("./loginpage.html", byteValue, 0644)
	check(fErr)

	log.Printf(resp.Status)
	//log.Println(resp.Cookies)
	log.Println(resp.Header)
	log.Println(resp.ContentLength)

	log.Println(resp.Header.Get("Server"))
	log.Println(resp.Header.Get("Content-Type"))
	log.Println(resp.Header.Get("Last-Modified"))
	log.Println(resp.Header.Get("ETag"))
	log.Println(resp.Header.Get("Content-Encoding"))

	log.Println("Number of cookies: ", len(resp.Cookies()))

	d := resp.Cookies()

	// c := []FPLCookie{}

	// log.Println("--- PRINTING COOKIES ---")
	// for _, cookie := range resp.Cookies() {
	// 	log.Printf(cookie.Name)
	// 	log.Println(cookie.Value)
	// 	log.Println(cookie.Domain)
	// 	cc := FPLCookie{
	// 		Name:       cookie.Name,
	// 		Value:      cookie.Value,
	// 		Domain:     cookie.Domain,
	// 		HttpOnly:   cookie.HttpOnly,
	// 		MaxAge:     cookie.MaxAge,
	// 		Path:       cookie.Path,
	// 		RawExpires: cookie.RawExpires,
	// 		Secure:     cookie.Secure,
	// 	}
	// 	c = append(c, cc)
	// }

	getMyTeamOld(d)
}

func getMyTeamOld(cookies []*http.Cookie) {
	apiURL := "https://fantasy.premierleague.com/api/my-team/1759299/"

	log.Println("--- PRINTING COOKIES AGAIN ---")
	for _, cookie := range cookies {
		log.Printf(cookie.Name)
		log.Println(cookie.Value)
		log.Println(cookie.Domain)
		log.Println(cookie.Path)
		log.Println(cookie.Secure)
		log.Println(cookie.RawExpires)
	}

	client := &http.Client{}
	r, _ := http.NewRequest(http.MethodGet, apiURL, nil)
	for _, cookie := range cookies {
		if cookie.Name != "elevate" {
			r.AddCookie(cookie)
			// r.AddCookie(&http.Cookie{
			// 	Name:   cookie.Name,
			// 	Value:  cookie.Value,
			// 	Domain: cookie.Domain,
			// 	Path:   cookie.Path,
			// })
		}
	}
	r.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.11; rv:40.0) Gecko/20100101 Firefox/40.0'")

	resp, respErr := client.Do(r)
	check(respErr)

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Panic("\nStatus Code: ", resp.StatusCode)
		log.Println(resp.Status)
		return
	}

	byteValue, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	fErr := ioutil.WriteFile("./myteam.json", byteValue, 0644)
	check(fErr)

}

func callGetBootstrap() map[string]interface{} {
	fplBootsrapURL := "https://fantasy.premierleague.com/api/bootstrap-static/"

	fplClient := http.Client{
		Timeout: time.Second * 5,
	}

	req, err := http.NewRequest(http.MethodGet, fplBootsrapURL, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("User-Agent", "cloudjonno")

	res, getErr := fplClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	byteValue, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	fErr := ioutil.WriteFile("./fpl-bootstrap.json", byteValue, 0644)
	check(fErr)

	var result map[string]interface{}
	json.Unmarshal([]byte(byteValue), &result)

	return result
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
