package transfer

import (
	"github.com/vl-mobitutor/Netology_GO_Task5/pkg/card"
	"testing"
)

func TestService_Card2Card(t *testing.T) {
	type fields struct {
		CardSvc *card.Service
		Fees    map[string]Fee
	}
	type args struct {
		fromNumber string
		toNumber   string
		amount     int64
	}
	tests := []struct {
		name         string
		fields       fields
		args         args
		wantTotalSum int64
		wantErr      error
	}{
		//-------------------------Параметры тестовых кейсов-------------------------------------
		{
			name:         "Tecт 1 - Карта своего банка -> Карта своего банка (денег достаточно)",
			fields:       fields{
				CardSvc: &card.Service{
					BankName: "Super Bank",
					Cards: []*card.Card{
						{
							Id:       1,
							Issuer:   "MasterCard",
							Balance: 10000_00,
							Currency: "RUR",
							Number:   "5106 2111 1111 0001",
						},
						{
							Id:       2,
							Issuer:   "MasterCard",
							Balance: 10000_00,
							Currency: "RUR",
							Number:   "5106 2111 1111 0002",
						},
						{
							Id:       3,
							Issuer:   "MasterCard",
							Balance: 10000_00,
							Currency: "RUR",
							Number:   "5106 2111 1111 0003",
						},
						{
							Id:       4,
							Issuer:   "Visa",
							Balance: 10000_00,
							Currency: "RUR",
							Number:   "5106 2111 1111 0004",
						},
						{
							Id:       5,
							Issuer:   "Visa",
							Balance: 10000_00,
							Currency: "RUR",
							Number:   "5106 2111 1111 0005",
						},
					},
				},

				Fees: map[string]Fee {
					"in-to-in": {
						Description: "С карты на карту внутри банка",
						FeePercentage: 0.0000, //0.00%
						FeeMinimum: 0, //Минимальная комиссия в копейках
					},
					"in-to-out": {
						Description: "С карты банка на внешнюю карту",
						FeePercentage: 0.0050, //0.5%
						FeeMinimum: 10_00, //Минимальная комиссия в копейках
					},
					"out-to-in": {
						Description: "С внешней карты на карту банка",
						FeePercentage: 0.0000, //0.0%
						FeeMinimum: 0, //Минимальная комиссия в копейках
					},
					"out-to-out": {
						Description: "С внешней карты на внешнюю карту",
						FeePercentage: 0.0150, //1.5%
						FeeMinimum: 30_00, //Минимальная комиссия в копейках
					},
				},
			},
			args:         args{
				fromNumber: "5106 2111 1111 0001",
				toNumber: "5106 2111 1111 0002",
				amount: 5000_00,
			},
			wantTotalSum: 5000_00,
			wantErr:       nil,
		},

		{
			//---------------------------------------------------------------------------------------
			name:         "Tecт 2 - Карта своего банка -> Карта своего банка (денег недостаточно)",
			fields:       fields{
				CardSvc: &card.Service{
					BankName: "Super Bank",
					Cards: []*card.Card{
						{
							Id:       1,
							Issuer:   "MasterCard",
							Balance: 10000_00,
							Currency: "RUR",
							Number:   "5106 2111 1111 0001",
						},
						{
							Id:       2,
							Issuer:   "MasterCard",
							Balance: 10000_00,
							Currency: "RUR",
							Number:   "5106 2111 1111 0002",
						},
						{
							Id:       3,
							Issuer:   "MasterCard",
							Balance: 10000_00,
							Currency: "RUR",
							Number:   "5106 2111 1111 0003",
						},
						{
							Id:       4,
							Issuer:   "Visa",
							Balance: 10000_00,
							Currency: "RUR",
							Number:   "5106 2111 1111 0004",
						},
						{
							Id:       5,
							Issuer:   "Visa",
							Balance: 10000_00,
							Currency: "RUR",
							Number:   "5106 2111 1111 0005",
						},
					},
				},

				Fees: map[string]Fee {
					"in-to-in": {
						Description: "С карты на карту внутри банка",
						FeePercentage: 0.0000, //0.00%
						FeeMinimum: 0, //Минимальная комиссия в копейках
					},
					"in-to-out": {
						Description: "С карты банка на внешнюю карту",
						FeePercentage: 0.0050, //0.5%
						FeeMinimum: 10_00, //Минимальная комиссия в копейках
					},
					"out-to-in": {
						Description: "С внешней карты на карту банка",
						FeePercentage: 0.0000, //0.0%
						FeeMinimum: 0, //Минимальная комиссия в копейках
					},
					"out-to-out": {
						Description: "С внешней карты на внешнюю карту",
						FeePercentage: 0.0150, //1.5%
						FeeMinimum: 30_00, //Минимальная комиссия в копейках
					},
				},
			},
			args:         args{
				fromNumber: "5106 2111 1111 0003",
				toNumber: "5106 2111 1111 0004",
				amount: 20_000_00,
			},
			wantTotalSum: 20_000_00,
			wantErr:      TransferError("Ошибка: на карте-источнике недостаточно средтв для перевода!"),
		},

		{
			//--------------------------------------------------------------------------------------
			name:         "Tecт 3 - Карта своего банка -> Карта чужого банка (денег достаточно)",
			fields:       fields{
				CardSvc: &card.Service{
					BankName: "Super Bank",
					Cards: []*card.Card{
						{
							Id:       1,
							Issuer:   "MasterCard",
							Balance: 10000_00,
							Currency: "RUR",
							Number:   "5106 2111 1111 0001",
						},
						{
							Id:       2,
							Issuer:   "MasterCard",
							Balance: 10000_00,
							Currency: "RUR",
							Number:   "5106 2111 1111 0002",
						},
						{
							Id:       3,
							Issuer:   "MasterCard",
							Balance: 10000_00,
							Currency: "RUR",
							Number:   "5106 2111 1111 0003",
						},
						{
							Id:       4,
							Issuer:   "Visa",
							Balance: 10000_00,
							Currency: "RUR",
							Number:   "5106 2111 1111 0004",
						},
						{
							Id:       5,
							Issuer:   "Visa",
							Balance: 10000_00,
							Currency: "RUR",
							Number:   "5106 2111 1111 0005",
						},
					},
				},

				Fees: map[string]Fee {
					"in-to-in": {
						Description: "С карты на карту внутри банка",
						FeePercentage: 0.0000, //0.00%
						FeeMinimum: 0, //Минимальная комиссия в копейках
					},
					"in-to-out": {
						Description: "С карты банка на внешнюю карту",
						FeePercentage: 0.0050, //0.5%
						FeeMinimum: 10_00, //Минимальная комиссия в копейках
					},
					"out-to-in": {
						Description: "С внешней карты на карту банка",
						FeePercentage: 0.0000, //0.0%
						FeeMinimum: 0, //Минимальная комиссия в копейках
					},
					"out-to-out": {
						Description: "С внешней карты на внешнюю карту",
						FeePercentage: 0.0150, //1.5%
						FeeMinimum: 30_00, //Минимальная комиссия в копейках
					},
				},
			},
			args:         args{
				fromNumber: "5106 2111 1111 0001",
				toNumber: "5106 2111 1111 0009",
				amount: 5_000_00,
			},
			wantTotalSum: 5_025_00,
			wantErr:       nil,
		},

		{
			//--------------------------------------------------------------------------------------
			name:         "Tecт 4 - Карта своего банка -> Карта чужого банка (денег недостаточно)",
			fields:       fields{
				CardSvc: &card.Service{
					BankName: "Super Bank",
					Cards: []*card.Card{
						{
							Id:       1,
							Issuer:   "MasterCard",
							Balance: 10000_00,
							Currency: "RUR",
							Number:   "5106 2111 1111 0001",
						},
						{
							Id:       2,
							Issuer:   "MasterCard",
							Balance: 10000_00,
							Currency: "RUR",
							Number:   "5106 2111 1111 0002",
						},
						{
							Id:       3,
							Issuer:   "MasterCard",
							Balance: 10000_00,
							Currency: "RUR",
							Number:   "5106 2111 1111 0003",
						},
						{
							Id:       4,
							Issuer:   "Visa",
							Balance: 10000_00,
							Currency: "RUR",
							Number:   "5106 2111 1111 0004",
						},
						{
							Id:       5,
							Issuer:   "Visa",
							Balance: 10000_00,
							Currency: "RUR",
							Number:   "5106 2111 1111 0005",
						},
					},
				},

				Fees: map[string]Fee {
					"in-to-in": {
						Description: "С карты на карту внутри банка",
						FeePercentage: 0.0000, //0.00%
						FeeMinimum: 0, //Минимальная комиссия в копейках
					},
					"in-to-out": {
						Description: "С карты банка на внешнюю карту",
						FeePercentage: 0.0050, //0.5%
						FeeMinimum: 10_00, //Минимальная комиссия в копейках
					},
					"out-to-in": {
						Description: "С внешней карты на карту банка",
						FeePercentage: 0.0000, //0.0%
						FeeMinimum: 0, //Минимальная комиссия в копейках
					},
					"out-to-out": {
						Description: "С внешней карты на внешнюю карту",
						FeePercentage: 0.0150, //1.5%
						FeeMinimum: 30_00, //Минимальная комиссия в копейках
					},
				},
			},
			args:         args{
				fromNumber: "5106 2111 1111 0001",
				toNumber: "5106 2111 1111 0009",
				amount: 20_000_00,
			},
			wantTotalSum: 20_100_00,
			wantErr:       TransferError("Ошибка: на карте-источнике недостаточно средтв для перевода!"),
		},

		{
			//--------------------------------------------------------------------------------------
			name:         "Tecт 5 - Карта чужого банка -> Карта своего банка",
			fields:       fields{
				CardSvc: &card.Service{
					BankName: "Super Bank",
					Cards: []*card.Card{
						{
							Id:       1,
							Issuer:   "MasterCard",
							Balance: 10000_00,
							Currency: "RUR",
							Number:   "5106 2111 1111 0001",
						},
						{
							Id:       2,
							Issuer:   "MasterCard",
							Balance: 10000_00,
							Currency: "RUR",
							Number:   "5106 2111 1111 0002",
						},
						{
							Id:       3,
							Issuer:   "MasterCard",
							Balance: 10000_00,
							Currency: "RUR",
							Number:   "5106 2111 1111 0003",
						},
						{
							Id:       4,
							Issuer:   "Visa",
							Balance: 10000_00,
							Currency: "RUR",
							Number:   "5106 2111 1111 0004",
						},
						{
							Id:       5,
							Issuer:   "Visa",
							Balance: 10000_00,
							Currency: "RUR",
							Number:   "5106 2111 1111 0005",
						},
					},
				},

				Fees: map[string]Fee {
					"in-to-in": {
						Description: "С карты на карту внутри банка",
						FeePercentage: 0.0000, //0.00%
						FeeMinimum: 0, //Минимальная комиссия в копейках
					},
					"in-to-out": {
						Description: "С карты банка на внешнюю карту",
						FeePercentage: 0.0050, //0.5%
						FeeMinimum: 10_00, //Минимальная комиссия в копейках
					},
					"out-to-in": {
						Description: "С внешней карты на карту банка",
						FeePercentage: 0.0000, //0.0%
						FeeMinimum: 0, //Минимальная комиссия в копейках
					},
					"out-to-out": {
						Description: "С внешней карты на внешнюю карту",
						FeePercentage: 0.0150, //1.5%
						FeeMinimum: 30_00, //Минимальная комиссия в копейках
					},
				},
			},
			args:         args{
				fromNumber: "5106 2111 1111 0009",
				toNumber: "5106 2111 1111 0001",
				amount: 5_000_00,
			},
			wantTotalSum: 5_000_00,
			wantErr:       nil,
		},

		{
			//--------------------------------------------------------------------------------------
			name:         "Tecт 6 - Карта чужого банка -> Карта чужого банка",
			fields:       fields{
				CardSvc: &card.Service{
					BankName: "Super Bank",
					Cards: []*card.Card{
						{
							Id:       1,
							Issuer:   "MasterCard",
							Balance: 10000_00,
							Currency: "RUR",
							Number:   "5106 2111 1111 0001",
						},
						{
							Id:       2,
							Issuer:   "MasterCard",
							Balance: 10000_00,
							Currency: "RUR",
							Number:   "5106 2111 1111 0002",
						},
						{
							Id:       3,
							Issuer:   "MasterCard",
							Balance: 10000_00,
							Currency: "RUR",
							Number:   "5106 2111 1111 0003",
						},
						{
							Id:       4,
							Issuer:   "Visa",
							Balance: 10000_00,
							Currency: "RUR",
							Number:   "5106 2111 1111 0004",
						},
						{
							Id:       5,
							Issuer:   "Visa",
							Balance: 10000_00,
							Currency: "RUR",
							Number:   "5106 2111 1111 0005",
						},
					},
				},

				Fees: map[string]Fee {
					"in-to-in": {
						Description: "С карты на карту внутри банка",
						FeePercentage: 0.0000, //0.00%
						FeeMinimum: 0, //Минимальная комиссия в копейках
					},
					"in-to-out": {
						Description: "С карты банка на внешнюю карту",
						FeePercentage: 0.0050, //0.5%
						FeeMinimum: 10_00, //Минимальная комиссия в копейках
					},
					"out-to-in": {
						Description: "С внешней карты на карту банка",
						FeePercentage: 0.0000, //0.0%
						FeeMinimum: 0, //Минимальная комиссия в копейках
					},
					"out-to-out": {
						Description: "С внешней карты на внешнюю карту",
						FeePercentage: 0.0150, //1.5%
						FeeMinimum: 30_00, //Минимальная комиссия в копейках
					},
				},
			},
			args:         args{
				fromNumber: "5106 2111 1111 0009",
				toNumber: "5106 2111 1111 0008",
				amount: 5_000_00,
			},
			wantTotalSum: 5_075_00,
			wantErr:       nil,
		},

		{
			//--------------------------------------------------------------------------------------
			name:         "Tecт 7 - Некорректный эмитент по карте - источнику",
			fields:       fields{
				CardSvc: &card.Service{
					BankName: "Super Bank",
					Cards: []*card.Card{
						{
							Id:       1,
							Issuer:   "MasterCard",
							Balance: 10000_00,
							Currency: "RUR",
							Number:   "5106 2111 1111 0001",
						},
						{
							Id:       2,
							Issuer:   "MasterCard",
							Balance: 10000_00,
							Currency: "RUR",
							Number:   "5106 2111 1111 0002",
						},
						{
							Id:       3,
							Issuer:   "MasterCard",
							Balance: 10000_00,
							Currency: "RUR",
							Number:   "5106 2111 1111 0003",
						},
						{
							Id:       4,
							Issuer:   "Visa",
							Balance: 10000_00,
							Currency: "RUR",
							Number:   "5106 2111 1111 0004",
						},
						{
							Id:       5,
							Issuer:   "Visa",
							Balance: 10000_00,
							Currency: "RUR",
							Number:   "5106 2111 1111 0005",
						},
					},
				},

				Fees: map[string]Fee {
					"in-to-in": {
						Description: "С карты на карту внутри банка",
						FeePercentage: 0.0000, //0.00%
						FeeMinimum: 0, //Минимальная комиссия в копейках
					},
					"in-to-out": {
						Description: "С карты банка на внешнюю карту",
						FeePercentage: 0.0050, //0.5%
						FeeMinimum: 10_00, //Минимальная комиссия в копейках
					},
					"out-to-in": {
						Description: "С внешней карты на карту банка",
						FeePercentage: 0.0000, //0.0%
						FeeMinimum: 0, //Минимальная комиссия в копейках
					},
					"out-to-out": {
						Description: "С внешней карты на внешнюю карту",
						FeePercentage: 0.0150, //1.5%
						FeeMinimum: 30_00, //Минимальная комиссия в копейках
					},
				},
			},
			args:         args{
				fromNumber: "1111 2111 1111 0009",
				toNumber: "5106 2111 1111 0008",
				amount: 5_000_00,
			},
			wantTotalSum: 0,
			wantErr:      TransferError("Ошибка по карте-источнику - Ошибка: некорректный эмитент!"),
		},

		{
			//--------------------------------------------------------------------------------------
			name:         "Tecт 8 - Некорректный эмитент по карте - получателю",
			fields:       fields{
				CardSvc: &card.Service{
					BankName: "Super Bank",
					Cards: []*card.Card{
						{
							Id:       1,
							Issuer:   "MasterCard",
							Balance: 10000_00,
							Currency: "RUR",
							Number:   "5106 2111 1111 0001",
						},
						{
							Id:       2,
							Issuer:   "MasterCard",
							Balance: 10000_00,
							Currency: "RUR",
							Number:   "5106 2111 1111 0002",
						},
						{
							Id:       3,
							Issuer:   "MasterCard",
							Balance: 10000_00,
							Currency: "RUR",
							Number:   "5106 2111 1111 0003",
						},
						{
							Id:       4,
							Issuer:   "Visa",
							Balance: 10000_00,
							Currency: "RUR",
							Number:   "5106 2111 1111 0004",
						},
						{
							Id:       5,
							Issuer:   "Visa",
							Balance: 10000_00,
							Currency: "RUR",
							Number:   "5106 2111 1111 0005",
						},
					},
				},

				Fees: map[string]Fee {
					"in-to-in": {
						Description: "С карты на карту внутри банка",
						FeePercentage: 0.0000, //0.00%
						FeeMinimum: 0, //Минимальная комиссия в копейках
					},
					"in-to-out": {
						Description: "С карты банка на внешнюю карту",
						FeePercentage: 0.0050, //0.5%
						FeeMinimum: 10_00, //Минимальная комиссия в копейках
					},
					"out-to-in": {
						Description: "С внешней карты на карту банка",
						FeePercentage: 0.0000, //0.0%
						FeeMinimum: 0, //Минимальная комиссия в копейках
					},
					"out-to-out": {
						Description: "С внешней карты на внешнюю карту",
						FeePercentage: 0.0150, //1.5%
						FeeMinimum: 30_00, //Минимальная комиссия в копейках
					},
				},
			},
			args:         args{
				fromNumber: "5106 2111 1111 0009",
				toNumber: "1111 2111 1111 0008",
				amount: 5_000_00,
			},
			wantTotalSum: 0,
			wantErr:      TransferError("Ошибка по карте-получателю - Ошибка: некорректный эмитент!"),
		},

	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				CardSvc: tt.fields.CardSvc,
				Fees:    tt.fields.Fees,
			}
			gotTotalSum, err := s.Card2Card(tt.args.fromNumber, tt.args.toNumber, tt.args.amount)
			if (err != nil) && (err!= tt.wantErr) {
				t.Errorf("Card2Card() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotTotalSum != tt.wantTotalSum {
				t.Errorf("Card2Card() gotTotalSum = %v, want %v", gotTotalSum, tt.wantTotalSum)
			}
		})
	}
}