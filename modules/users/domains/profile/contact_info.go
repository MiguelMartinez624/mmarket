package profile

type CommunicationChannel int8

const (
	Email      CommunicationChannel = 0
	MobilPhone                      = 1
	LocalPhone                      = 2
)

type ContactInfo struct {
	ID      string
	Channel CommunicationChannel
	Key     string
	ItsMain bool
}
