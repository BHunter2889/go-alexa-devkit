package alexa

import (
	"log"
	"fmt"
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

func NewPermissionsRequestResponse(skillName string) Response {
	var builder SSMLBuilder
	builder.Say(fmt.Sprintf("%s was unable to access your device's zip code and country information. ", skillName))
	builder.Pause("750")
	builder.Say(fmt.Sprintf("If you have not enabled %s to access this information, ", skillName))
	builder.Pause("150")
	builder.Say(fmt.Sprintf("Please check your Alexa App to grant permission for %s to access your zip code and country " +
				"information to enable full interaction capabilities. ", skillName))
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

func NewUnsupportedLocationResponse(skillName string) Response {
	var builder SSMLBuilder
	builder.Say(fmt.Sprintf("%s does not currently support devices in your location. ", skillName))
	builder.Pause("750")
	builder.Say("If you would like us to provide support for your locale, ")
	builder.Pause("150")
	builder.Say(fmt.Sprintf("Please leave feedback on the %s skill page in the Alexa App Skill Store with your country or locale information. ", skillName))
	builder.Pause("750")
	builder.Say("If you are presently in a supported locale, you may need to alter your device's settings in the Alexa App.")
	builder.Pause("500")
	builder.Say(fmt.Sprintf("The %s team apologizes for the inconvenience.", skillName))

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

func NewLaunchRequestGetPermissionsResponse(skillName string, features []string) Response {
	var builder SSMLBuilder
	builder.Say(fmt.Sprintf("Welcome to %s!", skillName))
	builder.Pause("1000")
	builder.Say(fmt.Sprintf("%s will need to access your device's zip code and country information. ", skillName))
	builder.Pause("750")
	builder.Say(fmt.Sprintf("If you have not already enabled %s to access this information, ", skillName))
	builder.Pause("150")
	builder.Say(fmt.Sprintf("Please check your Alexa App to grant permission for %s to access your zip code and country " +
				"information in order to have full feature capabilities. ", skillName))
	builder.Pause("750")
	builder.Say("Currently, ")
	builder.Pause("100")
	builder.Say("once you have granted this permission, ")
	builder.Pause("100")
	builder.Say(fmt.Sprintf("You can have Alexa ask %s %s, ", skillName, features[0]))
	builder.Pause("150")
	builder.Say(fmt.Sprintf("or %s, ", features[1]))
	builder.Pause("150")
	builder.Say(fmt.Sprintf("and %s. ", features[2]))
	builder.Pause("1000")
	builder.Say(fmt.Sprintf("We hope %s fulfills or exceeeds your expectations and appreciate any feedback through reviews on the skill page in the Alexa Skill Store! ")
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

func NewDefaultErrorResponse(skillName string) Response {
	var builder SSMLBuilder
	builder.Say(fmt.Sprintf("%s encountered some issues while processing your request. ", skillName))
	builder.Pause("750")
	builder.Say(fmt.Sprintf("The problem has been recorded and will be checked out by the %s team, ", skillName))
	builder.Pause("100")
	builder.Say("So, please accept our apologies for the inconvenience. ")
	builder.Pause("500")
	builder.Say(fmt.Sprintf("Please try %s again later.", skillName))

	r := Response{
		Version: "1.0",
		Body: ResBody{
			OutputSpeech: &Payload{
				Type: "SSML",
				SSML: builder.Build(),
			},
			Card: &Payload{
				Type:  "Simple",
				Title: fmt.Sprintf("%s Under Maintenance", skillName),
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
