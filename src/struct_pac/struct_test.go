package struct_pac

import "testing"

func TestName(t *testing.T) {
	t.Run("interface", func(t *testing.T) {
		areaTest := []struct {
			name string
			shape Shape
			hasArea float64
		}{
			{name: "Rectangle1", shape: Rectangle{Width:  2.0, Height: 2.0}, hasArea: 4.0},
			{name: "Rectangle2", shape: Rectangle{Width:  4.0, Height: 4.0}, hasArea: 16.0},
		}

		for _ , shape := range areaTest {
			t.Run(shape.name, func(t *testing.T) {
				got := shape.shape.Area()
				if got != shape.hasArea {
					t.Errorf("%#v got %g want %g", shape.shape, got, shape.hasArea)
				}
			})
		}

	})
}
