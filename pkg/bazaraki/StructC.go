package bazaraki

type C_Ads struct {
	Rubric          string `json:"rubric"`
	Listing         string `json:"listing"`
	Traits          string `json:"traits"`
	RubricCities    string `json:"rubric_cities"`
	RubricDistricts string `json:"rubric_districts"`
	NeedReset       bool   `json:"need_reset"`
	SeoText         string `json:"seo_text"`
	MetaInfo        struct {
		PageTitle       string `json:"page_title"`
		PageDescription string `json:"page_description"`
		PageKeywords    string `json:"page_keywords"`
		PageRobots      string `json:"page_robots"`
		H1Tag           string `json:"h1_tag"`
		PageH1Tag       string `json:"page_h1_tag"`
	} `json:"meta_info"`
	FullURL         string `json:"full_url"`
	Ordering        any    `json:"ordering"`
	ShowResetButton bool   `json:"show_reset_button"`
	PageRobots      string `json:"page_robots"`
	LastPage        bool   `json:"last_page"`
	QueryCount      int    `json:"query_count"`
	QueryCountFmt   string `json:"query_count_fmt"`
	AlgoliaParams   string `json:"algolia_params"`
}
