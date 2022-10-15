package power

type Model struct {
	ID        int64  `json:"id" gorm:"column:id;primary_key" `   //ID
	CreatedAt int64  `gorm:"column:created_at" json:"createdAt"` // 创建时间戳
	UpdatedAt int64  `gorm:"column:updated_at" json:"updatedAt"` // 更新时间戳
	Status    int8   `json:"status" gorm:"column:status"`        // 44 删除, 1 启用, 4 禁用
	Name      string `json:"name" gorm:"column:name"`            //名称
	Tag       string `json:"tag" gorm:"column:tag" `             //`权限标签`,唯一，eg: box-edit
	Type      int8   `json:"type" gorm:"column:type" `           //产品 类型, 默认 0
	Extended  string `json:"extended" gorm:"column:extended" `   // 扩展的
	Fid       int64  `json:"fid" gorm:"column:fid" `             // 父亲id
	Level     int    `json:"level" gorm:"column:level" `         // 层次 深度 [1,4]
	Order     int    `json:"order" gorm:"column:order" `         // 排序 默认 10000
}

func (sv *Model) TableName() string {
	return "rb_power"
}
