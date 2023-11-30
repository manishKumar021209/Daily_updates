package main

import (
	"fmt"
	ruleengine "rule_engine/rule_engine"

	"github.com/hyperjumptech/grule-rule-engine/logger"
)

type User struct {
	Name              string  `json:"name"`
	Username          string  `json:"username"`
	Email             string  `json:"email"`
	Age               int     `json:"age"`
	Gender            string  `json:"gender"`
	TotalOrders       int     `json:"total_orders"`
	AverageOrderValue float64 `json:"average_order_value"`
}

type OfferService interface {
	CheckOfferForUser(user User) bool
}

type OfferServiceClient struct {
	ruleEngineSvc *ruleengine.RuleEngineSvc
}

// CheckOfferForUser implements offerService.
func (svc OfferServiceClient) CheckOfferForUser(user User) bool {
	offerCard := ruleengine.NewUSerOfferContext()
	offerCard.UserOfferInput = &ruleengine.UserOfferInput{
		Name:              user.Name,
		Username:          user.Username,
		Email:             user.Email,
		Gender:            user.Gender,
		Age:               user.Age,
		TotalOrders:       user.TotalOrders,
		AverageOrderValue: user.AverageOrderValue,
	}

	err := svc.ruleEngineSvc.Execute(offerCard)
	if err != nil {
		logger.Log.Error("get user offer rule engine failed", err)
	}
	return offerCard.UserOfferOutput.IsOfferApplicable
}

func NewOfferService(ruleEngineSvc *ruleengine.RuleEngineSvc) OfferService {
	return &OfferServiceClient{
		ruleEngineSvc: ruleEngineSvc,
	}
}

func main() {
	ruleEngineSvc := ruleengine.NewRuleEngineSvc()
	offerSvc := NewOfferService(ruleEngineSvc)

	userA := User{
		Name:              "Harish",
		Username:          "hariya",
		Email:             "bhaihaikahi@gmail.com",
		Gender:            "Male",
		Age:               24,
		TotalOrders:       50,
		AverageOrderValue: 225,
	}

	fmt.Println("offer validity for user A: ", offerSvc.CheckOfferForUser(userA))

	userB := User{
		Name:              "Manish",
		Username:          "mk",
		Email:             "meriemail@gmail.com",
		Gender:            "Male",
		Age:               24,
		TotalOrders:       10,
		AverageOrderValue: 200,
	}

	fmt.Println("offer validity for user B: ", offerSvc.CheckOfferForUser(userB))
}
