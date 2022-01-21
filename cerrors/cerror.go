package cerrors

import (
	"errors"
	"github.com/rozturac/cerror"
)

func Is(source, target error) bool {
	cSource, cSourceOk := source.(cerror.Error)
	cTarget, cTargetOk := target.(cerror.Error)
	if cSourceOk && cTargetOk {
		return cSource.ErrorType() == cTarget.ErrorType() &&
			cSource.ErrorCode() == cTarget.ErrorCode() &&
			cSource.Error() == cTarget.Error()
	} else if !cSourceOk && !cTargetOk {
		return errors.Is(source, target)
	} else {
		return false
	}
}
