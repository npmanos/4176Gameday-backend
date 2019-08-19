package summary

// Report defines a report for a single team in a single match at a single event, which is
// just a list of name/values.
type Report []ReportField

// ReportField defines a single field of a report, a mapping of a name to a value.
type ReportField struct {
	Name  string
	Value float64
}

// ScoreBreakdown defines a TBA score breakdown (changes year to year) which is a mapping of
// strings to JSON values (float64, bool, string).
type ScoreBreakdown map[string]interface{}

// Match defines information relevant to summarizing matches (match key, reports, score
// breakdowns, alliances). RobotPosition should be the one-indexed position of the robot
// on the field, and the score breakdown should be the relevant score breakdown to the
// alliance the robot was on.
type Match struct {
	Key            string
	Reports        []Report
	RobotPosition  int
	ScoreBreakdown ScoreBreakdown
}

// Schema defines a list of schema fields for a schema. The Schema will outline how to summarize
// data from reports, TBA, and computed properties.
type Schema []SchemaField

// FieldDescriptor defines properties of a schema field that aren't related to how it should be
// summarized, but just information about the field (name, period, type).
type FieldDescriptor struct {
	Name string
}

// SchemaField is a singular schema field. Only specify one of: ReportReference, TBAReference,
// Sum, or AnyOf.
type SchemaField struct {
	FieldDescriptor
	ReportReference string
	TBAReference    string
	Sum             []FieldDescriptor
	AnyOf           []EqualExpression
}

// EqualExpression defines a reference that should equal some JSON value (float64, number,
// string).
type EqualExpression struct {
	FieldDescriptor
	Equals interface{}
}

// Summary defines a summarized list of matches.
type Summary []SummaryStat

// SummaryStat defines a single stat in a match.
type SummaryStat struct {
	FieldDescriptor
	Max     float64
	Average float64
}

// SummarizeTeam summarizes a singular team's performance in a list of matches.
func SummarizeTeam(team string, schema Schema, matches []Match) (Summary, error) {
	return Summary{}, nil
}
