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

var (
	activePlayers = make(map[string]bool)
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

// Gets the active players from the server
// sends a message to the channel if someone enters or leaves
func ChontaduroCron(session *discordgo.Session, channelID string) error {
	response, err := mcutil.Status(chontaduroAddress, 25565)
	if err != nil {
		return err
	}

	for _, player := range response.Players.Sample {
		if _, ok := activePlayers[player.NameClean]; !ok {
			activePlayers[player.NameClean] = true
			_, err = session.ChannelMessageSend(channelID, fmt.Sprintf("%s entró a Chontaduroland!", player.NameClean))
			if err != nil {
				return err
			}
		}
	}

	for player := range activePlayers {
		found := false
		for _, p := range response.Players.Sample {
			if p.NameClean == player {
				found = true
				break
			}
		}
		if !found {
			delete(activePlayers, player)
			_, err = session.ChannelMessageSend(channelID, fmt.Sprintf("%s salió de Chontaduroland!", player))
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func SearchImage(session *discordgo.Session, message *discordgo.MessageCreate) error {
	image, err := search.GetRandomImage(message.Content[1:])
	if err != nil {
		return err
	}
	_, err = session.ChannelMessageSend(message.ChannelID, image)
	return err
}
