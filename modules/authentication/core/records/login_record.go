package records

type LoginAttempStatus int8

var (
	Succeded = 0
	Failed   = 1
)

type LoginRecord struct {
	Status LoginAttempStatus
}
