package adapter

type ItemFromApi struct {
	AcceptsMercadopago bool `json:"accepts_mercadopago"`
	Attributes         []struct {
		AttributeGroupID   string      `json:"attribute_group_id"`
		AttributeGroupName string      `json:"attribute_group_name"`
		ID                 string      `json:"id"`
		Name               string      `json:"name"`
		ValueID            string      `json:"value_id"`
		ValueName          string      `json:"value_name"`
		ValueStruct        interface{} `json:"value_struct"`
	} `json:"attributes"`
	AutomaticRelist   bool          `json:"automatic_relist"`
	AvailableQuantity int64         `json:"available_quantity"`
	BasePrice         int64         `json:"base_price"`
	BuyingMode        string        `json:"buying_mode"`
	CatalogProductID  interface{}   `json:"catalog_product_id"`
	CategoryID        string        `json:"category_id"`
	Condition         string        `json:"condition"`
	CoverageAreas     []interface{} `json:"coverage_areas"`
	CurrencyID        string        `json:"currency_id"`
	DateCreated       string        `json:"date_created"`
	DealIds           []interface{} `json:"deal_ids"`
	Descriptions      []struct {
		ID string `json:"id"`
	} `json:"descriptions"`
	DifferentialPricing interface{} `json:"differential_pricing"`
	DomainID            interface{} `json:"domain_id"`
	Geolocation         struct {
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
	} `json:"geolocation"`
	Health                       interface{}   `json:"health"`
	HistoricalStartTime          string        `json:"historical_start_time"`
	ID                           string        `json:"id"`
	InitialQuantity              int64         `json:"initial_quantity"`
	InternationalDeliveryMode    string        `json:"international_delivery_mode"`
	LastUpdated                  string        `json:"last_updated"`
	ListingSource                string        `json:"listing_source"`
	ListingTypeID                string        `json:"listing_type_id"`
	Location                     struct{}      `json:"location"`
	NonMercadoPagoPaymentMethods []interface{} `json:"non_mercado_pago_payment_methods"`
	OfficialStoreID              interface{}   `json:"official_store_id"`
	OriginalPrice                interface{}   `json:"original_price"`
	ParentItemID                 interface{}   `json:"parent_item_id"`
	Permalink                    string        `json:"permalink"`
	Pictures                     []struct {
		ID        string `json:"id"`
		MaxSize   string `json:"max_size"`
		Quality   string `json:"quality"`
		SecureURL string `json:"secure_url"`
		Size      string `json:"size"`
		URL       string `json:"url"`
	} `json:"pictures"`
	Price           int64         `json:"price"`
	SaleTerms       []interface{} `json:"sale_terms"`
	SecureThumbnail string        `json:"secure_thumbnail"`
	SellerAddress   struct {
		AddressLine string `json:"address_line"`
		City        struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"city"`
		Comment string `json:"comment"`
		Country struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"country"`
		ID             int64   `json:"id"`
		Latitude       float64 `json:"latitude"`
		Longitude      float64 `json:"longitude"`
		SearchLocation struct {
			City struct {
				ID   string `json:"id"`
				Name string `json:"name"`
			} `json:"city"`
			Neighborhood struct {
				ID   string `json:"id"`
				Name string `json:"name"`
			} `json:"neighborhood"`
			State struct {
				ID   string `json:"id"`
				Name string `json:"name"`
			} `json:"state"`
		} `json:"search_location"`
		State struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"state"`
		ZipCode string `json:"zip_code"`
	} `json:"seller_address"`
	SellerContact interface{} `json:"seller_contact"`
	SellerID      int64       `json:"seller_id"`
	Shipping      struct {
		Dimensions   interface{}   `json:"dimensions"`
		FreeShipping bool          `json:"free_shipping"`
		LocalPickUp  bool          `json:"local_pick_up"`
		LogisticType string        `json:"logistic_type"`
		Methods      []interface{} `json:"methods"`
		Mode         string        `json:"mode"`
		StorePickUp  bool          `json:"store_pick_up"`
		Tags         []interface{} `json:"tags"`
	} `json:"shipping"`
	SiteID          string        `json:"site_id"`
	SoldQuantity    int64         `json:"sold_quantity"`
	StartTime       string        `json:"start_time"`
	Status          string        `json:"status"`
	StopTime        string        `json:"stop_time"`
	SubStatus       []interface{} `json:"sub_status"`
	Subtitle        interface{}   `json:"subtitle"`
	Tags            []string      `json:"tags"`
	Thumbnail       string        `json:"thumbnail"`
	Title           string        `json:"title"`
	TotalListingFee interface{}   `json:"total_listing_fee"`
	Variations      []interface{} `json:"variations"`
	VideoID         interface{}   `json:"video_id"`
	Warnings        []interface{} `json:"warnings"`
	Warranty        interface{}   `json:"warranty"`
}
