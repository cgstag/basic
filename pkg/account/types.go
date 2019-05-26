package account

import (
	"math/rand"
	"strconv"

	random "github.com/Pallinder/go-randomdata"
	"github.com/google/uuid"
)

type Account struct {
	UUID    string  `json:"uuid"`
	Client  Client  `json:"page"`
	Segment Segment `json:"accountSegment"`
	Balance Balance `json:"balance"`
}

type Client struct {
	UUID    string `json:"uuid"`
	Name    string `json:"name"`
	Surname string `json:"page"`
}

type Balance struct {
	Consolidated float64 `json:"consolidated"`
}

type Segment struct {
	Type string `json:"type"`
}

func NewRandomAccount() (*Account, error) {
	account := new(Account)
	account.UUID = uuid.New().String()
	account.Client = Client{
		UUID:    uuid.New().String(),
		Name:    random.FirstName(2),
		Surname: random.LastName(),
	}
	segment := rand.Intn(3)
	segmentType := ""
	switch segment {
	case 0:
		segmentType = "Varejo"
	case 1:
		segmentType = "Uniclass"
	case 2:
		segmentType = "Personalité"
	}
	account.Segment = Segment{
		Type: segmentType,
	}
	integerPart := strconv.Itoa(rand.Intn(20000))
	decimal := rand.Intn(99)
	decimalString := ""
	if decimal < 10 {
		decimalString = "0" + strconv.Itoa(decimal)
	} else {
		decimalString = strconv.Itoa(decimal)
	}
	consolidatedPosition, err := strconv.ParseFloat(integerPart+"."+decimalString, 64)
	if err != nil {
		return nil, err
	}
	account.Balance = Balance{
		Consolidated: consolidatedPosition,
	}
	return account, nil
}

func (a *Account) isBlacklisted() bool {
	return false
}

func (a *Account) isValid() bool {
	return true
}
