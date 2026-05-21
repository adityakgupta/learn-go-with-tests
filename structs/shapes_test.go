package structs

import "testing"

func TestArea(t *testing.T) {

	areaTest := []struct {
		name string
		s    Shape
		want float64
	}{
		{name: "Rectangle", s: Rectangle{Width: 20, Height: 40}, want: 800},
		{name: "Circle", s: Circle{Radius: 10}, want: 314.1592653589793},
		{name: "Triangle", s: Triangle{Base: 3, Height: 4}, want: 6},
	}

	for _, tt := range areaTest {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.s.Area()

			if got != tt.want {
				t.Errorf("For %#v got %g, want %g", tt.s, got, tt.want)
			}
		})
	}
}
