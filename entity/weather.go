package entity

type Response struct {
	Status    string      `json:"status"`
	Count     string      `json:"count"`
	Info      string      `json:"info"`
	Infocode  string      `json:"infocode"`
	Lives     []Lives     `json:"lives"`
	Forecasts []Forecasts `json:"forecasts"`
}

type Lives struct {
	Province      string `json:"province"`
	City          string `json:"city"`
	Adcode        string `json:"adcode"`
	Weather       string `json:"weather"`
	Temperature   string `json:"temperature"`
	Winddirection string `json:"winddirection"`
	Windpower     string `json:"windpower"`
	Humidity      string `json:"humidity"`
	Reporttime    string `json:"reporttime"`
}

type Forecasts struct {
	Province   string  `json:"province"`
	City       string  `json:"city"`
	Adcode     string  `json:"adcode"`
	Reporttime string  `json:"reporttime"`
	Casts      []Casts `json:"casts"`
}

type Casts struct {
	Date         string `json:"date"`
	Week         string `json:"week"`
	Dayweather   string `json:"dayweather"`
	Nightweather string `json:"nightweather"`
	Daytemp      string `json:"daytemp"`
	Nighttemp    string `json:"nighttemp"`
	Daywind      string `json:"daywind"`
	Nightwind    string `json:"nightwind"`
	Daypower     string `json:"daypower"`
	Nightpower   string `json:"nightpower"`
}
