package main

import (
	"context"
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
	resp, err := s.db.NamedExecContext(ctx, "INSERT INTO airline (code, name) VALUES (:code, :name)", &dto)
	if err != nil {
		var mysqlErr *mysql.MySQLError
		log.Println(err)
		if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
			return fmt.Errorf("Airline with passed code already exists")
		}
		return err
	}
	log.Println(resp.LastInsertId())
	return nil

}

func (s *Service) deleteAirlineByCode(ctx context.Context, code string) error {
	if _, err := s.db.ExecContext(ctx, "DELETE FROM airline WHERE code = ?", code); err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (s *Service) changeAirlineProviders(ctx context.Context) error {
	return nil
}

func (s *Service) addProvider(ctx context.Context) error {
	return nil
}

func (s *Service) deleteProviderById(ctx context.Context) error {
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

// tx := s.db.MustBegin()
// resp, err := tx.NamedExec("INSERT INTO airline (code, name) VALUES (:code, :name)", &dto)
// if err != nil {
// 	var mysqlErr *mysql.MySQLError
// 	log.Println(err)
// 	if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
// 		return fmt.Errorf("Airline with passed code already exists")
// 	}
// 	return err
// }
// tx.Commit()
