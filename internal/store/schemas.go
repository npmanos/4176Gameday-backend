package store

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"fmt"

	"github.com/lib/pq"
	"github.com/pkg/errors"
)

// Schema describes the statistics that reports should include
type Schema struct {
	ID      int64            `json:"id" db:"id"`
	Year    *int64           `json:"year,omitempty" db:"year"`
	RealmID *int64           `json:"realmId,omitempty" db:"realm_id"`
	Auto    StatDescriptions `json:"auto" db:"auto"`
	Teleop  StatDescriptions `json:"teleop" db:"teleop"`
}

// StatDescription describes a single statistic in a schema
type StatDescription struct {
	Name  string      `json:"name"`
	Type  string      `json:"type"`
	Sum   []Reference `json:"sum"`   // only non-nil on type "number" stats
	AnyOf []AnyOf     `json:"anyOf"` // only non-nil on type "boolean" stats
}

// Reference is used to refer to a TBA or report field.
type Reference struct {
	// TBA match score breakdown field name, you can use {{ robotIndex }}
	// to inject the index of the robot (1, 2, 3) for fields like "endgameRobot1"
	TBA string `json:"tba"`
	// Field is the name of the report field to use. Note you can only use TBA *or*
	// Field, not both.
	Field string `json:"field"`
}

// AnyOf is used for boolean stats. If the TBA ore report field value matches the
// given equals value, the boolean stat is true.
type AnyOf struct {
	Reference
	// Equals is a string, but the value it refers to can be a string, number, or
	// bool. If the reference refers to a number or bool the string will be parsed
	// and then compared.
	// We decided to go with this over using an interface or multiple fields for
	// each type.
	Equals string `json:"equals"`
}

// StatDescriptions holds multiple StatDescriptions for storing in one DB column
type StatDescriptions []StatDescription

// Value implements driver.Valuer to return JSON for the DB from StatDescription.
func (sd StatDescriptions) Value() (driver.Value, error) { return json.Marshal(sd) }

// Scan implements sql.Scanner to scan JSON from the DB into StatDescriptions.
func (sd *StatDescriptions) Scan(src interface{}) error {
	j, ok := src.([]byte)
	if !ok {
		return errors.New("got invalid type for StatDescriptions")
	}

	return json.Unmarshal(j, sd)
}

// CreateSchema creates a new schema
func (s *Service) CreateSchema(ctx context.Context, schema Schema) error {
	tx, err := s.db.BeginTxx(ctx, nil)
	if err != nil {
		return errors.Wrap(err, "unable to begin transaction")
	}

	_, err = tx.NamedExecContext(ctx, `
	INSERT
		INTO
			schemas (year, realm_id, auto, teleop)
		VALUES (:year, :realm_id, :auto, :teleop)
	`, schema)

	if err, ok := err.(*pq.Error); ok && err.Code == pgExists {
		_ = tx.Rollback()
		return &ErrExists{fmt.Errorf("schema already exists: %v", err.Error())}
	} else if err != nil {
		_ = tx.Rollback()
		return errors.Wrap(err, "unable to insert schema")
	}

	return errors.Wrap(tx.Commit(), "unable to commit transaction")
}

// GetSchemaByID retrieves a schema given its ID
func (s *Service) GetSchemaByID(ctx context.Context, id int64) (Schema, error) {
	var schema Schema

	err := s.db.GetContext(ctx, &schema, "SELECT * FROM schemas WHERE id = $1", id)
	if err == sql.ErrNoRows {
		return schema, ErrNoResults{fmt.Errorf("schema %d does not exist", schema.ID)}
	}

	return schema, errors.Wrap(err, "unable to retrieve schema")
}

// GetSchemaByYear retrieves the schema for a given year
func (s *Service) GetSchemaByYear(ctx context.Context, year int) (Schema, error) {
	var schema Schema

	err := s.db.GetContext(ctx, &schema, "SELECT * FROM schemas WHERE year = $1", year)
	if err == sql.ErrNoRows {
		return schema, ErrNoResults{fmt.Errorf("no schema for year %d exists", year)}
	}

	return schema, errors.Wrap(err, "unable to retrieve schema")
}

// GetVisibleSchemas retrieves schemas from the database frm a specific realm,
// from realms with public events, and standard FRC schemas. If the realm ID is
// nil, no private realms' schemas will be retrieved.
func (s *Service) GetVisibleSchemas(ctx context.Context, realmID *int64) ([]Schema, error) {
	schemas := []Schema{}
	var err error

	if realmID == nil {
		err = s.db.SelectContext(ctx, &schemas, `
		WITH public_realms AS (
			SELECT id FROM realms WHERE share_reports = true
		)
		SELECT * 
			FROM schemas
				WHERE year IS NULL OR realm_id IN (SELECT id FROM public_realms)
		`)
	} else {
		err = s.db.SelectContext(ctx, &schemas, `
		WITH public_realms AS (
			SELECT id FROM realms WHERE share_reports = true
		)
		SELECT * 
			FROM schemas
			    WHERE year IS NULL OR realm_id = $1 OR (SELECT id FROM public_realms)
		`, *realmID)
	}

	return schemas, errors.Wrap(err, "unable to retrieve schemas")
}

// GetSchemas retrieves all schemas from the database.
func (s *Service) GetSchemas(ctx context.Context) ([]Schema, error) {
	schemas := []Schema{}
	err := s.db.SelectContext(ctx, &schemas, "SELECT * FROM schemas")
	return schemas, errors.Wrap(err, "unable to retrieve schemas")
}
