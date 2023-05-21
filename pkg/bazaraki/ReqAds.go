package bazaraki

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Спарсить дело и вернуть ответ
func PageAds(AdsId int) (Ads, error) {

	// Выполнить запрос
	Response, ErrorGet := http.Get(fmt.Sprintf(AdsURL, AdsId))
	if ErrorGet != nil {
		return Ads{}, ErrorGet
	}
	defer Response.Body.Close()

	// Получить массив []byte из ответа
	BodyPage, ErrorReadAll := io.ReadAll(Response.Body)
	if ErrorReadAll != nil {
		return Ads{}, ErrorReadAll
	}

	// Распарсить полученный json в структуру
	var DataAds Ads
	ErrorUnmarshal := json.Unmarshal(BodyPage, &DataAds)
	if ErrorUnmarshal != nil {
		return Ads{}, ErrorUnmarshal
	}

	return DataAds, nil
}
