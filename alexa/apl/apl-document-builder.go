package apl

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

// Note: This won't work as the standard file also includes a sibling DataSources object
func ReadAplDocumentFromJsonFile(fileName string, out APLDocument) error {
	jsonFile, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer jsonFile.Close()

	bytes, _ := ioutil.ReadAll(jsonFile)

	if err := json.Unmarshal(bytes, &out); err != nil {
		return err
	}

	return nil
}

/**
This is the basic structure of the type of JSON document you will get from exporting code generated from the
APL Authoring Tool. This is likely not fully comprehensive of the available options. The intent is to add new options as
they are encountered.
TODO - Continue to add any new options available when using different templates, etc.

`json:"document,omitempty"`
*/
type APLDocument struct {
	Type    string `json:"type,omitempty"`
	Version string `json:"version,omitempty"`
	Theme   string `json:"theme,omitempty"`
	Import  []struct {
		Name    string `json:"name"`
		Version string `json:"version"`
	} `json:"import,omitempty"`
	Settings  struct{} `json:"settings,omitempty"`
	Resources []struct {
		Description string `json:"description,omitempty"`
		When        string `json:"when,omitempty"`
		Colors      struct {
			ColorTextPrimary       string `json:"colorTextPrimary,omitempty"`
			ColorBackgroundOverlay string `json:"colorBackgroundOverlay,omitempty"`
		} `json:"colors,omitempty"`
		Dimensions Dimensions `json:"dimensions,omitempty"`
	} `json:"resources,omitempty"`
	Styles struct {
		TextStyleBase struct {
			Description string `json:"description,omitempty"`
			Values      []struct {
				Color      string `json:"color,omitempty"`
				FontFamily string `json:"fontFamily,omitempty"`
			} `json:"values,omitempty"`
		} `json:"textStyleBase,omitempty"`
		TextStyleBase0 struct {
			Description string `json:"description,omitempty"`
			Extend      string `json:"extend,omitempty"`
			Values      struct {
				FontWeight string `json:"fontWeight,omitempty"`
			} `json:"values,omitempty"`
		} `json:"textStyleBase0,omitempty"`
		TextStyleBase1 struct {
			Description string `json:"description,omitempty"`
			Extend      string `json:"extend,omitempty"`
			Values      struct {
				FontWeight string `json:"fontWeight,omitempty"`
			} `json:"values,omitempty"`
		} `json:"textStyleBase1,omitempty"`
		MixinBody struct {
			Values struct {
				FontSize string `json:"fontSize,omitempty"`
			} `json:"values,omitempty"`
		} `json:"mixinBody,omitempty"`
		MixinBodySecondary struct {
			Values struct {
				FontSize   string `json:"fontSize"`
				FontWeight int    `json:"fontWeight"`
			} `json:"values"`
		} `json:"mixinBodySecondary,omitempty"`
		MixinPrimary struct {
			Values struct {
				FontSize string `json:"fontSize,omitempty"`
			} `json:"values,omitempty"`
		} `json:"mixinPrimary,omitempty"`
		MixinSecondary struct {
			Values struct {
				FontSize string `json:"fontSize,omitempty"`
			} `json:"values,omitempty"`
		} `json:"mixinSecondary,omitempty"`
		TextStylePrimary struct {
			Extend []string `json:"extend,omitempty"`
		} `json:"textStylePrimary,omitempty"`
		TextStyleSecondary struct {
			Extend []string `json:"extend,omitempty"`
		} `json:"textStyleSecondary,omitempty"`
		TextStyleBody struct {
			Extend []string `json:"extend,omitempty"`
		} `json:"textStyleBody,omitempty"`
		TextStyleBodySecondary struct {
			Extend []string `json:"extend,omitempty"`
		} `json:"textStyleBodySecondary,omitempty"`
		TextStyleSecondaryHint struct {
			Values struct {
				FontFamily string `json:"fontFamily,omitempty"`
				FontStyle  string `json:"fontStyle,omitempty"`
				FontSize   string `json:"fontSize,omitempty"`
				Color      string `json:"color,omitempty"`
			} `json:"values,omitempty"`
		} `json:"textStyleSecondaryHint,omitempty"`
	} `json:"styles,omitempty"`
	OnMount  []interface{} `json:"onMount"`
	Graphics struct {
	} `json:"graphics"`
	Commands struct {
	} `json:"commands"`
	Layouts struct {
	} `json:"layouts,omitempty"`
	MainTemplate struct {
		Description string   `json:"description,omitempty"`
		Parameters  []string `json:"parameters,omitempty"`
		Items       []Item   `json:"items,omitempty"`
	} `json:"mainTemplate,omitempty"`
}

