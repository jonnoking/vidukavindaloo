package main

import (
	//"bytes"
	"context"
	"crypto/tls"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/jonnoking/vidukavindaloo/utils/config"
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

func main() {

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

	// Print out environment variables
	// fmt.Println(conf.FPL.Me)
	// fmt.Println(conf.FPL.VVLeague)
	// fmt.Println(conf.DebugMode)
	// fmt.Println(conf.MaxUsers)

	// // Print out each role
	// for _, role := range conf.UserRoles {
	// 	fmt.Println(role)
	// }
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
