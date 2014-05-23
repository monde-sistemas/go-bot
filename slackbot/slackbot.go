package main

import (
	"bytes"
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
)

var (
	canal, mensagem string
	slackURL        = os.Getenv("SLACKBOT_URL")
	token           = os.Getenv("SLACKBOT_TOKEN")
)

func buildURL() *url.URL {
	u, err := url.Parse(slackURL)
	if err != nil {
		log.Fatal(err)
	}
	q := u.Query()
	q.Set("token", token)
	q.Set("channel", canal)
	u.RawQuery = q.Encode()

	log.Println("URL: ", u.String())
	return u
}

func postMessage(postURL *url.URL) {
	r, err := http.Post(postURL.String(), "text/plain", bytes.NewBufferString(mensagem))
	if err != nil {
		log.Fatalf("Erro ao postar: %s", err)
	}
	defer r.Body.Close()
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(b))
}

func main() {
	parseFlags()
	postMessage(buildURL())
}

func parseFlags() {
	flag.StringVar(&canal, "canal", "", "Canal. Ex: #random")
	flag.StringVar(&mensagem, "mensagem", "", "Mensagem para enviar ao canal")
	flag.Parse()

	if canal == "" || mensagem == "" {
		flag.Usage()
		os.Exit(1)
	}
}
