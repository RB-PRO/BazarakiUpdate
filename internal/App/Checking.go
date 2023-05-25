package app

import (
	"github.com/RB-PRO/BazarakiUpdate/pkg/bazaraki"
)

// Проверить основной массив в соответствии с новым добавленным и вернуть разницу, а именно новые элементы и удалённые элементы
func Checking(PagesResult *[]bazaraki.ResultsPage, AddPages []bazaraki.ResultsPage) ([]bazaraki.ResultsPage, []bazaraki.ResultsPage, []bazaraki.ResultsPage, bool) {

	// B-A
	NewPages, IsNew := Difference(AddPages, *PagesResult)

	// A-B
	DeletePages, IsDel := Difference(*PagesResult, AddPages)

	// Теперь ищем обновления.
	// В обновления могут попасть только те данные, которые не находятся в новых или удаляемых,
	// Соответственно сперва необходимо приготовить массив из новых сведений, которые будут содержать остающаюся информацию о объявлениях,
	// ведь именно в ней могут произойти обновления. Далее этот массив данных проверяем на обновы
	UpdateTecalPages := *PagesResult
	if IsNew { // Вычитаем новые
		UpdateTecalPages, _ = Difference(UpdateTecalPages, NewPages)
	}
	if IsDel { // Вычитаем старые
		UpdateTecalPages, _ = Difference(UpdateTecalPages, DeletePages)
	}
	UpdatePages, IsUpd := UpdateVlookup(UpdateTecalPages, AddPages)

	// Обновляем исходные данные
	*PagesResult = AddPages

	return NewPages, DeletePages, UpdatePages, IsNew || IsDel || IsUpd
}

// Вычесть множество из множества
func Difference(a, b []bazaraki.ResultsPage) (diff []bazaraki.ResultsPage, IsEdit bool) {
	m := make(map[int]bool)

	for _, item := range b {
		m[item.ID] = true
	}

	for _, item := range a {
		if _, ok := m[item.ID]; !ok {
			diff = append(diff, item)
			IsEdit = true
		}
	}
	return diff, IsEdit
}

// Сравниваем два массива:
//   - Исходный массив данных
//   - Новый массив данных
//
// Логика такая:
// `Если цена товара обновилась, то сохраняем эти сведения в return`
func UpdateVlookup(a, b []bazaraki.ResultsPage) (diff []bazaraki.ResultsPage, IsEdit bool) {
	type MapObj struct {
		IsEdit bool
		Price  string
	}

	m := make(map[int]MapObj)

	for _, item := range b {
		m[item.ID] = MapObj{IsEdit: true, Price: item.Price}
	}

	for _, item := range a {
		if TecalMapsObject, ok := m[item.ID]; ok && item.Price != TecalMapsObject.Price {
			diff = append(diff, item)
			IsEdit = true
		}
	}
	return diff, IsEdit
}
