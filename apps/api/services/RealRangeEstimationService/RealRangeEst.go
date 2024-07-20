package RealRangeEstimationService

import "gorm.io/gorm"

type RealRangeEstimationService struct {
	db *gorm.DB
}

func New(db *gorm.DB) *RealRangeEstimationService {
	return &RealRangeEstimationService{
		db: db,
	}
}
