/*
Package repository implements logics repository.
*/
package repository

import (
	"context"
)

// IECSRepository method about ecs repository
type IECSRepository interface {
	UpdateTask(context.Context) error
}
