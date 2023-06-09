package bazaraki

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

// Ошибка, возникаемая при выполнении запроса на получение параметра "C"
var ErrorPages error = errors.New("Pages: ")

// Пропарсить все страницы и вернуть список всех массив страниц
// На вход принимаем задержку между запросами и массив категорий
func Pages(WaintSeconds int, Rubrics []int) (PagesResult []ResultsPage, ErrorPages error) {
	for _, rubric := range Rubrics {
		c, ErrorParseC := CAds(rubric) // Пропарсить параметр "c"
		if ErrorParseC != nil {
			return nil, errors.Join(ErrorPages, ErrorParseC)
		}
		var IsNext bool = true                // Переменная, которая определяет, будет ли парситься следующая страница
		for PageInt := 1; IsNext; PageInt++ { // Цикл по всем-всем страницам

			// Пропарсить страницу
			TecalPages, IsNextTecal, ErrorPageOne := PageOne(PageInt, rubric, c)
			if ErrorPageOne != nil {
				return nil, errors.Join(ErrorPages, ErrorPageOne)
			}

			// Добавть результатыв слайс, который далее вернётся
			PagesResult = append(PagesResult, TecalPages.Results...)

			IsNext = IsNextTecal // Записать результат переменной, которая отвечает за продолжение парсинга страничек

			time.Sleep(time.Duration(WaintSeconds) * time.Second)
		}
	}
	return PagesResult, nil

}

// Ошибка, возникаемая при выполнении запроса на получение параметра "C"
var ErrorPageOne error = errors.New("PageOne: ")

// Спарсить одну страницу и вернуть ответ
func PageOne(PageInt, rubric, c int) (Page, bool, error) {
	url := fmt.Sprintf(PageURL, rubric, PageInt, c)
	// fmt.Println(url)

	// Выполнить запрос
	Response, ErrorGet := http.Get(url)
	if ErrorGet != nil {
		return Page{}, false, errors.Join(ErrorPageOne, ErrorGet)
	}
	defer Response.Body.Close()

	// Получить массив []byte из ответа
	BodyPage, ErrorReadAll := io.ReadAll(Response.Body)
	if ErrorReadAll != nil {
		return Page{}, false, errors.Join(ErrorPageOne, ErrorReadAll)
	}

	// Распарсить полученный json в структуру
	var DataPage Page
	ERrrorUnmarshal := json.Unmarshal(BodyPage, &DataPage)
	if ERrrorUnmarshal != nil {
		return Page{}, false, errors.Join(ErrorPageOne, ERrrorUnmarshal)
	}

	// Проверка существования ссылки на следующую страницу
	var IsNext bool
	if DataPage.Next != "" {
		IsNext = true
	}

	return DataPage, IsNext, nil
}
