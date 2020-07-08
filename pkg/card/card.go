package card

import (
	"fmt"
	"strconv"
	"strings"
)

type Service struct {
	BankName string
	Cards []*Card
}


type Card struct {
	Id int64
	Issuer string
	Balance int64
	Currency string
	Number string
}


func NewService(bankName string) *Service {
	return &Service{BankName: bankName}
}


//Функция выпуска карты
func (s *Service) IssueCard(id int64, issuer, currency, number string) *Card {
	card := &Card {
		Id: id,
		Issuer: issuer,
		Balance: 0,
		Currency: currency,
		Number: number,
	}
	s.Cards = append(s.Cards, card)
	return card
}


//Поиск карты по номеру в массиве собственных карт банка
type FindCardError string
func (e FindCardError) Error() string {
	return string(e)
}

func (s *Service) SearchByNumber (number string) (card *Card, ourCard bool, err error) {

	ourIssuer :="5106 21"
	if !strings.HasPrefix(number, ourIssuer) {
		return nil, false, FindCardError("Ошибка: некорректный эмитент!")
	}

	for _, card := range s.Cards {
		if card.Number == number {
			return card, true,nil
		}
	}
	return nil, false, nil
}




//Проверка корректности номера карты по алгоритму Луна
func IsValid (cardNumber string) bool {

	//превращаем строку номера в слайс из строк-цифр
	figureSlice := strings.Split(strings.ReplaceAll(cardNumber, " ", ""), "")

	//превращаем слайс из строк-цифр с слайс из чисел int
	var numberSlice []int
	for _, oneFigure := range figureSlice {
		oneNumber, err := strconv.Atoi(oneFigure);
		if  err != nil {
			return false
		}
		numberSlice = append(numberSlice, oneNumber)
	}

	//создаем контрольный слайс по алгоритму
	var controlNumberSlice []int
	for i, oneNumber := range numberSlice {
		//Для нечетных элементов массива, т.е. с индексами 0, 2, 4 и т.д. заменяем элементы в слайсе
		if i % 2 == 0 {
			if 2*oneNumber > 9 {
				controlNumberSlice = append(controlNumberSlice, 2*oneNumber - 9)
			} else {
				controlNumberSlice = append(controlNumberSlice, 2*oneNumber)
			}
			continue
		}
		//Для четных элементов массива, т.е. с индексами 1, 3, 5 и т.д. оставляем значения без изменений
		controlNumberSlice = append(controlNumberSlice, oneNumber)
	}

	//Проверяем сумму элементов контрольного слайса
	controlSum := 0
	for _, oneControlNumber := range controlNumberSlice {
		controlSum += oneControlNumber
	}
	if controlSum%10 ==0 {
		return true //Контрольная сумма корректна
	}

	fmt.Println("Контрольная сумма некорректна -", controlSum)
	return false  //Контрольная сумма некорректна
}