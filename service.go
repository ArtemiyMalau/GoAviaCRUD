package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Service struct {
	db *sqlx.DB
}

func NewService(db *sqlx.DB) *Service {
	return &Service{
		db: db,
	}
}

func (s *Service) addAirline(ctx context.Context, dto AirlineDTOAdd) error {
	if _, err := s.db.NamedExecContext(ctx, "INSERT INTO airline (code, name) VALUES (:code, :name)", &dto); err != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
			return fmt.Errorf("Airline with passed code already exists")
		}
		return err
	}
	return nil
}

func (s *Service) deleteAirlineByCode(ctx context.Context, code string) error {
	if _, err := s.db.ExecContext(ctx, "DELETE FROM airline WHERE code = ?", code); err != nil {
		return err
	}
	return nil
}

func (s *Service) changeAirlineProviders(ctx context.Context, dto AirlineDTOChangeProviders) error {
	tx := s.db.MustBeginTx(ctx, nil)
	err := func() error {
		var airlineM AirlineModel
		if err := tx.GetContext(ctx, &airlineM, "SELECT id, code, name FROM airline WHERE code = ?", dto.Code); err != nil {
			if err == sql.ErrNoRows {
				return fmt.Errorf("Airline with passed code not exist")
			}
			return err
		}

		// Replacing airline's providers by passed providers
		if _, err := tx.ExecContext(ctx, "DELETE FROM airline_provider WHERE airline_id = ?", airlineM.Id); err != nil {
			return err
		}

		query, args, _ := sqlx.In("SELECT id FROM provider WHERE code IN (?)", dto.Providers)
		var providerIds []int
		if err := tx.SelectContext(ctx, &providerIds, query, args...); err != nil {
			return err
		}

		for _, providerId := range providerIds {
			if _, err := tx.ExecContext(ctx, "INSERT INTO airline_provider (airline_id, provider_id) VALUES (?, ?)", airlineM.Id, providerId); err != nil {
				return nil
			}
		}

		return nil
	}()
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

func (s *Service) addProvider(ctx context.Context, dto ProviderDTOAdd) error {
	if _, err := s.db.NamedExecContext(ctx, "INSERT INTO provider (code, name) VALUES (:code, :name)", &dto); err != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
			return fmt.Errorf("Provider with passed code already exists")
		}
		return err
	}
	return nil
}

func (s *Service) deleteProviderByCode(ctx context.Context, code string) error {
	if _, err := s.db.ExecContext(ctx, "DELETE FROM provider WHERE code = ?", code); err != nil {
		return err
	}
	return nil
}

func (s *Service) getProviderAirlinesByCode(ctx context.Context, code string) (airlines []Airline, err error) {
	airlines = []Airline{}
	if err = s.db.SelectContext(ctx, &airlines,
		`SELECT airline.code, airline.name
		FROM provider
		JOIN airline_provider ON provider.id = airline_provider.provider_id
		JOIN airline ON airline_provider.airline_id = airline.id
		WHERE provider.code = ?`,
		code); err != nil {
		return
	}
	return
}

