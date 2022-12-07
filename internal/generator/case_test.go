package generator

import "testing"

func Test_CamelCase(t *testing.T) {
	tests := []struct {
		in   string
		want string
	}{
		{
			in:   "order_info",
			want: "OrderInfo",
		},
		{
			in:   "order_info_product_info",
			want: "OrderInfoProductInfo",
		},
		{
			in:   "order_info123",
			want: "OrderInfo123",
		},
		{
			in:   "order_info123@!#$^$^_test",
			want: "OrderInfo123Test",
		},
	}
	for _, tt := range tests {
		got := CamelCase(tt.in, true)
		if got != tt.want {
			t.Errorf("in(%s) got(%s) want(%s)", tt.in, got, tt.want)
		}
	}
}
