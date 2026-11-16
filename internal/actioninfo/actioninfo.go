package actioninfo

import (
	"log"
)

// DataParser описывает типы, умеющие разбирать данные и формировать отчет.
type DataParser interface {
	Parse(datastring string) error
	ActionInfo() (string, error)
}

// Info разбирает набор записей и выводит информацию по каждой из них.
func Info(dataset []string, dp DataParser) {
	for _, data := range dataset {
		if err := dp.Parse(data); err != nil {
			log.Printf("ошибка разбора данных %q: %v", data, err)
			continue
		}

		info, err := dp.ActionInfo()
		if err != nil {
			log.Printf("не удалось сформировать информацию для %q: %v", data, err)
			continue
		}
		log.Print(info)
	}
}