func (s *Service) addSchema(ctx context.Context, dto SchemaDTOAdd) error {
	tx := s.db.MustBeginTx(ctx, nil)
	err := func() error {
		// Inserting scheme
		resp, err := tx.ExecContext(ctx, "INSERT INTO scheme (name) VALUES (?)", dto.Name)
		if err != nil {
			return err
		}
		schemeId, _ := resp.LastInsertId()

		// Inserting scheme's providers
		query, args, _ := sqlx.In("SELECT id FROM provider WHERE code IN (?)", dto.Providers)
		var providerIds []int
		if err := tx.SelectContext(ctx, &providerIds, query, args...); err != nil {
			return err
		}

		for _, providerId := range providerIds {
			if _, err := tx.ExecContext(ctx, "INSERT INTO scheme_provider (scheme_id, provider_id) VALUES (?, ?)", schemeId, providerId); err != nil {
				return err
			}
		}

		return nil
	}()
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

func (s *Service) getSchemeByName(ctx context.Context, name string) (schema Schema, err error) {
	schema.Providers = []providerCode{}

	tx := s.db.MustBeginTx(ctx, nil)
	err = func() error {
		// Getting scheme table data
		if err := tx.GetContext(ctx, &schema, "SELECT scheme.id, scheme.name FROM scheme WHERE scheme.name = ?", name); err != nil {
			return err
		}

		// Getting scheme's providers
		if err := tx.SelectContext(ctx, &schema.Providers,
			`SELECT provider.code
			FROM scheme_provider
			JOIN provider ON scheme_provider.provider_id = provider.id
			WHERE scheme_provider.scheme_id = ?`, schema.Id); err != nil {
			return err
		}

		return nil
	}()
	if err != nil {
		return
	}
	tx.Commit()

	return
}

func (s *Service) updateSchemeById(ctx context.Context, dto SchemaDTOUpdate) error {
	tx := s.db.MustBeginTx(ctx, nil)
	err := func() error {
		if err := tx.QueryRowContext(ctx, "SELECT id FROM scheme WHERE id = ?", dto.Id).Scan(&dto.Id); err != nil {
			if err == sql.ErrNoRows {
				return fmt.Errorf("scheme with passed id not exist")
			}
			return err

		}

		if dto.Name != "" {
			if _, err := tx.ExecContext(ctx, "UPDATE scheme SET name = ? WHERE id = ?", dto.Name, dto.Id); err != nil {
				return err
			}
		}

		if len(dto.Providers) > 0 {
			// Replace scheme's providers by passed providers
			if _, err := tx.ExecContext(ctx, "DELETE FROM scheme_provider WHERE scheme_id = ?", dto.Id); err != nil {
				return err
			}

			query, args, _ := sqlx.In("SELECT id FROM provider WHERE code IN (?)", dto.Providers)
			var providerIds []int
			if err := tx.SelectContext(ctx, &providerIds, query, args...); err != nil {
				return err
			}

			for _, providerId := range providerIds {
				if _, err := tx.ExecContext(ctx, "INSERT INTO scheme_provider (scheme_id, provider_id) VALUES (?, ?)", dto.Id, providerId); err != nil {
					return err
				}
			}
		}

		return nil
	}()
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

func (s *Service) deleteSchemeById(ctx context.Context, id int) error {
	// Сhecking whether the scheme is assigned to any account perform on db level
	if _, err := s.db.ExecContext(ctx, "DELETE FROM scheme WHERE id = ?", id); err != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) && mysqlErr.Number == 1451 {
			return fmt.Errorf("cannot delete scheme assigned to any account")
		}
		return err
	}
	return nil
}

func (s *Service) addAccount(ctx context.Context, dto AccountDTOAdd) error {
	// Сhecking whether the passed schema actually exists perform on db level
	if _, err := s.db.ExecContext(ctx, "INSERT INTO account (scheme_id) VALUES (?)", dto.SchemaId); err != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) && mysqlErr.Number == 1452 {
			return fmt.Errorf("passed scheme not exists")
		}
		return err
	}
	return nil
}

func (s *Service) setAccountScheme(ctx context.Context, dto AccountDTOSetScheme) error {
	if _, err := s.db.NamedExecContext(ctx, "UPDATE account SET scheme_id = :scheme_id WHERE id = :id", &dto); err != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) && mysqlErr.Number == 1452 {
			return fmt.Errorf("passed scheme not exists")
		}
		return err
	}
	return nil
}

func (s *Service) deleteAccountById(ctx context.Context, id int) error {
	if _, err := s.db.ExecContext(ctx, "DELETE FROM account WHERE id = ?", id); err != nil {
		return err
	}
	return nil
}

func (s *Service) getAccountAirlinesById(ctx context.Context, id int) (airlines []Airline, err error) {
	airlines = []Airline{}
	if err = s.db.SelectContext(ctx, &airlines,
		`SELECT DISTINCT airline.code, airline.name
		FROM account
		JOIN scheme ON account.scheme_id = scheme.id
		JOIN scheme_provider ON scheme.id = scheme_provider.scheme_id
		JOIN provider ON scheme_provider.provider_id = provider.id
		JOIN airline_provider ON provider.id = airline_provider.provider_id
		JOIN airline ON airline_provider.airline_id = airline.id
		WHERE account.id = ?`,
		id); err != nil {
		return
	}
	return
}
