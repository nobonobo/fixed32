package fixed32

import (
	"reflect"
	"testing"
)

func TestFixed32_Pow(t *testing.T) {
	type args struct {
		exponent Fixed32
	}
	tests := []struct {
		name string
		f    Fixed32
		args args
		want Fixed32
	}{
		{"1^1", FromFloat32(1), args{FromFloat32(1)}, FromFloat32(1)},
		{"2^2", FromFloat32(2), args{FromFloat32(2)}, FromFloat32(3.999195)},
		{"3^2", FromFloat32(3), args{FromFloat32(2)}, FromFloat32(8.997205)},
		{"4^2", FromFloat32(4), args{FromFloat32(2)}, FromFloat32(15.98607)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.f.Pow(tt.args.exponent); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Fixed32.Pow() = %v, want %v", got, tt.want)
			}
		})
	}
}
