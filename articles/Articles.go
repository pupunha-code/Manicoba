package articles

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

// Article Mapeamento do dados do json para a struct
type Article *struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	URL         string `json:"url"`
	ImageURL    string `json:"social_image"`
}

// FetchArticles Função base que retorna o artigo que sera enviado
func FetchArticles(tags string) (Article, error) {

	//Faz a solicitação http
	resp, err := http.Get("https://dev.to/api/articles?tags=" + tags + "&top=1")
	if err != nil {
		return nil, err
	}

	// Fecha a conexão com o endpoint no final da função
	defer resp.Body.Close()

	// trata codigos http diferentes de 200
	if resp.StatusCode != 200 {
		return nil, errors.New("http error: " + resp.Status)
	}

	// Acessa os dados da response e guarda na na slice do tipo []byte
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Faz o unmarshal do bodyBytes, salva na struct, e trata possiveis erros
	var articles []Article
	err = json.Unmarshal(bodyBytes, &articles)
	if err != nil {
		return nil, err
	}

	// tratamento para retorno de 0 artigos
	if len(articles) == 0 {
		return nil, errors.New("articles not found")
	}

	return articles[0], nil
}
