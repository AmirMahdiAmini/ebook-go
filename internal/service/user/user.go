package user

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"time"

	"github.com/amirmahdiamini/ebook-go/internal/dto"
	"github.com/amirmahdiamini/ebook-go/internal/model"
	"github.com/amirmahdiamini/ebook-go/internal/repository"
	"github.com/amirmahdiamini/ebook-go/pkg"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

type (
	UserService interface {
		Signup(user *model.User) error
		Signin(data *dto.Signin) (*string, error)
		VerifyAccount(verification *dto.VerifyAccount) error
		ForgotPassword(data *dto.ForgotPassword) error
		ChangePassword(data *dto.ChangePassowrd) error
	}
	userService struct {
		userRepository repository.UserRepository
		codeRepository repository.CodeRepository
		JWT            pkg.JWTService
	}
)

func NewUserService(userRepository repository.UserRepository, codeRepository repository.CodeRepository, JWT pkg.JWTService) UserService {
	return &userService{
		userRepository: userRepository,
		codeRepository: codeRepository,
		JWT:            JWT,
	}
}

func (ctx *userService) Signup(user *model.User) error {
	match, err := regexp.MatchString("^[A-Za-z]{3,}[A-Za-z0-9]{2,}$", user.UserName)
	if err != nil {
		return errors.New("invalid Username")
	}
	if !match {
		return errors.New("invalid Username")
	}
	username, _ := ctx.userRepository.FindByUsername(user.UserName)
	if username != nil {
		return errors.New("username already exists")
	}
	phone, _ := ctx.userRepository.FindByPhone(user.Phone)
	if phone != nil {
		return errors.New("phone already exists")
	}
	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("something went wrong #111")
	}
	user.Password = string(password)
	user.SID = pkg.SID(35)
	user.Created_at = time.Now().Format("2006/01/02 15:04:05")
	user.Is_verified = false
	user.Notifications = []string{fmt.Sprint(time.Now().Format("2006/01/02 15:04:05"), " ", "you created your account")}
	if err := ctx.userRepository.Insert(user); err != nil {
		return err
	}
	code := pkg.VerifyCode(234456, 987564)
	if err := ctx.codeRepository.SetCode(user.Phone, code, 2*time.Minute); err != nil {
		return err
	}
	if err := pkg.NewSMS(fmt.Sprint("کد تایید : \n", code), user.Phone).SendSMS(); err != nil {
		return err
	}
	return nil
}
func (ctx *userService) Signin(data *dto.Signin) (*string, error) {
	user, err := ctx.userRepository.FindByUsernameOrPhone(data.Username_or_phone)
	if err != nil {
		return nil, err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data.Password)); err != nil {
		return nil, errors.New("invalid password")
	}
	if !user.Is_verified {
		code := pkg.VerifyCode(234456, 987564)
		if err := ctx.codeRepository.SetCode(user.Phone, code, 2*time.Minute); err != nil {
			return nil, err
		}
		if err := pkg.NewSMS(fmt.Sprint("کد تایید : \n", code), user.Phone).SendSMS(); err != nil {
			return nil, err
		}
		return nil, errors.New("user is not verified")
	}
	ctx.userRepository.Notification(user.Phone, fmt.Sprint(time.Now().Format("2006/01/02 15:04:05"), " ", "you logged in"))
	token := ctx.JWT.GenerateToken(user.SID)
	return &token, nil
}

func (ctx *userService) VerifyAccount(verification *dto.VerifyAccount) error {
	code, err := ctx.codeRepository.GetCode(verification.Phone)
	if err != nil {
		return err
	}
	if code != verification.Code {
		return errors.New("invalid code")
	}
	user, err := ctx.userRepository.FindByPhone(verification.Phone)
	if err != nil {
		return err
	}
	if user.Is_verified {
		ctx.codeRepository.DeleteCode(verification.Phone)
		authorization_code := pkg.ChangePassword()
		i, err := strconv.Atoi(authorization_code)
		if err != nil {
			return err
		}

		if err := ctx.codeRepository.SetCode(user.Phone, int32(i), 2*time.Minute); err != nil {
			fmt.Println(err)
			return err
		}
		return nil
	}
	if err := ctx.userRepository.VerifyAccount(verification.Phone); err != nil {
		return err
	}
	ctx.codeRepository.DeleteCode(verification.Phone)
	return nil
}

func (ctx *userService) ForgotPassword(data *dto.ForgotPassword) error {
	user, err := ctx.userRepository.FindByPhone(data.Phone)
	if err != nil {
		return err
	}
	code := pkg.VerifyCode(234456, 987564)
	if err := ctx.codeRepository.SetCode(user.Phone, code, 2*time.Minute); err != nil {
		return err
	}
	if err := pkg.NewSMS(fmt.Sprint("تغییر رمز : \n", code), user.Phone).SendSMS(); err != nil {
		return err
	}
	return nil
}
func (ctx *userService) ChangePassword(data *dto.ChangePassowrd) error {
	user, err := ctx.userRepository.FindByPhone(data.Phone)
	if err != nil {
		return err
	}
	code, err := ctx.codeRepository.GetCode(user.Phone)
	if err != nil {
		return err
	}
	if code != data.Code {
		return errors.New("invalid code")
	}
	password, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("something went wrong #111")
	}
	user.Password = string(password)
	if err := ctx.userRepository.UpdateOne(bson.M{"phone": user.Phone}, bson.M{"$set": bson.M{"password": user.Password}}); err != nil {
		return err
	}
	ctx.codeRepository.DeleteCode(user.Phone)
	return nil
}
