package account

import (
	"math/rand"
	"strconv"

	random "github.com/Pallinder/go-randomdata"
	"github.com/google/uuid"
)

type Account struct {
	AccountUUID    string
	Client         Client
	AccountSegment Segment
	Balance        Balance
}

type Client struct {
	ClientUUID    string
	ClientName    string
	ClientSurname string
}

type Balance struct {
	Consolidated float64
}

type Segment struct {
	SegmentType string
}

func NewRandomAccount() (*Account, error) {
	account := new(Account)
	account.AccountUUID = uuid.New().String()
	account.Client = Client{
		ClientUUID:    uuid.New().String(),
		ClientName:    random.FirstName(2),
		ClientSurname: random.LastName(),
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
	account.AccountSegment = Segment{
		SegmentType: segmentType,
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
