package app

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	database "github.com/RB-PRO/BazarakiUpdate/pkg/DataBase"
	"github.com/RB-PRO/BazarakiUpdate/pkg/bazaraki"
)

// Ошибка неверного ввода данных
var ErrorInput error = errors.New(`error: Неверный запуск программы.
Введите: ./main [Файл настройки БД] [Токен Telegram] [Чат ID Telegram]`)

func Start() { // go run cmd/main/main.go bd token -848128665
	if len(os.Args) != 4 {
		panic(ErrorInput)
	}
	var DB_FileName string
	if len(os.Args) == 1 {
		DB_FileName = "root"
	} else {
		DB_FileName = os.Args[1]
	}

	TelegramToken_FileName := os.Args[2]
	ChatID := os.Args[3]

	// # Запускаем Телеграм Нотификации
	// Загрузить данные из файла
	notifToken, ErrorDatafileNotif := dataFile(TelegramToken_FileName)
	if ErrorDatafileNotif != nil {
		panic(ErrorDatafileNotif)
	}

	// Запускаем утилиту уведомлений
	notif, ErrorNotif := NewNotification(notifToken, ChatID)
	if ErrorNotif != nil {
		panic(ErrorNotif)
	}

	// # Подключаемся к БД
	// Загрузить данные из файла
	DataDB, LoadFileDataDB := dataFile(DB_FileName)
	if LoadFileDataDB != nil {
		panic(LoadFileDataDB)
	}

	// Подключаемся к БД
	DB, ErrorBD := database.New(DataDB)
	if ErrorBD != nil {
		panic(ErrorBD)
	}

	// Получить все записи. Потом будем сранивать именно с этими данными
	Base, ErrorSelects := DB.Selects()
	if ErrorSelects != nil {
		panic(ErrorSelects)
	}
	MainBase := make([]bazaraki.ResultsPage, len(Base))
	for index := range MainBase { // Тут заполняем первый промежуточный массив, с которым и будем сравнивать значения
		MainBase[index].ID = Base[index].ID
		MainBase[index].Price = fmt.Sprintf("%.0f.00", Base[index].Price)
	}

	for {
		log.Println("Проверка")

		// Парсим все страницы
		NewBase, ErrorPage := bazaraki.Pages(1)
		if ErrorPage != nil {
			panic(ErrorPage)
		}

		// Сохраняем данные в xlsx
		ErrorSave := bazaraki.SaveXlsx(NewBase)
		if ErrorSave != nil {
			panic(ErrorSave)
		}

		// Решающее правило. Обновляем БД или нет:
		if NewP, DelP, UpdP, IsEqual := Checking(&MainBase, NewBase); IsEqual {
			if len(NewP) != 0 {
				for _, ads := range NewP {
					Message := fmt.Sprintf("Добавляю товар с ID: %d, ценой %s, и ссылкой https://www.bazaraki.com/adv%d_%s/", ads.ID, ads.Price, ads.ID, ads.Slug)
					log.Println(Message)
					notif.Sends(Message)
					DB.Incert(NewP2Data(ads))
				}
			}
			if len(DelP) != 0 {
				for _, ads := range DelP {
					Message := fmt.Sprintf("Добавляю товар с ID: %d, ценой %s, и ссылкой https://www.bazaraki.com/adv%d_%s/", ads.ID, ads.Price, ads.ID, ads.Slug)
					log.Println(Message)
					notif.Sends(Message)
					DB.Delete(ads.ID)
				}
			}
			if len(UpdP) != 0 {
				for _, ads := range UpdP {
					Message := fmt.Sprintf("Добавляю товар с ID: %d, ценой %s, и ссылкой https://www.bazaraki.com/adv%d_%s/", ads.ID, ads.Price, ads.ID, ads.Slug)
					log.Println(Message)
					notif.Sends(Message)
					DB.UpdatePrice(ads.ID, ads.Price)
				}
			}
		}

		time.Sleep(time.Minute)
	}

}

// Функция перевода структура в структуру.
//   - bazaraki.ResultsPage - структура, получаемая в результате парсинга
//   - database.Data - структура для заполнения БД
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
