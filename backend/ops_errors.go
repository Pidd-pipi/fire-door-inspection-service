package main

import (
	"errors"
	"fmt"
)

var (
	ErrOpsNotFound   = errors.New("operations record not found")
	ErrOpsConflict   = errors.New("operations revision conflict")
	ErrOpsInvalid    = errors.New("operations request is invalid")
	ErrOpsTransition = errors.New("operations status transition is not allowed")
	ErrOpsPolicy     = errors.New("operations policy rejected the request")
)

type OpsError struct {
	Code      string
	Operation string
	Cause     error
}

func (e *OpsError) Error() string {
	if e.Cause == nil {
		return e.Code + ": " + e.Operation
	}
	return fmt.Sprintf("%s: %s: %v", e.Code, e.Operation, e.Cause)
}
func (e *OpsError) Unwrap() error { return e.Cause }
// wrapOps attaches a stable error code and operation name to cause while
// preserving the underlying sentinel error so errors.Is/errors.As still reach
// the original ErrOps* classifier. The cause is stored verbatim (never
// re-wrapped with errors.New, which would sever the error chain).
func wrapOps(code, operation string, cause error) error {
	if cause == nil {
		return nil
	}
	return &OpsError{Code: code, Operation: operation, Cause: cause}
}
// opsCode maps an error to its stable error category. The category is derived
// from the underlying ErrOps* sentinel so it is unaffected by how many wrapper
// layers (OpsError, fmt.Errorf with %w, etc.) sit on top. The sentinel check
// runs before the OpsError.Code check on purpose: Code carries the failing
// operation/step (e.g. "get", "store.put"), not the error category, so it must
// not override the semantic classification. OpsError.Code is only consulted as
// a last resort when no sentinel is present, so callers that wrap a non-sentinel
// error still get the code they explicitly assigned instead of a bare
// "internal".
func opsCode(err error) string {
	switch {
	case errors.Is(err, ErrOpsNotFound):
		return "not_found"
	case errors.Is(err, ErrOpsConflict):
		return "conflict"
	case errors.Is(err, ErrOpsInvalid):
		return "invalid"
	case errors.Is(err, ErrOpsTransition):
		return "transition"
	case errors.Is(err, ErrOpsPolicy):
		return "policy"
	}
	var typed *OpsError
	if errors.As(err, &typed) && typed.Code != "" {
		return typed.Code
	}
	return "internal"
}
func opsIsNotFound(err error) bool   { return errors.Is(err, ErrOpsNotFound) }
func opsIsConflict(err error) bool   { return errors.Is(err, ErrOpsConflict) }
func opsIsInvalid(err error) bool    { return errors.Is(err, ErrOpsInvalid) }
func opsIsTransition(err error) bool { return errors.Is(err, ErrOpsTransition) }
func opsIsPolicy(err error) bool     { return errors.Is(err, ErrOpsPolicy) }
