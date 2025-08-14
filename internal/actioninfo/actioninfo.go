package actioninfo

import (
	"log"
	"fmt"
)

type DataParser interface {
	Parse(string) (error)
	ActionInfo() (string, error)
}

func Info(dataset []string, dp DataParser) {
	//перебор датасета
	for _,v := range dataset {
		err := dp.Parse(v)
		if err != nil {
			log.Println(err)
			continue
		}
		s, errAI := dp.ActionInfo()
		if errAI != nil {
			log.Println(errAI)
		}
		fmt.Println(s)
	}
}
