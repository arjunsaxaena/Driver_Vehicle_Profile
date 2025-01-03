package controllers

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/huandu/go-sqlbuilder"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	_ "github.com/lib/pq"

	"github.com/arjunsaxaena/driver_vehicle_profile/model"
)

type DBVehicleStore struct {
	db *sqlx.DB
}

func NewDBVehicleStore(db *sqlx.DB) *DBVehicleStore {
	return &DBVehicleStore{db: db}
}

func (s *DBVehicleStore) CreateVehicle(v *model.Vehicle) error {
	if v.ID == uuid.Nil {
		v.ID = uuid.New()
	}

	if v.DriverHelperID != uuid.Nil {
		var driverHelperExists bool
		err := s.db.Get(&driverHelperExists, "SELECT EXISTS (SELECT 1 FROM driver_helpers WHERE id = $1)", v.DriverHelperID)
		if err != nil {
			return fmt.Errorf("failed to check if driver helper exists: %w", err)
		}
		if !driverHelperExists {
			return fmt.Errorf("driver helper with ID %s does not exist", v.DriverHelperID)
		}
	}

	if v.SeatsAvailable < 0 || v.SeatsAvailable > v.TotalStudentsCapacity {
		return fmt.Errorf("invalid seats_available: %d; must be between 0 and total_students_capacity", v.SeatsAvailable)
	}

	sb := sqlbuilder.NewInsertBuilder()
	sb.SetFlavor(sqlbuilder.PostgreSQL)
	sb.InsertInto("vehicles").
		Cols("id", "vehicle_number", "route_number", "total_students_capacity", "seats_available",
			"driver_helper_id", "insurance_number", "insurance_expiry_date", "pollution_certificate_number",
			"pollution_certificate_expiry_date", "fitness_certificate_number", "fitness_certificate_expiry_date",
			"vehicle_document_path")

	if v.DriverHelperID != uuid.Nil {
		sb.Values(v.ID, v.VehicleNumber, v.RouteNumber, v.TotalStudentsCapacity, v.SeatsAvailable,
			v.DriverHelperID, v.InsuranceNumber, v.InsuranceExpiryDate, v.PollutionCertificateNumber,
			v.PollutionCertificateExpiryDate, v.FitnessCertificateNumber, v.FitnessCertificateExpiryDate,
			v.VehicleDocumentPath)
	} else {
		sb.Values(v.ID, v.VehicleNumber, v.RouteNumber, v.TotalStudentsCapacity, v.SeatsAvailable,
			nil, v.InsuranceNumber, v.InsuranceExpiryDate, v.PollutionCertificateNumber,
			v.PollutionCertificateExpiryDate, v.FitnessCertificateNumber, v.FitnessCertificateExpiryDate,
			v.VehicleDocumentPath)
	}

	query, args := sb.Build()
	_, err := s.db.Exec(query, args...)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "unique_violation":
				return fmt.Errorf("vehicle with number %s already exists", v.VehicleNumber)
			default:
				return fmt.Errorf("failed to insert vehicle: %w", err)
			}
		}
		return fmt.Errorf("failed to insert vehicle: %w", err)
	}

	return nil
}

