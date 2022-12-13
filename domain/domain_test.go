package domain_test

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/magartifex/employcity/constants"
	"github.com/magartifex/employcity/domain"
	"github.com/magartifex/employcity/repository/inmemory"
)

func TestDomain_Delete(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		d := domain.New(inmemory.New())

		entity, err := d.Set(context.Background(), "first data")
		if err != nil {
			t.Errorf("Set() error = %v", err)
			return
		}

		err = d.Delete(context.Background(), entity.Key)
		if err != nil {
			t.Errorf("Delete() error = %v", err)
			return
		}

		_, err = d.GetByKey(context.Background(), entity.Key)
		if !errors.Is(err, constants.ErrKeyNotExists) {
			t.Errorf("GetByKey() error = %v", err)
			return
		}
	})
}

func TestDomain_Get(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		d := domain.New(inmemory.New())

		firstEntity, err := d.Set(context.Background(), "first data")
		if err != nil {
			t.Errorf("Set() error = %v", err)
			return
		}

		secondEntity, err := d.Set(context.Background(), "second data")
		if err != nil {
			t.Errorf("Set() error = %v", err)
			return
		}

		lastEntity, err := d.Set(context.Background(), "last data")
		if err != nil {
			t.Errorf("Set() error = %v", err)
			return
		}

		got, err := d.Get(context.Background())
		if err != nil {
			t.Errorf("Get() error = %v", err)
			return
		}

		if len(got) != 3 {
			t.Error("Get() count list != 3")
			return
		}

		for _, entity := range got {
			switch {
			case
				reflect.DeepEqual(entity, firstEntity),
				reflect.DeepEqual(entity, secondEntity),
				reflect.DeepEqual(entity, lastEntity):

			default:
				t.Errorf("GetByKey() got = %#v", got)
				return
			}
		}
	})
}

func TestDomain_GetByKey(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		d := domain.New(inmemory.New())

		entity, err := d.Set(context.Background(), "first data")
		if err != nil {
			t.Errorf("Set() error = %v", err)
			return
		}

		got, err := d.GetByKey(context.Background(), entity.Key)
		if err != nil {
			t.Errorf("GetByKey() error = %v", err)
			return
		}

		if !reflect.DeepEqual(got, entity) {
			t.Errorf("GetByKey() got = %v, entity %v", got, entity)
		}
	})

	t.Run("key not exists", func(t *testing.T) {
		d := domain.New(inmemory.New())

		_, err := d.GetByKey(context.Background(), "not key")
		if !errors.Is(err, constants.ErrKeyNotExists) {
			t.Errorf("GetByKey() error = %v", err)
			return
		}
	})
}
