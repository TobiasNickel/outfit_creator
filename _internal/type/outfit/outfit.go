package outfit

import "time"

const (
	GENDER_MALE                       = "MALE"
	GENDER_FEMALE                     = "FEMALE"
	COUNTRY_DE                        = "de"
	COUNTRY_EN                        = "en"
	SECTION_UPPER                     = "UPPER"
	SECTION_UNDER                     = "UNDER"
	SECTION_ACCESSORIES               = "ACCESSORIES"
	WEB_CATEGPRU_FEMALE_ACCESSORIES   = "Accessoires,WCA01156,WCA01159,WCA01155,WCA01152,WCA01158,WCA01153,WCA01157,WCA01154"
	WEB_CATEGPRU_FEMALE_UPPER_GARMENT = "Blusen,WCA00122,WCA00121,tops_t-shirts,WCA00111,WCA00112,WCA00110,Sweatshirts,WCA00132,WCA00131"
	WEB_CATEGPRU_FEMALE_UNDER_GARMENT = "Hosen,WCA00172,WCA00173,WCA00171,roecke,WCA00161,WCA00162,WCA00163"
	WEB_CATEGPRU_MALE_ACCESSORIES     = "Accessoires,WCA02305,WCA02306,WCA02304,WCA02303,WCA02308,WCA02309,WCA02307,WCA02301,WCA02302"
	WEB_CATEGPRU_MALE_UPPER_GARMENT   = "Hemden,WCA02211,T-Shirts,WCA00221,WCA00223,WCA00222,WCA00220,sweatshirts_pullover,WCA02222,WCA02223,WCA02224,WCA02221"
	WEB_CATEGPRU_MALE_UNDER_GARMENT   = "Hosen,WCA02252,WCA02251,WCA02253,denim,WCA02246,WCA02242,WCA02241,WCA02243,WCA02245,WCA02244"
)

// https://api.newyorker.de/csp/products/public/query?filters[country]=de&filters[gender]=FEMALE&filters[web_category]=Blusen,WCA00122,WCA00121,tops_t-shirts,WCA00111,WCA00112,WCA00110,Sweatshirts,WCA00132,WCA00131&limit=1

// https://api.newyorker.de/csp/products/public/product/03.01.022.2981?country=de

// https://api.newyorker.de/csp/images/image/public/7cd1b6608654321666c34e7fc38aede6.png?res=higher&frame=2_3
// https://api.newyorker.de/csp/images/image/public/9d08a34ee13c8dec8a268546aa2f64d8.png?res=low&frame=1_1

// generated at https://transform.tools/json-to-go
type OutfitAPIResponse struct {
	TotalCount int               `json:"totalCount"`
	Items      []APIResponseItem `json:"items"`
}

type APIResponseItem struct {
	GlobalItemID     string `json:"global_item_id"`
	ID               string `json:"id"`
	Country          string `json:"country"`
	MaintenanceGroup string `json:"maintenance_group"`
	WebCategoryID    string `json:"web_category_id"`
	WebCategory      string `json:"web_category"`
	Brand            string `json:"brand"`
	SalesUnit        int    `json:"sales_unit"`
	CustomerGroup    string `json:"customer_group"`
	Variants         []struct {
		ID               string    `json:"id"`
		ProductID        string    `json:"product_id"`
		PublishDate      time.Time `json:"publish_date"`
		NewIn            bool      `json:"new_in"`
		ComingSoon       bool      `json:"coming_soon"`
		Sale             bool      `json:"sale"`
		Highlight        bool      `json:"highlight"`
		ColorName        string    `json:"color_name"`
		PantoneColor     string    `json:"pantone_color"`
		PantoneColorName string    `json:"pantone_color_name"`
		Red              int       `json:"red"`
		Green            int       `json:"green"`
		Blue             int       `json:"blue"`
		ColorGroup       string    `json:"color_group"`
		BasicColor       string    `json:"basic_color"`
		Currency         string    `json:"currency"`
		OriginalPrice    float64   `json:"original_price"`
		CurrentPrice     float64   `json:"current_price"`
		RedPriceChange   bool      `json:"red_price_change"`
		Sizes            []struct {
			SizeValue string `json:"size_value"`
			SizeName  string `json:"size_name"`
			BarCode   string `json:"bar_code"`
		} `json:"sizes"`
		Images []struct {
			Key          string `json:"key"`
			Type         string `json:"type"`
			Angle        string `json:"angle"`
			HasThumbnail bool   `json:"has_thumbnail"`
			Position     int    `json:"position"`
		} `json:"images"`
	} `json:"variants"`
	Descriptions []struct {
		Language    string `json:"language"`
		Description string `json:"description"`
	} `json:"descriptions"`
}
