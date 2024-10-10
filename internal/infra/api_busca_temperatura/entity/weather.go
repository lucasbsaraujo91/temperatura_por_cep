package entity

type WeatherResponse struct {
	Location Location `json:"location"`
	Current  Current  `json:"current"`
}

type Location struct {
	Name        string  `json:"name"`
	Region      string  `json:"region"`
	Country     string  `json:"country"`
	Lat         float64 `json:"lat"`
	Lon         float64 `json:"lon"`
	TzID        string  `json:"tz_id"`
	Localtime   string  `json:"localtime"`
	LocaltimeEp int64   `json:"localtime_epoch"`
}

type Current struct {
	LastUpdatedEp int64     `json:"last_updated_epoch"`
	LastUpdated   string    `json:"last_updated"`
	TempC         float64   `json:"temp_c"`
	TempF         float64   `json:"temp_f"`
	TempK         float64   `json:"temp_k"`
	IsDay         int       `json:"is_day"`
	Condition     Condition `json:"condition"`
	WindMph       float64   `json:"wind_mph"`
	WindKph       float64   `json:"wind_kph"`
	WindDegree    int       `json:"wind_degree"`
	WindDir       string    `json:"wind_dir"`
	PressureMb    float64   `json:"pressure_mb"`
	PressureIn    float64   `json:"pressure_in"`
	PrecipMm      float64   `json:"precip_mm"`
	PrecipIn      float64   `json:"precip_in"`
	Humidity      int       `json:"humidity"`
	Cloud         int       `json:"cloud"`
	FeelsLikeC    float64   `json:"feelslike_c"`
	FeelsLikeF    float64   `json:"feelslike_f"`
	WindChillC    float64   `json:"windchill_c"`
	WindChillF    float64   `json:"windchill_f"`
	HeatIndexC    float64   `json:"heatindex_c"`
	HeatIndexF    float64   `json:"heatindex_f"`
	DewPointC     float64   `json:"dewpoint_c"`
	DewPointF     float64   `json:"dewpoint_f"`
	VisibilityKm  float64   `json:"vis_km"`
	VisibilityMi  float64   `json:"vis_miles"`
	UV            float64   `json:"uv"`
	GustMph       float64   `json:"gust_mph"`
	GustKph       float64   `json:"gust_kph"`
}

type Condition struct {
	Text string `json:"text"`
	Icon string `json:"icon"`
	Code int    `json:"code"`
}

type FullWeather struct {
	Location Location `json:"location"`
	Current  Current  `json:"current"`
}
