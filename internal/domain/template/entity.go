package template

import (
	"time"
)

type DataTemplate struct {
	ID          string      `json:"id" bson:"_id"`
	Title       string      `json:"title" bson:"title"`
	Description string      `json:"description" bson:"description"`
	Language    string      `json:"language" bson:"language"`
	Type        string      `json:"type" bson:"type"`
	Channel     string      `json:"channel" bson:"channel"`
	CreatedAt   time.Time   `json:"created_at" bson:"created_at"`
	UpdatedAt   time.Time   `json:"updated_at" bson:"updated_at"`
	DeletedAt   interface{} `json:"deleted_at" bson:"deleted_at"`
	OwnerID     string      `json:"owner_id" bson:"owner_id"`
}

type Template struct {
	ID          string      `json:"id" bson:"_id"`
	Title       string      `json:"title" bson:"title"`
	Description string      `json:"description" bson:"description"`
	Language    string      `json:"language" bson:"language"`
	Type        string      `json:"type" bson:"type"`
	Channel     string      `json:"channel" bson:"channel"`
	OwnerID     string      `json:"owner_id" bson:"ownerID"`
	CreatedAt   interface{} `json:"created_at" bson:"created_at"`
	UpdatedAt   interface{} `json:"updated_at" bson:"updated_at"`
	DeletedAt   interface{} `json:"deleted_at" bson:"deleted_at"`
	Stories     []Stories   `json:"stories" bson:"stories"`
}
type Position struct {
	X int `json:"x" bson:"x"`
	Y int `json:"y" bson:"y"`
}
type Text struct {
	Body string `json:"body" bson:"body"`
}

type Header struct {
	Type string `json:"type" bson:"type"`
	Text Text   `json:"text" bson:"text"`
}
type Body struct {
	Text string `json:"text" bson:"text"`
}
type Footer struct {
	Text string `json:"text" bson:"text"`
}
type Reply struct {
	ID    string `json:"id" bson:"id"`
	Title string `json:"title" bson:"title"`
}
type Buttons struct {
	Type  string `json:"type" bson:"type"`
	Reply Reply  `json:"reply" bson:"reply"`
}

type Rows struct {
	ID          string `json:"id" bson:"id"`
	Title       string `json:"title" bson:"title"`
	Description string `json:"description" bson:"description"`
}
type Sections struct {
	Title string `json:"title,omitempty" bson:"title,omitempty"`
	Rows  []Rows `json:"rows,omitempty" bson:"rows,omitempty"`
}
type Action struct {
	Button   string      `json:"button" bson:"button"`
	Sections *[]Sections `json:"sections,omitempty" bson:"sections,omitempty"`
	Buttons  *[]Buttons  `json:"buttons,omitempty" bson:"buttons,omitempty"`
}

type Interactive struct {
	Type   string `json:"type" bson:"type"`
	Header Header `json:"header,omitempty" bson:"header,omitempty"`
	Body   Body   `json:"body,omitempty" bson:"body,omitempty"`
	Footer Footer `json:"footer,omitempty" bson:"footer,omitempty"`
	Action Action `json:"action,omitempty" bson:"action,omitempty"`
}

type Payload struct {
	Type        string       `json:"type" bson:"type"`
	Text        *Text        `json:"text,omitempty" bson:"text,omitempty"`
	Interactive *Interactive `json:"interactive,omitempty" bson:"interactive,omitempty"`
}

type Contents struct {
	ID      string  `json:"id" bson:"id"`
	Label   string  `json:"label" bson:"label"`
	Type    string  `json:"type" bson:"type"`
	Payload Payload `json:"payload,omitempty" bson:"payload,omitempty"`
}
type Blocks struct {
	ID          string      `json:"id" bson:"id"`
	Label       string      `json:"label" bson:"label"`
	Placeholder interface{} `json:"placeholder" bson:"placeholder"`
	Type        string      `json:"type" bson:"type"`
	Position    Position    `json:"position" bson:"position"`
	Contents    []Contents  `json:"contents" bson:"contents"`
}
type Edges struct {
	ID            string      `json:"id" bson:"id"`
	SourceBlockID string      `json:"source_block_id" bson:"source_block_id"`
	TargetBlockID string      `json:"target_block_id" bson:"target_block_id"`
	Status        interface{} `json:"status" bson:"status"`
}
type Stories struct {
	ID     string   `json:"id" bson:"id"`
	Title  string   `json:"title" bson:"title"`
	Blocks []Blocks `json:"blocks" bson:"blocks"`
	Edges  []Edges  `json:"edges" bson:"edges"`
}
