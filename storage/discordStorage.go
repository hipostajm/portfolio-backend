package storage

import (
	"github.com/clinet/discordgo-embed"
	"github.com/bwmarrin/discordgo"
)

type DiscordStorage struct{
	discord *discordgo.Session	
	channelID string
}

func (s *DiscordStorage) AddMessage(messageTitle string, messageContent string) error{
	em := embed.NewGenericEmbed(messageTitle , messageContent)
	_, err := s.discord.ChannelMessageSendComplex(s.channelID, &discordgo.MessageSend{Content: "@everyone", Embeds: []*discordgo.MessageEmbed{em}})
	return err
}

func NewDiscordStorage(token string, channelID string) (*DiscordStorage,error){
	discord, err := discordgo.New("Bot " + token)
	if err != nil{
		return nil,err
	}

	err = discord.Open()
	if err != nil{
		return nil, err
	}

	return &DiscordStorage{discord: discord, channelID: channelID}, nil
}
