package models
type Abonemen struct {
	UserID     string  `json:"userID"`
	SubscriberID string `json:"subscriberID"`
	Notification bool `json:"notification"`
}
