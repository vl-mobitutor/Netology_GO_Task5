package card

import "testing"

func TestIsValid(t *testing.T) {
	type args struct {
		cardNumber string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Тест 1 - корректный номер карты",
			args: args{
				cardNumber: "4561 2612 1234 5467",
			},
			want: true,
		},
		{
			name: "Тест 2 - некорректный номер карты",
			args: args{
				cardNumber: "4561 2612 1234 5464",
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsValid(tt.args.cardNumber); got != tt.want {
				t.Errorf("IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}