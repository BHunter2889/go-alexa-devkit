package alexa

import (
	"encoding/json"
	"github.com/BHunter2889/go-alexa-devkit/alexa/apl"
	"io/ioutil"
	"os"
)

type transformer string

const renderDirectiveType = "Alexa.Presentation.APL.RenderDocument"

const (
	SSMLToSpeech transformer = "ssmlToSpeech"
	SSMLToText   transformer = "ssmlToText"
)

type Directives struct {
	Directives []Directive
}

func (directives *Directives) NewBasicRenderDocumentDirective(token string, document apl.APLDocument, sources DataSources) {
	if len(directives.Directives) == 0 || &directives.Directives == nil {
		directives.Directives = make([]Directive, 1)
	}
	directives.Directives = append(directives.Directives, Directive{
		Type:        "Alexa.Presentation.APL.RenderDocument",
		Token:       token,
		Document:    document,
		DataSources: sources,
	})
}

func NewBasicAPLDirectives(token string, document apl.APLDocument, sources DataSources) Directives {
	d := Directives{}
	d.NewBasicRenderDocumentDirective(token, document, sources)
	return d
}

// Read from JSON File. Can't store the JSON in the binary so this is here in case you
// want to fetch the file from somewhere else.
func ExtractNewRenderDocDirectiveFromJson(token string, fileName string, out *Directive) error {
	jsonFile, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer jsonFile.Close()

	bytes, _ := ioutil.ReadAll(jsonFile)

	if err := json.Unmarshal(bytes, &out); err != nil {
		return err
	}

	out.Token = token
	out.Type = renderDirectiveType

	return nil
}

func ExtractNewRenderDocDirectiveFromString(token string, jsonString string, out *Directive) error {
	if err := json.Unmarshal([]byte(jsonString), &out); err != nil {
		return err
	}
	out.Token = token
	out.Type = renderDirectiveType

	return nil
}

func NewDirectivesList(title string, opts ...Directive) []Directive {
	dl := make([]Directive, 0)

	for _, opt := range opts {
		opt.DataSources.BodyTemplateData.Title = title
		dl = append(dl, opt)
	}

	return dl
}

type Directive struct {
	Type          string          `json:"type"`               // i.e. "Alexa.Presentation.APL.RenderDocument"
	Token         string          `json:"token"`              // i.e. "a-document" - string reference used to invoke subsequent directives like ExecuteCommands
	Document      apl.APLDocument `json:"document,omitempty"` // There may be other types of documents that can go here - TODO - generify the type if this becomes apparent.
	DataSources   DataSources     `json:"datasources,omitempty"`
	SlotToElicit  string          `json:"slotToElicit,omitempty"`
	UpdatedIntent *UpdatedIntent  `json:"UpdatedIntent,omitempty"`
	PlayBehavior  string          `json:"playBehavior,omitempty"`
	AudioItem     struct {
		Stream struct {
			Token                string `json:"token,omitempty"`
			URL                  string `json:"url,omitempty"`
			OffsetInMilliseconds int    `json:"offsetInMilliseconds,omitempty"`
		} `json:"stream,omitempty"`
	} `json:"audioItem,omitempty"`
}

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
			Title   string `json:"title,omitempty"`
			LogoURL string `json:"logoUrl,omitempty"`
			Image   string `json:"image,omitempty"`
			SSML    string `json:"ssml,omitempty"`
		} `json:"properties"`
		Transformers []Transformer `json:"transformers,omitempty"`
	} `json:"templateData,omitempty"`
	BodyTemplateData struct {
		Type            string      `json:"type"`
		ObjectID        interface{} `json:"objectId,omitempty"`
		BackgroundImage APLImage    `json:"backgroundImage,omitempty"`
		Title           string      `json:"title,omitempty"` // Intent Response title Heading to display
		TextContent     struct {
			Title       TextElement `json:"title,omitempty"`
			SubTitle    TextElement `json:"subtitle,omitempty"`
			PrimaryText TextElement `json:"primaryText,omitempty"`
			BulletPoint TextElement `json:"bulletPoint,omitempty"` // Must add the bullet character (i.e.: "•") yourself.
		} `json:"textContent,omitempty"`
		LogoURL string `json:"logoUrl,omitempty"`
	} `json:"bodyTemplateData,omitempty"` // NOTE: Depending on the template used, i.e. from the  Alexa Developer Portal APL template generator tool, this may have a different name.
	// TODO - create dynamic extraction/unmarshalling of this inconsistently named object source.
}

type Transformer struct {
	InputPath   string      `json:"inputPath"`
	OutputName  string      `json:"outputName"`
	Transformer transformer `json:"transformer"`
}

// Provided for maintaining consistency, ease of use. You can still implement Transformer separately.
func NewSSMLToSpeechTransformer() Transformer {
	return Transformer{
		InputPath: "ssml",
		OutputName: "speech",
		Transformer: SSMLToSpeech,
	}
}

// Provided for maintaining consistency, ease of use. You can still implement Transformer separately.
func NewSSMLToTextTransformer() Transformer {
	return Transformer{
		InputPath: "ssml",
		OutputName: "text",
		Transformer: SSMLToText,
	}
}

type TextElement struct {
	Type string `json:"type,omitempty"`
	Text string `json:"text,omitempty"` // The text to display. Dynamically populate after reading into structs, unless always returning a single static response from your template.
}

type APLImage struct {
	ContentDescription string        `json:"contentDescription,omitempty"` // For Screen Readers. Should always be included but not "required".
	SmallSourceURL     string        `json:"smallSourceUrl,omitempty"`
	MediumSourceURL    string        `json:"mediumSourceUrl,omitempty"`
	LargeSourceURL     string        `json:"largeSourceUrl,omitempty"`
	Sources            []ImageSource `json:"sources,omitempty"` // TODO - create builder to append new Sources/Images.
}

type ImageSource struct {
	URL          string `json:"url"`
	Size         string `json:"size"`
	WidthPixels  int    `json:"widthPixels,omitempty"`
	HeightPixels int    `json:"heightPixels,omitempty"`
}

type UpdatedIntent struct {
	Name               string                 `json:"name,omitempty"`
	ConfirmationStatus string                 `json:"confirmationStatus,omitempty"`
	Slots              map[string]interface{} `json:"slots,omitempty"`
}

// Experimental - Probably not worth the time
//type DirectiveOption func(token string, fileName string) (Directive, error)

//func (dir *Directives) BuildDirectives(out *[]Directive, opts ...DirectiveOption) error {
//	for _, opt := range opts {
//		in, err :=
//		out := append(*out, in)
//	}
//}
//
//func ExtractRenderDocDirectiveOption(directives *Directives) DirectiveOption {
//	return func(token string , fileName string) (Directive, error) {
//		out := Directive{
//			Type: renderDirectiveType,
//			Token: token,
//		}
//
//		jsonFile, err := os.Open(fileName)
//		if err != nil {
//			return Directive{}, err
//		}
//		defer jsonFile.Close()
//
//		bytes, _ := ioutil.ReadAll(jsonFile)
//
//		if err := json.Unmarshal(bytes, &out); err != nil {
//			return Directive{}, err
//		}
//
//		//directives = Directives{
//		//	Directives: NewDirectivesList(out),
//		//}
//
//		return Directive{}, nil
//	}
//}
