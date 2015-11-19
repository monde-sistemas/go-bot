package main

import (
	"os"

	"github.com/monde-sistemas/go-bot/Godeps/_workspace/src/github.com/go-chat-bot/bot/slack"
	_ "github.com/monde-sistemas/go-bot/Godeps/_workspace/src/github.com/go-chat-bot/plugins/catfacts"
	_ "github.com/monde-sistemas/go-bot/Godeps/_workspace/src/github.com/go-chat-bot/plugins/catgif"
	_ "github.com/monde-sistemas/go-bot/Godeps/_workspace/src/github.com/go-chat-bot/plugins/chucknorris"
	_ "github.com/monde-sistemas/go-bot/Godeps/_workspace/src/github.com/go-chat-bot/plugins/cnpj"
	_ "github.com/monde-sistemas/go-bot/Godeps/_workspace/src/github.com/go-chat-bot/plugins/cotacao"
	_ "github.com/monde-sistemas/go-bot/Godeps/_workspace/src/github.com/go-chat-bot/plugins/cpf"
	_ "github.com/monde-sistemas/go-bot/Godeps/_workspace/src/github.com/go-chat-bot/plugins/crypto"
	_ "github.com/monde-sistemas/go-bot/Godeps/_workspace/src/github.com/go-chat-bot/plugins/dilma"
	_ "github.com/monde-sistemas/go-bot/Godeps/_workspace/src/github.com/go-chat-bot/plugins/encoding"
	_ "github.com/monde-sistemas/go-bot/Godeps/_workspace/src/github.com/go-chat-bot/plugins/gif"
	_ "github.com/monde-sistemas/go-bot/Godeps/_workspace/src/github.com/go-chat-bot/plugins/godoc"
	_ "github.com/monde-sistemas/go-bot/Godeps/_workspace/src/github.com/go-chat-bot/plugins/guid"
	_ "github.com/monde-sistemas/go-bot/Godeps/_workspace/src/github.com/go-chat-bot/plugins/megasena"
	_ "github.com/monde-sistemas/go-bot/Godeps/_workspace/src/github.com/go-chat-bot/plugins/puppet"
	_ "github.com/monde-sistemas/go-bot/Godeps/_workspace/src/github.com/go-chat-bot/plugins/treta"
)

func main() {
	slack.Run(os.Getenv("SLACK_API_TOKEN"))
}
