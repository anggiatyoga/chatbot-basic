package request

import (
	"chatbotbasic/internal/domain/template"
)

type TemplateParamRequest struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Language    string    `json:"language"`
	Type        string    `json:"type"`
	Channel     string    `json:"channel"`
	OwnerID     string    `json:"owner_id"`
	Stories     []Stories `json:"stories"`
}

func (t *TemplateParamRequest) ToEntity() template.Template {
	var stories []template.Stories
	for _, s := range t.Stories {
		stories = append(stories, s.ToEntity())
	}

	return template.Template{
		ID:          t.ID,
		Title:       t.Title,
		Description: t.Description,
		Language:    t.Language,
		Type:        t.Type,
		Channel:     t.Channel,
		OwnerID:     t.OwnerID,
		Stories:     stories,
	}
}

type Position struct {
	X int `json:"x"`
	Y int `json:"y"`
}

func (p *Position) ToEntity() template.Position {
	return template.Position{
		X: p.X,
		Y: p.Y,
	}
}

type Text struct {
	Body *string `json:"body"`
}

type Header struct {
	Type string `json:"type"`
	Text Text   `json:"text"`
}
type Body struct {
	Text string `json:"text"`
}
type Footer struct {
	Text string `json:"text"`
}
type Reply struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}
type Buttons struct {
	Type  string `json:"type"`
	Reply Reply  `json:"reply"`
}

func (b *Buttons) ToEntity() template.Buttons {
	return template.Buttons{
		Type: b.Type,
		Reply: template.Reply{
			ID:    b.Reply.ID,
			Title: b.Reply.Title,
		},
	}
}

type Rows struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

func (r *Rows) ToEntity() template.Rows {
	return template.Rows{
		ID:          r.ID,
		Title:       r.Title,
		Description: r.Description,
	}
}

type Sections struct {
	Title string `json:"title"`
	Rows  []Rows `json:"rows"`
}

func (s *Sections) ToEntity() template.Sections {
	var rows []template.Rows
	for _, r := range s.Rows {
		rows = append(rows, r.ToEntity())
	}

	return template.Sections{
		Title: s.Title,
		Rows:  rows,
	}
}

type Action struct {
	Button   *string     `json:"button"`
	Sections *[]Sections `json:"sections"`
	Buttons  *[]Buttons  `json:"buttons"`
}

type Interactive struct {
	Type   string `json:"type"`
	Header Header `json:"header"`
	Body   Body   `json:"body"`
	Footer Footer `json:"footer"`
	Action Action `json:"action"`
}

type Payload struct {
	Type        *string      `json:"type,omitempty"`
	Text        *Text        `json:"text,omitempty"`
	Interactive *Interactive `json:"interactive"`
}

type Contents struct {
	ID      string  `json:"id"`
	Label   string  `json:"label"`
	Type    string  `json:"type"`
	Payload Payload `json:"payload,omitempty"`
	//PayloadButtons *PayloadButtons `json:"payload,omitempty"`
	//PayloadList    *PayloadList    `json:"payload,omitempty"`
}

func (c *Contents) ToEntity() template.Contents {

	if c.Type == "TEXT" {
		return template.Contents{
			ID:    c.ID,
			Label: c.Label,
			Type:  c.Type,
			Payload: template.Payload{
				Type: *c.Payload.Type,
				Text: &template.Text{Body: *c.Payload.Text.Body},
			},
		}
	} else if c.Type == "BUTTON" {
		var actionButtons []template.Buttons
		for _, ab := range *c.Payload.Interactive.Action.Buttons {
			actionButtons = append(actionButtons, ab.ToEntity())
		}

		return template.Contents{
			ID:    c.ID,
			Label: c.Label,
			Type:  c.Type,
			Payload: template.Payload{
				Type: *c.Payload.Type,
				Interactive: &template.Interactive{
					Type: c.Payload.Interactive.Type,
					Header: template.Header{
						Type: c.Payload.Interactive.Type,
						Text: template.Text{Body: *c.Payload.Interactive.Header.Text.Body},
					},
					Body:   template.Body{Text: c.Payload.Interactive.Body.Text},
					Footer: template.Footer{Text: c.Payload.Interactive.Footer.Text},
					Action: template.Action{
						Buttons: &actionButtons,
					},
				},
			},
		}
	} else if c.Type == "LIST" {

		var sections []template.Sections
		for _, sc := range *c.Payload.Interactive.Action.Sections {
			sections = append(sections, sc.ToEntity())
		}

		return template.Contents{
			ID:    c.ID,
			Label: c.Label,
			Type:  c.Type,
			Payload: template.Payload{
				Type: c.Payload.Interactive.Type,
				Interactive: &template.Interactive{
					Type: c.Payload.Interactive.Type,
					Header: template.Header{
						Type: c.Payload.Interactive.Header.Type,
						Text: template.Text{Body: *c.Payload.Interactive.Header.Text.Body},
					},
					Body:   template.Body{Text: c.Payload.Interactive.Body.Text},
					Footer: template.Footer{Text: c.Payload.Interactive.Footer.Text},
					Action: template.Action{
						Button:   *c.Payload.Interactive.Action.Button,
						Sections: &sections,
					},
				},
			},
		}
	}

	return template.Contents{
		ID:    c.ID,
		Label: c.Label,
		Type:  c.Type,
	}
}

type Blocks struct {
	ID          string      `json:"id"`
	Label       string      `json:"label"`
	Placeholder interface{} `json:"placeholder"`
	Type        string      `json:"type"`
	Position    Position    `json:"position"`
	Contents    []Contents  `json:"contents"`
}

func (b *Blocks) ToEntity() template.Blocks {
	var contents []template.Contents
	for _, c := range b.Contents {
		contents = append(contents, c.ToEntity())
	}

	return template.Blocks{
		ID:          b.ID,
		Label:       b.Label,
		Placeholder: b.Placeholder,
		Type:        b.Type,
		Position:    b.Position.ToEntity(),
		Contents:    contents,
	}

}

type Edges struct {
	ID            string      `json:"id"`
	SourceBlockID string      `json:"source_block_id"`
	TargetBlockID string      `json:"target_block_id"`
	Status        interface{} `json:"status"`
}

func (e Edges) ToEntity() template.Edges {
	return template.Edges{
		ID:            e.ID,
		SourceBlockID: e.SourceBlockID,
		TargetBlockID: e.TargetBlockID,
		Status:        e.Status,
	}
}

type Stories struct {
	ID     string   `json:"id"`
	Title  string   `json:"title"`
	Blocks []Blocks `json:"blocks"`
	Edges  []Edges  `json:"edges"`
}

func (s *Stories) ToEntity() template.Stories {
	var blocks []template.Blocks
	for _, b := range s.Blocks {
		blocks = append(blocks, b.ToEntity())
	}

	var edges []template.Edges
	for _, e := range s.Edges {
		edges = append(edges, e.ToEntity())
	}

	return template.Stories{
		ID:     s.ID,
		Title:  s.Title,
		Blocks: blocks,
		Edges:  edges,
	}
}
