package main

import (
	"os"

	"github.com/go-chat-bot/bot/slack"
	_ "github.com/go-chat-bot/plugins/catfacts"
	_ "github.com/go-chat-bot/plugins/catgif"
	_ "github.com/go-chat-bot/plugins/chucknorris"
	_ "github.com/go-chat-bot/plugins-br/cnpj"
	_ "github.com/go-chat-bot/plugins-br/cotacao"
	_ "github.com/go-chat-bot/plugins-br/cpf"
	_ "github.com/go-chat-bot/plugins/crypto"
	_ "github.com/go-chat-bot/plugins-br/dilma"
	_ "github.com/go-chat-bot/plugins/encoding"
	_ "github.com/go-chat-bot/plugins/gif"
	_ "github.com/go-chat-bot/plugins/godoc"
	_ "github.com/go-chat-bot/plugins/guid"
	_ "github.com/go-chat-bot/plugins-br/megasena"
	_ "github.com/go-chat-bot/plugins/puppet"
	_ "github.com/go-chat-bot/plugins/treta"
)

func main() {
	slack.Run(os.Getenv("SLACK_API_TOKEN"))
}
