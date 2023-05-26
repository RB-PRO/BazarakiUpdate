package bazaraki

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

// Ошибка, возникаемая при выполнении запроса на получение параметра "C"
var ErrorPageAds error = errors.New("PageAds: ")

// Спарсить дело и вернуть ответ
func PageAds(AdsId int) (Ads, error) {

	// Выполнить запрос
	Response, ErrorGet := http.Get(fmt.Sprintf(AdsURL, AdsId))
	if ErrorGet != nil {
		return Ads{}, errors.Join(ErrorPageAds, ErrorGet)
	}
	defer Response.Body.Close()

	// Получить массив []byte из ответа
	BodyPage, ErrorReadAll := io.ReadAll(Response.Body)
	if ErrorReadAll != nil {
		return Ads{}, errors.Join(ErrorPageAds, ErrorReadAll)
	}

	// Распарсить полученный json в структуру
	var DataAds Ads
	ErrorUnmarshal := json.Unmarshal(BodyPage, &DataAds)
	if ErrorUnmarshal != nil {
		return Ads{}, errors.Join(ErrorPageAds, ErrorUnmarshal)
	}

	return DataAds, nil
}
