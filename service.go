package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"

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
		log.Println(err)
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
		log.Println(err)
		return err
	}
	return nil
}

func (s *Service) changeAirlineProviders(ctx context.Context, dto AirlineDTOChangeProviders) error {
	tx := s.db.MustBeginTx(ctx, nil)
	err := func() error {
		var airlineM AirlineModel
		if err := tx.GetContext(ctx, &airlineM, "SELECT * FROM airline WHERE code = ?", dto.Code); err != nil {
			if err == sql.ErrNoRows {
				return fmt.Errorf("Airline with passed code not exist")
			}
			return err
		}

		if _, err := tx.ExecContext(ctx, "DELETE FROM airline_provider WHERE airline_id = ?", airlineM.Id); err != nil {
			return err
		}

		query, args, err := sqlx.In("SELECT * FROM provider WHERE code IN (?)", dto.Providers)
		if err != nil {
			return err
		}
		providerMs := []ProviderModel{}
		if err := tx.SelectContext(ctx, &providerMs, query, args...); err != nil {
			return err
		}

		for _, providerM := range providerMs {
			tx.ExecContext(ctx, "INSERT INTO airline_provider (airline_id, provider_id) VALUES (?, ?)", airlineM.Id, providerM.Id)
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
	if _, err := s.db.NamedExecContext(ctx, "INSERT INTO provider (code, name) VALUES (:id, :name)", &dto); err != nil {
		log.Println(err)
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
		log.Println(err)
		return err
	}
	return nil
}

func (s *Service) getProviderAirlines(ctx context.Context) error {
	return nil
}

func (s *Service) addSchema(ctx context.Context) error {
	return nil
}

func (s *Service) getSchemeByName(ctx context.Context) error {
	return nil
}

func (s *Service) updateSchemeById(ctx context.Context) error {
	return nil
}

func (s *Service) deleteSchemeById(ctx context.Context) error {
	return nil
}

func (s *Service) addAccount(ctx context.Context) error {
	return nil
}

func (s *Service) setAccountScheme(ctx context.Context) error {
	return nil
}

func (s *Service) deleteAccountById(ctx context.Context) error {
	return nil
}

func (s *Service) getAccountAirlines(ctx context.Context) error {
	return nil
}
