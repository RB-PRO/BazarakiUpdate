package app

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	database "github.com/RB-PRO/BazarakiUpdate/pkg/DataBase"
	"github.com/RB-PRO/BazarakiUpdate/pkg/bazaraki"
)

func Start() {
	var filename string
	if len(os.Args) == 1 {
		filename = "root"
	} else {
		filename = os.Args[1]
	}
	// Подключается к БД
	DB, ErrorBD := database.New(filename)
	if ErrorBD != nil {
		panic(ErrorBD)
	}

	// Получить все записи. Потом будем сранивать именно с этими данными
	Base, ErrorSelects := DB.Selects()
	if ErrorSelects != nil {
		panic(ErrorSelects)
	}

	MainBase := make([]bazaraki.ResultsPage, len(Base))
	for index := range MainBase {
		MainBase[index].ID = Base[index].ID
	}

	for {
		log.Println("Проверка")

		// Парсим все страницы
		NewBase, ErrorPage := bazaraki.Pages()
		if ErrorPage != nil {
			panic(ErrorPage)
		}

		// Сохраняем данные в xlsx
		ErrorSave := bazaraki.SaveXlsx(NewBase)
		if ErrorSave != nil {
			panic(ErrorSave)
		}

		// Решающее правило. Обновляем БД или нет:
		if NewP, DelP, IsEqual := Checking(&MainBase, NewBase); IsEqual {
			if len(NewP) != 0 {
				for _, ads := range NewP {
					log.Println("Добавляю товар с ID", ads)
					DB.Incert(NewP2Data(ads))
				}
			}
			if len(DelP) != 0 {
				for _, ads := range DelP {
					log.Println("Удаляю товар с ID", ads)
					DB.Incert(NewP2Data(ads))
				}
			}
		}

		time.Sleep(15 * time.Second)
	}

}

func NewP2Data(PageResult bazaraki.ResultsPage) database.Data {
	Price, _ := strconv.ParseFloat(PageResult.Price, 64)            // Цена
	Area := float64(PageResult.Attrs.AttrsArea)                     // Площадь
	TimeCreate, _ := time.Parse(time.RFC1123, PageResult.CreatedDt) // Дата создания объявления

	return database.Data{
		ID:         PageResult.ID,
		Name:       PageResult.Title,
		Link:       fmt.Sprintf("https://www.bazaraki.com/adv/%d_%s/", PageResult.ID, PageResult.Slug),
		Price:      Price,
		Area:       Area,
		TimeCreate: TimeCreate,
	}
}
