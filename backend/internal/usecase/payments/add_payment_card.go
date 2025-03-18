package useCasePayments

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
)

func (uc *UseCasePayments) AddUserPaymentCard(ctx context.Context, userFullCard entity.UserFullPaymentCard) (uuid.UUID, error) {
	log.Printf("usecasePayments.AddUserPaymentCard: user uid %s", userFullCard.UserUid)

	if uuid.Equal(uuid.UUID{}, userFullCard.UserUid) {
		log.Printf("failed to add user payment card: user uid is empty")
		return uuid.UUID{}, errors.New("ID юзера пуст")
	}
	if len(userFullCard.Number) != 16 {
		log.Printf("failed to add user payment card: invalid card number")
		return uuid.UUID{}, errors.New("неправильный формат номера карты")
	}

	if err := validateCardExpired(userFullCard.Expired); err != nil {
		log.Printf("failed to add user payment card: %v", err)
		return uuid.UUID{}, err
	}

	if len(userFullCard.CVV) != 3 {
		log.Printf("failed to add user payment card: invalid card cvv")
		return uuid.UUID{}, errors.New("неправильный формат СVV")
	}

	_, isFound, err := uc.usersService.GetUser(ctx, userFullCard.UserUid)
	if err != nil {
		log.Printf("failed to add user (%s) payment card: %v", userFullCard.UserUid, err)
		return uuid.UUID{}, errors.New("ошибка получения пользователя")
	}

	if !isFound {
		log.Printf("failed to add user (%s) payment card: user not found", userFullCard.UserUid)
		return uuid.UUID{}, errors.New("пользователь с таким ID не найден")
	}

	userFullCard.Uid = uuid.NewV4()

	userCard := entity.UserPaymentCard{
		Uid:         userFullCard.Uid,
		UserUid:     userFullCard.UserUid,
		ExternalUid: uuid.NewV4(), // пока рандом
		Number:      string([]byte(userFullCard.Number[len(userFullCard.Number)-4:])),
		Expired:     userFullCard.Expired,
	}

	if err := uc.paymentsRepository.CreateUserPaymentCard(ctx, userCard); err != nil {
		log.Printf("failed to add user (%s) payment card: %v", userFullCard.UserUid, err)
		return uuid.UUID{}, errors.WithStack(err)
	}

	return userFullCard.Uid, nil
}

func validateCardExpired(expired string) error {
	fmt.Println(expired)
	if len(expired) != 5 {
		return errors.New("неправильный формат даты окончания действия карты")
	}

	expiredParts := strings.Split(expired, "/")
	if len(expiredParts) != 2 {
		return errors.New("неправильный формат даты окончания действия карты")
	}

	month, err := strconv.Atoi(expiredParts[0])
	if err != nil || month <= 0 || month > 12 {
		log.Printf("неправильный формат даты окончания действия карты: недествительный месяц: %v", err)
		return errors.New("неправильный формат даты окончания действия карты: недествительный месяц")
	}
	yearSuffix, err := strconv.Atoi(expiredParts[1])
	if err != nil {
		log.Printf("неправильный формат даты окончания действия карты: недествительный год: %v", err)
		return errors.New("неправильный формат даты окончания действия карты: недествительный год")
	}

	bracketYearSuffix := time.Now().Add(10*365*24*time.Hour).Year() % 100

	fmt.Println("YEAR SUFFIX: ", yearSuffix)
	fmt.Println("BRACKET: ", bracketYearSuffix)

	if yearSuffix > bracketYearSuffix {
		log.Printf("неправильный формат даты окончания действия карты: недествительный год")
		return errors.New("неправильный формат даты окончания действия карты: недествительный год: срок действия карты должен быть меньше 10 лет")
	}
	if yearSuffix < time.Now().Year()%100 {
		log.Printf("неправильный формат даты окончания действия карты: недествительный год")
		return errors.New("неправильный формат даты окончания действия карты: недествительный год: истек срок действия")
	}

	return nil
}
