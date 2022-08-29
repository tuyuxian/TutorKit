// Code generated by ent, DO NOT EDIT.

package ent

import (
	"backend/ent/entcourse"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// EntCourse is the model entity for the EntCourse schema.
type EntCourse struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// CourseUrl holds the value of the "courseUrl" field.
	CourseUrl string `json:"courseUrl,omitempty"`
	// PaymentMethod holds the value of the "paymentMethod" field.
	PaymentMethod entcourse.PaymentMethod `json:"paymentMethod,omitempty"`
	// PaymentAmount holds the value of the "paymentAmount" field.
	PaymentAmount float64 `json:"paymentAmount,omitempty"`
	// StartDate holds the value of the "startDate" field.
	StartDate time.Time `json:"startDate,omitempty"`
	// EndDate holds the value of the "endDate" field.
	EndDate time.Time `json:"endDate,omitempty"`
	// Monday holds the value of the "monday" field.
	Monday bool `json:"monday,omitempty"`
	// Tuesday holds the value of the "tuesday" field.
	Tuesday bool `json:"tuesday,omitempty"`
	// Wednesday holds the value of the "wednesday" field.
	Wednesday bool `json:"wednesday,omitempty"`
	// Thursday holds the value of the "thursday" field.
	Thursday bool `json:"thursday,omitempty"`
	// Friday holds the value of the "friday" field.
	Friday bool `json:"friday,omitempty"`
	// Saturday holds the value of the "saturday" field.
	Saturday bool `json:"saturday,omitempty"`
	// Sunday holds the value of the "sunday" field.
	Sunday bool `json:"sunday,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the EntCourseQuery when eager-loading is set.
	Edges EntCourseEdges `json:"edges"`
}

// EntCourseEdges holds the relations/edges for other nodes in the graph.
type EntCourseEdges struct {
	// CourseOwner holds the value of the courseOwner edge.
	CourseOwner []*EntUser `json:"courseOwner,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// CourseOwnerOrErr returns the CourseOwner value or an error if the edge
// was not loaded in eager-loading.
func (e EntCourseEdges) CourseOwnerOrErr() ([]*EntUser, error) {
	if e.loadedTypes[0] {
		return e.CourseOwner, nil
	}
	return nil, &NotLoadedError{edge: "courseOwner"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*EntCourse) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case entcourse.FieldMonday, entcourse.FieldTuesday, entcourse.FieldWednesday, entcourse.FieldThursday, entcourse.FieldFriday, entcourse.FieldSaturday, entcourse.FieldSunday:
			values[i] = new(sql.NullBool)
		case entcourse.FieldPaymentAmount:
			values[i] = new(sql.NullFloat64)
		case entcourse.FieldID:
			values[i] = new(sql.NullInt64)
		case entcourse.FieldName, entcourse.FieldCourseUrl, entcourse.FieldPaymentMethod:
			values[i] = new(sql.NullString)
		case entcourse.FieldStartDate, entcourse.FieldEndDate:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type EntCourse", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the EntCourse fields.
func (ec *EntCourse) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case entcourse.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ec.ID = int(value.Int64)
		case entcourse.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				ec.Name = value.String
			}
		case entcourse.FieldCourseUrl:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field courseUrl", values[i])
			} else if value.Valid {
				ec.CourseUrl = value.String
			}
		case entcourse.FieldPaymentMethod:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field paymentMethod", values[i])
			} else if value.Valid {
				ec.PaymentMethod = entcourse.PaymentMethod(value.String)
			}
		case entcourse.FieldPaymentAmount:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field paymentAmount", values[i])
			} else if value.Valid {
				ec.PaymentAmount = value.Float64
			}
		case entcourse.FieldStartDate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field startDate", values[i])
			} else if value.Valid {
				ec.StartDate = value.Time
			}
		case entcourse.FieldEndDate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field endDate", values[i])
			} else if value.Valid {
				ec.EndDate = value.Time
			}
		case entcourse.FieldMonday:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field monday", values[i])
			} else if value.Valid {
				ec.Monday = value.Bool
			}
		case entcourse.FieldTuesday:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field tuesday", values[i])
			} else if value.Valid {
				ec.Tuesday = value.Bool
			}
		case entcourse.FieldWednesday:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field wednesday", values[i])
			} else if value.Valid {
				ec.Wednesday = value.Bool
			}
		case entcourse.FieldThursday:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field thursday", values[i])
			} else if value.Valid {
				ec.Thursday = value.Bool
			}
		case entcourse.FieldFriday:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field friday", values[i])
			} else if value.Valid {
				ec.Friday = value.Bool
			}
		case entcourse.FieldSaturday:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field saturday", values[i])
			} else if value.Valid {
				ec.Saturday = value.Bool
			}
		case entcourse.FieldSunday:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field sunday", values[i])
			} else if value.Valid {
				ec.Sunday = value.Bool
			}
		}
	}
	return nil
}

