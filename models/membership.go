package models

type MembershipDetail struct {
	MembershipID string `json:"membershipId"`
	UserID     string      `json:"userId"`
	Date       string      `json:"date"`
	Bill       int         `json:"bill"`
}