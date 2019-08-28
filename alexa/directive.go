package alexa

import (
	"encoding/json"
	"github.com/BHunter2889/go-alexa-devkit/alexa/apl"
	"io/ioutil"
	"os"
)

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