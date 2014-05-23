package main

import (
	"bytes"
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

var (
	URL, token, canal, mensagem string
)

func buildURL() *url.URL {
	u, err := url.Parse(URL)
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
	flag.StringVar(&URL, "url", "https://cadena-monde.slack.com/services/hooks/slackbot", "URL do serviço do slackbot")
	flag.StringVar(&token, "token", "", "Slack token")
	flag.StringVar(&canal, "canal", "", "Canal, exemplo: #random")
	flag.StringVar(&mensagem, "mensagem", "", "Mensagem para enviar ao canal")
	flag.Parse()

	if mensagem == "" {
		log.Fatal("Informe a mensagem")
	}
}
