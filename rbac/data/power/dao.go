package power

import (
	"fmt"
	"time"

	"github.com/go-liam/util/conv"
	"github.com/go-liam/util/response"
	"gorm.io/gorm"
)

var Server = new(SrvPower)

type SrvPower struct {
	DB *gorm.DB
}

func (e *SrvPower) Create(item *Model) (int64, error) {
	item.CreatedAt = time.Now().Unix()
	item.UpdatedAt = item.CreatedAt
	v := e.DB.Create(item)
	return v.RowsAffected, v.Error
}

func (e *SrvPower) FindOne(id int64) (*Model, error) {
	item := new(Model)
	sql := "select * from rb_power where id =?  and status < 44 limit 1 "
	v := e.DB.Raw(sql, id).Scan(item)
	return item, v.Error
}

func (e *SrvPower) FindMultiByNil() ([]*Model, error) {
	var result []*Model
	sql := "select * from rb_power where status < 44 "
	v := e.DB.Raw(sql).Scan(&result)
	return result, v.Error
}

func (e *SrvPower) FindMulti(p *response.Pagination, s *response.ListParameter) ([]*Model, error) {
	var result []*Model
	sqlLimit := fmt.Sprintf(" limit %d , %d  ", (p.Current-1)*p.PageSize, p.PageSize)
	sqlWhere := " status < 44 " + s.WhereSt
	sql := "select * from rb_power where " + sqlWhere + s.OrderSt + sqlLimit
	var Total int64 = 0
	e.DB.Model(&Model{}).Where(sqlWhere).Count(&Total)
	p.Total = int(Total)
	v := e.DB.Raw(sql).Scan(&result)
	return result, v.Error
}

func (e *SrvPower) Update(item *Model) (int64, error) {
	item.UpdatedAt = time.Now().Unix()
	sql := "update rb_power set `name` = ?,extended=?,updated_at=?,`tag`=?,`type`=?,`status`=? where `id` = ? "
	v := e.DB.Exec(sql, item.Name, item.Extended, item.UpdatedAt, item.Tag, item.Type, item.Status, item.ID)
	return v.RowsAffected, v.Error
}

func (e *SrvPower) UpdateStatus(item *Model) (int64, error) {
	item.UpdatedAt = time.Now().Unix()
	sql := "update rb_power set `status` = ?,updated_at=? where `id` = ? "
	v := e.DB.Exec(sql, item.Status, item.UpdatedAt, item.ID)
	return v.RowsAffected, v.Error
}

func (e *SrvPower) UpdateStatusByIDs(status int, ls []int64) (int64, error) {
	updatedAt := time.Now().Unix()
	ids := conv.ArrayToString(ls, ",")
	sql := fmt.Sprintf("update rb_power set `status` = ?,updated_at=? where `id` in (%s) ", ids)
	v := e.DB.Exec(sql, status, updatedAt)
	return v.RowsAffected, v.Error
}

func (e *SrvPower) FindMultiByIDs(ids []int64) ([]*Model, error) {
	result := make([]*Model, 0)
	st := conv.ArrayToString(ids, ",")
	sql := fmt.Sprintf("select * from rb_power where id in (%v)  and status < 44  ", st)
	v := e.DB.Raw(sql).Scan(&result)
	return result, v.Error
}

func GetTagByList(list []*Model) []string {
	ls := make([]string, 0)
	for _, v := range list {
		ls = append(ls, v.Tag)
	}
	return ls
}

func FindTopLevel(list []*Model) []*Model {
	result := make([]*Model, 0)
	for _, i := range list {
		if i.Fid == 0 {
			result = append(result, i)
		}
	}
	return result
}

// 根据ID查找下1级子元素
func FindSonByID(id int64, list []*Model) []*Model {
	result := make([]*Model, 0)
	for _, i := range list {
		if i.Fid == id {
			result = append(result, i)
		}
	}
	return result
}

func FindSonIDsByID(id int64, list []*Model) []int64 {
	result := make([]int64, 0)
	for _, i := range list {
		if i.Fid == id {
			result = append(result, i.ID)
		}
	}
	return result
}

func FindAllSonIDByID(id int64, list []*Model) []int64 {
	ids := make([]int64, 0)
	if id == 0 {
		return ids
	}
	//ids = append(ids, id)
	son := FindSonIDsByID(id, list) // level 2
	if len(son) <= 0 {
		return ids
	}
	ids = append(ids, son...)
	for _, i := range son {
		son2 := FindSonIDsByID(i, list) // level 3
		if len(son2) <= 0 {
			break
		}
		ids = append(ids, son2...)
		for _, j := range son2 {
			son3 := FindSonIDsByID(j, list) //level4
			if len(son3) <= 0 {
				break
			}
			ids = append(ids, son3...)
		}
	}
	return ids
}
