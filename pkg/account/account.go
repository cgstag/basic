package account

import (
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/Pallinder/go-randomdata"

	"github.com/google/uuid"

	"github.com/labstack/echo/v4"
)

type Account struct {
	UUID      string    `json:"uuid" dynamo:"UUID,hash" index:"Seq-ID-index,range"`
	Client    Client    `json:"page" dynamo:"Client"`
	Segment   Segment   `json:"accountSegment" dynamo:"Segment"`
	Balance   Balance   `json:"balance" dynamo:"Balance"`
	CreatedAt time.Time `json:"createdAt" dynamo:"CreatedAt,range"`
	UpdatedAt time.Time `json:"balance" dynamo:"UpdatedAt"`
	Seq       int64     `localIndex:"ID-Seq-index,range" index:"Seq-ID-index,hash"`
}

type Client struct {
	UUID    string `json:"uuid"`
	Name    string `json:"name"`
	Surname string `json:"page"`
}

type Segment struct {
	Type string `json:"type"`
}

// Generate a random Bank Account
func (r *resource) random(c echo.Context) error {
	account, err := NewRandomAccount()
	if err != nil {
		r.log.Error("Error generating random account")
		c.JSON(http.StatusInternalServerError, err.Error())
	}
	if account != nil && account.isBlacklisted() {
		return c.JSON(http.StatusForbidden, account)
	}
	return c.JSON(http.StatusOK, account)
}

/**
 * Create a Bank Account
 * @param Account
 */
func (r *resource) create(c echo.Context) error {
	// Bind Account Payload
	acc := new(Account)
	if err := c.Bind(acc); err != nil {
		return c.JSON(http.StatusBadRequest, "Couldn't bind JSON payload to account Struct")
	}

	// Verify Data
	if !acc.isValid() {
		return c.JSON(http.StatusBadRequest, "This account is invalid")
	}
	if acc.isBlacklisted() {
		return c.JSON(http.StatusForbidden, "This account is blacklisted")
	}

	// Create Account

	return c.JSON(http.StatusOK, "Account created")
}

/**
 * Read a Bank Account
 * @param :uuid string
 */
func (r *resource) read(c echo.Context) error {
	return c.JSON(http.StatusOK, "Account Detail")
}

/**
 * Update a Bank Account
 * @param Account
 */
func (r *resource) update(c echo.Context) error {
	return c.JSON(http.StatusOK, "Account Detail")
}

/**
 * Delete a Bank Account
 * @param :uuid string
 */
func (r *resource) delete(c echo.Context) error {
	return c.JSON(http.StatusOK, "Consolidated Position")
}

func NewRandomAccount() (*Account, error) {
	account := new(Account)
	account.UUID = uuid.New().String()
	account.Client = Client{
		UUID:    uuid.New().String(),
		Name:    randomdata.FirstName(2),
		Surname: randomdata.LastName(),
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
