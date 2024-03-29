// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"novel/internal/dao/internal"
)

// internalWriterDao is internal type for wrapping internal DAO implements.
type internalWriterDao = *internal.WriterDao

// writerDao is the data access object for table nv_writer.
// You can define custom methods on it to extend its functionality as you wish.
type writerDao struct {
	internalWriterDao
}

var (
	// Writer is globally public accessible object for table nv_writer operations.
	Writer = writerDao{
		internal.NewWriterDao(),
	}
)

// Fill with you ideas below.
