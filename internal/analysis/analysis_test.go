package analysis

import (
	"testing"

	"github.com/Pigmice2733/peregrine-backend/internal/store"
)

func newInt(a int) *int {
	return &a
}

func newBool(a bool) *bool {
	return &a
}

func TestAnalsyzeReports(t *testing.T) {
	reports := []store.Report{
		store.Report{
			ID:      0,
			TeamKey: "frc2471",
			Data: store.ReportData{
				Teleop: []store.Stat{
					{Name: "Cubes", Attempts: newInt(5), Successes: newInt(3)},
				},
				Auto: []store.Stat{
					{Name: "Line", Attempted: newBool(true), Succeeded: newBool(false)},
				},
			},
		},
		store.Report{
			ID:      1,
			TeamKey: "frc2733",
			Data: store.ReportData{
				Teleop: []store.Stat{
					{Name: "Cubes", Attempts: newInt(8), Successes: newInt(6)},
				},
				Auto: []store.Stat{
					{Name: "Line", Attempted: newBool(true), Succeeded: newBool(true)},
				},
			},
		},
		store.Report{
			ID:      2,
			TeamKey: "frc2733",
			Data: store.ReportData{
				Teleop: []store.Stat{
					{Name: "Line", Attempted: newBool(true), Succeeded: newBool(false)},
				},
				Auto: []store.Stat{
					{Name: "Cubes", Attempts: newInt(5), Successes: newInt(3)},
				},
			},
		},
		store.Report{
			ID:      3,
			TeamKey: "frc2471",
			Data: store.ReportData{
				Teleop: []store.Stat{
					{Name: "Cubes", Attempts: newInt(2), Successes: newInt(2)},
				},
				Auto: []store.Stat{
					{Name: "Line", Attempted: newBool(true), Succeeded: newBool(true)},
				},
			},
		},
	}

	schema := store.Schema{
		Auto:   []store.StatDescription{{Name: "Line", Type: "boolean"}},
		Teleop: []store.StatDescription{{Name: "Cubes", Type: "number"}},
	}

	analyzedStats, err := AnalyzeReports(schema, reports)
	if err != nil {
		t.Errorf("analysis failed with error: %v", err)
	}

	if _, ok := analyzedStats["frc2733"]; !ok {
		t.Errorf("analysis doesn't contain team frc2733")
	}
	if _, ok := analyzedStats["frc2471"]; !ok {
		t.Errorf("analysis doesn't contain team frc2471")
	}

	if analyzedStats["frc2733"].Team != "frc2733" {
		t.Errorf("analysis for team frc2733 has wrong team key")
	}
	if analyzedStats["frc2471"].Team != "frc2471" {
		t.Errorf("analysis for team frc2471 has wrong team key")
	}

	for name, stat := range analyzedStats["frc2471"].AutoBoolean {
		if stat.Name != name {
			t.Errorf("stat %s has wrong stat name: %s", name, stat.Name)
		}

		if stat.Name != "Line" {
			if stat.Attempts != 2 || stat.Successes != 1 {
				t.Errorf("analysis for frc2471 'Line' is wrong")
			}
		}
	}

	for name, stat := range analyzedStats["frc2471"].TeleopNumeric {
		if stat.Name != name {
			t.Errorf("stat %s has wrong stat name: %s", name, stat.Name)
		}

		if stat.Name == "Cubes" {
			attempts := MaxAvg{
				Max:     5,
				Avg:     3.5,
				Total:   7,
				Matches: 2,
			}

			successes := MaxAvg{
				Max:     3,
				Avg:     2.5,
				Total:   5,
				Matches: 2,
			}
			if stat.Attempts != attempts || stat.Successes != successes {
				t.Errorf("analysis for frc2471 'Cubes' is wrong")
			}
		}
	}
}
