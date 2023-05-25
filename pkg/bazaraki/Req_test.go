package bazaraki_test

import (
	"testing"

	"github.com/RB-PRO/BazarakiUpdate/pkg/bazaraki"
)

// Прпоарсить одну страницу
func TestPageOne(t *testing.T) {
	Pages, IsNext, ErrorPageOne := bazaraki.PageOne(1)
	if ErrorPageOne != nil {
		t.Error(ErrorPageOne)
	}
	if !IsNext {
		t.Error("PageOne: Параметр, который овтвечает за наличие/отсутствие следующей страницы пернул 'false'")
	}
	if len(Pages.Results) == 0 {
		t.Error("PageOne: Найдено всего 0 объявлений")
	}
	t.Log("Всего объявлений:", len(Pages.Results))
}

func TestPages(t *testing.T) {
	Pages, ErrorPage := bazaraki.Pages(0)
	if ErrorPage != nil {
		t.Error(ErrorPage)
	}
	if len(Pages) == 0 {
		t.Error("Pages: Найдено всего 0 объявлений")
	}
	t.Log("Всего объявлений:", len(Pages))
}

func TestPageAds(t *testing.T) {
	id := 4497557
	Ads, ErrorAds := bazaraki.PageAds(id)
	if ErrorAds != nil {
		t.Error(ErrorAds)
	}
	if Ads.ID != id {
		t.Error("PageAds: Для ID", id, "было найдено неверное значение Ads")
	}
}
