package controllers

import (
	"github.com/ramasapto/clean-architecture/usecases"
)

// ctrl struct with value interface Usecases
type ctrl struct {
	uc usecases.Usecases
}

// Controllers represent the Controllers contract
type Controllers interface {
	// Authentication controllers
}

/*NewCtrl will create an object that represent the Controllers interface (Controllers)
 * @parameter
 * r - Repository Interface
 *
 * @represent
 * interface Controllers
 *
 * @return
 * uc struct with value interface Usecases
 */
func NewCtrl(u usecases.Usecases) Controllers {
	return &ctrl{uc: u}
}
