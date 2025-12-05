package models

import (
	"time"

	"gorm.io/gorm"
)

// User 人员模型
type User struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	Name         string `gorm:"type:varchar(50);not null" json:"name"`           // 姓名
	DepartmentID uint   `gorm:"not null;index" json:"department_id"`             // 所属部门ID
	Ip           string `gorm:"type:varchar(50)" json:"ip"`                      // IP地址
	EmployeeID   string `gorm:"type:varchar(50);uniqueIndex" json:"employee_id"` // 工号（可选）
	Position     string `gorm:"type:varchar(100)" json:"position"`               // 岗位名称/描述

	// 关联关系
	Department Department `gorm:"foreignKey:DepartmentID" json:"department,omitempty"`
	Uploads    []Upload   `gorm:"foreignKey:UserID" json:"uploads,omitempty"`
}

// TableName 指定表名
func (User) TableName() string {
	return "users"
}