func (s *DBVehicleStore) UpdateVehicle(v *model.Vehicle) error {
	var driverHelperExists bool
	err := s.db.Get(&driverHelperExists, "SELECT EXISTS (SELECT 1 FROM driver_helpers WHERE id = $1)", v.DriverHelperID)
	if err != nil {
		return fmt.Errorf("failed to check if driver helper exists: %w", err)
	}
	if !driverHelperExists {
		return fmt.Errorf("driver helper with ID %s does not exist", v.DriverHelperID)
	}

	sb := sqlbuilder.NewUpdateBuilder()
	sb.SetFlavor(sqlbuilder.PostgreSQL)
	sb.Update("vehicles").Set(
		sb.Assign("vehicle_number", v.VehicleNumber),
		sb.Assign("route_number", v.RouteNumber),
		sb.Assign("total_students_capacity", v.TotalStudentsCapacity),
		sb.Assign("seats_available", v.SeatsAvailable),
		sb.Assign("driver_helper_id", v.DriverHelperID),
		sb.Assign("insurance_number", v.InsuranceNumber),
		sb.Assign("insurance_expiry_date", v.InsuranceExpiryDate),
		sb.Assign("pollution_certificate_number", v.PollutionCertificateNumber),
		sb.Assign("pollution_certificate_expiry_date", v.PollutionCertificateExpiryDate),
		sb.Assign("fitness_certificate_number", v.FitnessCertificateNumber),
		sb.Assign("fitness_certificate_expiry_date", v.FitnessCertificateExpiryDate),
		sb.Assign("vehicle_document_path", v.VehicleDocumentPath),
		sb.Assign("updated_at", time.Now()),
	).Where(sb.Equal("id", v.ID))

	query, args := sb.Build()
	_, err = s.db.Exec(query, args...)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "23514" {
			return fmt.Errorf("failed to update vehicle: constraint violation. Please check the input values (e.g., seats_available > total_students_capacity)")
		}
		return fmt.Errorf("failed to update vehicle: %w", err)
	}

	return nil
}

func (s *DBVehicleStore) DeleteVehicle(id uuid.UUID) error {
	sb := sqlbuilder.NewDeleteBuilder()
	sb.SetFlavor(sqlbuilder.PostgreSQL)
	sb.DeleteFrom("vehicles").Where(sb.Equal("id", id))

	query, args := sb.Build()
	_, err := s.db.Exec(query, args...)
	return err
}

func (s *DBVehicleStore) Vehicles() ([]model.Vehicle, error) {
	var vehicles []model.Vehicle
	sb := sqlbuilder.NewSelectBuilder()
	sb.SetFlavor(sqlbuilder.PostgreSQL)
	sb.Select("*").From("vehicles")

	query, args := sb.Build()
	err := s.db.Select(&vehicles, query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch vehicles: %w", err)
	}

	return vehicles, nil
}

func (s *DBVehicleStore) VehicleByID(id uuid.UUID) (model.Vehicle, error) {
	var v model.Vehicle
	sb := sqlbuilder.NewSelectBuilder()
	sb.SetFlavor(sqlbuilder.PostgreSQL)
	sb.Select("*").From("vehicles").Where(sb.Equal("id", id))

	query, args := sb.Build()
	err := s.db.Get(&v, query, args...)
	return v, err
}

func (s *DBVehicleStore) VehiclesByDriverHelperID(driverHelperID uuid.UUID) ([]model.Vehicle, error) {
	var vehicles []model.Vehicle
	sb := sqlbuilder.NewSelectBuilder()
	sb.SetFlavor(sqlbuilder.PostgreSQL)
	sb.Select("*").From("vehicles").Where(sb.Equal("driver_helper_id", driverHelperID))

	query, args := sb.Build()
	err := s.db.Select(&vehicles, query, args...)
	return vehicles, err
}

func (s *DBVehicleStore) VehiclesByRouteNumber(routeNumber string) ([]model.Vehicle, error) {
	var vehicles []model.Vehicle
	sb := sqlbuilder.NewSelectBuilder()
	sb.SetFlavor(sqlbuilder.PostgreSQL)
	sb.Select("*").From("vehicles").Where(sb.Equal("route_number", routeNumber))

	query, args := sb.Build()
	err := s.db.Select(&vehicles, query, args...)
	return vehicles, err
}

func (s *DBVehicleStore) ExpiredCertificatesVehicles() ([]model.Vehicle, error) {
	var vehicles []model.Vehicle
	sb := sqlbuilder.NewSelectBuilder()
	sb.SetFlavor(sqlbuilder.PostgreSQL)
	sb.Select("*").From("vehicles").Where(
		sb.LessThan("insurance_expiry_date", time.Now()),
		sb.Or(
			sb.LessThan("pollution_certificate_expiry_date", time.Now()),
			sb.LessThan("fitness_certificate_expiry_date", time.Now()),
		),
	)

	query, args := sb.Build()
	err := s.db.Select(&vehicles, query, args...)
	return vehicles, err
}
