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



type TransferError string

func (e TransferError) Error() string {
	return string(e)
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
func (s *Service) Card2Card(fromNumber, toNumber string, amount int64) (totalSum int64, err error) {
	var myFee int64

	if fromNumber == toNumber { //Проверка на совпадение номеров карты-источника и карты получателя
		return 0, TransferError("Ошибка: номера карты-источника и карты-получателя совпадают!")
	}

	if amount <= 0 { //Проверка корректности суммы перевода
		return 0, TransferError("Ошибка: сумма перевода отрицательна либо равна нулю!")
	}

	//Определяем по номерам чьи карты
	cardFrom, ourFrom, errFrom := s.CardSvc.SearchByNumber(fromNumber)
	cardTo, ourTo, errTo := s.CardSvc.SearchByNumber(toNumber)

	if errFrom != nil {

		return 0, TransferError("Ошибка по карте-источнику - " + errFrom.Error())
	}

	if errTo != nil {
		return 0, TransferError("Ошибка по карте-получателю - " + errTo.Error())
	}

	//-----------------------------Блок, если обе карты "наши"---------------------------------
	if (ourFrom == true) && (ourTo == true) {

		totalSum = amount + s.FeeCalculation("in-to-in", amount) //Полная сумма списания с карты-источника

		if totalSum > cardFrom.Balance {
			return totalSum, TransferError("Ошибка: на карте-источнике недостаточно средтв для перевода!")
		}

		cardFrom.Balance -= totalSum //Списание с карты источника суммы перевода  + комиссия
		cardTo.Balance += amount //Зачисление на карту-получатель суммы перевода (без комиссии)

		fmt.Println("Тип перевода - внутрибанковский платеж")
	}


	//-------------------------Блок, если с "нашей" карты на внешнюю карту---------------------
	if (ourFrom == true ) && (ourTo == false) {

		totalSum = amount + s.FeeCalculation("in-to-out", amount) //Полная сумма списания с карты-источника

		if totalSum > cardFrom.Balance {
			return totalSum, TransferError("Ошибка: на карте-источнике недостаточно средтв для перевода!")
		}

		cardFrom.Balance -= totalSum //Списание с карты источника суммы перевода  + комиссия
		fmt.Println("Тип перевода - с карты банка на внешнюю карту")
	}


	//------------------------Блок, если с внешней карты на карту банка------------------------
	if (ourFrom == false) && (ourTo == true) {

		totalSum = amount + s.FeeCalculation("out-to-in", amount) //Полная сумма списания с карты-источника

		totalSum = amount + myFee //Полная сумма списания с карты-источника
		cardTo.Balance += amount //Зачисление на карту-получатель суммы перевода (без комиссии)

		fmt.Println("Тип перевода - с внешней карты на карту банка")
	}

	//------------------------Блок, если с внешней карты на внешнюю банка---------------------
	if (ourFrom == false) && (ourTo == false) {

		totalSum = amount + s.FeeCalculation("out-to-out", amount) //Полная сумма списания с карты-источника

		fmt.Println("Тип перевода - с внешней карты на внешнюю карту")
	}

	return totalSum, nil
}
