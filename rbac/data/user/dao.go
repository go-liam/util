package user

import (
	"errors"
	"fmt"

	"github.com/go-liam/util/conv"
	"github.com/go-liam/util/response"
	"gorm.io/gorm"

	//models "grape/internal/pkg/model"
	"time"
)

var Server = new(SrvUser)

type SrvUser struct {
	DB *gorm.DB
}

func (e *SrvUser) UpdateOrCreate(item *Model) (int64, error) {
	ck, _ := e.FindOne(item.ID)
	if ck.ID > 0 {
		ck.RoleIDs = item.RoleIDs
		return e.Update(ck)
	}
	return e.Create(item)
}

func (e *SrvUser) CheckAndCreate(item *Model) (int64, error) {
	ck, _ := e.FindOne(item.ID)
	if ck.ID > 0 {
		return 0, errors.New("had data")
	}
	return e.Create(item)
}

func (e *SrvUser) Create(item *Model) (int64, error) {
	item.CreatedAt = time.Now().Unix()
	item.UpdatedAt = item.CreatedAt
	v := e.DB.Create(item)
	return v.RowsAffected, v.Error
}

func (e *SrvUser) FindOne(id int64) (*Model, error) {
	item := new(Model)
	sql := "select * from rb_user where id =?  and status < 44 limit 1 "
	v := e.DB.Raw(sql, id).Scan(item)
	return item, v.Error
}

// whereSt: " and id>1 " ,orderSt =" order by ID "
func (e *SrvUser) FindMulti(p *response.Pagination, s *response.ListParameter) ([]*Model, error) {
	var result []*Model
	sqlLimit := fmt.Sprintf(" limit %d , %d  ", (p.Current-1)*p.PageSize, p.PageSize)
	sqlWhere := " status < 44 " + s.WhereSt
	sql := "select * from rb_user where " + sqlWhere + s.OrderSt + sqlLimit
	var Total int64 = 0
	e.DB.Model(&Model{}).Where(sqlWhere).Count(&Total)
	p.Total = int(Total)
	v := e.DB.Raw(sql).Scan(&result)
	return result, v.Error
}

func (e *SrvUser) Update(item *Model) (int64, error) {
	item.UpdatedAt = time.Now().Unix()
	sql := "update rb_user set extended=?,updated_at=?,flag=?,remark=?,role_ids= ? where `id` = ? "
	v := e.DB.Exec(sql, item.Extended, item.UpdatedAt, item.Flag, item.Remark, item.RoleIDs,
		item.ID)
	return v.RowsAffected, v.Error
}

func (e *SrvUser) UpdateStatus(item *Model) (int64, error) {
	item.UpdatedAt = time.Now().Unix()
	sql := "update rb_user set `status` = ?,updated_at=? where `id` = ? "
	v := e.DB.Exec(sql, item.Status, item.UpdatedAt, item.ID)
	return v.RowsAffected, v.Error
}

func (e *SrvUser) UpdateStatusByIDs(status int, ls []int64) (int64, error) {
	updatedAt := time.Now().Unix()
	ids := conv.ArrayToString(ls, ",")
	sql := fmt.Sprintf("update rb_user set `status` = ?,updated_at=? where `id` in (%s) ", ids)
	v := e.DB.Exec(sql, status, updatedAt)
	return v.RowsAffected, v.Error
}

func (e *SrvUser) FindMultiByIDs(ids []int64) ([]*Model, error) {
	result := make([]*Model, 0)
	if len(ids) == 0 {
		return result, nil
	}
	st := conv.ArrayToString(ids, ",")
	sql := fmt.Sprintf("select * from rb_user where id in (%v)  and status < 44  ", st)
	v := e.DB.Raw(sql).Scan(&result)
	return result, v.Error
}

func (e *SrvUser) FindMoreByList(list []*Model, ids []int64) []*Model {
	ls := make([]*Model, 0)
	for _, v := range list {
		if conv.ArrayInt64Contains(ids, v.ID) {
			ls = append(ls, v)
		}
	}
	return ls
}
