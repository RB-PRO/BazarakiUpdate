package app

import (
	"fmt"
	"testing"

	"github.com/RB-PRO/BazarakiUpdate/pkg/bazaraki"
)

func TestChecking(t *testing.T) {

	Pages := []bazaraki.ResultsPage{
		{ID: 1}, {ID: 2}, {ID: 3}, {ID: 4}, {ID: 5}, {ID: 6}, {ID: 7}, {ID: 8},
	}
	LenPages := len(Pages)

	// Создаём новый массив, куда и добавляем новые элементы
	// AddPage := Pages
	AddPage := make([]bazaraki.ResultsPage, len(Pages)-3)
	copy(AddPage, Pages[:len(Pages)-3])
	// AddPage = AddPage[:len(AddPage)-3]
	AddPage = append(AddPage, bazaraki.ResultsPage{ID: 123})

	fmt.Println("\nИсходные данные:")
	fmt.Print("Pages ")
	prints(Pages)
	fmt.Print("AddPage ")
	prints(AddPage)
	fmt.Println("\nОбработанные данные:")
	//
	NewP, DelP, IsEqual := Checking(&Pages, AddPage)
	fmt.Print("Pages ")
	prints(Pages)
	fmt.Print("NewP ")
	prints(NewP)
	fmt.Print("DelP ")
	prints(DelP)
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
func prints(data []bazaraki.ResultsPage) {
	for _, dd := range data {
		fmt.Print(dd.ID, " ")
	}
	fmt.Println()
}
