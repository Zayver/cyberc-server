package model

type ComplaintType uint
const (
	AGAINST_PERSON ComplaintType = iota
	AGAINST_PROPERTY
	SEXUAL
	FINANCIAL
	CYBERNETIC
	AGAINST_PUBLIC_HEALTH
	AGAINST_STATE
)

type ComplaintStatus uint

const (
	CREATED ComplaintStatus = iota
	IN_PROCESS
	FINALIZED
)