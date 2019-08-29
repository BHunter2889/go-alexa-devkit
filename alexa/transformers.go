package alexa

type transformer string // Do not export so that the Transformer field values are limited to constants exported below.

const (
	SSMLToSpeech transformer = "ssmlToSpeech"
	SSMLToText   transformer = "ssmlToText"
) // Exported constants for the TemplateData.Transformers[i].Transformer value. These are the only options the AVS accepts.

type Transformer struct {
	InputPath   string      `json:"inputPath"`
	OutputName  string      `json:"outputName"`
	Transformer transformer `json:"transformer"`
}

// Provided for maintaining consistency, ease of use. You can still implement Transformer separately.
func NewSSMLToSpeechTransformer() Transformer {
	return Transformer{
		InputPath:   "ssml",
		OutputName:  "speech",
		Transformer: SSMLToSpeech,
	}
}

// Provided for maintaining consistency, ease of use. You can still implement Transformer separately.
func NewSSMLToTextTransformer() Transformer {
	return Transformer{
		InputPath:   "ssml",
		OutputName:  "text",
		Transformer: SSMLToText,
	}
}

func NewSSMLTransformerList() []Transformer {
	tl := make([]Transformer, 0)
	tl = append(tl, NewSSMLToSpeechTransformer())
	tl = append(tl, NewSSMLToTextTransformer())
	return tl
}