# `go-alexa-devkit`

## A developer kit for generating simple to complex Alexa Skill Responses
_More Documentation Forthcomming._

##### **Pull Requests & other contributions Welcome!**

### Payload Structs & Builders
- The `alexa.Response{}` `struct` should fit the majority of use cases for even some more complex resposnses. 
  - `alexa/directives-builder.go` provides additional structs and functions related to building directives for more advanced Alexa Skill features.
  - `ssml-response-builder.go` adds functions, structs, and builder support for constructing SSML enriched responses.
- `request.go` contains the structs for parsing an incoming skill request payload.
- The `apl` package provides support for the Alexa Presentation Language (APL) for devices with screens.
  - Provides support for parsing an apl template `.json` file.
    -`apl` package function is broken. Use the `alexa.ExtractNewRenderDocDirectiveFromJson()` function in `alexa/directives-builder.go`
  - Structs for working with an APL Document.
- Additional features are also supported. 

### **If you see a missing feature, please feel free to create an issue and/or a PR to request the change. Currently only being added as the need/request arises.**
