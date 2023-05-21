package bazaraki

const AdsURL string = "https://www.bazaraki.com/api/items/%d/"

//	type Ads struct {
//		ID          int      // ID данного объявления
//		Name        string   // Название объявления
//		Link        string   // Ссылка на объявление
//		Description string   // Описание объявления
//		Adress      string   // Адрес
//		Number      string   // Номер телефона
//		Price       float64  // Цена (€)
//		Area        float64  // Площадь (m²)
//		Type        string   // Тип здания
//		Bedrooms    int      // К-во комнат
//		Images      []string // Ссылки на картинки
//	}
type Ads struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Slug        string `json:"slug"`
	Rubric      int    `json:"rubric"`
	Description string `json:"description"`
	City        int    `json:"city"`
	User        struct {
		ID           int    `json:"id"`
		Name         string `json:"name"`
		HasEmail     bool   `json:"has_email"`
		Verified     bool   `json:"verified"`
		AccountType  string `json:"account_type"`
		CustomFields []any  `json:"custom_fields"`
		Joined       string `json:"joined"`
		Phone        any    `json:"phone"`
		Logo         string `json:"logo"`
		CompanyName  any    `json:"company_name"`
		LegalName    string `json:"legal_name"`
		Website      string `json:"website"`
		ContactPhone any    `json:"contact_phone"`
	} `json:"user"`
	Images []struct {
		ID         int    `json:"id"`
		URL        string `json:"url"`
		Orig       string `json:"orig"`
		IsFlatplan bool   `json:"is_flatplan"`
	} `json:"images"`
	Attrs struct {
		AttrsArea              int    `json:"attrs__area"`
		AttrsType              string `json:"attrs__type"`
		AttrsParking           string `json:"attrs__parking"`
		AttrsCondition         int    `json:"attrs__condition"`
		AttrsOnlineViewing     int    `json:"attrs__online-viewing"`
		AttrsAirConditioning   string `json:"attrs__air-conditioning"`
		AttrsEnergyEfficiency  int    `json:"attrs__energy-efficiency"`
		AttrsNumberOfBedrooms  string `json:"attrs__number-of-bedrooms"`
		AttrsSquareMeterPrice  string `json:"attrs__square-meter-price"`
		AttrsNumberOfBathrooms int    `json:"attrs__number-of-bathrooms"`
	} `json:"attrs"`
	Price      string `json:"price"`
	StartPrice string `json:"start_price"`
	Contacts   struct {
	} `json:"contacts"`
	HitCount         int    `json:"hit_count"`
	Currency         string `json:"currency"`
	PhoneHitcount    int    `json:"phone_hitcount"`
	RaiseDt          string `json:"raise_dt"`
	CreatedDt        string `json:"created_dt"`
	OwnerAdvertCount int    `json:"owner_advert_count"`
	Coordinates      any    `json:"coordinates"`
	Zoom             any    `json:"zoom"`
	NegotiablePrice  bool   `json:"negotiable_price"`
	Exchange         bool   `json:"exchange"`
	PriceDescription string `json:"price_description"`
	InTop            bool   `json:"in_top"`
	InPremium        bool   `json:"in_premium"`
	IsEditable       bool   `json:"is_editable"`
	IsFavorite       bool   `json:"is_favorite"`
	CityDistricts    []int  `json:"city_districts"`
	Flatplan         bool   `json:"flatplan"`
	VideoLink        any    `json:"video_link"`
	CreditType       any    `json:"credit_type"`
	CreditAttrs      any    `json:"credit_attrs"`
	CreditLink       any    `json:"credit_link"`
	Breadcrumbs      []struct {
		ID   int    `json:"id"`
		Path string `json:"path"`
		Name string `json:"name"`
	} `json:"breadcrumbs"`
	TemplatedTitle  string `json:"templated_title"`
	CloudinaryVideo struct {
	} `json:"cloudinary_video"`
	Whatsapp              any    `json:"whatsapp"`
	Viber                 any    `json:"viber"`
	ImeiChecked           bool   `json:"imei_checked"`
	ImeiInfo              []any  `json:"imei_info"`
	PhoneBenchmarkResults []any  `json:"phone_benchmark_results"`
	ExternalID            string `json:"external_id"`
	ItemLink              string `json:"item_link"`
	VirtualTourLink       string `json:"virtual_tour_link"`
	SquareMeterPrice      string `json:"square_meter_price"`
	IsCarcheck            bool   `json:"is_carcheck"`
	Delivery              bool   `json:"delivery"`
	HasOnlineViewing      bool   `json:"has_online_viewing"`
	HasCarcheckReport     bool   `json:"has_carcheck_report"`
	HasFreeCarcheckReport bool   `json:"has_free_carcheck_report"`
	CategoryType          string `json:"category_type"`
	NewInStockLabel       bool   `json:"new_in_stock_label"`
	NewToOrderLabel       bool   `json:"new_to_order_label"`
	CarcheckReport        any    `json:"carcheck_report"`
	PriceFrom             bool   `json:"price_from"`
	ShowSendForm          bool   `json:"show_send_form"`
}
