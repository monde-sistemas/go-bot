package main

import (
	"os"

	"github.com/fabioxgn/go-bot"
	_ "github.com/fabioxgn/go-bot/commands/catfacts"
	_ "github.com/fabioxgn/go-bot/commands/catgif"
	_ "github.com/fabioxgn/go-bot/commands/chucknorris"
	_ "github.com/fabioxgn/go-bot/commands/cotacao"
	_ "github.com/fabioxgn/go-bot/commands/dilma"
	_ "github.com/fabioxgn/go-bot/commands/gif"
	_ "github.com/fabioxgn/go-bot/commands/godoc"
	_ "github.com/fabioxgn/go-bot/commands/megasena"
	_ "github.com/fabioxgn/go-bot/commands/puppet"
)

func main() {
	bot.RunSlack(os.Getenv("SLACK_API_TOKEN"))
}
