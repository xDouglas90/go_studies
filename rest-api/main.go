package main

import (
	"rest-api/models"
	"rest-api/routes"
)

func main() {
	models.PoaPersonalities = []models.Personality{
		{Name: "Daiane dos Santos", History: "Daiane dos Santos encantou o Brasil e o mundo com o seu Brasileirinho no solo da ginástica artística. A porto-alegrense é a primeira atleta brasileira - entre homens e mulheres - a ser campeã mundial na ginástica, feito obtido em 2003, nos Estados Unidos."},
		{Name: "Ronaldinho Gaúcho", History: "Ronaldinho Gaúcho leva Rio Grande do Sul no nome e Porto Alegre para todo o mundo. Talvez, o jogador brasileiro mais famoso do mundo depois de Pelé, Ronaldinho foi duas vezes eleito o melhor jogador do planeta (2004 e 2005) e, apesar de não ser muito querido pelo lado azul dos porto-alegrenses, é uma marca da cidade."},
		{Name: "Mário Quintana", History: "Mário Quintana é um dos maiores poetas brasileiros e um dos mais importantes nomes da literatura gaúcha. Nascido em Porto Alegre, o poeta foi um dos fundadores do movimento modernista brasileiro e um dos principais nomes da poesia gaúcha. Mário Quintana morreu em 1994, mas sua obra continua viva e é estudada em todo o Brasil."},
		{Name: "Mário de Andrade", History: "Mário de Andrade é um dos maiores nomes da literatura brasileira e um dos mais importantes escritores do modernismo. Nascido em Porto Alegre, o escritor foi um dos fundadores do movimento modernista brasileiro e um dos principais nomes da poesia gaúcha. Mário de Andrade morreu em 1945, mas sua obra continua viva e é estudada em todo o Brasil."},
	}

	routes.HandleRequests()
}
