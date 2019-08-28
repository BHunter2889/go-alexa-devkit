package alexa

import (
	"encoding/json"
	"github.com/BHunter2889/da-fish-alexa/alexa/apl"
	"io/ioutil"
	"os"
)

const renderDirectiveType = "Alexa.Presentation.APL.RenderDocument"

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
		opt.DataSources.BodyTemplate1Data.Title = title
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
		Properties struct {
			BackgroundImage struct {
				Sources []struct {
					URL string `json:"url"`
				} `json:"sources"`
			} `json:"backgroundImage"`
		} `json:"properties"`
	} `json:"templateData,omitempty"`
	BodyTemplate1Data struct {
		Type            string      `json:"type"`
		ObjectID        interface{} `json:"objectId,omitempty"`
		BackgroundImage struct {
			ContentDescription string     `json:"contentDescription,omitempty"` // For Screen Readers. Should always be included but not "required".
			SmallSourceURL     string     `json:"smallSourceUrl,omitempty"`
			MediumSourceURL    string     `json:"mediumSourceUrl,omitempty"`
			LargeSourceURL     string     `json:"largeSourceUrl,omitempty"`
			Sources            []struct { // TODO - Add Source struct and create builder to append new Sources.
				URL          string `json:"url"`
				Size         string `json:"size"`
				WidthPixels  int    `json:"widthPixels,omitempty"`
				HeightPixels int    `json:"heightPixels,omitempty"`
			} `json:"sources,omitempty"`
		} `json:"backgroundImage,omitempty"`
		Title       string `json:"title,omitempty"` // Intent Response title Heading to display
		TextContent struct {
			PrimaryText struct {
				Type string `json:"type,omitempty"`
				Text string `json:"text,omitempty"` // The text to display. Dynamically populate after reading into structs, unless always returning a single static response from your template.
			} `json:"primaryText,omitempty"`
		} `json:"textContent,omitempty"`
		LogoURL string `json:"logoUrl,omitempty"`
	} `json:"bodyTemplate1Data,omitempty"`
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
