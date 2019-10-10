package defs

type Err struct {
	Error     string `json:"error"`
	ErrorCode string `json:"error_code"`
}

type ErrResponse struct {
	HttpSC int
	Error  Err
}

var (
	ErrorDBError = ErrResponse{
		HttpSC: 500,
		Error: Err{
			Error:     "DB ops failed",
			ErrorCode: "003",
		},
	}

	ErrorInternalFaults = ErrResponse{
		HttpSC: 500,
		Error: Err{
			Error:     "Internal service error",
			ErrorCode: "004",
		},
	}
)
