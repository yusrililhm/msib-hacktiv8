package dto

type WaterResponse struct {
	Water  int    `json:"water"`
	Status string `json:"status"`
}

type WindResponse struct {
	Wind   int    `json:"wind"`
	Status string `json:"status"`
}

type GetWeatherResponse struct {
	WaterResponses WaterResponse `json:"waterResponse"`
	WindResponses  WindResponse  `json:"windResponse"`
}

func (w *WaterResponse) WaterStatus() {
	if w.Water <= 5 {
		w.Status = "aman"
	} else if w.Water <= 8 {
		w.Status = "siaga"
	} else {
		w.Status = "bahaya"
	}
}

func (w *WindResponse) WindStatus() {
	if w.Wind <= 6 {
		w.Status = "aman"
	} else if w.Wind <= 15 {
		w.Status = "siaga"
	} else {
		w.Status = "bahaya"
	}
}
