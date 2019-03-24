package errors

type M map[string]interface{}

type Error struct {
	Code        string
	Status      string
	ErrorType   string
	InnerErrors map[string]string
}

func (e *Error) SetValidationErrors(errs map[string]string) {
	e.Code = codeList["controllers-unexpected"]["code"]
	e.Status = codeList["controllers-unexpected"]["status"]
	e.ErrorType = codeList["controllers-unexpected"]["status_text"]
	e.InnerErrors = errs

}
var codeList = map[string]map[string]string{
	"controllers-unexpected": {
		"code":        "Y5001",
		"status":      "500",
		"status_text": "Validation-Errors",
	},
	"controllers-merge_params": {
		"code":        "Y5002",
		"status":      "500",
		"status_text": "Internal Server Error",
	},
	"controllers-params_binding": {
		"code":        "Y5003",
		"status":      "422",
		"status_text": "Unprocessable Entity",
		"reason":      "Validation of <no value> failed",
	},
	"models-not_found": {
		"code":        "Y5004",
		"status":      "404",
		"status_text": "Not Found",
		"reason":      "Could not find {{.resource}} with id `{{.id}}`",
	},
}
