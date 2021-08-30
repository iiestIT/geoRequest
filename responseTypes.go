package geoRequest

type languages struct {
	Code string `json:"code"`
	Name string `json:"name"`
	Native string `json:"native"`
}

type locationResponse struct {
	GeoNameID int `json:"geoname_id"`
	Capital string `json:"capital"`
	Languages []languages `json:"languages"`
	CountryFlag string `json:"country_flag"`
	CountryFlagEmoji string `json:"country_flag_emoji"`
	CountryFlagEmojiUnicode string `json:"country_flag_emoji_unicode"`
	CallingCode string `json:"calling_code"`
	IsEU bool `json:"is_eu"`
}

type timeZone struct {
	ID string `json:"id"`
	CurrentTime string `json:"current_time"`
	GmtOffset int `json:"gmt_offset"`
	Code string `json:"code"`
	IsDaylightSaving bool `json:"is_daylight_saving"`
}

type currency struct {
	Code string `json:"code"`
	Name string `json:"name"`
	Plural string `json:"plural"`
	Symbol string `json:"symbol"`
	SymbolNative string `json:"symbol_native"`
}

type connection struct {
	ASN string `json:"asn"`
	ISP string `json:"isp"`
}

type security struct {
	IsProxy bool `json:"is_proxy"`
	ProxyType string `json:"proxy_type"`
	IsCrawler bool `json:"is_crawler"`
	CrawlerName string `json:"crawler_name"`
	CrawlerType string `json:"crawler_type"`
	IsTor bool `json:"is_tor"`
	ThreadLevel string `json:"thread_level"`
	ThreadTypes string `json:"thread_types"`
}

type baseResponse struct {
	IP string `json:"ip"`
	Type string `json:"type"`
	ContinentCode string `json:"continent_code"`
	ContinentName string `json:"continent_name"`
	CountryCode string `json:"country_code"`
	CountryName string `json:"country_name"`
	RegionCode string `json:"region_code"`
	RegionName string `json:"region_name"`
	City string `json:"city"`
	Zip string `json:"zip"`
	Latitude float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Location locationResponse `json:"location"`
	TimeZone timeZone `json:"time_zone"`
	Currency currency `json:"currency"`
	Connection connection `json:"connection"`
	Security security `json:"security"`
}

type resError struct {
	Code int `json:"code"`
	Type string `json:"type"`
	Info string `json:"info"`
}

type responseError struct {
	Success bool `json:"success"`
	Error resError `json:"error"`
}
