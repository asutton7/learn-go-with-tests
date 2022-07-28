package perimeter

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	received := Perimeter(rectangle)
	expected := 40.0

	if received != expected {
		t.Errorf("got %.2f expected %.2f", received, expected)
	}
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		name   string
		shape  Shape
		expect float64
	}{
		{name: "Rectangle", shape: Rectangle{Length: 12, Width: 6}, expect: 72.0},
		{name: "Circle", shape: Circle{Radius: 10}, expect: 314.1592653589793},
		{name: "Triangle", shape: Triangle{Base: 12, Height: 6}, expect: 36.0},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.expect {
				t.Errorf("%#v got %g want %g", tt.shape, got, tt.expect)
			}
		})
	}

}
