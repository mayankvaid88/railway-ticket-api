package error

import (
	"google.golang.org/grpc/codes"
)

type errorResponse struct {
	StatusCode codes.Code
	Code       ErrorCode
	Message    string
}

func (e errorResponse) Error() string {
	return e.Message
}

var errorResponseMap = map[ErrorCode]errorResponse{
	SeatNotAvailable: {
		StatusCode: codes.ResourceExhausted,
		Code:       SeatNotAvailable,
		Message:    "seat not available",
	},
	TicketNotFound: {
		StatusCode: codes.NotFound,
		Code:       TicketNotFound,
		Message:    "ticket not found",
	},
	InternalServerError: {
		StatusCode: codes.Internal,
		Code:       InternalServerError,
		Message:    "internal server error",
	},
	SeatNotAllocationForSection: {
		StatusCode: codes.NotFound,
		Code:       SeatNotAllocationForSection,
		Message:    "seats are not allocated for given section",
	},
	BadRequest: {
		StatusCode: codes.InvalidArgument,
		Code: BadRequest,
		Message: "request is invalid.",
	},
}

func GetErrorResponseByCode(errCode ErrorCode) errorResponse {
	return errorResponseMap[errCode]
}

func GetErrorResponse(code codes.Code,errCode ErrorCode,errorMessage string) errorResponse {
	return errorResponse{

		Code: errCode,
		Message: errorMessage,
	}
}

type ErrorCode string

const (
	SeatNotAvailable            ErrorCode = "SEAT_NOT_AVAILABLE"
	TicketNotFound              ErrorCode = "TICKET_NOT_FOUND"
	SeatNotAllocationForSection ErrorCode = "SEAT_NOT_ALLOCATED_FOR_SECTION"
	InternalServerError         ErrorCode = "INTERNAL_SERVER_ERROR"
	BadRequest									ErrorCode = "BAD_REQUEST"
)
