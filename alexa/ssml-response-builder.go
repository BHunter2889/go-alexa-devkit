package alexa

import "strings"

func NewSSMLResponse(ssml string) Response {
	r := Response{
		Version: "1.0",
		Body: ResBody{
			OutputSpeech: &Payload{
				Type: "SSML",
				SSML: ssml,
			},
			ShouldEndSession: true,
		},
	}
	return r
}

func NewAPLResponse(ssml string, directives []Directive) Response {
	r := Response{
		Version: "1.0",
		Body: ResBody{
			OutputSpeech: &Payload{
				Type: "SSML",
				SSML: ssml,
			},
			Directives:       directives,
			ShouldEndSession: true,
		},
	}
	return r
}

type SSML struct {
	text  string
	pause string
}

type SSMLBuilder struct {
	SSML []SSML
}

func ParseString(text string) string {
	text = strings.Replace(text, "&", "and", -1)
	text = strings.Replace(text, "+", "plus", -1)
	text = strings.Replace(text, "@", "at", -1)
	text = strings.Replace(text, "w/", "with", -1)
	text = strings.Replace(text, "-", "", -1)
	text = strings.Replace(text, "™", "", -1)
	text = strings.Replace(text, "  ", " ", -1)
	return text
}

func (builder *SSMLBuilder) Say(text string) {
	text = ParseString(text)
	builder.SSML = append(builder.SSML, SSML{text: text})
}

func (builder *SSMLBuilder) Pause(pause string) {
	builder.SSML = append(builder.SSML, SSML{pause: pause})
}

func (builder *SSMLBuilder) Build() string {
	var response string
	for index, ssml := range builder.SSML {
		if ssml.text != "" {
			response += ssml.text + " "
		} else if ssml.pause != "" && index != len(builder.SSML)-1 {
			response += "<break time='" + ssml.pause + "ms'/> "
		}
	}
	return "<speak>" + response + "</speak>"
}
