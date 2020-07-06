package main

import (
	"fmt"
	"github.com/vl-mobitutor/Netology_GO_Task5/pkg/card"
	"github.com/vl-mobitutor/Netology_GO_Task5/pkg/transfer"
)

func main() {

	//Массив карт к "выпуску"
	myCards := []card.Card {
		{
			Id: 1,
			Issuer: "MasterCard",
			Currency: "RUR",
			Number: "1111 1111 1111 0001",
		},
		{
			Id: 2,
			Issuer: "MasterCard",
			Currency: "RUR",
			Number: "1111 1111 1111 0002",
		},
		{
			Id: 3,
			Issuer: "MasterCard",
			Currency: "RUR",
			Number: "1111 1111 1111 0003",
		},
		{
			Id: 4,
			Issuer: "Visa",
			Currency: "RUR",
			Number: "1111 1111 1111 0004",
		},
		{
			Id: 5,
			Issuer: "Visa",
			Currency: "RUR",
			Number: "1111 1111 1111 0005",
		},
	}

	//Выпускаем карты и кладем каждой карте на счет по 10_000 рублей
	svc := card.NewService("Super Bank")
	for index, newCard := range myCards {
		svc.IssueCard(newCard.Id, newCard.Issuer, newCard.Currency, newCard.Number)
		svc.Cards[index].Balance = 10_000_00
	}


	fmt.Println("-------------------Балансы собственных карт банка до операции перевода---------------------------")
	for _, value := range svc.Cards {
		fmt.Println(*value)
	}
	fmt.Println("------------------------------------------------------------------------------")
	fmt.Println()

	//Настройки комиссий
	feeSet := map[string]transfer.Fee {
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
	}

	//Инициация и первоначальная настройка перевода
	trf := transfer.NewService(svc, feeSet)

	//Выполнение перевода
	fromNumber := "1111 1111 1111 0001" //Протестировать свои карты - меням последнюю цифру от 1 до 5
	toNumber := "2111 1111 1111 0002"   //Протестировать внешние карты - меняем первые цифры
	amount := 20000_00
	totalAmount, transferOk :=trf.Card2Card(fromNumber, toNumber, int64(amount))

	if transferOk {
		fmt.Printf("Перевод c карты %s успешно выполнен: \n", fromNumber)
		fmt.Printf("Сумма перевода - %d \n", amount)
		fmt.Printf("Полная сумма списания с комиссией - %d \n", totalAmount)
		fmt.Println()
		fmt.Println("-------------------Балансы собственных карт банка после операции перевода------------------------")
		for _, value := range svc.Cards {
			fmt.Println(*value)
		}
	} else {
		fmt.Printf("Перевод на сумму %d не выполнен!", totalAmount)
	}

}
