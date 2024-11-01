package pkgerr

import (
	"net/http"
	"transactor-server/pkg/db/ent"
)

// WrapDAOError wraps certain ent errors to specific status codes
func WrapDAOError(err error) error {
	if ent.IsConstraintError(err) {
		return NewServiceError("db", "constraint", http.StatusBadRequest, err.Error())
	} else if ent.IsNotFound(err) {
		return NewServiceError("db", "not_found", http.StatusNotFound, err.Error())
	}
	return NewServiceError("db", "unkown", http.StatusInternalServerError, err.Error())
}
