package app

import (
	"testing"

	"github.com/RB-PRO/BazarakiUpdate/pkg/bazaraki"
)

func TestChecking(t *testing.T) {
	// Pages, ErrorPage := bazaraki.Pages()
	// if ErrorPage != nil {
	// 	t.Error(ErrorPage)
	// }
	// if len(Pages) == 0 {
	// 	t.Error("Pages: Найдено всего 0 объявлений")
	// }

	Pages := []bazaraki.ResultsPage{
		{ID: 1}, {ID: 2}, {ID: 3}, {ID: 4}, {ID: 5}, {ID: 6}, {ID: 7}, {ID: 8},
	}
	LenPages := len(Pages)

	// Создаём новый массив, куда и добавляем новые элементы
	AddPage := Pages
	AddPage = AddPage[:len(AddPage)-3]
	AddPage = append(AddPage, bazaraki.ResultsPage{ID: 123})

	//
	NewP, DelP, IsEqual := Checking(&Pages, AddPage)
	if len(NewP) != 1 {
		t.Error("Checking: После нового объявления должны были получить новых товаров 1, а получили", len(NewP))
	}
	if len(DelP) != 3 {
		t.Error("Checking: После нового объявления должны были получить удалённых товаров 3, а получили", len(DelP))
	}
	if len(Pages) != LenPages-2 {
		t.Error("Checking: После нового объявления старый длина старого массива должна была быть равна", LenPages-2, "а получена", len(Pages))
	}
	if !IsEqual {
		t.Error("Checking: Параметр bool должен был быть true,тк есть изменения в массиве")
	}

}
