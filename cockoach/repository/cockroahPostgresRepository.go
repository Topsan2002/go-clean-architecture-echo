package repository

import (
	"github.com/labstack/gommon/log"

	"go-clean-ex/cockoach/entity"
	"go-clean-ex/database"
)

type cockoachPosgresRepository struct {
	db database.Database
}

func NewCockroachPostgresRepository(db database.Database) CockroachRepository {
	return &cockoachPosgresRepository{db: db}
}

// InsertCockoachData implements CockroachRepository.
func (c *cockoachPosgresRepository) InsertCockoachData(in *entity.InsertCockroachDto) error {
	data := &entity.Cockroach{
		Amount: in.Amount,
	}

	result := c.db.GetDb().Create(data)
	if result.Error != nil {
		// log.Errorf()
		log.Errorf("InsertCockroachData %v", result.Error)
		return result.Error
	}

	log.Debugf("InsertCockroachData %v", result.RowsAffected)
	return nil
}
