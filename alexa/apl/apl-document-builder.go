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
	Resources []struct {
		Description string `json:"description,omitempty"`
		When        string `json:"when,omitempty"`
		Colors      struct {
			ColorTextPrimary       string `json:"colorTextPrimary,omitempty"`
			ColorBackgroundOverlay string `json:"colorBackgroundOverlay,omitempty"`
		} `json:"colors,omitempty"`
		Dimensions struct {
			TextSizeBody          int `json:"textSizeBody,omitempty"`
			TextSizePrimary       int `json:"textSizePrimary,omitempty"`
			TextSizeSecondary     int `json:"textSizeSecondary,omitempty"`
			TextSizeSecondaryHint int `json:"textSizeSecondaryHint,omitempty"`
		} `json:"dimensions,omitempty"`
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
		TextStyleSecondaryHint struct {
			Values struct {
				FontFamily string `json:"fontFamily,omitempty"`
				FontStyle  string `json:"fontStyle,omitempty"`
				FontSize   string `json:"fontSize,omitempty"`
				Color      string `json:"color,omitempty"`
			} `json:"values,omitempty"`
		} `json:"textStyleSecondaryHint,omitempty"`
	} `json:"styles,omitempty"`
	Layouts struct {
	} `json:"layouts,omitempty"`
	MainTemplate struct {
		Description string   `json:"description,omitempty"`
		Parameters  []string `json:"parameters,omitempty"`
		Items       []struct {
			When      string `json:"when,omitempty"`
			Type      string `json:"type"`
			Direction string `json:"direction,omitempty"`
			Width     string `json:"width,omitempty"`
			Height    string `json:"height,omitempty"`
			Items     []struct {
				Type                   string `json:"type"`
				Source                 string `json:"source,omitempty"`
				OverlayColor           string `json:"overlayColor,omitempty"`
				Position               string `json:"position,omitempty"`
				Width                  string `json:"width,omitempty"`
				Height                 string `json:"height,omitempty"`
				Scale                  string `json:"scale,omitempty"`
				HeaderTitle            string `json:"headerTitle,omitempty"`
				HeaderAttributionImage string `json:"headerAttributionImage,omitempty"`
				Grow                   int    `json:"grow,omitempty"`
				PaddingLeft            string `json:"paddingLeft,omitempty"`
				PaddingRight           string `json:"paddingRight,omitempty"`
				PaddingBottom          string `json:"paddingBottom,omitempty"`
				Items                  []struct {
					Type     string `json:"type"`
					Text     string `json:"text,omitempty"`
					FontSize string `json:"fontSize,omitempty"`
					Spacing  string `json:"spacing,omitempty"`
					Style    string `json:"style,omitempty"`
				} `json:"items,omitempty"`
			} `json:"items,omitempty"`
		} `json:"items,omitempty"`
	} `json:"mainTemplate,omitempty"`
}
