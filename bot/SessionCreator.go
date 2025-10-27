package bot

import (
	"github.com/bwmarrin/discordgo"
)

// SessionCreator Função que cria e retorna a sessão que sera usada nas demais funções
func SessionCreator(token string) (*discordgo.Session, error) {
	//Cria a sessão com o bot
	session, err := discordgo.New("Bot " + token)
	// Tratamento de erros de sessão
	if err != nil {
		return nil, err
	}

	return session, nil
}
