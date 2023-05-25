package bazaraki

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// Пропарсить все страницы и вернуть список всех массив страниц
func Pages(waits int) (PagesResult []ResultsPage, ErrorPages error) {
	var IsNext bool = true                // Переменная, которая определяет, будет ли парситься следующая страница
	for PageInt := 1; IsNext; PageInt++ { // Цикл по всем-всем страницам

		// Пропарсить страницу
		TecalPages, IsNextTecal, ErrorPageOne := PageOne(PageInt)
		if ErrorPageOne != nil {
			return nil, ErrorPageOne
		}

		// Добавть результатыв слайс, который далее вернётся
		PagesResult = append(PagesResult, TecalPages.Results...)

		IsNext = IsNextTecal // Записать результат переменной, которая отвечает за продолжение парсинга страничек

		time.Sleep(time.Duration(waits) * time.Second)
	}

	return PagesResult, nil

}

// Спарсить одну страницу и вернуть ответ
func PageOne(PageInt int) (Page, bool, error) {

	// Выполнить запрос
	Response, ErrorGet := http.Get(fmt.Sprintf(PageURL, PageInt))
	if ErrorGet != nil {
		return Page{}, false, ErrorGet
	}
	defer Response.Body.Close()

	// Получить массив []byte из ответа
	BodyPage, ErrorReadAll := io.ReadAll(Response.Body)
	if ErrorReadAll != nil {
		return Page{}, false, ErrorReadAll
	}

	// Распарсить полученный json в структуру
	var DataPage Page
	ERrrorUnmarshal := json.Unmarshal(BodyPage, &DataPage)
	if ERrrorUnmarshal != nil {
		return Page{}, false, ERrrorUnmarshal
	}

	// Проверка существования ссылки на следующую страницу
	var IsNext bool
	if DataPage.Next != "" {
		IsNext = true
	}

	return DataPage, IsNext, nil
}
