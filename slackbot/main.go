package main

import (
	"github.com/cadena-monde/slackbot"
	"os"
	"flag"
)

var (
	canal, mensagem string
	slackURL        = os.Getenv("SLACKBOT_URL")
	token           = os.Getenv("SLACKBOT_TOKEN")
)

func main() {
	parseFlags()
	postMessage();
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

func postMessage() {
	b := slackbot.New(slackURL, token)
	b.PostMessage(canal, mensagem)
}
