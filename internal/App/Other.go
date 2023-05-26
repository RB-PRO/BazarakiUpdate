package app

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"time"

	database "github.com/RB-PRO/BazarakiUpdate/pkg/DataBase"
	"github.com/RB-PRO/BazarakiUpdate/pkg/bazaraki"
)

// Фунция получения конфигурационных файлов
//
// Получение значение из файла filename
func dataFile(filename string) (string, error) {
	// Открыть файл
	fileToken, errorToken := os.Open(filename)
	if errorToken != nil {
		return "", errorToken
	}

	// Прочитать значение файла
	data := make([]byte, 512)
	n, err := fileToken.Read(data)
	if err == io.EOF { // если конец файла
		return "", errorToken
	}
	fileToken.Close() // Закрытие файла

	return string(data[:n]), nil
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
		Rubric:     PageResult.Rubric,
		TimeCreate: TimeCreate,
	}
}
