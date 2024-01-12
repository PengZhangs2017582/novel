// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"novel/internal/dao/internal"
)

// internalShortStoryDao is internal type for wrapping internal DAO implements.
type internalShortStoryDao = *internal.ShortStoryDao

// shortStoryDao is the data access object for table nv_short_story.
// You can define custom methods on it to extend its functionality as you wish.
type shortStoryDao struct {
	internalShortStoryDao
}

var (
	// ShortStory is globally public accessible object for table nv_short_story operations.
	ShortStory = shortStoryDao{
		internal.NewShortStoryDao(),
	}
)

// Fill with you ideas below.
