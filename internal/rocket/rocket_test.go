package rocket

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestRocketService(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	t.Run("GetRocketByID", func(t *testing.T) {
		rocketStoreMock := NewMockStore(mockCtrl)
		id := "UUID-1"
		rocketStoreMock.EXPECT().GetRocketByID(id).Return(Rocket{
			ID: id,
		}, nil)
		rocketService := New(rocketStoreMock)
		rkt, err := rocketService.GetRocketByID(
			context.Background(), id,
		)
		assert.NoError(t, err)
		assert.Equal(t, id, rkt.ID)
	})

	t.Run("InsertRocket", func(t *testing.T) {
		rocketStoreMock := NewMockStore(mockCtrl)
		id := "UUID-1"
		rocket := Rocket{
			ID: id,
		}
		rocketStoreMock.EXPECT().InsertRocket(rocket).Return(rocket, nil)
		rocketService := New(rocketStoreMock)
		rkt, err := rocketService.InsertRocket(
			context.Background(), rocket,
		)
		assert.NoError(t, err)
		assert.Equal(t, id, rkt.ID)
	})

	t.Run("DeleteRocket", func(t *testing.T) {
		rocketStoreMock := NewMockStore(mockCtrl)
		id := "UUID-1"
		rocketStoreMock.EXPECT().DeleteRocket(id).Return(nil)
		rocketService := New(rocketStoreMock)
		err := rocketService.DeleteRocket(
			context.Background(), id,
		)
		assert.NoError(t, err)
	})
}
