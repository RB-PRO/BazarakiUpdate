package app

import "github.com/RB-PRO/BazarakiUpdate/pkg/bazaraki"

// Проверить основной массив в соответствии с новым добавленным и вернуть разницу, а именно новые элементы и удалённые элементы
func Checking(PagesResult *[]bazaraki.ResultsPage, AddPages []bazaraki.ResultsPage) ([]bazaraki.ResultsPage, []bazaraki.ResultsPage, bool) {
	NewPages := make([]bazaraki.ResultsPage, 0) // Массив новых
	DeletePages := make([]bazaraki.ResultsPage, 0)

	// Малоэффективный алгоритм поиска различий между двумя массивами
	for _, AddP := range AddPages {
		for _, ResP := range *PagesResult {

			// Если такой элемент есть, то пропускаем иттерации
			if ResP.ID == AddP.ID {
				continue
			}
		}
	}

	return NewPages, DeletePages, false
}
