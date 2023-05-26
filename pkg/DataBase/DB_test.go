package database_test

import (
	"errors"
	"fmt"
	"testing"
	"time"

	database "github.com/RB-PRO/BazarakiUpdate/pkg/DataBase"
)

func TestDB(t *testing.T) {
	// Подключаемся к таблице
	DB, ErrorDB := database.New("bd")
	if ErrorDB != nil {
		t.Error(ErrorDB)
	}

	// Создаём тестовый запрос
	DB_TestData := database.Data{
		ID:         1234,
		Name:       "Name",
		Link:       "Link",
		Area:       100.0,
		Price:      200.0,
		Rubric:     123123,
		TimeCreate: time.Date(2023, time.May, 23, 15, 30, 45, 0, time.UTC),
	}

	// Записать даные в бд
	ErrorIncert := DB.Incert(DB_TestData)
	if ErrorIncert != nil {
		t.Error(ErrorIncert)
	}

	// Получить данные из БД
	FindTestIncert, ErrorSelect := DB.Select(DB_TestData.ID)
	if ErrorSelect != nil {
		t.Error(errors.Join(errors.New("Select"), ErrorSelect))
	}
	if FindTestIncert.Name != DB_TestData.Name {
		fmt.Println(FindTestIncert)
		t.Error("Select: Найдена некорректаная информация. Во всяком случае имя не совпадает.")
	}

	// Удалить данные из БД
	ErrorDelete := DB.Delete(DB_TestData.ID)
	if ErrorDelete != nil {
		t.Error(ErrorDelete)
	}

	// Попытка получить данные из БД
	FindTestIncert2, ErrorSelect2 := DB.Select(DB_TestData.ID)
	if ErrorSelect2 == nil {
		t.Error(errors.New("Select2: Не должны были найти данные, однако нашли..."))
	}
	if FindTestIncert2.Name == DB_TestData.Name {
		t.Error("Select2: Найдена корректаная информация. Не должно быть так получиться.")
	}

}

// func TestSelect(t *testing.T) {

// 	// Подключаемся к таблице
// 	DB, ErrorDB := database.New()
// 	if ErrorDB != nil {
// 		t.Error(ErrorDB)
// 	}
// 	defer DB.Close()

// 	FindTestIncert2, ErrorSelect2 := DB.Select(123)
// 	if ErrorSelect2 != nil {
// 		t.Error(errors.New("Select2: Не должны были найти данные, однако нашли..."))
// 	}
// 	fmt.Println("Получили:", FindTestIncert2)

// }

// func TestSelects(t *testing.T) {

// 	// Подключаемся к таблице
// 	DB, ErrorDB := database.New()
// 	if ErrorDB != nil {
// 		t.Error(ErrorDB)
// 	}
// 	defer DB.Close()

// 	FindTestIncert2, ErrorSelect2 := DB.Selects()
// 	if ErrorSelect2 != nil {
// 		t.Error(errors.New("Select2: Не должны были найти данные, однако нашли..."))
// 	}
// 	fmt.Println("Получили:", FindTestIncert2)

// }
