package service

import (
	"testing"

	"github.com/Wastoids/boxesandthingsbackend/storage"
	"github.com/smartystreets/goconvey/convey"
)

func TestThingService_GetThingsByBox(t *testing.T) {
	convey.Convey("Given a box id", t, func() {

		convey.Convey("When i need to get the things inside a box", func() {
			boxID := "60df17a9-3439-4f8b-a58a-2623e4d95601"
			thingSvc := NewThingService(storage.NewRepository())
			things, err := thingSvc.GetThingsByBox(boxID)

			convey.Convey("Then the things must not be empty", func() {
				convey.So(len(things), convey.ShouldResemble, 2)
				convey.So(err, convey.ShouldResemble, nil)
			})
		})
	})
}
