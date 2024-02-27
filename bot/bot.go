package bot

import (
	"fmt"

	"github.com/Sourjaya/go-discord-ping/config"
	"github.com/bwmarrin/discordgo"
)

var BotID string

//var discord *discordgo.Session

func Start() {
	discord, err := discordgo.New("Bot " + config.Token)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	user, err := discord.User("@me")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	BotID = user.ID
	discord.AddHandler(messageHandler)
	err = discord.Open()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Bot is running!")
}
func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == BotID {
		return
	}
	if m.Content == "ping" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "pong")
	}
}
