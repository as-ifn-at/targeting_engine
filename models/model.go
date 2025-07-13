package models

type Campaign struct {
	CampaignId string `json:"cid"`
	Name       string `json:"name"`
	Image      string `json:"img"`
	CTA        string `json:"cta"`
	Status     string `json:"status"`
}

type TargetRules struct {
	CampaignId string `json:"cid"`
	Rules      map[string]any
}

type UrlList struct {
	App     string
	Country string
	OS      string
}