// `json:"dimensions,omitempty"`
type Dimensions struct {
	TextSizeBody          int `json:"textSizeBody,omitempty"`
	TextSizeBodySecondary int `json:"textSizeBodySecondary,omitempty"`
	TextSizePrimary       int `json:"textSizePrimary,omitempty"`
	TextSizeSecondary     int `json:"textSizeSecondary,omitempty"`
	TextSizeSecondaryHint int `json:"textSizeSecondaryHint,omitempty"`
	SpacingThin           int `json:"spacingThin,omitempty"`
	SpacingSmall          int `json:"spacingSmall,omitempty"`
	SpacingMedium         int `json:"spacingMedium,omitempty"`
	SpacingLarge          int `json:"spacingLarge,omitempty"`
	SpacingExtraLarge     int `json:"spacingExtraLarge,omitempty"`
	MarginTop             int `json:"marginTop,omitempty"`
	MarginLeft            int `json:"marginLeft,omitempty"`
	MarginRight           int `json:"marginRight,omitempty"`
	MarginBottom          int `json:"marginBottom,omitempty"`
	PaddingLeft           int `json:"paddingLeft,omitempty"`
	PaddingRight          int `json:"paddingRight,omitempty"`
	PaddingBottom         int `json:"paddingBottom,omitempty"`
	PaddingTop            int `json:"paddingTop,omitempty"`
	PaddingLeft40         int `json:"paddingLeft40,omitempty"`
	PaddingRight72        int `json:"paddingRight72,omitempty"`
	PaddingTop40          int `json:"paddingTop40,omitempty"`
	PaddingTop50          int `json:"paddingTop50,omitempty"`
	PrimaryTextPaddingTop int `json:"primaryTextPaddingTop,omitempty"`
	BulletPointPaddingTop int `json:"bulletPointPaddingTop,omitempty"`
}

type Item struct {
	Type                   string `json:"type"`
	When                   string `json:"when,omitempty"`
	Direction              string `json:"direction,omitempty"`
	Source                 string `json:"source,omitempty"`
	OverlayColor           string `json:"overlayColor,omitempty"`
	Position               string `json:"position,omitempty"`
	Align                  string `json:"align,omitempty"`
	AlignItems             string `json:"alignItems,omitempty"`
	Width                  string `json:"width,omitempty"`
	Height                 string `json:"height,omitempty"`
	Scale                  string `json:"scale,omitempty"`
	HeaderTitle            string `json:"headerTitle,omitempty"`
	HeaderAttributionImage string `json:"headerAttributionImage,omitempty"`
	Grow                   int    `json:"grow,omitempty"`
	Shrink                 int    `json:"shrink,omitempty"`
	Spacing                string `json:"spacing,omitempty"`
	PaddingLeft            string `json:"paddingLeft,omitempty"`
	PaddingRight           string `json:"paddingRight,omitempty"`
	PaddingBottom          string `json:"paddingBottom,omitempty"`
	PaddingTop             string `json:"paddingTop,omitempty"`
	Text                   string `json:"text,omitempty"`
	TextAlign              string `json:"textAlign,omitempty"`
	FontSize               string `json:"fontSize,omitempty"`
	Style                  string `json:"style,omitempty"`
	Color                  string `json:"color,omitempty"`
	Items                  []Item `json:"items,omitempty"`
	Item                   []Item `json:"item,omitempty"` // Yes, at present they expect this to be singular when the array contains one item. -_-
}
