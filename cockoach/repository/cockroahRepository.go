package repository

import "go-clean-ex/cockoach/entity"

type CockroachRepository interface {
	InsertCockoachData(in *entity.InsertCockroachDto) error
}
