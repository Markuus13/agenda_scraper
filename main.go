package main

// "github.com/gocolly/colly"

func main() {
	// c := colly.NewCollector(
	// 	colly.AllowedDomains("antigo.economia.gov.br", "www.gov.br"),
	// )

	// c.OnHTML("div.dados-agenda", func(e *colly.HTMLElement) {
	// 	fmt.Println(e.ChildText(".agenda-autoridade"))
	// 	fmt.Println(e.ChildText(".agenda-dia"))
	// 	fmt.Println()
	// })

	// c.OnHTML("ul.list-compromissos", func(e *colly.HTMLElement) {
	// 	e.ForEach("li.item-compromisso", func(_ int, el *colly.HTMLElement) {
	// 		fmt.Println("Horário:", el.ChildText(".comprimisso-inicio"))
	// 		fmt.Println("Compromisso:", el.ChildText(".comprimisso-titulo"))
	// 		fmt.Println("Local:", el.ChildText(".comprimisso-local"))
	// 		fmt.Println()
	// 	})
	// })

	// c.Visit("http://antigo.economia.gov.br/Economia/agendas/gabinete-do-ministro/ministro-da-economia/paulo-guedes/2022-03-17?month:int=3&year:int=2022")
}
