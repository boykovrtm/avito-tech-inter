package facade

import (
	"avito-tech-inter/api"
	"avito-tech-inter/domain/auth"
	"avito-tech-inter/domain/user"
	"context"
	"github.com/jameskeane/bcrypt"
)

type Controller struct {
	userRepo user.Repository
}

func NewController(userRepo user.Repository) *Controller {
	return &Controller{userRepo: userRepo}
}

func (s *Controller) APIAuthPost(ctx context.Context, req *api.AuthRequest) (api.APIAuthPostRes, error) {
	hash, err := s.userRepo.FindHashByUsername(ctx, req.Username)
	if err != nil {
		return nil, err
	}

	if hash == nil {
		return &api.APIAuthPostBadRequest{}, nil
	}

	if bcrypt.Match(req.Password, *hash) {

	} else {
		return &api.APIAuthPostUnauthorized{}, nil
	}

}

// why so POST-ish method is GET
func (s *Controller) APIBuyItemGet(ctx context.Context, params api.APIBuyItemGetParams) (api.APIBuyItemGetRes, error) {
	//Передать товар юзеру
	panic("implement me")
}

func (s *Controller) APIInfoGet(ctx context.Context) (api.APIInfoGetRes, error) {
	//выдать информацию о купленных товарах
	return api.InfoResponse{
		Coins: api.OptInt{Value: 98},
		Inventory: []api.InfoResponseInventoryItem{api.InfoResponseInventoryItem{
			Type:     api.OptString{},
			Quantity: api.OptInt{Value: 0, Set: true},
		}},
		CoinHistory: api.OptInfoResponseCoinHistory{
			Value: api.InfoResponseCoinHistory{
				Received: []api.InfoResponseCoinHistoryReceivedItem{
					api.InfoResponseCoinHistoryReceivedItem{
						FromUser: api.OptString{},
						Amount:   api.OptInt{},
					},
				},
				Sent: nil,
			},
		},
	}, nil
}

func (s *Controller) APISendCoinPost(ctx context.Context, req *api.SendCoinRequest) (api.APISendCoinPostRes, error) {
	// проверка на наличие монет

	// если хватает - списываем

	// не хватает выдаем ошибку
	panic("implement me")
}

func (s *Controller) HandleBearerAuth(ctx context.Context, operationName api.OperationName, t api.BearerAuth) (context.Context, error) {
}
