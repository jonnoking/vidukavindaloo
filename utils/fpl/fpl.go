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
	r.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.11; rv:40.0) Gecko/20100101 Firefox/40.0'")
	r.Header.Add("Connection", "keep-alive")
	r.Header.Add("Accept", "*/*")
	//r.Header.Add("Accept-Encoding", "gzip, deflate")
	r.Header.Add("Cache-Control", "no-cache")

	resp, respErr := client.Do(r)
	check(respErr)
	defer resp.Body.Close()

	byteValue, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	// zipReader, zErr := zip.NewReader(bytes.NewReader(byteValue), int64(len(byteValue)))
	// check(zErr)

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

	log.Println("--- PRINTING COOKIES ---")
	for _, cookie := range resp.Cookies() {
		log.Printf(cookie.Name)
		log.Println(cookie.Value)
	}

	for _, h := range resp.Request.Cookies() {
		log.Printf("%v - %v", h.Name, h.Value)
	}
}

func login() {

	url := "https://users.premierleague.com/accounts/login"
	//url := "https://fantasy.premierleague.com/api/my-team/1759299/"

	fplClient := http.Client{
		Timeout: time.Second * 5,
	}
	req, rErr := http.NewRequest(http.MethodPost, url, nil)
	check(rErr)

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("password", "")
	req.Header.Add("login", "jonno.king@gmail.com")
	req.Header.Add("redirect_uri", "https://fantasy.premierleague.com/")
	req.Header.Add("app", "plfpl-web")

	res, resErr := fplClient.Do(req)
	check(resErr)

	byteValue, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	var result map[string]interface{}
	json.Unmarshal([]byte(byteValue), &result)

	log.Printf("%v - %v", res.StatusCode, res.Status)
	log.Println(res.Cookies)
	log.Println(res.Header)
	log.Println(res.Body)
	log.Print(result)

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
