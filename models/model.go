package models

import "strings"

type Campaign struct {
	CampaignId string `json:"cid" binding:"required"`
	Name       string `json:"name"`
	Image      string `json:"img"`
	CTA        string `json:"cta"`
	Status     string `json:"status"`
}

type RuleSet struct {
	IncludeCountry []string `json:"include_country,omitempty"`
	ExcludeCountry []string `json:"exclude_country,omitempty"`
	IncludeOS      []string `json:"include_os,omitempty"`
	ExcludeOS      []string `json:"exclude_os,omitempty"`
	IncludeApp     []string `json:"include_app,omitempty"`
	ExcludeApp     []string `json:"exclude_app,omitempty"`
}

type TargetRules struct {
	CampaignId string  `json:"cid" binding:"required"`
	Rules      RuleSet `json:"rules"`
}

type UrlFields struct {
	App     string
	Country string
	OS      string
}

type DeliverResponse struct {
	CampaignId string `json:"cid" binding:"required"`
	Image      string `json:"img"`
	CTA        string `json:"cta"`
}

func (r *RuleSet) Normalize() {
	r.IncludeCountry = toLowerSlice(r.IncludeCountry)
	r.ExcludeCountry = toLowerSlice(r.ExcludeCountry)
	r.IncludeOS = toLowerSlice(r.IncludeOS)
	r.ExcludeOS = toLowerSlice(r.ExcludeOS)
	r.IncludeApp = toLowerSlice(r.IncludeApp)
	r.ExcludeApp = toLowerSlice(r.ExcludeApp)
}

func toLowerSlice(s []string) []string {
	for i := range s {
		s[i] = strings.ToLower(s[i])
	}
	return s
}
