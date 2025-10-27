package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	"github.com/pupunha-code/Manicoba/articles"
	"github.com/pupunha-code/Manicoba/bot"
	"github.com/robfig/cron/v3"
)

var (
	discordToken string
	channelID    string
)

func main() {

	// Carrega as variaveis de ambiente
	godotenv.Load(".env")

	discordToken = os.Getenv("DISCORD_TOKEN")
	channelID = os.Getenv("DISCORD_CHANNEL_ID")

	// Valida as variaveis de ambiente
	if discordToken == "" {
		log.Fatal("ERRO: DISCORD_TOKEN não foi definido no ambiente!")
	}
	if channelID == "" {
		log.Fatal("ERRO: DISCORD_CHANNEL_ID não foi definido no ambiente!")
	}

	//cria a sessão
	session, err := bot.SessionCreator(discordToken)
	if err != nil {
		log.Fatal("Não foi possivel criar a seção: ", err)
	}

	log.Println("Bot iniciado")

	c := cron.New() // cria o agendador

	c.AddFunc("0 9 * * *", func() { // manda artigos de frontend
		log.Println("Buscando artigos de Frontend")
		article, err := articles.FetchArticles("javascript,typescript,react,nextjs,vue,angular,svelte,tailwindcss,vite&top=1")
		if err != nil {
			log.Println("Erro ao buscar artigo: ", err)
		} else {
			bot.ArticleSender(session, channelID, article)
		}
	})

	c.AddFunc("30 12 * * *", func() { // manda artigos de backend
		log.Println("Buscando artigos de Bakcend")
		article, err := articles.FetchArticles("node,python,go,java,rust,springboot,django,fastapi,laravel,graphql,postgres,redis,backend,architecture")
		if err != nil {
			log.Println("Erro ao buscar artigo: ", err)
		} else {
			bot.ArticleSender(session, channelID, article)
		}
	})

	c.AddFunc("0 18 * * *", func() { // manda artigos de devops e cloud
		log.Println("Buscando artigos de Devops e Cloud")
		article, err := articles.FetchArticles("devops,aws,azure,gcp,docker,kubernetes,terraform,ansible,githubactions,cicd,prometheus,grafana,observability,security,linux")
		if err != nil {
			log.Println("Erro ao buscar artigo: ", err)
		} else {
			bot.ArticleSender(session, channelID, article)
		}
	})

	c.Start() // inicia o agendador

	log.Println("Agendador iniciado. Pressione Ctrl+C para sair.")

	channel := make(chan os.Signal, 1)

	signal.Notify(channel, syscall.SIGINT, syscall.SIGTERM)

	// Bloqueia a main
	<-channel

	log.Println("Desligando o bot...")

	c.Stop() // Para o agendador

	log.Println("Bot desligado.")
}
