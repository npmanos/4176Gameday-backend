package summary

import (
	"sort"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestSummarizeTeam(t *testing.T) {
	actualSummary, err := SummarizeTeam(testSchema, testMatches)

	if err != nil {
		t.Errorf("did not expect error but got: %v\n", err)
	}

	sort.Slice(actualSummary, func(i, j int) bool {
		return actualSummary[i].Name < actualSummary[j].Name
	})

	sort.Slice(testSummary, func(i, j int) bool {
		return testSummary[i].Name < testSummary[j].Name
	})

	if !cmp.Equal(actualSummary, testSummary) {
		t.Errorf("expected actual summary to equal test summary but got diff: %v\n", cmp.Diff(actualSummary, testSummary))
	}
}

var testSchema Schema = []SchemaField{
	{
		FieldDescriptor: FieldDescriptor{Name: "Cargo Placed"},
		ReportReference: "Cargo Placed",
	},
	{
		FieldDescriptor: FieldDescriptor{Name: "Hatches Placed"},
		ReportReference: "Hatches Placed",
	},
	{
		FieldDescriptor: FieldDescriptor{Name: "Cargo Ship Hatches"},
		ReportReference: "Cargo Ship Hatches",
	},
	{
		FieldDescriptor: FieldDescriptor{Name: "Cargo Ship Cargo"},
		ReportReference: "Cargo Ship Cargo",
	},
	{
		FieldDescriptor: FieldDescriptor{Name: "Rocket Hatches Lvl 1"},
		ReportReference: "Rocket Hatches Lvl 1",
	},
	{
		FieldDescriptor: FieldDescriptor{Name: "Rocket Cargo Lvl 1"},
		ReportReference: "Rocket Cargo Lvl 1",
	},
	{
		FieldDescriptor: FieldDescriptor{Name: "Rocket Hatches Lvl 2"},
		ReportReference: "Rocket Hatches Lvl 2",
	},
	{
		FieldDescriptor: FieldDescriptor{Name: "Rocket Cargo Lvl 2"},
		ReportReference: "Rocket Cargo Lvl 2",
	},
	{
		FieldDescriptor: FieldDescriptor{Name: "Rocket Hatches Lvl 3"},
		ReportReference: "Rocket Hatches Lvl 3",
	},
	{
		FieldDescriptor: FieldDescriptor{Name: "Rocket Cargo Lvl 3"},
		ReportReference: "Rocket Cargo Lvl 3",
	},
	{
		FieldDescriptor: FieldDescriptor{Name: "endgame"},
		TBAReference:    "endgameRobot{{.RobotPosition}}",
	},
	{
		FieldDescriptor: FieldDescriptor{Name: "Climbed Lvl 1"},
		AnyOf: []EqualExpression{
			{
				FieldDescriptor: FieldDescriptor{Name: "endgame"},
				Equals:          "HabLevel1",
			},
		},
	},
	{
		FieldDescriptor: FieldDescriptor{Name: "Climbed Lvl 2"},
		AnyOf: []EqualExpression{
			{
				FieldDescriptor: FieldDescriptor{Name: "endgame"},
				Equals:          "HabLevel2",
			},
		},
	},
	{
		FieldDescriptor: FieldDescriptor{Name: "Climbed Lvl 3"},
		AnyOf: []EqualExpression{
			{
				FieldDescriptor: FieldDescriptor{Name: "endgame"},
				Equals:          "HabLevel3",
			},
		},
	},
	{
		FieldDescriptor: FieldDescriptor{Name: "Climbed Lvl 1+"},
		AnyOf: []EqualExpression{
			{
				FieldDescriptor: FieldDescriptor{Name: "Climbed Lvl 1"},
				Equals:          true,
			},
			{
				FieldDescriptor: FieldDescriptor{Name: "Climbed Lvl 2"},
				Equals:          true,
			},
			{
				FieldDescriptor: FieldDescriptor{Name: "Climbed Lvl 3"},
				Equals:          true,
			},
		},
	},
	{
		FieldDescriptor: FieldDescriptor{Name: "Climbed Lvl 2+"},
		AnyOf: []EqualExpression{
			{
				FieldDescriptor: FieldDescriptor{Name: "Climbed Lvl 2"},
				Equals:          true,
			},
			{
				FieldDescriptor: FieldDescriptor{Name: "Climbed Lvl 3"},
				Equals:          true,
			},
		},
	},
	{
		FieldDescriptor: FieldDescriptor{Name: "Assisted Climb Points"},
		ReportReference: "Assisted Climb Points",
	},
	{
		FieldDescriptor: FieldDescriptor{Name: "Teleop Hatches"},
		Sum: []FieldDescriptor{
			{Name: "Rocket Hatches Lvl 1"},
			{Name: "Rocket Hatches Lvl 2"},
			{Name: "Rocket Hatches Lvl 3"},
			{Name: "Cargo Ship Hatches"},
		},
	},
	{
		FieldDescriptor: FieldDescriptor{Name: "Teleop Cargo"},
		Sum: []FieldDescriptor{
			{Name: "Rocket Cargo Lvl 1"},
			{Name: "Rocket Cargo Lvl 2"},
			{Name: "Rocket Cargo Lvl 3"},
			{Name: "Cargo Ship Cargo"},
		},
	},
	{
		FieldDescriptor: FieldDescriptor{Name: "Teleop Gamepieces"},
		Sum: []FieldDescriptor{
			{Name: "Teleop Hatches"},
			{Name: "Teleop Cargo"},
		},
	},
}

var testSummary Summary = []SummaryStat{
	{
		FieldDescriptor: FieldDescriptor{Name: "Cargo Placed"},
		Average:         0,
		Max:             0,
	},
	{
		FieldDescriptor: FieldDescriptor{Name: "Hatches Placed"},
		Average:         12.0 / 9.0,
		Max:             2,
	},
	{
		FieldDescriptor: FieldDescriptor{Name: "Cargo Ship Hatches"},
		Average:         2.0 / 9.0,
		Max:             1,
	},
	{
		FieldDescriptor: FieldDescriptor{Name: "Cargo Ship Cargo"},
		Average:         8.0 / 9.0,
		Max:             4,
	},
	{
		FieldDescriptor: FieldDescriptor{Name: "Rocket Hatches Lvl 1"},
		Average:         7.0 / 9.0,
		Max:             2,
	},
	{
		FieldDescriptor: FieldDescriptor{Name: "Rocket Cargo Lvl 1"},
		Average:         5.0 / 3.0,
		Max:             2,
	},
	{
		FieldDescriptor: FieldDescriptor{Name: "Rocket Hatches Lvl 2"},
		Average:         16.0 / 9.0,
		Max:             2,
	},
	{
		FieldDescriptor: FieldDescriptor{Name: "Rocket Cargo Lvl 2"},
		Average:         16.0 / 9.0,
		Max:             2,
	},
	{
		FieldDescriptor: FieldDescriptor{Name: "Rocket Hatches Lvl 3"},
		Average:         4.0 / 3.0,
		Max:             2,
	},
	{
		FieldDescriptor: FieldDescriptor{Name: "Rocket Cargo Lvl 3"},
		Average:         1.0,
		Max:             2,
	},
	{
		FieldDescriptor: FieldDescriptor{Name: "Climbed Lvl 1"},
		Average:         0.4375,
		Max:             1,
	},
	{
		FieldDescriptor: FieldDescriptor{Name: "Climbed Lvl 1+"},
		Average:         0.9375,
		Max:             1,
	},
	{
		FieldDescriptor: FieldDescriptor{Name: "Climbed Lvl 2"},
		Average:         0.0625,
		Max:             1,
	},
	{
		FieldDescriptor: FieldDescriptor{Name: "Climbed Lvl 2+"},
		Average:         0.5,
		Max:             1,
	},
	{
		FieldDescriptor: FieldDescriptor{Name: "Climbed Lvl 3"},
		Average:         0.4375,
		Max:             1,
	},
	{
		FieldDescriptor: FieldDescriptor{Name: "Assisted Climb Points"},
		Average:         0,
		Max:             0,
	},
	{
		FieldDescriptor: FieldDescriptor{Name: "Teleop Hatches"},
		Average:         37.0 / 9.0,
		Max:             6,
	},
	{
		FieldDescriptor: FieldDescriptor{Name: "Teleop Cargo"},
		Average:         48.0 / 9.0,
		Max:             7,
	},
	{
		FieldDescriptor: FieldDescriptor{Name: "Teleop Gamepieces"},
		Average:         85.0 / 9.0,
		Max:             12,
	},
	{FieldDescriptor: FieldDescriptor{Name: "endgame"}},
}

var testMatches = []Match{
	{
		Key:           "2019tur_f1m1",
		Reports:       []Report{},
		RobotPosition: 3,
		ScoreBreakdown: map[string]interface{}{
			"endgameRobot1": "HabLevel3",
			"endgameRobot2": "HabLevel3",
			"endgameRobot3": "HabLevel1",
		},
	},
	{
		Key:           "2019tur_f1m2",
		Reports:       []Report{},
		RobotPosition: 3,
		ScoreBreakdown: map[string]interface{}{
			"endgameRobot1": "HabLevel1",
			"endgameRobot2": "HabLevel3",
			"endgameRobot3": "HabLevel3",
		},
	},
	{
		Key:           "2019tur_qf1m1",
		Reports:       []Report{},
		RobotPosition: 3,
		ScoreBreakdown: map[string]interface{}{
			"endgameRobot1": "HabLevel3",
			"endgameRobot2": "HabLevel3",
			"endgameRobot3": "HabLevel1",
		},
	},
	{
		Key:           "2019tur_qf1m2",
		Reports:       []Report{},
		RobotPosition: 3,
		ScoreBreakdown: map[string]interface{}{
			"endgameRobot1": "HabLevel3",
			"endgameRobot2": "HabLevel3",
			"endgameRobot3": "HabLevel1",
		},
	},
	{
		Key:           "2019tur_qm106",
		Reports:       []Report{},
		RobotPosition: 3,
		ScoreBreakdown: map[string]interface{}{
			"endgameRobot1": "HabLevel1",
			"endgameRobot2": "HabLevel1",
			"endgameRobot3": "HabLevel3",
		},
	},
	{
		Key: "2019tur_qm16",
		Reports: []Report{
			{
				{Name: "Sandstorm 1", Value: 1},
				{Name: "Sandstorm 2", Value: 0},
				{Name: "Cargo Placed", Value: 0},
				{Name: "Hatches Placed", Value: 1},
				{Name: "Cargo Ship Hatches", Value: 0},
				{Name: "Cargo Ship Cargo", Value: 0},
				{Name: "Rocket Hatches Lvl 1", Value: 1},
				{Name: "Rocket Cargo Lvl 1", Value: 2},
				{Name: "Rocket Hatches Lvl 2", Value: 2},
				{Name: "Rocket Cargo Lvl 2", Value: 2},
				{Name: "Rocket Hatches Lvl 3", Value: 2},
				{Name: "Rocket Cargo Lvl 3", Value: 0},
				{Name: "Climbed Lvl 1", Value: 0},
				{Name: "Climbed Lvl 2", Value: 0},
				{Name: "Climbed Lvl 3", Value: 1},
				{Name: "Assisted Climb Points", Value: 0},
				{Name: "Total Hatches", Value: 6},
				{Name: "Total Cargo", Value: 4},
			},
			{
				{Name: "Sandstorm 1", Value: 1},
				{Name: "Sandstorm 2", Value: 0},
				{Name: "Cargo Placed", Value: 0},
				{Name: "Hatches Placed", Value: 1},
				{Name: "Cargo Ship Hatches", Value: 0},
				{Name: "Cargo Ship Cargo", Value: 0},
				{Name: "Rocket Hatches Lvl 1", Value: 1},
				{Name: "Rocket Cargo Lvl 1", Value: 2},
				{Name: "Rocket Hatches Lvl 2", Value: 2},
				{Name: "Rocket Cargo Lvl 2", Value: 2},
				{Name: "Rocket Hatches Lvl 3", Value: 2},
				{Name: "Rocket Cargo Lvl 3", Value: 0},
				{Name: "Climbed Lvl 1", Value: 0},
				{Name: "Climbed Lvl 2", Value: 0},
				{Name: "Climbed Lvl 3", Value: 1},
				{Name: "Assisted Climb Points", Value: 0},
				{Name: "Total Hatches", Value: 6},
				{Name: "Total Cargo", Value: 4},
			},
		},
		RobotPosition: 2,
		ScoreBreakdown: map[string]interface{}{
			"endgameRobot1": "HabLevel2",
			"endgameRobot2": "HabLevel3",
			"endgameRobot3": "HabLevel3",
		},
	},
	{
		Key: "2019tur_qm26",
		Reports: []Report{
			{
				{Name: "Sandstorm 1", Value: 0},
				{Name: "Sandstorm 2", Value: 1},
				{Name: "Cargo Placed", Value: 0},
				{Name: "Hatches Placed", Value: 1},
				{Name: "Cargo Ship Hatches", Value: 0},
				{Name: "Cargo Ship Cargo", Value: 4},
				{Name: "Rocket Hatches Lvl 1", Value: 1},
				{Name: "Rocket Cargo Lvl 1", Value: 2},
				{Name: "Rocket Hatches Lvl 2", Value: 0},
				{Name: "Rocket Cargo Lvl 2", Value: 0},
				{Name: "Rocket Hatches Lvl 3", Value: 0},
				{Name: "Rocket Cargo Lvl 3", Value: 1},
				{Name: "Climbed Lvl 1", Value: 0},
				{Name: "Climbed Lvl 2", Value: 0},
				{Name: "Climbed Lvl 3", Value: 1},
				{Name: "Assisted Climb Points", Value: 0},
				{Name: "Total Hatches", Value: 2},
				{Name: "Total Cargo", Value: 7},
			},
		},
		RobotPosition: 2,
		ScoreBreakdown: map[string]interface{}{
			"endgameRobot1": "HabLevel2",
			"endgameRobot2": "HabLevel2",
			"endgameRobot3": "None",
		},
	},
	{
		Key: "2019tur_qm36",
		Reports: []Report{
			{
				{Name: "Sandstorm 1", Value: 1},
				{Name: "Sandstorm 2", Value: 0},
				{Name: "Cargo Placed", Value: 0},
				{Name: "Hatches Placed", Value: 2},
				{Name: "Cargo Ship Hatches", Value: 0},
				{Name: "Cargo Ship Cargo", Value: 0},
				{Name: "Rocket Hatches Lvl 1", Value: 0},
				{Name: "Rocket Cargo Lvl 1", Value: 2},
				{Name: "Rocket Hatches Lvl 2", Value: 2},
				{Name: "Rocket Cargo Lvl 2", Value: 2},
				{Name: "Rocket Hatches Lvl 3", Value: 2},
				{Name: "Rocket Cargo Lvl 3", Value: 2},
				{Name: "Climbed Lvl 1", Value: 0},
				{Name: "Climbed Lvl 2", Value: 0},
				{Name: "Climbed Lvl 3", Value: 0},
				{Name: "Assisted Climb Points", Value: 0},
				{Name: "Total Hatches", Value: 6},
				{Name: "Total Cargo", Value: 6},
			},
		},
		RobotPosition: 2,
		ScoreBreakdown: map[string]interface{}{
			"endgameRobot1": "HabLevel3",
			"endgameRobot2": "None",
			"endgameRobot3": "HabLevel1",
		},
	},
	{
		Key: "2019tur_qm47",
		Reports: []Report{
			{
				{Name: "Sandstorm 1", Value: 1},
				{Name: "Sandstorm 2", Value: 0},
				{Name: "Cargo Placed", Value: 0},
				{Name: "Hatches Placed", Value: 1},
				{Name: "Cargo Ship Hatches", Value: 1},
				{Name: "Cargo Ship Cargo", Value: 1},
				{Name: "Rocket Hatches Lvl 1", Value: 1},
				{Name: "Rocket Cargo Lvl 1", Value: 0},
				{Name: "Rocket Hatches Lvl 2", Value: 2},
				{Name: "Rocket Cargo Lvl 2", Value: 2},
				{Name: "Rocket Hatches Lvl 3", Value: 0},
				{Name: "Rocket Cargo Lvl 3", Value: 0},
				{Name: "Climbed Lvl 1", Value: 0},
				{Name: "Climbed Lvl 2", Value: 0},
				{Name: "Climbed Lvl 3", Value: 1},
				{Name: "Assisted Climb Points", Value: 0},
				{Name: "Total Hatches", Value: 5},
				{Name: "Total Cargo", Value: 3},
			},
		},
		RobotPosition: 3,
		ScoreBreakdown: map[string]interface{}{
			"endgameRobot1": "HabLevel1",
			"endgameRobot2": "HabLevel1",
			"endgameRobot3": "HabLevel3",
		},
	},
	{
		Key: "2019tur_qm66",
		Reports: []Report{
			{
				{Name: "Sandstorm 1", Value: 1},
				{Name: "Sandstorm 2", Value: 0},
				{Name: "Cargo Placed", Value: 0},
				{Name: "Hatches Placed", Value: 2},
				{Name: "Cargo Ship Hatches", Value: 0},
				{Name: "Cargo Ship Cargo", Value: 3},
				{Name: "Rocket Hatches Lvl 1", Value: 0},
				{Name: "Rocket Cargo Lvl 1", Value: 2},
				{Name: "Rocket Hatches Lvl 2", Value: 2},
				{Name: "Rocket Cargo Lvl 2", Value: 2},
				{Name: "Rocket Hatches Lvl 3", Value: 0},
				{Name: "Rocket Cargo Lvl 3", Value: 0},
				{Name: "Climbed Lvl 1", Value: 0},
				{Name: "Climbed Lvl 2", Value: 0},
				{Name: "Climbed Lvl 3", Value: 1},
				{Name: "Assisted Climb Points", Value: 0},
				{Name: "Total Hatches", Value: 4},
				{Name: "Total Cargo", Value: 7},
			},
		},
		RobotPosition: 1,
		ScoreBreakdown: map[string]interface{}{
			"endgameRobot1": "HabLevel3",
			"endgameRobot2": "HabLevel2",
			"endgameRobot3": "HabLevel1",
		},
	},
	{
		Key: "2019tur_qm7",
		Reports: []Report{
			{
				{Name: "Sandstorm 1", Value: 1},
				{Name: "Sandstorm 2", Value: 0},
				{Name: "Cargo Placed", Value: 0},
				{Name: "Hatches Placed", Value: 0},
				{Name: "Cargo Ship Hatches", Value: 0},
				{Name: "Cargo Ship Cargo", Value: 0},
				{Name: "Rocket Hatches Lvl 1", Value: 2},
				{Name: "Rocket Cargo Lvl 1", Value: 2},
				{Name: "Rocket Hatches Lvl 2", Value: 2},
				{Name: "Rocket Cargo Lvl 2", Value: 2},
				{Name: "Rocket Hatches Lvl 3", Value: 2},
				{Name: "Rocket Cargo Lvl 3", Value: 2},
				{Name: "Climbed Lvl 1", Value: 1},
				{Name: "Climbed Lvl 2", Value: 0},
				{Name: "Climbed Lvl 3", Value: 0},
				{Name: "Assisted Climb Points", Value: 0},
				{Name: "Total Hatches", Value: 6},
				{Name: "Total Cargo", Value: 6},
			},
		},
		RobotPosition: 1,
		ScoreBreakdown: map[string]interface{}{
			"endgameRobot1": "HabLevel1",
			"endgameRobot2": "None",
			"endgameRobot3": "HabLevel1",
		},
	},
	{
		Key: "2019tur_qm73",
		Reports: []Report{
			{
				{Name: "Sandstorm 1", Value: 0},
				{Name: "Sandstorm 2", Value: 0},
				{Name: "Cargo Placed", Value: 0},
				{Name: "Hatches Placed", Value: 2},
				{Name: "Cargo Ship Hatches", Value: 0},
				{Name: "Cargo Ship Cargo", Value: 0},
				{Name: "Rocket Hatches Lvl 1", Value: 1},
				{Name: "Rocket Cargo Lvl 1", Value: 2},
				{Name: "Rocket Hatches Lvl 2", Value: 2},
				{Name: "Rocket Cargo Lvl 2", Value: 2},
				{Name: "Rocket Hatches Lvl 3", Value: 2},
				{Name: "Rocket Cargo Lvl 3", Value: 0},
				{Name: "Climbed Lvl 1", Value: 1},
				{Name: "Climbed Lvl 2", Value: 0},
				{Name: "Climbed Lvl 3", Value: 0},
				{Name: "Assisted Climb Points", Value: 0},
				{Name: "Total Hatches", Value: 7},
				{Name: "Total Cargo", Value: 4},
			},
		},
		RobotPosition: 2,
		ScoreBreakdown: map[string]interface{}{
			"endgameRobot1": "HabLevel3",
			"endgameRobot2": "HabLevel1",
			"endgameRobot3": "HabLevel2",
		},
	},
	{
		Key: "2019tur_qm86",
		Reports: []Report{
			{
				{Name: "Sandstorm 1", Value: 1},
				{Name: "Sandstorm 2", Value: 0},
				{Name: "Cargo Placed", Value: 0},
				{Name: "Hatches Placed", Value: 1},
				{Name: "Cargo Ship Hatches", Value: 0},
				{Name: "Cargo Ship Cargo", Value: 0},
				{Name: "Rocket Hatches Lvl 1", Value: 1},
				{Name: "Rocket Cargo Lvl 1", Value: 1},
				{Name: "Rocket Hatches Lvl 2", Value: 2},
				{Name: "Rocket Cargo Lvl 2", Value: 2},
				{Name: "Rocket Hatches Lvl 3", Value: 2},
				{Name: "Rocket Cargo Lvl 3", Value: 2},
				{Name: "Climbed Lvl 1", Value: 0},
				{Name: "Climbed Lvl 2", Value: 0},
				{Name: "Climbed Lvl 3", Value: 1},
				{Name: "Assisted Climb Points", Value: 0},
				{Name: "Total Hatches", Value: 6},
				{Name: "Total Cargo", Value: 5},
			},
		},
		RobotPosition: 1,
		ScoreBreakdown: map[string]interface{}{
			"endgameRobot1": "HabLevel3",
			"endgameRobot2": "HabLevel3",
			"endgameRobot3": "HabLevel2",
		},
	},
	{
		Key: "2019tur_qm93",
		Reports: []Report{
			{
				{Name: "Sandstorm 1", Value: 1},
				{Name: "Sandstorm 2", Value: 0},
				{Name: "Cargo Placed", Value: 0},
				{Name: "Hatches Placed", Value: 2},
				{Name: "Cargo Ship Hatches", Value: 1},
				{Name: "Cargo Ship Cargo", Value: 0},
				{Name: "Rocket Hatches Lvl 1", Value: 0},
				{Name: "Rocket Cargo Lvl 1", Value: 2},
				{Name: "Rocket Hatches Lvl 2", Value: 2},
				{Name: "Rocket Cargo Lvl 2", Value: 2},
				{Name: "Rocket Hatches Lvl 3", Value: 2},
				{Name: "Rocket Cargo Lvl 3", Value: 2},
				{Name: "Climbed Lvl 1", Value: 1},
				{Name: "Climbed Lvl 2", Value: 0},
				{Name: "Climbed Lvl 3", Value: 0},
				{Name: "Assisted Climb Points", Value: 0},
				{Name: "Total Hatches", Value: 7},
				{Name: "Total Cargo", Value: 6},
			},
		},
		RobotPosition: 3,
		ScoreBreakdown: map[string]interface{}{
			"endgameRobot1": "HabLevel3",
			"endgameRobot2": "HabLevel2",
			"endgameRobot3": "HabLevel1",
		},
	},
	{
		Key:           "2019tur_sf1m1",
		Reports:       []Report{},
		RobotPosition: 3,
		ScoreBreakdown: map[string]interface{}{
			"endgameRobot1": "HabLevel3",
			"endgameRobot2": "HabLevel1",
			"endgameRobot3": "HabLevel1",
		},
	},
	{
		Key:           "2019tur_sf1m2",
		Reports:       []Report{},
		RobotPosition: 3,
		ScoreBreakdown: map[string]interface{}{
			"endgameRobot1": "None",
			"endgameRobot2": "HabLevel3",
			"endgameRobot3": "HabLevel3",
		},
	},
}
