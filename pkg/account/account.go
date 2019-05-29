package account

import (
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/guregu/dynamo"

	"github.com/Pallinder/go-randomdata"

	"github.com/google/uuid"

	"github.com/labstack/echo/v4"
)

type Account struct {
	UUID      string    `json:"uuid" dynamo:"UUID,hash"`
	CPF       string    `json:"cpf" index:"CPF,hash"`
	Name      string    `json:"name" dynamo:"Name"`
	Surname   string    `json:"page" dynamo:"Surname"`
	Segment   string    `json:"accountSegment" dynamo:"Segment,range"`
	Balance   float64   `json:"balance" dynamo:"Balance"`
	CreatedAt int64     `json:"createdAt" dynamo:"CreatedAt"`
	UpdatedAt time.Time `json:"balance" dynamo:"UpdatedAt"`
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
	table := r.db.Table("Account")
	err = table.Put(account).Run()
	if err != nil {
		r.log.Error("Error inserting account into DynamoDB")
		c.JSON(http.StatusInternalServerError, err.Error())
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
	result := new(Account)
	table := r.db.Table("Account")
	err := table.Get("UUID", c.Param("uuid")).Range("Segment", dynamo.Equal, "Varejo").One(&result)
	if err != nil {
		r.log.Error("Error inserting account into DynamoDB")
		c.JSON(http.StatusNoContent, err.Error())
	}
	return c.JSON(http.StatusOK, result)
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
	account.UUID = uuid.New().String()
	account.Name = randomdata.FirstName(2)
	account.Surname = randomdata.LastName()
	account.CreatedAt = time.Now().Unix()
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
	account.Segment = segmentType
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
	account.Balance = consolidatedPosition
	return account, nil
}

func (a *Account) isBlacklisted() bool {
	return false
}

func (a *Account) isValid() bool {
	return true
}
