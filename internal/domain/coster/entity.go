package coster

import "time"

type TemplateCoster struct {
	ID                  string             `json:"id"`
	TemplateName        string             `json:"template_name"`
	TemplateDescription string             `json:"template_description"`
	ClientID            string             `json:"client_id"`
	ChannelID           string             `json:"channel_id"`
	AccountID           string             `json:"account_id"`
	AccountAlias        string             `json:"account_alias"`
	DivisionID          string             `json:"division_id"`
	IsDeleted           bool               `json:"is_deleted"`
	OwnerID             string             `json:"owner_id"`
	UpdatedAt           time.Time          `json:"updated_at"`
	CreatedAt           time.Time          `json:"created_at"`
	TemplateContents    []TemplateContents `json:"template_contents"`
}

type OptionPosition struct {
	X int `json:"x,omitempty"`
	Y int `json:"y,omitempty"`
}

type ParentIds struct {
	ParentID       string         `json:"parent_id,omitempty"`
	Option         string         `json:"option,omitempty"`
	OptionLabel    string         `json:"option_label,omitempty"`
	OptionPosition OptionPosition `json:"option_position,omitempty"`
}

// lihat dari []edge
// matching target dan source
// source merupakan parent_ids

type Text struct {
	Body string `json:"body"`
}

type Payloads struct {
	Type string `json:"type"`
	Text Text   `json:"text"`
}

type TemplateContents struct {
	ID        string      `json:"id"`
	ParentIds []ParentIds `json:"parent_ids"`
	Label     string      `json:"label"`
	Payloads  []Payloads  `json:"payloads"`
}
