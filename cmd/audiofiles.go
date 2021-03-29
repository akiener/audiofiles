package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"
	"github.com/zmb3/spotify"
	"golang.org/x/oauth2/clientcredentials"

	"muse/database"
)

func main() {
	db, err := database.New()
	if err != nil {
		log.WithError(err).Fatal("failed to connect to database")
	}
	defer db.Close()

	songs, err := db.Test()
	if err != nil {
		log.WithError(err).Fatal("failed to test database connection")
	}
	log.Infof("found %d songs in the database!", len(*songs))

	config := &clientcredentials.Config{
		ClientID:     os.Getenv("AUDIOFILES_SPOTIFY_CLIENT_ID"),
		ClientSecret: os.Getenv("AUDIOFILES_SPOTIFY_CLIENT_SECRET"),
		TokenURL:     spotify.TokenURL,
	}
	token, err := config.Token(context.Background())
	if err != nil {
		log.Fatalf("couldn't get token: %v", err)
	}

	client := spotify.Authenticator{}.NewClient(token)
	search, err := client.Search("aLIEz", spotify.SearchTypeTrack)
	if err != nil {
		log.Fatalf("couldn't get features playlists: %v", err)
	}
	println(search)

	discord, err := discordgo.New("Bot " + os.Getenv("AUDIOFILES_BOT_TOKEN"))
	if err != nil {
		log.WithError(err).Fatal("failed to connect to discord")
	}

	discord.AddHandler(messageCreate)

	// In this example, we only care about receiving message events.
	discord.Identify.Intents = discordgo.IntentsDirectMessages

	// Open a websocket connection to Discord and begin listening.
	err = discord.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}
	defer discord.Close()

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	log.Info(m.Content)

	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}
	// If the message is "ping" reply with "Pong!"
	if m.Content == "ping" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "Pong!")
	}

	// If the message is "pong" reply with "Ping!"
	if m.Content == "pong" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "Ping!")
	}
}
