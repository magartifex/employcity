package api_test

import (
	"context"
	"reflect"
	"testing"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/magartifex/employcity/api"
	"github.com/magartifex/employcity/domain"
	v1 "github.com/magartifex/employcity/proto/v1"
	"github.com/magartifex/employcity/repository/inmemory"
)

func TestApi_Delete(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		a := api.New(domain.New(inmemory.New()))

		resp, err := a.Set(context.Background(), &v1.Set_Request{Value: "first data"})
		if err != nil {
			t.Errorf("Set() error = %v", err)
			return
		}

		_, err = a.Delete(context.Background(), &v1.Delete_Request{Key: resp.GetData().GetKey()})
		if err != nil {
			t.Errorf("Delete() error = %v", err)
			return
		}

		_, err = a.GetByKey(context.Background(), &v1.GetByKey_Request{Key: resp.GetData().GetKey()})
		if status.Convert(err).Code() != codes.NotFound {
			t.Errorf("GetByKey() error = %v", err)
			return
		}
	})
}

func TestApi_Get(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		a := api.New(domain.New(inmemory.New()))

		firstEntity, err := a.Set(context.Background(), &v1.Set_Request{Value: "first data"})
		if err != nil {
			t.Errorf("Set() error = %v", err)
			return
		}

		secondEntity, err := a.Set(context.Background(), &v1.Set_Request{Value: "second data"})
		if err != nil {
			t.Errorf("Set() error = %v", err)
			return
		}

		lastEntity, err := a.Set(context.Background(), &v1.Set_Request{Value: "last data"})
		if err != nil {
			t.Errorf("Set() error = %v", err)
			return
		}

		res, err := a.Get(context.Background(), &v1.Empty{})
		if err != nil {
			t.Errorf("Get() error = %v", err)
			return
		}

		if len(res.GetList()) != 3 {
			t.Error("Get() count list != 3")
			return
		}

		for _, entity := range res.GetList() {
			switch {
			case
				reflect.DeepEqual(entity, firstEntity.GetData()),
				reflect.DeepEqual(entity, secondEntity.GetData()),
				reflect.DeepEqual(entity, lastEntity.GetData()):

			default:
				t.Errorf("GetByKey() got = %#v", res.GetList())
				return
			}
		}
	})
}

func TestApi_GetByKey(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		a := api.New(domain.New(inmemory.New()))

		entity, err := a.Set(context.Background(), &v1.Set_Request{Value: "first data"})
		if err != nil {
			t.Errorf("Set() error = %v", err)
			return
		}

		got, err := a.GetByKey(context.Background(), &v1.GetByKey_Request{Key: entity.GetData().GetKey()})
		if err != nil {
			t.Errorf("GetByKey() error = %v", err)
			return
		}

		if !reflect.DeepEqual(got.GetData(), entity.GetData()) {
			t.Errorf("GetByKey() got = %v, entity %v", got.GetData(), entity.GetData())
		}
	})

	t.Run("key not exists", func(t *testing.T) {
		a := api.New(domain.New(inmemory.New()))

		_, err := a.GetByKey(context.Background(), &v1.GetByKey_Request{Key: "not key"})
		if status.Convert(err).Code() != codes.NotFound {
			t.Errorf("GetByKey() error = %v", err)
			return
		}
	})
}
