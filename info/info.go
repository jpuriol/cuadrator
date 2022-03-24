package info

const (
	participantsFile = "participants.yaml"
	schemaFile       = "schema.yaml"
	quadrantFile     = "quadrant.yaml"
)

type quadrant map[int]shift
type shift map[int][]team
type team []string
