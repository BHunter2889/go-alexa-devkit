package alexa

import (
	"log"
)

func NewSimpleResponse(title string, text string) Response {
	r := Response{
		Version: "1.0",
		Body: ResBody{
			OutputSpeech: &Payload{
				Type: "PlainText",
				Text: text,
			},
			Card: &Payload{
				Type:    "Simple",
				Title:   title,
				Content: text,
			},
			ShouldEndSession: true,
		},
	}
	return r
}

func NewPermissionsRequestResponse() Response {
	var builder SSMLBuilder
	builder.Say("Bug Caster was unable to access your device's zip code and country information. ")
	builder.Pause("750")
	builder.Say("If you have not enabled Bug Caster to access this information, ")
	builder.Pause("150")
	builder.Say("Please check your Alexa App to grant permission for Bug Caster to access your zip code and country " +
		"information so that the fishing forecast for your area may be determined. ")
	r := Response{
		Version: "1.0",
		Body: ResBody{
			OutputSpeech: &Payload{
				Type: "SSML",
				SSML: builder.Build(),
			},
			Card: &Payload{
				Type:        "AskForPermissionsConsent",
				Permissions: []string{"read::alexa:device:all:address:country_and_postal_code"},
			},
			ShouldEndSession: true,
		},
	}
	log.Print(r)
	return r
}

func NewUnsupportedLocationResponse() Response {
	var builder SSMLBuilder
	builder.Say("Bug Caster does not currently support device locales listed outside the United States or Canada. ")
	builder.Pause("750")
	builder.Say("If you would like us to provide support for your locale, ")
	builder.Pause("150")
	builder.Say("Please leave feedback on the Bug Caster skill page in the Alexa App Skill Store with your country or locale information. ")
	builder.Pause("750")
	builder.Say("If you are presently in a supported locale, you may need to alter your device's settings in the Alexa App.")
	builder.Pause("500")
	builder.Say("Bug Caster apologizes for the inconvenience.")

	r := Response{
		Version: "1.0",
		Body: ResBody{
			OutputSpeech: &Payload{
				Type: "SSML",
				SSML: builder.Build(),
			},
			Card: &Payload{
				Type:  "Simple",
				Title: "Unsupported Locale",
				Text:  "Supported Device Locales: United States & Canada",
			},
			ShouldEndSession: true,
		},
	}
	log.Print(r)
	return r
}

func NewLaunchRequestGetPermissionsResponse() Response {
	var builder SSMLBuilder
	builder.Say("Welcome to Bug Caster!")
	builder.Pause("1000")
	builder.Say("Bug Caster uses solunar theory and applied analytics to determine how probable fish activity translates to quality of fishing by the hour.")
	builder.Pause("1000")
	builder.Say("Bug Caster will need to access your device's zip code and country information. ")
	builder.Pause("750")
	builder.Say("If you have not already enabled Bug Caster to access this information, ")
	builder.Pause("150")
	builder.Say("Please check your Alexa App to grant permission for Bug Caster to access your zip code and country " +
		"information so that the fishing forecast for your area may be determined. ")
	builder.Pause("750")
	builder.Say("Currently, ")
	builder.Pause("100")
	builder.Say("once you have granted this permission, ")
	builder.Pause("100")
	builder.Say("You can have Alexa ask Bug Caster for your fishing forecast, ")
	builder.Pause("150")
	builder.Say("or how the fishing is, ")
	builder.Pause("150")
	builder.Say("and get the best time to go fishing over the next couple of hours with a summarized rating and projected wind speed. ")
	builder.Pause("1000")
	builder.Say("New features will be coming soon, ")
	builder.Pause("150")
	builder.Say("including the ability to ask for a forecast for a specific time and location, ")
	builder.Pause("150")
	builder.Say("the best time during a specified range or normal daylight hours, ")
	builder.Pause("150")
	builder.Say("and potentially premium content such as a weekly forecast summary with graphic display. ")
	builder.Pause("1000")
	builder.Say("We hope Bug Caster improves your fishing experiences and appreciate any feedback through reviews on the skill page in the Alexa Skill Store! ")
	r := Response{
		Version: "1.0",
		Body: ResBody{
			OutputSpeech: &Payload{
				Type: "SSML",
				SSML: builder.Build(),
			},
			Card: &Payload{
				Type:        "AskForPermissionsConsent",
				Permissions: []string{"read::alexa:device:all:address:country_and_postal_code"},
			},
			ShouldEndSession: true,
		},
	}
	log.Print(r)
	return r
}

func NewDefaultErrorResponse() Response {
	var builder SSMLBuilder
	builder.Say("Bug Caster caught a snag downstream while processing your request. ")
	builder.Pause("750")
	builder.Say("I can't blame the wind, ")
	builder.Pause("100")
	builder.Say("So please accept my apologies for the inconvenience. ")
	builder.Pause("500")
	builder.Say("Please try Bug Caster again later.")

	r := Response{
		Version: "1.0",
		Body: ResBody{
			OutputSpeech: &Payload{
				Type: "SSML",
				SSML: builder.Build(),
			},
			Card: &Payload{
				Type:  "Simple",
				Title: "BugCaster Under Maintenance",
				Text:  builder.Build(),
			},
			ShouldEndSession: true,
		},
	}
	return r
}

func (r *Response) AddDirectives(d []Directive) {
	r.Body.Directives = d
}

type Response struct {
	Version           string                 `json:"version"`
	SessionAttributes map[string]interface{} `json:"sessionAttributes,omitempty"`
	Body              ResBody                `json:"response"`
}

type ResBody struct {
	OutputSpeech     *Payload    `json:"outputSpeech,omitempty"`
	Card             *Payload    `json:"card,omitempty"`
	Reprompt         *Reprompt   `json:"reprompt,omitempty"`
	Directives       []Directive `json:"directives,omitempty"`
	ShouldEndSession bool        `json:"shouldEndSession"`
}

type Reprompt struct {
	OutputSpeech Payload `json:"outputSpeech,omitempty"`
}

type Image struct {
	SmallImageURL string `json:"smallImageUrl,omitempty"`
	LargeImageURL string `json:"largeImageUrl,omitempty"`
}

type Payload struct {
	Type        string   `json:"type,omitempty"`
	Title       string   `json:"title,omitempty"`
	Text        string   `json:"text,omitempty"`
	SSML        string   `json:"ssml,omitempty"`
	Content     string   `json:"content,omitempty"`
	Image       Image    `json:"image,omitempty"`
	Permissions []string `json:"permissions,omitempty"`
}

// Response(s) from requests made back to the Alexa Api

type DeviceLocationResponse struct {
	CountryCode string `json:"countryCode"`
	PostalCode  string `json:"postalCode"`
}
