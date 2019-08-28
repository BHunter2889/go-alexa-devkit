package alexa

import "fmt"


const PlainText = "PlainText"

type TextElement struct {
	Type string `json:"type,omitempty"`
	Text string `json:"text,omitempty"` // The text to display. Dynamically populate after reading into structs, unless always returning a single static response from your template.
}

func NewBulletPoint(bulletText string) string {
	return fmt.Sprintf("• %s", bulletText)
}

func (d *Directive) SetBodyContentTitleText(title string) {
	d.DataSources.TemplateData.Properties.Title = title
	d.DataSources.BodyTemplateData.TextContent.Title.Text = title
	d.DataSources.BodyTemplateData.TextContent.Title.Type = PlainText
}

func (d *Directive) SetBodyContentPrimaryText(primaryText string) {
	d.DataSources.TemplateData.Properties.PrimaryText = primaryText
	d.DataSources.BodyTemplateData.TextContent.PrimaryText.Text = primaryText
	d.DataSources.BodyTemplateData.TextContent.PrimaryText.Type = PlainText
}

func (d *Directive) SetBodyContentSubtitle(subtitle string) {
	d.DataSources.TemplateData.Properties.Subtitle = subtitle
	d.DataSources.BodyTemplateData.TextContent.SubTitle.Text = subtitle
	d.DataSources.BodyTemplateData.TextContent.SubTitle.Type = PlainText
}

func (d *Directive) AddBodyContentBullets(bulletStrings ...string) {
	btel := make([]TextElement, 0)
	bl := make([]string, 0)

	for _, bullet := range bulletStrings {
		nb := NewBulletPoint(bullet)
		btel = append(btel, TextElement{Text: nb, Type: PlainText})
		bl = append(bl, nb)
	}
	d.DataSources.TemplateData.Properties.BulletList = bl
	d.DataSources.BodyTemplateData.TextContent.BulletPoints = btel
}
