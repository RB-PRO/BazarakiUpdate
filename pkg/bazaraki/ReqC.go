package bazaraki

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Получить параметр C по запросу
func CAds(rubric int) (c int, ErrorCAds error) {
	client := &http.Client{}
	req, ErrorNewRequest := http.NewRequest(http.MethodGet, fmt.Sprintf("https://www.bazaraki.com/ajax-items-list/?rubric=%v&lat=34.51120606379305&lng=33.16486102832196&radius=30000&attrs__area_min=50&attrs__area_max=250", rubric), nil)
	if ErrorNewRequest != nil {
		return 0, ErrorNewRequest
	}
	req.Header.Add("authority", "www.bazaraki.com")
	req.Header.Add("accept", "*/*")
	req.Header.Add("accept-language", "ru,en;q=0.9,lt;q=0.8,it;q=0.7")
	req.Header.Add("cookie", "sessionid=p7272y7ftt2rwbjj18isiz073a34adrl; csrftoken=D27O7L7MgtYB7WypcihlGkz3M7cuoi1IsmEqh2QxMUwvl7t4Cr8MSKQeY4gDIbCo; _ALGOLIA=anonymous-4d4c78c5-73ea-47a4-af75-376bc363b9bd; _ym_uid=168431676563533153; _ym_d=1684316765; agree=1; _setLocation=1; last_user_radius=30000; _ga=GA1.1.806844570.1684316765; privacy_cookies=all; _ym_isad=1; last_user_coordinates={\"lat\":34.51120606379305,\"lng\":33.16486102832196}; _ym_visorc=b; django_language=en; _ga_R4K7DWBDXW=GS1.1.1685088918.11.1.1685093970.0.0.0")
	req.Header.Add("sec-ch-ua", "\"Chromium\";v=\"110\", \"Not A(Brand\";v=\"24\", \"YaBrowser\";v=\"23\"")
	req.Header.Add("sec-ch-ua-mobile", "?0")
	req.Header.Add("sec-ch-ua-platform", "\"Windows\"")
	req.Header.Add("sec-fetch-dest", "empty")
	req.Header.Add("sec-fetch-mode", "cors")
	req.Header.Add("sec-fetch-site", "same-origin")
	req.Header.Add("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 YaBrowser/23.3.4.603 Yowser/2.5 Safari/537.36")
	req.Header.Add("x-requested-with", "XMLHttpRequest")

	res, ErrorDo := client.Do(req)
	if ErrorDo != nil {
		return 0, ErrorDo
	}
	defer res.Body.Close()

	// Получить массив []byte из ответа
	BodyPage, ErrorReadAll := io.ReadAll(res.Body)
	if ErrorReadAll != nil {
		return 0, ErrorReadAll
	}

	// Распарсить полученный json в структуру
	var DataAds C_Ads
	ErrorUnmarshal := json.Unmarshal(BodyPage, &DataAds)
	if ErrorUnmarshal != nil {
		return 0, ErrorUnmarshal
	}

	return DataAds.QueryCount, nil
}
