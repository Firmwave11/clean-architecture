package repository

import (
	"errors"

	"gorm.io/gorm"
)

var (
	// ErrRecordNotFound record not found error
	ErrRecordNotFound = errors.New("record not found")
	// ErrInvalidTransaction invalid transaction when you are trying to `Commit` or `Rollback`
	ErrInvalidTransaction = errors.New("no valid transaction")
	// ErrNotImplemented not implemented
	ErrNotImplemented = errors.New("not implemented")
	// ErrMissingWhereClause missing where clause
	ErrMissingWhereClause = errors.New("WHERE conditions required")
	// ErrUnsupportedRelation unsupported relations
	ErrUnsupportedRelation = errors.New("unsupported relations")
	// ErrPrimaryKeyRequired primary keys required
	ErrPrimaryKeyRequired = errors.New("primary key required")
	// ErrModelValueRequired model value required
	ErrModelValueRequired = errors.New("model value required")
	// ErrInvalidData unsupported data
	ErrInvalidData = errors.New("unsupported data")
	// ErrUnsupportedDriver unsupported driver
	ErrUnsupportedDriver = errors.New("unsupported driver")
	// ErrRegistered registered
	ErrRegistered = errors.New("registered")
	// ErrInvalidField invalid field
	ErrInvalidField = errors.New("invalid field")
	// ErrEmptySlice empty slice found
	ErrEmptySlice = errors.New("empty slice found")
	// ErrDryRunModeUnsupported dry run mode unsupported
	ErrDryRunModeUnsupported = errors.New("dry run mode unsupported")
	// ErrConflict for error if data conflict
	ErrConflict = errors.New("data conflict")
	// ErrBadRequest for error bad request
	ErrBadRequest = errors.New("bad request")
	// ErrUnouthorized for error authorization
	ErrUnouthorized = errors.New("unouthorized")
)

// repo struct with value mysqldb connection
type repo struct {
	db *gorm.DB
}

// Repo represent the Repository contract
type Repo interface {
}

/*NewRepo will create an object that represent the Repository interface (Repo)
 * @parameter
 * db - mysql database connection
 *
 * @represent
 * interface Repo
 *
 * @return
 * repo struct with value db (mysql database connection)
 */
func NewRepo(db *gorm.DB) Repo {
	return &repo{db: db}
}
