package alexa

// Credit - Nic Raboy: Modified version of Arien Malec's work
// https://github.com/nraboy/alexa-slick-dealer/blob/master/alexa/request.go
// https://github.com/arienmalec/alexa-go
// https://medium.com/@amalec/alexa-skills-with-go-54db0c21e758

const (
	HelpIntent   = "AMAZON.HelpIntent"
	CancelIntent = "AMAZON.CancelIntent"
	StopIntent   = "AMAZON.StopIntent"
)

type Request struct {
	Version string  `json:"version"`
	Session Session `json:"session"`
	Body    ReqBody `json:"request"`
	Context Context `json:"context"`
}

type Session struct {
	New         bool   `json:"new"`
	SessionID   string `json:"sessionId"`
	Application struct {
		ApplicationID string `json:"applicationId"`
	} `json:"application"`
	Attributes map[string]interface{} `json:"attributes"`
	User       struct {
		UserID      string `json:"userId"`
		AccessToken string `json:"accessToken,omitempty"`
	} `json:"user"`
}

type Context struct {
	AudioPlayer struct {
		PlayerActivity string `json:"playerActivity,omitempty"` // i.e. 'IDLE'
	} `json:"AudioPlayer,omitempty"`
	Display struct {
		Token string `json:"token,omitempty"`
	} `json:"Display,omitempty"`
	Viewport struct {
		Experiences []struct {
			CanRotate bool `json:"canRotate"`
			CanResize bool `json:"canResize"`
		} `json:"experiences,omitempty"`
		Shape              string   `json:"shape"`              // supported: "RECTANGLE" (i.e. Echo Show), & "ROUND" (i.e. Echo Spot)
		PixelWidth         int      `json:"pixelWidth"`         // maximum viewport value
		PixelHeight        int      `json:"pixelHeight"`        // maximum viewport value
		Width              int      `json:"width"`              // Width of the viewport in dp.
		Height             int      `json:"height"`             // Height of the viewport in dp.
		Dpi                int      `json:"dpi"`                // The pixel density of the viewport.
		CurrentPixelWidth  int      `json:"currentPixelWidth"`  // viewport width that is currently in use
		CurrentPixelHeight int      `json:"currentPixelHeight"` // viewport height that is currently in use
		Theme              string   `json:"theme"`              // supported: "LIGHT" or "DARK" - The basic color scheme in use
		Touch              []string `json:"touch,omitempty"`    // i.e. ["SINGLE"]
		Keyboard           []string `json:"keyboard,omitempty"` // i.e. ["DIRECTION"]
	} `json:"Viewport,omitempty"`
	System struct {
		APIAccessToken string `json:"apiAccessToken"`
		APIEndpoint    string `json:"apiEndpoint"`
		Device         struct {
			DeviceID            string              `json:"deviceId,omitempty"`
			SupportedInterfaces SupportedInterfaces `json:"supportedInterfaces,omitempty"`
		} `json:"device,omitempty"`
		Application struct {
			ApplicationID string `json:"applicationId,omitempty"`
		} `json:"application,omitempty"`
	} `json:"System,omitempty"`
}

// Interfaces Supported by the User's device. This is not comprehensive.
type SupportedInterfaces struct {
	APL struct {
		Runtime struct {
			MaxVersion string `json:"maxVersion,omitempty"`
		} `json:"runtime,omitempty"`
	} `json:"Alexa.Presentation.APL,omitempty"`
	AudioPlayer struct{} `json:"AudioPlayer,omitempty"` // This appears to always be an empty object
}

/**
APL Document UserEvents
see: https://developer.amazon.com/docs/alexa-presentation-language/apl-support-for-your-skill.html#listen-for-apl-userevents-from-alexa

Usage: `json:"event,omitempty"`
*/
type Event struct {
	Source struct {
		Type    string      `json:"type,omitempty"`
		Handler string      `json:"handler,omitempty"`
		ID      string      `json:"id,omitempty"`
		Value   interface{} `json:"value,omitempty"`
	} `json:"source,omitempty"`
	Arguments []string `json:"arguments,omitempty"`
}

type ReqBody struct {
	Type        string `json:"type"`
	RequestID   string `json:"requestId"`
	Timestamp   string `json:"timestamp"`
	Locale      string `json:"locale"`
	Token       string `json:"token,omitempty"`
	Event       Event  `json:"event,omitempty"`
	Intent      Intent `json:"intent,omitempty"`
	Reason      string `json:"reason,omitempty"`
	DialogState string `json:"dialogState,omitempty"`
}

type Intent struct {
	Name  string          `json:"name"`
	Slots map[string]Slot `json:"slots"`
}

type Slot struct {
	Name        string      `json:"name"`
	Value       string      `json:"value"`
	Resolutions Resolutions `json:"resolutions"`
}

type Resolutions struct {
	ResolutionPerAuthority []struct {
		Values []struct {
			Value struct {
				Name string `json:"name"`
				Id   string `json:"id"`
			} `json:"value"`
		} `json:"values"`
	} `json:"resolutionsPerAuthority"`
}
