package _TemplatePackageName

import "errors"

// StopIterationError ...
type StopIterationError struct {
	msg string
}

func (sie *StopIterationError) Error() string {
	return sie.msg
}

// PartialResultError ...
type PartialResultError struct {
	msg string
}

func (pre *PartialResultError) Error() string {
	return pre.msg
}

// SearherInitializationError ...
type SearherInitializationError struct {
	msg string
}

func (csie *SearherInitializationError) Error() string {
	return csie.msg
}

var (
	errEmptyAdnID = errors.New("adnID is empty")
	//errAdTypesOrAdSettingsError = errors.New("ad_types or ad_settings should be in the query")
)
