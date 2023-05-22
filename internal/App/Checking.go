package app

import "github.com/RB-PRO/BazarakiUpdate/pkg/bazaraki"

// Проверить основной массив в соответствии с новым добавленным и вернуть разницу, а именно новые элементы и удалённые элементы
func Checking(PagesResult *[]bazaraki.ResultsPage, AddPages []bazaraki.ResultsPage) ([]bazaraki.ResultsPage, []bazaraki.ResultsPage, bool) {

	// B-A
	NewPages, IsNew := Difference(AddPages, *PagesResult)

	// A-B
	DeletePages, IsDel := Difference(*PagesResult, AddPages)

	// Обновляем исходные данные
	*PagesResult = AddPages

	return NewPages, DeletePages, IsNew || IsDel
}

// Set Difference: A - B
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
