package bazaraki

// Ссылка на страницы по определённым фильтрам
const PageURL string = "https://www.bazaraki.com/api/items/?rubric=%d&page=%d&c=%d&ordering=&q=&attrs__area_min=50&attrs__area_max=250&lat=34.50498196022126&lng=33.14846821577484&radius=30000"

type Page struct {
	Count    int           `json:"count"`
	Next     string        `json:"next"`
	Previous string        `json:"previous"`
	Results  []ResultsPage `json:"results"`
}

// Структура результата по всем страницам
type ResultsPage struct {
	ID            int    `json:"id"`
	Title         string `json:"title"`
	Slug          string `json:"slug"`
	Rubric        int    `json:"rubric"`
	Description   string `json:"description"`
	City          int    `json:"city"`
	CityDistricts []int  `json:"city_districts"`
	User          struct {
		ID       int    `json:"id"`
		Phone    any    `json:"phone"`
		Name     string `json:"name"`
		Joined   string `json:"joined"`
		HasEmail bool   `json:"has_email"`
		Verified bool   `json:"verified"`
		TypeNew  bool   `json:"type_new"`
	} `json:"user"`
	Images []struct {
		ID         int    `json:"id"`
		URL        string `json:"url"`
		Orig       string `json:"orig"`
		IsFlatplan bool   `json:"is_flatplan"`
	} `json:"images"`
	Attrs struct {
		AttrsArea             int    `json:"attrs__area"`
		AttrsType             string `json:"attrs__type"`
		AttrsCondition        int    `json:"attrs__condition"`
		AttrsEnergyEfficiency int    `json:"attrs__energy-efficiency"`
		AttrsSquareMeterPrice string `json:"attrs__square-meter-price"`
	} `json:"attrs,omitempty"`
	Price            string `json:"price" csv:"Цена"`
	HitCount         int    `json:"hit_count"`
	PhoneHitcount    int    `json:"phone_hitcount"`
	Currency         string `json:"currency"`
	CreatedDt        string `json:"created_dt"`
	RaiseDt          string `json:"raise_dt"`
	OwnerAdvertCount int    `json:"owner_advert_count"`
	PhoneHide        bool   `json:"phone_hide"`
	Coordinates      struct {
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
	} `json:"coordinates"`
	Zoom             any    `json:"zoom"`
	NegotiablePrice  bool   `json:"negotiable_price"`
	Exchange         bool   `json:"exchange"`
	ImeiChecked      bool   `json:"imei_checked"`
	PriceDescription string `json:"price_description"`
	InTop            bool   `json:"in_top"`
	InPremium        bool   `json:"in_premium"`
	IsEditable       bool   `json:"is_editable"`
	IsFavorite       bool   `json:"is_favorite"`
	VideoLink        any    `json:"video_link"`
	CloudinaryVideo  struct {
	} `json:"cloudinary_video"`
	AllImages       []any  `json:"all_images"`
	TemplatedTitle  string `json:"templated_title"`
	CreditType      any    `json:"credit_type"`
	CreditAttrs     any    `json:"credit_attrs"`
	CreditLink      any    `json:"credit_link"`
	Flatplan        bool   `json:"flatplan"`
	VirtualTourLink string `json:"virtual_tour_link"`
	IsCarcheck      bool   `json:"is_carcheck"`
	NewInStockLabel bool   `json:"new_in_stock_label"`
	NewToOrderLabel bool   `json:"new_to_order_label"`
	PriceFrom       bool   `json:"price_from"`
	PriceShort      string `json:"price_short"`
}
