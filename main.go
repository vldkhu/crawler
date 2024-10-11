package main

import (
	"crawler/checker"
)

func main() {
	urls := []string{
		"https://job.mts.ru/programs/start/stazher-golang-razrabotchik-cdn-432815515610649066", // Замените на реальные URL
		"https://education.tbank.ru/start/go/",
		"https://start.avito.ru/tech",
		"https://internship.vk.company/vacancy/819",
		"https://route256.ozon.ru/?utm_source=ozontech&utm_medium=site&utm_campaign=intern#popup:infoblock",
		"https://internship.vk.company/vacancy/950",
		"https://tech.wildberries.ru/#golang",
	}

	for _, url := range urls {
		checker.CheckForStajirovki(url)
	}
}
