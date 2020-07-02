package card

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


//Функция поиска карты по номеру в массиве собственных карт банка
func (s *Service) SearchByNumber (number string) *Card {
	for _, card := range s.Cards {
		if card.Number == number {
			return card
		}
	}
	return nil
}