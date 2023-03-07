package handlers

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/javtor/renato/search"
	"github.com/mcstatus-io/mcutil"
)

const (
	chontaduroAddress = "chontaduroland.sparked.miami"
)

func Pong(session *discordgo.Session, message *discordgo.MessageCreate) error {
	_, err := session.ChannelMessageSend(message.ChannelID, "pong")
	return err
}

func Chontaduro(session *discordgo.Session, message *discordgo.MessageCreate) error {
	response, err := mcutil.Status(chontaduroAddress, 25565)

	if err != nil {
		return err
	}

	msg := fmt.Sprintf("Hay %d jugadores activos:\n", response.Players.Online)
	for _, player := range response.Players.Sample {
		msg += fmt.Sprintf("- %s\n", player.NameClean)
	}

	_, err = session.ChannelMessageSend(message.ChannelID, msg)
	return err
}

func SearchImage(session *discordgo.Session, message *discordgo.MessageCreate) error {
	image, err := search.GetRandomImage(message.Content[1:])
	if err != nil {
		return err
	}
	_, err = session.ChannelMessageSend(message.ChannelID, image)
	return err
}
