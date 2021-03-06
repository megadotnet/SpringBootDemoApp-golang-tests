package main

import (
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"testing"
	sw "testingappdemo/swagger"
	"time"

	"github.com/stretchr/testify/assert"
	"golang.org/x/net/context"
)

const ACCOUNTPREFIX string = "account"

//HTTP GET /api/account
func TestGetAccount(t *testing.T) {
	timen := time.Now() //It will return time.Time object with current timestamp
	resp, r, err := GetAccount(t, timen.UnixNano())
	assert := assert.New(t)
	if err != nil {
		t.Errorf("Error while get GetAccount api ")
		t.Log(err)
	} else {
		assert.NotNil(resp, "Result should  not be null")
		assert.Equal("admin", resp.Login, "incorrect login account name.")
	}
	if r.StatusCode != 200 {
		t.Log(err)
	}
}

//Register Account
func TestRegisterAccount(t *testing.T) {
	timen := time.Now() //It will return time.Time object with current timestamp

	resp, r, err := CreateRegisterProc(t, timen.UnixNano())
	assert := assert.New(t)
	if err != nil {
		t.Errorf("Error while TestRegisterAccount api")
		t.Log(err)
	} else {
		assert.NotNil(resp, "Result should not be null")
	}
	if r.StatusCode != 201 {
		t.Log(err)
	}
}

//create random customer accout
//response body if success:
//HTTP/1.1 201 Created
//Content-Length: 0
func CreateRegisterProc(t *testing.T, timen int64) (sw.ResponseEntity, *http.Response, error) {

	registerModel := sw.ManagedUserVm{
		Email:     RandStringBytesMaskImprSrcSB(16) + strconv.FormatInt(int64(time.Nanosecond)*timen/int64(time.Microsecond), 10) + "@gmail.com",
		Password:  "testpwd1234",
		Login:     RandStringBytesMaskImprSrcSB(16) + strconv.FormatInt(int64(time.Nanosecond)*timen/int64(time.Microsecond), 10),
		ImageUrl:  "IMAGE.png",
		LangKey:   "en-us",
		LastName:  "LastName",
		FirstName: "FirstName"}

	t.Log(registerModel.Login)

	return client.UserAccountControllerApi.RegisterAccountUsingPOST(context.Background(), registerModel)
}

/**
 * POST /account : update the current user information.
 *
 * @param userDTO
 *            the current user information
 * @return the ResponseEntity with status 200 (OK), or status 400 (Bad
 *         Request) or 500 (Internal Server Error) if the user couldn't be
 *         updated
 */
func TestSaveUserAccount(t *testing.T) {

	userDtoModel := sw.UserDto{
		Email:     "admin@gmail.com",
		Login:     "admin",
		ImageUrl:  "IMAGE1.png",
		LangKey:   "en-us",
		LastName:  "LastName",
		FirstName: "FirstName"}

	resp, r, err := client.UserAccountControllerApi.SaveAccountUsingPOST(context.Background(), userDtoModel)
	assert := assert.New(t)
	if err != nil {
		t.Errorf("Error while TestSaveUserAccount api")
		t.Log(err)
	} else {
		assert.NotNil(resp, "Result should not be null")
	}
	if r.StatusCode != 200 {
		t.Log(err)
	}
}

/**
 * POST /account/change_password : changes the current user's password
 *
 * @param password
 *            the new password
 * @return the ResponseEntity with status 200 (OK), or status 400 (Bad
 *         Request) if the new password is not strong enough
 */
func TestChangePassword(t *testing.T) {

	newPwd := "admin"

	resp, r, err := client.UserAccountControllerApi.ChangePasswordUsingPOST(context.Background(), newPwd)
	assert := assert.New(t)
	if err != nil {
		t.Errorf("Error while TestChangePassword api")
		t.Log(err)
	} else {
		assert.NotNil(resp, "Result should not be null")
	}
	if r.StatusCode != 200 {
		t.Log(err)
	}
}

//Get Account API Concurrency
func TestGetAccountConcurrency(t *testing.T) {
	errc := make(chan error)
	sum := 0
	timen := time.Now() //It will return time.Time object with current timestamp
	for i := 0; i <= 100; i++ {
		go func() {
			resp, r, err := GetAccount(t, timen.UnixNano())
			if err != nil {
				t.Errorf("Error while get account")
				t.Log(err, r, resp)
			}
			errc <- err
		}()
		sum++

	}
	waitOnFunctions(t, errc, sum)
}

/**
 * POST /account/reset_password/init : Send an email to reset the password
 * of the user
 *
 * @param mail
 *            the mail of the user
 * @return the ResponseEntity with status 200 (OK) if the email was sent, or
 *         status 400 (Bad Request) if the email address is not registered
 */
func TestRequestPasswordReset(t *testing.T) {

	testEmail := "admin@gmail.com"

	resp, r, err := client.UserAccountControllerApi.RequestPasswordResetUsingPOST(context.Background(), testEmail)
	assert := assert.New(t)
	if err != nil {
		t.Errorf("Error while RequestPasswordReset api")
		t.Log(err)
	} else {
		assert.NotNil(resp, "Result should not be null")
	}
	if r.StatusCode != 200 {
		t.Log(err)
	}
}

func GetAccount(t *testing.T, timen int64) (sw.UserDto, *http.Response, error) {

	return client.UserAccountControllerApi.GetAccountUsingGET(context.Background())
}

func GeneratePhoneNumber() string {
	return ACCOUNTPREFIX + strconv.Itoa(GenerateRandInt(1, 9999999))
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

var src = rand.NewSource(time.Now().UnixNano())

func RandStringBytesMaskImprSrcSB(n int) string {
	sb := strings.Builder{}
	sb.Grow(n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			sb.WriteByte(letterBytes[idx])
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return sb.String()
}

func GenerateRandInt(min, max int) int {
	rand.Seed(time.Now().UnixNano()) //????????????
	return rand.Intn(max-min) + min
}
