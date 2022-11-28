package structs

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %2.f", got, want)
	}
}

func TestArea(t *testing.T) {
	// v1
	// checkArea := func(t testing.TB, shape Shape, want float64) {
	// 	t.Helper()
	// 	got := shape.Area()
	// 	if got != want {
	// 		t.Errorf("got %.2f want %2.f", got, want)
	// 	}

	// }

	// t.Run("rectangles", func(t *testing.T) {
	// 	rectangle := Rectangle{12.0, 6.0}
	// 	checkArea(t, rectangle, 72.0)
	// })

	// t.Run("circles", func(t *testing.T) {
	// 	circle := Circle{10}
	// 	checkArea(t, circle, 314.1592653589793)
	// })

	// TABLE DRIVEN TESTS
	// Declaring slice of structs by using []struct with 2 fields
	// Fill flice with cases
	areaTests := []struct {
		name  string
		shape Shape
		want  float64
	}{
		// {Rectangle{12, 6}, 72.0},
		// {Circle{10}, 314.1592653589793},
		// {Triangle{12, 6}, 6.0},

		// Optionally name fields
		{name: "Rectangle", shape: Rectangle{Width: 12, Height: 6}, want: 72.0},
		{name: "Cirlce", shape: Circle{Radius: 10}, want: 314.1592653589793},
		{name: "Triangle", shape: Triangle{Base: 12, Height: 6}, want: 36.0},
	}

	// Iterate over the slices, using the struct fields to run out test
	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()

			if got != tt.want {
				t.Errorf("got %.2f want %2.f", got, tt.want)
			}
		})
	}

}
