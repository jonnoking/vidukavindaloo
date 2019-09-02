package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

func main() {
	loginFPL()
}

type FPLLogin struct {
	Password    string `json:"password"`
	Login       string `json:"login"`
	RedirectURL string `json:"redirect_uri"`
	App         string `json:"app"`
}

type FPLCookie struct {
	Name       string `json:"name"`
	Value      string `json:"value"`
	Path       string `json:"path"`
	Domain     string `json:"domain"`
	HttpOnly   bool   `json:"http_only"`
	Secure     bool   `json:"secure"`
	MaxAge     int    `json:"max_age"`
	RawExpires string `json:"raw_expires"`
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

	getMyTeam(d)
}

func getMyTeam(cookies []*http.Cookie) {
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
