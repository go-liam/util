package user

import "github.com/go-liam/util/response"

type Service interface {
	CheckAndCreate(item *Model) (int64, error)
	Create(item *Model) (int64, error)
	FindOne(id int64) (*Model, error)
	FindMulti(p *response.Pagination, s *response.ListParameter) ([]*Model, error)
	Update(item *Model) (int64, error)
	UpdateStatusByIDs(status int, ls []int64) (int64, error)
	//cache
	CacheOne(id int64) (*Model, error)
}