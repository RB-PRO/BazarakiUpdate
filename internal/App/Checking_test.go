package app

import (
	"fmt"
	"testing"

	"github.com/RB-PRO/BazarakiUpdate/pkg/bazaraki"
)

func TestChecking(t *testing.T) {

	Pages := []bazaraki.ResultsPage{
		{ID: 1, Price: "1.0"}, {ID: 2, Price: "2.0"}, {ID: 3, Price: "3.0"}, {ID: 4, Price: "4.0"}, {ID: 5, Price: "5.0"}, {ID: 6, Price: "6.0"}, {ID: 7, Price: "7.0"}, {ID: 8, Price: "8.0"},
	}
	LenPages := len(Pages)

	// Создаём новый массив, куда и добавляем новые элементы
	// AddPage := Pages
	AddPage := make([]bazaraki.ResultsPage, len(Pages)-3)
	copy(AddPage, Pages[:len(Pages)-3])
	// AddPage = AddPage[:len(AddPage)-3]
	AddPage = append(AddPage, bazaraki.ResultsPage{ID: 123})
	AddPage[3].Price = "2288.00" // Изменить элемент 4-й

	fmt.Println("\nИсходные данные:")
	fmt.Print("Pages ")
	prints(Pages)
	fmt.Print("AddPage ")
	prints(AddPage)
	fmt.Println("\nОбработанные данные:")
	//
	NewP, DelP, UpdP, IsEqual := Checking(&Pages, AddPage)
	fmt.Print("Pages ")
	prints(Pages)
	fmt.Print("NewP ")
	prints(NewP)
	fmt.Print("DelP ")
	prints(DelP)
	fmt.Print("UpdP ")
	prints(UpdP)
	if len(NewP) != 1 {
		t.Error("Checking: NewP: После нового объявления должны были получить новых товаров 1, а получили", len(NewP))
	}
	if len(DelP) != 3 {
		t.Error("Checking: DelP: После нового объявления должны были получить удалённых товаров 3, а получили", len(DelP))
	}
	if len(UpdP) != 1 {
		t.Error("Checking: UpdP: После нового объявления должны были получить обновленных товаров 1, а получили", len(UpdP))
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
