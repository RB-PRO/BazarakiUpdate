package parsing

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/RB-PRO/BazarakiUpdate/pkg/bazaraki"
	"github.com/xuri/excelize/v2"
)

// Сохранить результаты в Excel
func SaveXlsx(PagesResult []bazaraki.ResultsPage) error {
	file := excelize.NewFile()
	headers := []string{"ID", "Название", "Ссылка", "Площадь", "Цена", "Дата создания объявления"}
	SheetName := "Sheet1"
	for i, header := range headers {
		file.SetCellValue(SheetName, fmt.Sprintf("%s%d", string(rune(65+i)), 1), header)
	}

	for j, PageResult := range PagesResult {
		j += 2

		Price, _ := strconv.ParseFloat(PageResult.Price, 64)            // Цена
		Area := float64(PageResult.Attrs.AttrsArea)                     // Площадь
		TimeCreate, _ := time.Parse(time.RFC1123, PageResult.CreatedDt) // Дата создания объявления

		file.SetCellValue(SheetName, fmt.Sprintf("%s%d", string(rune(65+0)), j), PageResult.ID)
		file.SetCellValue(SheetName, fmt.Sprintf("%s%d", string(rune(65+1)), j), PageResult.Title)
		file.SetCellValue(SheetName, fmt.Sprintf("%s%d", string(rune(65+2)), j), fmt.Sprintf("https://www.bazaraki.com/adv/%d_%s/", PageResult.ID, PageResult.Slug))
		file.SetCellValue(SheetName, fmt.Sprintf("%s%d", string(rune(65+3)), j), Area)
		file.SetCellValue(SheetName, fmt.Sprintf("%s%d", string(rune(65+4)), j), Price)
		file.SetCellValue(SheetName, fmt.Sprintf("%s%d", string(rune(65+5)), j), TimeCreate)
	}

	if err := file.SaveAs("Bazaraki.xlsx"); err != nil {
		log.Fatal(err)
	}
	return nil
}
