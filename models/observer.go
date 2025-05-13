package models

// Observer defines the method that observers (participants) must implement
type Observer interface {
	Update(eventType string, group *Group)
}

// Subject defines only the NotifyAll method
type Subject interface {
	NotifyAll(eventType string)
}
