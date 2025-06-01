package winpush

import "errors"

var (
	ErrExecuteToast error = errors.New("winPush: error execute toast")
	ErrCreateFile   error = errors.New("winPush: error create file for xml push")
	ErrReadXML      error = errors.New("winPush: error read xml for push")
)
