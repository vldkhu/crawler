package checker

import (
	"crawler/parsers"
	"fmt"
)

func CheckForStajirovki(url string) {
	stajirovki, err := parsers.ParsePage(url)
	if err != nil {
		fmt.Println("Error parsing page:", err)
		return
	}

	if len(stajirovki) > 0 {
		fmt.Println("Найдены стажировки на", url)
		for _, s := range stajirovki {
			fmt.Printf("Название: %s, Ссылка: %s\n", s.Title, s.Link)
		}
	} else {
		fmt.Println("Стажировок не найдено на", url)
	}
}
