package rule_engine

import (
	"github.com/hyperjumptech/grule-rule-engine/ast"
	"github.com/hyperjumptech/grule-rule-engine/builder"
	"github.com/hyperjumptech/grule-rule-engine/engine"
	"github.com/hyperjumptech/grule-rule-engine/pkg"
)

// knowledgelibrary is a container for rule-related information.
var knowledgelibrary = *ast.NewKnowledgeLibrary()

type RuleInput interface {
	DataKey() string
}

type RuleOutput interface {
	DataKey() string
}

type RuleConfig interface {
	RuleName() string
	RuleInput() RuleInput
	RuleOutput() RuleOutput
}

type RuleEngineSvc struct {
	// This struct doesn't have any fields
	//but serves as a receiver for methods
	//related to the rule engine service.
}

func NewRuleEngineSvc() *RuleEngineSvc {
	buildRuleEngine()
	return &RuleEngineSvc{}
	//this is setting up rule engine
}

func buildRuleEngine() {
	ruleBuilder := builder.NewRuleBuilder(&knowledgelibrary)

	ruleFile := pkg.NewFileResource("rules.grl")
	err := ruleBuilder.BuildRuleFromResource("Rules", "0.0.1", ruleFile)

	if err != nil {
		panic(err)
	}

	//initializes a rule builder using knowledge library
	//then build rules written in the file rules.grl

}

func (svc *RuleEngineSvc) Execute(ruleconf RuleConfig) error {
	knowledgeBase, _ := knowledgelibrary.NewKnowledgeBaseInstance("Rules", "0.0.1")
	dataCtx := ast.NewDataContext()

	err := dataCtx.Add(ruleconf.RuleInput().DataKey(), ruleconf.RuleInput())
	if err != nil {
		return err
	}

	err = dataCtx.Add(ruleconf.RuleOutput().DataKey(), ruleconf.RuleOutput())
	if err != nil {
		return err
	}

	ruleEngine := engine.NewGruleEngine()
	err = ruleEngine.Execute(dataCtx, knowledgeBase)
	if err != nil {
		return err
	}
	return nil
}
