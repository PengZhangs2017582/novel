// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// WriterDao is the data access object for table nv_writer.
type WriterDao struct {
	table   string        // table is the underlying table name of the DAO.
	group   string        // group is the database configuration group name of current DAO.
	columns WriterColumns // columns contains all the column names of Table for convenient usage.
}

// WriterColumns defines and stores column names for table nv_writer.
type WriterColumns struct {
	Id           string //
	PenName      string // 笔名
	Phone        string // 手机号
	Password     string // 密码
	Introduction string // 简介
	Avatar       string // 头像
	Qq           string //
	RealName     string // 真名
	IdNo         string // 身份证号
	IsAdmin      string // 是否是管理员
	RoleIds      string // 角色id
	CreatedAt    string //
	UpdatedAt    string //
	DeletedAt    string //
}

// writerColumns holds the columns for table nv_writer.
var writerColumns = WriterColumns{
	Id:           "id",
	PenName:      "pen_name",
	Phone:        "phone",
	Password:     "password",
	Introduction: "introduction",
	Avatar:       "avatar",
	Qq:           "qq",
	RealName:     "real_name",
	IdNo:         "id_no",
	IsAdmin:      "is_admin",
	RoleIds:      "role_ids",
	CreatedAt:    "created_at",
	UpdatedAt:    "updated_at",
	DeletedAt:    "deleted_at",
}

// NewWriterDao creates and returns a new DAO object for table data access.
func NewWriterDao() *WriterDao {
	return &WriterDao{
		group:   "default",
		table:   "nv_writer",
		columns: writerColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *WriterDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *WriterDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *WriterDao) Columns() WriterColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *WriterDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *WriterDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *WriterDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
