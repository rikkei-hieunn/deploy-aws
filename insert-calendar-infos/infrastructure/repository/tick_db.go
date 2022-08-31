package repository

import (
	"context"
	"github.com/rs/zerolog/log"
	"insert-calendar-infos/infrastructure"
	"insert-calendar-infos/infrastructure/db/rdb"
)

// TickDBRepository Structure of repository DB
type TickDBRepository struct {
	db rdb.ITickDBHandler
}

// NewTickDBRepository Initialize a Repository DB
func NewTickDBRepository(infra *infrastructure.Infra) ITickDBRepository {
	return &TickDBRepository{
		infra.TickDB,
	}
}

// ExecWithTx exec in transaction
func (r *TickDBRepository) ExecWithTx(ctx context.Context, sql, kei string, args []interface{}) error {
	err := r.db.ExecWithTx(ctx, sql, kei, args)
	if err != nil {
		return err
	}

	return nil
}

// RollbackTx rollback the change if error
func (r *TickDBRepository) RollbackTx(kei string) {
	err := r.db.RollbackTx(kei)
	if err != nil {
		log.Error().Msg(err.Error())

		return
	}
}

// CommitTx Commit the change if all queries ran successfully
func (r *TickDBRepository) CommitTx(kei string) {
	err := r.db.CommitTx(kei)
	if err != nil {
		log.Error().Msg(err.Error())

		return
	}
}

// InitTx create a transaction
func (r *TickDBRepository) InitTx(ctx context.Context, kei string) error {
	err := r.db.InitTx(ctx, kei)
	if err != nil {
		return err
	}

	return err
}

// Close connection
func (t TickDBRepository) Close() error {
	return t.db.Close()
}
