package bot

import (
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/pupunha-code/Manicoba/articles"
)

// ArticleSender Função que envia o artigo para o chat do discord
func ArticleSender(session *discordgo.Session, channelID string, article articles.Article) {

	// chama a função de mapping
	createMessage := CreateDiscordMessage(article)

	// empacota a mensagem para ser usada pela função sender do discord
	finalMessage := &discordgo.MessageSend{Embeds: []*discordgo.MessageEmbed{createMessage}}

	_, err := session.ChannelMessageSendComplex(channelID, finalMessage)

	// tratamento de erros do envio das mensagens
	if err != nil {
		log.Print("Erro ao enviar o artigo  para o canal", err)
	}
}
