package alexa

// `json:"datasources,omitempty"`
type DataSources struct {
	TemplateData struct {
		// These can have any name, any number of props you want, and aren't required.
		// It only matters how you reference them from the APL document.
		// Using common basic elements for now so as not to overcomplicate.
		Properties struct {
			BackgroundImage struct {
				Sources []ImageSource `json:"sources"`
			} `json:"backgroundImage,omitempty"`
			Title       string   `json:"title,omitempty"`
			Subtitle    string   `json:"subtitle,omitempty"`
			PrimaryText string   `json:"primaryText,omitempty"`
			BulletList  []string `json:"bulletList,omitempty"`
			LogoURL     string   `json:"logoUrl,omitempty"`
			Image       string   `json:"image,omitempty"`
			SSML        string   `json:"ssml,omitempty"`
			Text        string   `json:"text,omitempty"`
		} `json:"properties"`
		Transformers []Transformer `json:"transformers,omitempty"`
	} `json:"templateData,omitempty"`
	BodyTemplateData struct {
		Type            string      `json:"type"`
		ObjectID        interface{} `json:"objectId,omitempty"`
		BackgroundImage APLImage    `json:"backgroundImage,omitempty"`
		Title           string      `json:"title,omitempty"` // Intent Response title Heading to display
		TextContent     struct {
			Title        TextElement   `json:"title,omitempty"`
			SubTitle     TextElement   `json:"subtitle,omitempty"`
			PrimaryText  TextElement   `json:"primaryText,omitempty"`
			BulletPoints []TextElement `json:"bulletPoints,omitempty"` // Must add the bullet character (i.e.: "•") yourself.
		} `json:"textContent,omitempty"`
		LogoURL string `json:"logoUrl,omitempty"`
	} `json:"bodyTemplateData,omitempty"` // NOTE: Depending on the template used, i.e. from the  Alexa Developer Portal APL template generator tool, this may have a different name.
	// TODO - create dynamic extraction/unmarshalling of this inconsistently named object source.
}
