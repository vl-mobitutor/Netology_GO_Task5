package transfer

import (
	"fmt"
	"github.com/vl-mobitutor/Netology_GO_Task5/pkg/card"
	"math"
)

type Service struct {
	CardSvc *card.Service
	Fees map[string]Fee
}

type Fee struct {
	Description string
	FeePercentage float64 //Размер комиссия в % - указывается в виде десятичной дроби - т.е. 1.50% => 0.0150
	FeeMinimum int64 //Минимальная комиссия - указывается в копейках
}

//Функция-конструктор сервиса
func NewService(cardSvc *card.Service, feeSet map[string]Fee) *Service {
	return &Service {
		CardSvc: cardSvc,
		Fees: feeSet,
	}
}


//Функция расчета комиссии
func (s *Service) FeeCalculation (operationType string, operationAmount int64) (fee int64) {
	fee = int64( math.Round(s.Fees[operationType].FeePercentage * float64(operationAmount)))

	if fee < s.Fees[operationType].FeeMinimum {
		fee = s.Fees[operationType].FeeMinimum
	}

	return fee
}


//Функция перевода по номеру карты
func (s *Service) Card2Card(fromNumber, toNumber string, amount int64) (totalSum int64, ok bool) {
	var myFee int64

	if fromNumber == toNumber { //Проверка на совпадение номеров карты-источника и карты получателя
		fmt.Println("Номера карты-источника и карты-получателя совпадают!")
		return amount, false
	}

	if amount <= 0 { //Проверка корректности суммы перевода
		fmt.Println("Некорректная сумма перевода!")
		return amount, false
	}

	//Определяем по номерам чьи карты
	cardFrom := s.CardSvc.SearchByNumber(fromNumber)
	cardTo := s.CardSvc.SearchByNumber(toNumber)


	//-----------------------------Блок, если обе карты "наши"---------------------------------
	if (cardFrom != nil) && (cardTo != nil) {

		totalSum = amount + s.FeeCalculation("in-to-in", amount) //Полная сумма списания с карты-источника

		if totalSum > cardFrom.Balance {
			fmt.Printf("На карте %s недостаточно средств для перевода! \n", fromNumber)
			ok = false
			return
		}

		cardFrom.Balance -= totalSum //Списание с карты источника суммы перевода  + комиссия
		cardTo.Balance += amount //Зачисление на карту-получатель суммы перевода (без комиссии)

		fmt.Println("Тип перевода - внутрибанковский платеж")
		ok = true
	}


	//-------------------------Блок, если с "нашей" карты на внешнюю карту---------------------
	if (cardFrom !=nil ) && (cardTo == nil) {

		totalSum = amount + s.FeeCalculation("in-to-out", amount) //Полная сумма списания с карты-источника

		if totalSum > cardFrom.Balance {
			fmt.Printf("На карте %s недостаточно средств для перевода! \n", fromNumber)
			ok = false
			return
		}

		cardFrom.Balance -= totalSum //Списание с карты источника суммы перевода  + комиссия

		fmt.Println("Тип перевода - с карты банка на внешнюю карту")
		ok = true
	}


	//------------------------Блок, если с внешней карты на карту банка------------------------
	if (cardFrom == nil) && (cardTo != nil) {

		totalSum = amount + s.FeeCalculation("out-to-in", amount) //Полная сумма списания с карты-источника

		totalSum = amount + myFee //Полная сумма списания с карты-источника
		cardTo.Balance += amount //Зачисление на карту-получатель суммы перевода (без комиссии)

		fmt.Println("Тип перевода - с внешней карты на карту банка")
		ok = true
	}

	//------------------------Блок, если с внешней карты на внешнюю банка---------------------
	if (cardFrom == nil) && (cardTo == nil) {

		totalSum = amount + s.FeeCalculation("out-to-out", amount) //Полная сумма списания с карты-источника

		fmt.Println("Тип перевода - с внешней карты на внешнюю карту")
		ok = true
	}

	return totalSum, ok
}
