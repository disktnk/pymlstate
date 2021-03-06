package mnist

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/sensorbee/sensorbee.v0/bql"
	"gopkg.in/sensorbee/sensorbee.v0/core"
	"gopkg.in/sensorbee/sensorbee.v0/data"
	"os"
	"testing"
)

func TestCreateSource(t *testing.T) {
	ctx := &core.Context{}
	ioParams := &bql.IOParams{}
	Convey("Given a MNIST data source creator", t, func() {
		dc := DataSourceCreator{}
		Convey("When get parameters which a lack required value", func() {
			baseParams := data.Map{
				"images_file_name":   data.String("_test_train_image"),
				"labels_file_name":   data.String("_test_train_label"),
				"data_size":          data.Int(1),
				"image_element_size": data.Int(1),
			}
			requiredParam := []string{
				"images_file_name", "labels_file_name", "data_size",
			}
			for _, v := range requiredParam {
				v := v
				msg := fmt.Sprintf("Then should return '%v' value is not found error", v)
				Convey(msg, func() {
					p := baseParams[v]
					delete(baseParams, v)

					s, err := dc.CreateSource(ctx, ioParams, baseParams)
					So(err, ShouldNotBeNil)
					So(err.Error(), ShouldContainSubstring, v)
					So(s, ShouldBeNil)

					baseParams[v] = p
				})
			}
		})
		Convey("When get parameters which set not exist image file name", func() {
			params := data.Map{
				"images_file_name": data.String("_test_train_image_"),
				"labels_file_name": data.String("_test_train_label"),
				"data_size":        data.Int(1),
			}
			Convey("Then the creator should return not found error", func() {
				s, err := dc.CreateSource(ctx, ioParams, params)
				So(os.IsNotExist(err), ShouldBeTrue)
				So(s, ShouldBeNil)
			})
		})
		Convey("When get parameters which set not exist label file name", func() {
			params := data.Map{
				"images_file_name": data.String("_test_train_image"),
				"labels_file_name": data.String("_test_train_label_"),
				"data_size":        data.Int(1),
			}
			Convey("Then the creator should return not found error", func() {
				s, err := dc.CreateSource(ctx, ioParams, params)
				So(os.IsNotExist(err), ShouldBeTrue)
				So(s, ShouldBeNil)
			})
		})

		// not error case
		Convey("When get full value parameters", func() {
			params := data.Map{
				"images_file_name":   data.String("_test_train_image"),
				"labels_file_name":   data.String("_test_train_label"),
				"data_size":          data.Int(1),
				"image_element_size": data.Int(1),
			}
			Convey("Then the creator should return not found error", func() {
				s, err := createMNISTDataSource(ctx, ioParams, params)
				So(err, ShouldBeNil)
				So(s, ShouldNotBeNil)

				ms, ok := s.(*mnistDataSource)
				So(ok, ShouldBeTrue)
				So(len(ms.data), ShouldEqual, 1)
				So(len(ms.data[0]), ShouldEqual, 1)
				So(len(ms.target), ShouldEqual, 1)
				So(ms.imageElemSize, ShouldEqual, 1)
				So(ms.dataSize, ShouldEqual, 1)
			})
		})
		Convey("When get parameters which set only required values", func() {
			params := data.Map{
				"images_file_name": data.String("_test_train_image"),
				"labels_file_name": data.String("_test_train_label"),
				"data_size":        data.Int(1),
			}
			Convey("Then the creator should return not found error", func() {
				s, err := createMNISTDataSource(ctx, ioParams, params)
				So(err, ShouldBeNil)
				So(s, ShouldNotBeNil)

				ms, ok := s.(*mnistDataSource)
				So(ok, ShouldBeTrue)
				So(len(ms.data), ShouldEqual, 1)
				So(len(ms.data[0]), ShouldEqual, 28*28)
				So(len(ms.target), ShouldEqual, 1)
				So(ms.imageElemSize, ShouldEqual, 28*28)
				So(ms.dataSize, ShouldEqual, 1)
			})
		})
	})
}
