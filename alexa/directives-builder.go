package alexa

import (
	"github.com/BHunter2889/go-alexa-devkit/alexa/apl"
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

func NewDirectivesList(title string, opts ...Directive) []Directive {
	dl := make([]Directive, 0)

	for _, opt := range opts {
		opt.DataSources.BodyTemplateData.Title = title
		dl = append(dl, opt)
	}

	return dl
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
