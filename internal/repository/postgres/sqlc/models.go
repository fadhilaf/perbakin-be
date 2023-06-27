// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0

package postgres

import (
	"database/sql/driver"
	"fmt"

	"github.com/jackc/pgx/v5/pgtype"
)

type Stage0Status string

const (
	Stage0Status1 Stage0Status = "1"
	Stage0Status2 Stage0Status = "2"
	Stage0Status3 Stage0Status = "3"
	Stage0Status4 Stage0Status = "4"
	Stage0Status5 Stage0Status = "5"
	Stage0Status6 Stage0Status = "6"
)

func (e *Stage0Status) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = Stage0Status(s)
	case string:
		*e = Stage0Status(s)
	default:
		return fmt.Errorf("unsupported scan type for Stage0Status: %T", src)
	}
	return nil
}

type NullStage0Status struct {
	Stage0Status Stage0Status
	Valid        bool // Valid is true if Stage0Status is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullStage0Status) Scan(value interface{}) error {
	if value == nil {
		ns.Stage0Status, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.Stage0Status.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullStage0Status) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.Stage0Status), nil
}

type Stage13Status string

const (
	Stage13Status1 Stage13Status = "1"
	Stage13Status2 Stage13Status = "2"
	Stage13Status3 Stage13Status = "3"
	Stage13Status4 Stage13Status = "4"
	Stage13Status5 Stage13Status = "5"
	Stage13Status6 Stage13Status = "6"
	Stage13Status7 Stage13Status = "7"
)

func (e *Stage13Status) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = Stage13Status(s)
	case string:
		*e = Stage13Status(s)
	default:
		return fmt.Errorf("unsupported scan type for Stage13Status: %T", src)
	}
	return nil
}

type NullStage13Status struct {
	Stage13Status Stage13Status
	Valid         bool // Valid is true if Stage13Status is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullStage13Status) Scan(value interface{}) error {
	if value == nil {
		ns.Stage13Status, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.Stage13Status.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullStage13Status) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.Stage13Status), nil
}

type Stage246Status string

const (
	Stage246Status1 Stage246Status = "1"
	Stage246Status2 Stage246Status = "2"
	Stage246Status3 Stage246Status = "3"
	Stage246Status4 Stage246Status = "4"
)

func (e *Stage246Status) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = Stage246Status(s)
	case string:
		*e = Stage246Status(s)
	default:
		return fmt.Errorf("unsupported scan type for Stage246Status: %T", src)
	}
	return nil
}

type NullStage246Status struct {
	Stage246Status Stage246Status
	Valid          bool // Valid is true if Stage246Status is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullStage246Status) Scan(value interface{}) error {
	if value == nil {
		ns.Stage246Status, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.Stage246Status.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullStage246Status) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.Stage246Status), nil
}

type Stage5Status string

const (
	Stage5Status1 Stage5Status = "1"
	Stage5Status2 Stage5Status = "2"
	Stage5Status3 Stage5Status = "3"
)

func (e *Stage5Status) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = Stage5Status(s)
	case string:
		*e = Stage5Status(s)
	default:
		return fmt.Errorf("unsupported scan type for Stage5Status: %T", src)
	}
	return nil
}

type NullStage5Status struct {
	Stage5Status Stage5Status
	Valid        bool // Valid is true if Stage5Status is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullStage5Status) Scan(value interface{}) error {
	if value == nil {
		ns.Stage5Status, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.Stage5Status.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullStage5Status) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.Stage5Status), nil
}

type Stages string

const (
	Stages0 Stages = "0"
	Stages1 Stages = "1"
	Stages2 Stages = "2"
	Stages3 Stages = "3"
	Stages4 Stages = "4"
	Stages5 Stages = "5"
	Stages6 Stages = "6"
	Stages7 Stages = "7"
)

func (e *Stages) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = Stages(s)
	case string:
		*e = Stages(s)
	default:
		return fmt.Errorf("unsupported scan type for Stages: %T", src)
	}
	return nil
}

type NullStages struct {
	Stages Stages
	Valid  bool // Valid is true if Stages is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullStages) Scan(value interface{}) error {
	if value == nil {
		ns.Stages, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.Stages.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullStages) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.Stages), nil
}

type Admin struct {
	ID     pgtype.UUID
	UserID pgtype.UUID
	ExamID pgtype.UUID
}

type Exam struct {
	ID        pgtype.UUID
	SuperID   pgtype.UUID
	Name      string
	Location  string
	Organizer string
	Begin     pgtype.Date
	Finish    pgtype.Date
	Active    bool
	CreatedAt pgtype.Timestamp
	UpdatedAt pgtype.Timestamp
}