// QueryCourseOwner queries the "courseOwner" edge of the EntCourse entity.
func (ec *EntCourse) QueryCourseOwner() *EntUserQuery {
	return (&EntCourseClient{config: ec.config}).QueryCourseOwner(ec)
}

// Update returns a builder for updating this EntCourse.
// Note that you need to call EntCourse.Unwrap() before calling this method if this EntCourse
// was returned from a transaction, and the transaction was committed or rolled back.
func (ec *EntCourse) Update() *EntCourseUpdateOne {
	return (&EntCourseClient{config: ec.config}).UpdateOne(ec)
}

// Unwrap unwraps the EntCourse entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ec *EntCourse) Unwrap() *EntCourse {
	_tx, ok := ec.config.driver.(*txDriver)
	if !ok {
		panic("ent: EntCourse is not a transactional entity")
	}
	ec.config.driver = _tx.drv
	return ec
}

// String implements the fmt.Stringer.
func (ec *EntCourse) String() string {
	var builder strings.Builder
	builder.WriteString("EntCourse(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ec.ID))
	builder.WriteString("name=")
	builder.WriteString(ec.Name)
	builder.WriteString(", ")
	builder.WriteString("courseUrl=")
	builder.WriteString(ec.CourseUrl)
	builder.WriteString(", ")
	builder.WriteString("paymentMethod=")
	builder.WriteString(fmt.Sprintf("%v", ec.PaymentMethod))
	builder.WriteString(", ")
	builder.WriteString("paymentAmount=")
	builder.WriteString(fmt.Sprintf("%v", ec.PaymentAmount))
	builder.WriteString(", ")
	builder.WriteString("startDate=")
	builder.WriteString(ec.StartDate.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("endDate=")
	builder.WriteString(ec.EndDate.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("monday=")
	builder.WriteString(fmt.Sprintf("%v", ec.Monday))
	builder.WriteString(", ")
	builder.WriteString("tuesday=")
	builder.WriteString(fmt.Sprintf("%v", ec.Tuesday))
	builder.WriteString(", ")
	builder.WriteString("wednesday=")
	builder.WriteString(fmt.Sprintf("%v", ec.Wednesday))
	builder.WriteString(", ")
	builder.WriteString("thursday=")
	builder.WriteString(fmt.Sprintf("%v", ec.Thursday))
	builder.WriteString(", ")
	builder.WriteString("friday=")
	builder.WriteString(fmt.Sprintf("%v", ec.Friday))
	builder.WriteString(", ")
	builder.WriteString("saturday=")
	builder.WriteString(fmt.Sprintf("%v", ec.Saturday))
	builder.WriteString(", ")
	builder.WriteString("sunday=")
	builder.WriteString(fmt.Sprintf("%v", ec.Sunday))
	builder.WriteByte(')')
	return builder.String()
}

// EntCourses is a parsable slice of EntCourse.
type EntCourses []*EntCourse

func (ec EntCourses) config(cfg config) {
	for _i := range ec {
		ec[_i].config = cfg
	}
}