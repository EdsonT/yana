package errors

type M map[string]interface{}

type Error struct {
	Code        string
	Status      string
	ErrorType   string
	InnerErrors map[string]string
}

var ie map[string]string

func (e *Error) SetValidationErrors(errs map[string]string) {
	e.Code = codeList["model-validation"]["code"]
	e.Status = codeList["model-validation"]["status"]
	e.ErrorType = codeList["model-validation"]["status_text"]
	e.InnerErrors = errs
}

func (e *Error) SetNotFoundOptionError(key string) {
	e.Code = codeList["model-not-found"]["code"]
	e.Status = codeList["model-not-found"]["status"]
	e.ErrorType = codeList["model-not-found"]["status_text"]
	ie[key] = "Is not a valid option or could not be found"
	e.InnerErrors = ie
}

func (e *Error) SetRepositoryError(errs map[string]string) {
	e.Code = codeList["repository-errors"]["code"]
	e.Status = codeList["repository-errors"]["status"]
	e.ErrorType = codeList["repository-errors"]["status_text"]
	e.InnerErrors = errs
}

var codeList = map[string]map[string]string{
	"model-validation": {
		"code":        "Y5001",
		"status":      "500",
		"status_text": "validation-errors",
	},
	"controllers-merge_params": {
		"code":        "Y5002",
		"status":      "500",
		"status_text": "internal-server-errors",
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
		"status_text": "not-found",
		"reason":      "Could not find {{.resource}} with id `{{.id}}`",
	},
	"repository-error": {
		"code":        "Y5005",
		"status":      "404",
		"status_text": "repository-errors",
		// "reason":      "Could not find {{.resource}} with id `{{.id}}`",
	},
}