type Result struct {
	ID        pgtype.UUID
	ShooterID pgtype.UUID
	Failed    bool
	Stage     Stages
	CreatedAt pgtype.Timestamp
	UpdatedAt pgtype.Timestamp
}

type Scorer struct {
	ID        pgtype.UUID
	UserID    pgtype.UUID
	ExamID    pgtype.UUID
	ImagePath string
}

type Session struct {
	Token  string
	Data   []byte
	Expiry pgtype.Timestamptz
}

type Shooter struct {
	ID        pgtype.UUID
	ScorerID  pgtype.UUID
	Name      string
	ImagePath string
	Province  string
	Club      string
	CreatedAt pgtype.Timestamp
	UpdatedAt pgtype.Timestamp
}

type Stage0Result struct {
	ID          pgtype.UUID
	ResultID    pgtype.UUID
	Status      Stage0Status
	Series1     string
	Series2     string
	Series3     string
	Series4     string
	Series5     string
	Checkmarks  string
	ShooterSign pgtype.Text
	ScorerSign  pgtype.Text
	CreatedAt   pgtype.Timestamp
	UpdatedAt   pgtype.Timestamp
}

type Stage13Try struct {
	ID         pgtype.UUID
	Status     Stage13Status
	No1        string
	No2        string
	No3        string
	No4        string
	No5        string
	No6        string
	Checkmarks string
}

type Stage1Result struct {
	ID          pgtype.UUID
	ResultID    pgtype.UUID
	Try1ID      pgtype.UUID
	Try2ID      pgtype.UUID
	IsTry2      bool
	ShooterSign pgtype.Text
	ScorerSign  pgtype.Text
	CreatedAt   pgtype.Timestamp
	UpdatedAt   pgtype.Timestamp
}

type Stage2Result struct {
	ID          pgtype.UUID
	ResultID    pgtype.UUID
	Try1ID      pgtype.UUID
	Try2ID      pgtype.UUID
	IsTry2      bool
	ShooterSign pgtype.Text
	ScorerSign  pgtype.Text
	CreatedAt   pgtype.Timestamp
	UpdatedAt   pgtype.Timestamp
}

type Stage2Try struct {
	ID         pgtype.UUID
	Status     Stage246Status
	No1        string
	No2        string
	No3        string
	Checkmarks string
}

type Stage3Result struct {
	ID          pgtype.UUID
	ResultID    pgtype.UUID
	Try1ID      pgtype.UUID
	Try2ID      pgtype.UUID
	IsTry2      bool
	ShooterSign pgtype.Text
	ScorerSign  pgtype.Text
	CreatedAt   pgtype.Timestamp
	UpdatedAt   pgtype.Timestamp
}

type Stage46Try struct {
	ID         pgtype.UUID
	Status     Stage246Status
	No1        string
	No2        string
	No3        string
	Checkmarks string
}

type Stage4Result struct {
	ID          pgtype.UUID
	ResultID    pgtype.UUID
	Try1ID      pgtype.UUID
	Try2ID      pgtype.UUID
	IsTry2      bool
	ShooterSign pgtype.Text
	ScorerSign  pgtype.Text
	CreatedAt   pgtype.Timestamp
	UpdatedAt   pgtype.Timestamp
}

type Stage5Result struct {
	ID          pgtype.UUID
	ResultID    pgtype.UUID
	Try1ID      pgtype.UUID
	Try2ID      pgtype.UUID
	IsTry2      bool
	ShooterSign pgtype.Text
	ScorerSign  pgtype.Text
	CreatedAt   pgtype.Timestamp
	UpdatedAt   pgtype.Timestamp
}

type Stage5Try struct {
	ID         pgtype.UUID
	Status     Stage5Status
	No1        string
	No2        string
	Checkmarks string
}

type Stage6Result struct {
	ID          pgtype.UUID
	ResultID    pgtype.UUID
	Try1ID      pgtype.UUID
	Try2ID      pgtype.UUID
	IsTry2      bool
	ShooterSign pgtype.Text
	ScorerSign  pgtype.Text
	CreatedAt   pgtype.Timestamp
	UpdatedAt   pgtype.Timestamp
}

type Super struct {
	ID     pgtype.UUID
	UserID pgtype.UUID
}

type User struct {
	ID        pgtype.UUID
	Username  string
	Password  string
	Name      string
	CreatedAt pgtype.Timestamp
	UpdatedAt pgtype.Timestamp
}
