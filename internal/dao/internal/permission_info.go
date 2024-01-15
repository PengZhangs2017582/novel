// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// PermissionInfoDao is the data access object for table nv_permission_info.
type PermissionInfoDao struct {
	table   string                // table is the underlying table name of the DAO.
	group   string                // group is the database configuration group name of current DAO.
	columns PermissionInfoColumns // columns contains all the column names of Table for convenient usage.
}

// PermissionInfoColumns defines and stores column names for table nv_permission_info.
type PermissionInfoColumns struct {
	Id        string //
	Name      string // 权限名称
	Path      string // 路径
	CreatedAt string //
	UpdatedAt string //
	DeletedAt string //
}

// permissionInfoColumns holds the columns for table nv_permission_info.
var permissionInfoColumns = PermissionInfoColumns{
	Id:        "id",
	Name:      "name",
	Path:      "path",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	DeletedAt: "deleted_at",
}

// NewPermissionInfoDao creates and returns a new DAO object for table data access.
func NewPermissionInfoDao() *PermissionInfoDao {
	return &PermissionInfoDao{
		group:   "default",
		table:   "nv_permission_info",
		columns: permissionInfoColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *PermissionInfoDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *PermissionInfoDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *PermissionInfoDao) Columns() PermissionInfoColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *PermissionInfoDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *PermissionInfoDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *PermissionInfoDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
