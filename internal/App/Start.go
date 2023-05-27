package app

import (
	"errors"
	"fmt"
	"log"
	"os"
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
		MainBase[index].Link = Base[index].Link
	}

	for {
		log.Println("Проверка")

		// Парсим все страницы
		NewBase, ErrorPage := bazaraki.Pages(1, []int{2405, 2408})
		if ErrorPage != nil {
			// panic(ErrorPage)
			log.Println("При выполнении возникла ошибка:", ErrorPage)
		}

		// Сохраняем данные в xlsx
		ErrorSave := bazaraki.SaveXlsx(NewBase)
		if ErrorSave != nil {
			panic(ErrorSave)
		}

		// Решающее правило. Обновляем БД или нет:
		if NewP, DelP, UpdP, IsEqual := Checking(&MainBase, NewBase); IsEqual {
			// fmt.Printf("NewP %+v\nDelP %+v\nUpdP %+v\n", NewP, DelP, UpdP)
			if len(NewP) != 0 {
				var MessageTG string
				for AdsIndex, ads := range NewP {
					// Message := fmt.Sprintf("%d. Добавляю товар с ID: %d, ценой %s, и ссылкой https://www.bazaraki.com/adv/%d_%s/", AdsIndex, ads.ID, ads.Price, ads.ID, ads.Slug)
					Message := fmt.Sprintf("%d. Добавляю: %d - https://www.bazaraki.com/adv/%d_%s/", AdsIndex+1, ads.ID, ads.ID, ads.Slug)
					log.Println(Message)
					MessageTG += Message + "\n"
					DB.Incert(NewP2Data(ads))
				}
				notif.Sends(MessageTG)
			}
			if len(DelP) != 0 {
				var MessageTG string
				for AdsIndex, ads := range DelP {
					// Message := fmt.Sprintf("%d. Добавляю товар с ID: %d, ценой %s, и ссылкой https://www.bazaraki.com/adv/%d_%s/", AdsIndex, ads.ID, ads.Price, ads.ID, ads.Slug)
					Message := fmt.Sprintf("%d. Удаляю: %d - %s", AdsIndex+1, ads.ID, ads.Link)
					log.Println(Message)
					MessageTG += Message + "\n"
					DB.Delete(ads.ID)
				}
				notif.Sends(MessageTG)
			}
			if len(UpdP) != 0 {
				var MessageTG string
				for AdsIndex, ads := range UpdP {
					// Message := fmt.Sprintf("%d. Добавляю товар с ID: %d, ценой %s, и ссылкой https://www.bazaraki.com/adv/%d_%s/", AdsIndex, ads.ID, ads.Price, ads.ID, ads.Slug)
					Message := fmt.Sprintf("%d. Обновляю: %d - %s", AdsIndex+1, ads.ID, ads.Link)
					log.Println(Message)
					MessageTG += Message + "\n"
					DB.UpdatePrice(ads.ID, ads.Price)
				}
				notif.Sends(MessageTG)
			}
		}

		time.Sleep(time.Minute)
	}

}
