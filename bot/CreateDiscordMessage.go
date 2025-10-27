package bot

import (
	"github.com/bwmarrin/discordgo"
	"github.com/pupunha-code/Manicoba/articles"
)

// CreateDiscordMessage Mapeia a struct do artigo para a de mensagem do discord
func CreateDiscordMessage(article articles.Article) *discordgo.MessageEmbed {

	// limita a descrição para no maximo 2048 caracteres
	if len(article.Description) > 2048 {
		article.Description = article.Description[:2045] + "..."
	}

	// Mepea o campos da struct montada em Articles.go para a struct usada para mandar mensagens na lib
	messageMapper := &discordgo.MessageEmbed{
		Title:       article.Title,
		Description: article.Description,
		URL:         article.URL,
		Image: &discordgo.MessageEmbedImage{
			URL: article.ImageURL,
		},
		Footer: &discordgo.MessageEmbedFooter{
			Text:    "Pupunha Code",
			IconURL: "https://media2.dev.to/dynamic/image/width=256,height=256,fit=cover,gravity=auto,format=auto/https%3A%2F%2Fdev-to-uploads.s3.amazonaws.com%2Fuploads%2Forganization%2Fprofile_image%2F11169%2Fe3381ceb-c88e-48c8-a22e-8f8f33846ad0.png", // A imagem e o footer são tratados com objetos.
		},
	}

	return messageMapper
}
