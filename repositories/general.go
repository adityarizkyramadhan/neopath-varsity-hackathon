package repositories

import (
	"fmt"

	"gorm.io/gorm"
)

type GeneralRepositoryImpl struct {
	DB *gorm.DB
}

func NewGeneralRepositoryImpl(db *gorm.DB) *GeneralRepositoryImpl {
	return &GeneralRepositoryImpl{
		DB: db,
	}
}

func (g *GeneralRepositoryImpl) Create(data interface{}) error {
	return g.DB.Transaction(func(tx *gorm.DB) error {
		return tx.Create(data).Error
	})
}

func (g *GeneralRepositoryImpl) FindById(id int, data interface{}) error {
	return g.DB.First(data, id).Error
}

func (g *GeneralRepositoryImpl) FindAll(data interface{}) error {
	return g.DB.Find(data).Error
}

func (g *GeneralRepositoryImpl) Delete(id int, data interface{}) error {
	return g.DB.Model(data).Delete(" ID = ?", id).Error
}

func (g *GeneralRepositoryImpl) Update(id int, data interface{}) error {
	return g.DB.Where("id = ?", id).Save(data).Error
}

func (g *GeneralRepositoryImpl) FindByColumn(column, value string, data interface{}) error {
	query := fmt.Sprintf("%s = ?", column)
	return g.DB.Where(query, value).Find(data).Error
}
