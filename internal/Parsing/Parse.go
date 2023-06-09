// Сохранить пропарсить данные и сохранить в xlsx
package parsing

import (
	"fmt"

	"github.com/RB-PRO/BazarakiUpdate/pkg/bazaraki"
)

// Пропарсить все страницы и сохранить данные в Xlsx
func ParsePages() {

	// Парсим все страницы
	Pages, ErrorPage := bazaraki.Pages(0, []int{2408, 2405})
	if ErrorPage != nil {
		panic(ErrorPage)
	}
	fmt.Println(len(Pages))

	// Сохраняем данные в xlsx
	ErrorSave := bazaraki.SaveXlsx(Pages)
	if ErrorSave != nil {
		panic(ErrorSave)
	}

}
