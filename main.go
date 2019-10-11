package main

import (
	"context"
	"crypto/tls"
	"fmt"
	api "github.com/jonnoking/vidukavindaloo/utils/fpl/api"
	"github.com/jonnoking/vidukavindaloo/utils/fpl/models"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/jonnoking/vidukavindaloo/utils/config"
	"github.com/jonnoking/vidukavindaloo/utils/fpl"

	//"./utils/config"
	"golang.org/x/crypto/acme/autocert"
	// "html/template"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"time"
)

// init is invoked before main()
func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

var VVLeague map[int]*models.Entry = map[int]*models.Entry{}

func getEntry(entryID int, total int, rank int, wg *sync.WaitGroup) {
	e, err := api.GetEntry(entryID)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%d %s\n", total, e.Name)
	VVLeague[entryID] = e

	wg.Done()
}

func GetLeagueParticipantsAysnc() {
	// Get the entries of a league using goroutines
	league, lErr := api.GetClassicLeague(1132753)
	if lErr != nil {
		panic(lErr)
	}

	var wg sync.WaitGroup
	wg.Add(len(league.Standings.Results))

	for _, e := range league.Standings.Results {
		go getEntry(e.Entry, e.Total, e.Rank, &wg)

	}
	wg.Wait()

	fmt.Println("Finished getting league participants")
	fmt.Printf("%s\n", VVLeague[4764698].PlayerFirstName)
}

type CR struct {
	EntryID int
	Entry   *models.Entry
}

func GetLeaguesChannels() {

	entries := make(chan CR)

	league, lErr := api.GetClassicLeague(1132753)
	if lErr != nil {
		panic(lErr)
	}

	var wg sync.WaitGroup
	wg.Add(len(league.Standings.Results))

	for _, e := range league.Standings.Results {

		go func(entryID int) {
			defer wg.Done()
			e, err := api.GetEntry(entryID)
			if err != nil {
				panic(err)
			}
			entries <- CR{EntryID: entryID, Entry: e}
		}(e.Entry)

	}

	// as channel finishes
	go func() {
		for entryResponse := range entries {
			VVLeague[entryResponse.EntryID] = entryResponse.Entry
			fmt.Printf("Entry: %s\n", entryResponse.Entry.Name)
		}
	}()

	wg.Wait()

	fmt.Println("Finished getting league participants")
	fmt.Printf("%s\n", VVLeague[4764698].PlayerFirstName)

}

func main() {

	fpl.LoadFromLive()
	entry, _ := api.GetEntry(4719576)
	fmt.Printf("Leagues: %d\n", len(entry.Leagues))
	fmt.Println(entry.Name)

	ef, _ := api.GetCompleteEntry(1759299)
	fmt.Printf("Leagues: %d\n", len(ef.Leagues))
	fmt.Println(ef)

	//	GetLeagueParticipantsAysnc()
	GetLeaguesChannels()

	// var input string
	// fmt.Scanln(&input)

	// t, _ := fpl.Teams.GetTeamByName("Southampton")
	// log.Println(t)

	// log.Printf("Teams: %d \n", len(fpl.Teams.Teams))
	// log.Printf("Phases: %d \n", len(fpl.Phases.Phases))
	// log.Printf("Player Types: %d \n", len(fpl.PlayerTypes.PlayerTypes))
	// log.Printf("Players: %d \n", len(fpl.Players.Players))
	// log.Printf("Events: %d \n", len(fpl.Events.Events))

	// mt, _ := fpl.GetMyTeam(1759299, fpl.Players, fpl.Teams, fpl.PlayerTypes)

	// for _, v := range mt.Picks {
	// 	fmt.Printf("%v [%s] %s\n", v.Player.GetFullName(), v.Position, v.Team.Name)
	// 	//fmt.Println(v.Player)
	// }

	//fmt.Printf("%+v", mt)

	// for code, team := range fpl.Teams.Teams {
	// 	log.Println(code, team.Name)
	// }

	// for code, player := range fpl.Players.Players {
	// 	//		log.Println(code, player.TeamCode, player.WebName)
	// 	if player.TeamCode == 1 {
	// 		log.Println(code, player.GetFullName())
	// 		log.Println(fmt.Sprintf("%+v", player))
	// 	}
	// }
	//will, _ := fpl.Players.GetPlayerByFullName("Will Norris")
	// wn := fpl.Players.Players[168399]
	// //168399
	// log.Println(fmt.Sprintf("%+v", wn))

	//fpl.GetScreenshot()

	//runServer()

}

func runServer() {
	conf := config.New()

	serverConfig := HttpServerConfig{
		Host:         conf.HTTP.HTTPHost,
		Port:         conf.HTTP.HTTPPort,
		HTTPSDomains: conf.HTTP.HTTPSDomains,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	httpServer := Start(serverConfig)
	defer httpServer.Stop()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	<-sigChan

	fmt.Printf("\nmain : shutting down")
}

//HttpServerConfig configuration for http server
type HttpServerConfig struct {
	Host         string
	Port         int
	HTTPSDomains []string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type HttpServer struct {
	server *http.Server
	wg     sync.WaitGroup
}

// HomeHandler Will serve default index.html page
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./static/html/index.html")
}

// PingHandler Will serve default ping.html page
func PingHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./static/html/ping.html")
}

func Start(cfg HttpServerConfig) *HttpServer {

	_, cancel := context.WithCancel(context.Background())
	defer cancel()

	router := mux.NewRouter()
	router.HandleFunc("/", HomeHandler)
	router.HandleFunc("/home", HomeHandler)
	router.HandleFunc("/ping", PingHandler)
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static", http.FileServer(http.Dir("./static"))))

	hostPolicy := func(ctx context.Context, host string) error {
		allowedHost := "cloudjonno.com"
		if host == allowedHost {
			return nil
		}
		return fmt.Errorf("\nacme/autocert: only %s host is allowed", allowedHost)
	}

	m := &autocert.Manager{
		Cache:      autocert.DirCache("vv-autocert"),
		Prompt:     autocert.AcceptTOS,
		HostPolicy: hostPolicy,
	}
	//autocert.HostWhitelist("54apenwith.com", "www.54apenwith.com"),

	httpServer := HttpServer{
		server: &http.Server{
			Addr:              ":" + strconv.Itoa(cfg.Port),
			TLSConfig:         &tls.Config{GetCertificate: m.GetCertificate},
			Handler:           router,
			ReadHeaderTimeout: cfg.ReadTimeout,
			WriteTimeout:      cfg.WriteTimeout,
			MaxHeaderBytes:    1 << 20,
		},
	}

	httpServer.wg.Add(1)

	go func() {
		fmt.Printf("\nHTTPServer : Service started : Host=%v\n", cfg.Host)
		httpServer.server.ListenAndServe()
		//httpServer.server.ListenAndServeTLS("", "")
		httpServer.wg.Done()
	}()

	return &httpServer
}

func (httpServer *HttpServer) Stop() error {
	const timeout = 5 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	fmt.Printf("\nHTTPServer : Service stopping\n")

	if err := httpServer.server.Shutdown(ctx); err != nil {
		if err := httpServer.server.Close(); err != nil {
			fmt.Printf("\nHTTPServer : Service stopping : Error=%v", err)
			return err
		}
	}

	httpServer.wg.Wait()
	fmt.Printf("\nHTTPServer : Stopped\n")
	return nil
}
