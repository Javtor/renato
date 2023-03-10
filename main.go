package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/go-co-op/gocron"
	"github.com/javtor/renato/handlers"
	"github.com/joho/godotenv"
)

var (
	token     string
	channelID string
)

func goDotEnvVariable(key string) string {
	// load .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	return os.Getenv(key)
}

func messageCreate(session *discordgo.Session, message *discordgo.MessageCreate) {
	if message.Author.ID == session.State.User.ID {
		return
	}

	if len(message.Content) == 0 {
		return
	}

	if message.Content == "ping" {
		err := handlers.Pong(session, message)
		if err != nil {
			log.Println(err)
		}
		log.Println("Ponged")
		return
	}

	if message.Content == "!chontaduro" {
		err := handlers.Chontaduro(session, message)
		if err != nil {
			log.Println(err)
		}
		log.Println("Chontaduro")
		return
	}

	if message.Content[0] == '!' {
		err := handlers.SearchImage(session, message)
		if err != nil {
			log.Println(err)
		}
		log.Println("Search")
		return
	}
}

func main() {
	token := goDotEnvVariable("BOT_TOKEN")
	channelID = goDotEnvVariable("CHANNEL_ID")

	discord, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Fatal(err)
	}

	discord.AddHandler(messageCreate)

	discord.Identify.Intents = discordgo.IntentsGuildMessages

	err = discord.Open()
	if err != nil {
		log.Fatal(err)
	}

	chontaduroCron := gocron.NewScheduler(time.UTC)
	chontaduroCron.Every(5).Seconds().Do(handlers.ChontaduroCron, discord, channelID)
	chontaduroCron.StartAsync()

	log.Println("Bot is running. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	discord.Close()
}
