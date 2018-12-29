package gov

import (
	"fmt"

	"github.com/irisnet/irishub/codec"
	"github.com/irisnet/irishub/modules/params"
	stakeTypes "github.com/irisnet/irishub/modules/stake/types"
	sdk "github.com/irisnet/irishub/types"
	"strconv"
	"time"
)

const (
	CRITICAL_DEPOSIT     = 4000
	IMPORTANT_DEPOSIT    = 2000
	NORMAL_DEPOSIT       = 1000
	CRITICAL             = "Critical"
	IMPORTANT            = "Important"
	NORMAL               = "Normal"
	LOWER_BOUND_AMOUNT   = 10
	UPPER_BOUND_AMOUNT   = 10000
	THREE_DAYS           = 3 * 3600 * 24
	SIXTY_HOURS          = 60 * 3600
	TWO_DAYS             = 2 * 3600 * 24 //
	ONE_DAY              = 1 * 3600 * 24
	STABLE_CRITIACAL_NUM = 1
	MIN_IMPORTANT_NUM    = 2
	MIN_NORMAL_NUM       = 1
)

var _ params.ParamSet = (*GovParams)(nil)

// default paramspace for params keeper
const (
	DefaultParamSpace = "gov"
)

//Parameter store key
var (
	KeyCriticalDepositPeriod = []byte(CRITICAL + "DepositPeriod")
	KeyCriticalMinDeposit    = []byte(CRITICAL + "MinDeposit")
	KeyCriticalVotingPeriod  = []byte(CRITICAL + "VotingPeriod")
	KeyCriticalMaxNum        = []byte(CRITICAL + "MaxNum")
	KeyCriticalThreshold     = []byte(CRITICAL + "Threshold")
	KeyCriticalVeto          = []byte(CRITICAL + "Veto")
	KeyCriticalParticipation = []byte(CRITICAL + "Participation")
	KeyCriticalPenalty       = []byte(CRITICAL + "Penalty")

	KeyImportantDepositPeriod = []byte(IMPORTANT + "DepositPeriod")
	KeyImportantMinDeposit    = []byte(IMPORTANT + "MinDeposit")
	KeyImportantVotingPeriod  = []byte(IMPORTANT + "VotingPeriod")
	KeyImportantMaxNum        = []byte(IMPORTANT + "MaxNum")
	KeyImportantThreshold     = []byte(IMPORTANT + "Threshold")
	KeyImportantVeto          = []byte(IMPORTANT + "Veto")
	KeyImportantParticipation = []byte(IMPORTANT + "Participation")
	KeyImportantPenalty       = []byte(IMPORTANT + "Penalty")

	KeyNormalDepositPeriod = []byte(NORMAL + "DepositPeriod")
	KeyNormalMinDeposit    = []byte(NORMAL + "MinDeposit")
	KeyNormalVotingPeriod  = []byte(NORMAL + "VotingPeriod")
	KeyNormalMaxNum        = []byte(NORMAL + "MaxNum")
	KeyNormalThreshold     = []byte(NORMAL + "Threshold")
	KeyNormalVeto          = []byte(NORMAL + "Veto")
	KeyNormalParticipation = []byte(NORMAL + "Participation")
	KeyNormalPenalty       = []byte(NORMAL + "Penalty")
)

// ParamTable for mint module
func ParamTypeTable() params.TypeTable {
	return params.NewTypeTable().RegisterParamSet(&GovParams{})
}

// mint parameters
type GovParams struct {
	CriticalDepositPeriod time.Duration `json:"critical_deposit_period"` //  Maximum period for Atom holders to deposit on a proposal. Initial value: 2 months
	CriticalMinDeposit    sdk.Coins     `json:"critical_min_deposit"`    //  Minimum deposit for a critical proposal to enter voting period.
	CriticalVotingPeriod  time.Duration `json:"critical_voting_period"`  //  Length of the critical voting period.
	CriticalMaxNum        uint64        `json:"critical_max_num"`
	CriticalThreshold     sdk.Dec       `json:"critical_threshold"`     //  Minimum propotion of Yes votes for proposal to pass. Initial value: 0.5
	CriticalVeto          sdk.Dec       `json:"critical_veto"`          //  Minimum value of Veto votes to Total votes ratio for proposal to be vetoed. Initial value: 1/3
	CriticalParticipation sdk.Dec       `json:"critical_participation"` //
	CriticalPenalty       sdk.Dec       `json:"critical_penalty"`       //  Penalty if validator does not vote

	ImportantDepositPeriod time.Duration `json:"important_deposit_period"` //  Maximum period for Atom holders to deposit on a proposal. Initial value: 2 months
	ImportantMinDeposit    sdk.Coins     `json:"important_min_deposit"`    //  Minimum deposit for a important proposal to enter voting period.
	ImportantVotingPeriod  time.Duration `json:"important_voting_period"`  //  Length of the important voting period.
	ImportantMaxNum        uint64        `json:"important_max_num"`
	ImportantThreshold     sdk.Dec       `json:"important_threshold"`     //  Minimum propotion of Yes votes for proposal to pass. Initial value: 0.5
	ImportantVeto          sdk.Dec       `json:"important_veto"`          //  Minimum value of Veto votes to Total votes ratio for proposal to be vetoed. Initial value: 1/3
	ImportantParticipation sdk.Dec       `json:"important_participation"` //
	ImportantPenalty       sdk.Dec       `json:"important_penalty"`       //  Penalty if validator does not vote

	NormalDepositPeriod time.Duration `json:"normal_deposit_period"` //  Maximum period for Atom holders to deposit on a proposal. Initial value: 2 months
	NormalMinDeposit    sdk.Coins     `json:"normal_min_deposit"`    //  Minimum deposit for a normal proposal to enter voting period.
	NormalVotingPeriod  time.Duration `json:"normal_voting_period"`  //  Length of the normal voting period.
	NormalMaxNum        uint64        `json:"normal_max_num"`
	NormalThreshold     sdk.Dec       `json:"normal_threshold"`     //  Minimum propotion of Yes votes for proposal to pass. Initial value: 0.5
	NormalVeto          sdk.Dec       `json:"normal_veto"`          //  Minimum value of Veto votes to Total votes ratio for proposal to be vetoed. Initial value: 1/3
	NormalParticipation sdk.Dec       `json:"normal_participation"` //
	NormalPenalty       sdk.Dec       `json:"normal_penalty"`       //  Penalty if validator does not vote
}

// Implements params.ParamStruct
func (p *GovParams) GetParamSpace() string {
	return DefaultParamSpace
}

func (p *GovParams) KeyValuePairs() params.KeyValuePairs {
	return params.KeyValuePairs{
		{KeyCriticalDepositPeriod, &p.CriticalDepositPeriod},
		{KeyCriticalMinDeposit, &p.CriticalMinDeposit},
		{KeyCriticalVotingPeriod, &p.CriticalVotingPeriod},
		{KeyCriticalMaxNum, &p.CriticalMaxNum},
		{KeyCriticalThreshold, &p.CriticalThreshold},
		{KeyCriticalVeto, &p.CriticalVeto},
		{KeyCriticalParticipation, &p.CriticalParticipation},
		{KeyCriticalPenalty, &p.CriticalPenalty},

		{KeyImportantDepositPeriod, &p.ImportantDepositPeriod},
		{KeyImportantMinDeposit, &p.ImportantMinDeposit},
		{KeyImportantVotingPeriod, &p.ImportantVotingPeriod},
		{KeyImportantMaxNum, &p.ImportantMaxNum},
		{KeyImportantThreshold, &p.ImportantThreshold},
		{KeyImportantVeto, &p.ImportantVeto},
		{KeyImportantParticipation, &p.ImportantParticipation},
		{KeyImportantPenalty, &p.ImportantPenalty},

		{KeyNormalDepositPeriod, &p.NormalDepositPeriod},
		{KeyNormalMinDeposit, &p.NormalMinDeposit},
		{KeyNormalVotingPeriod, &p.NormalVotingPeriod},
		{KeyNormalMaxNum, &p.NormalMaxNum},
		{KeyNormalThreshold, &p.NormalThreshold},
		{KeyNormalVeto, &p.NormalVeto},
		{KeyNormalParticipation, &p.NormalParticipation},
		{KeyNormalPenalty, &p.NormalPenalty},
	}
}

func (p *GovParams) Validate(key string, value string) (interface{}, sdk.Error) {
	return nil, nil
}

func (p *GovParams) StringFromBytes(cdc *codec.Codec, key string, bytes []byte) (string, error) {
	switch key {
	case string(KeyCriticalDepositPeriod):
		err := cdc.UnmarshalJSON(bytes, &p.CriticalDepositPeriod)
		return p.CriticalDepositPeriod.String(), err
	case string(KeyCriticalMinDeposit):
		err := cdc.UnmarshalJSON(bytes, &p.CriticalMinDeposit)
		return p.CriticalMinDeposit.String(), err
	case string(KeyCriticalVotingPeriod):
		err := cdc.UnmarshalJSON(bytes, &p.CriticalVotingPeriod)
		return p.CriticalDepositPeriod.String(), err
	case string(KeyCriticalMaxNum):
		err := cdc.UnmarshalJSON(bytes, &p.CriticalMaxNum)
		return strconv.FormatUint(p.CriticalMaxNum, 10), err
	case string(KeyCriticalThreshold):
		err := cdc.UnmarshalJSON(bytes, &p.CriticalThreshold)
		return p.CriticalThreshold.String(), err
	case string(KeyCriticalVeto):
		err := cdc.UnmarshalJSON(bytes, &p.CriticalVeto)
		return p.CriticalThreshold.String(), err
	case string(KeyCriticalParticipation):
		err := cdc.UnmarshalJSON(bytes, &p.CriticalParticipation)
		return p.CriticalParticipation.String(), err
	case string(KeyCriticalPenalty):
		err := cdc.UnmarshalJSON(bytes, &p.CriticalPenalty)
		return p.CriticalPenalty.String(), err
	default:
		return "", fmt.Errorf("%s is not existed", key)
	}
}

// default minting module parameters
func DefaultParams() GovParams {
	var criticalMinDeposit, _ = sdk.NewDefaultCoinType(stakeTypes.StakeDenomName).ConvertToMinCoin(fmt.Sprintf("%d%s", CRITICAL_DEPOSIT, stakeTypes.StakeDenomName))
	var importantMinDeposit, _ = sdk.NewDefaultCoinType(stakeTypes.StakeDenomName).ConvertToMinCoin(fmt.Sprintf("%d%s", IMPORTANT_DEPOSIT, stakeTypes.StakeDenomName))
	var normalMinDeposit, _ = sdk.NewDefaultCoinType(stakeTypes.StakeDenomName).ConvertToMinCoin(fmt.Sprintf("%d%s", NORMAL_DEPOSIT, stakeTypes.StakeDenomName))

	return GovParams{
		CriticalDepositPeriod: time.Duration(ONE_DAY) * time.Second,
		CriticalMinDeposit:    sdk.Coins{criticalMinDeposit},
		CriticalVotingPeriod:  time.Duration(THREE_DAYS) * time.Second,
		CriticalMaxNum:        STABLE_CRITIACAL_NUM,
		CriticalThreshold:     sdk.NewDecWithPrec(834, 3),
		CriticalVeto:          sdk.NewDecWithPrec(334, 3),
		CriticalParticipation: sdk.NewDecWithPrec(8572, 4),
		CriticalPenalty:       sdk.NewDecWithPrec(9, 4),

		ImportantDepositPeriod: time.Duration(ONE_DAY) * time.Second,
		ImportantMinDeposit:    sdk.Coins{importantMinDeposit},
		ImportantVotingPeriod:  time.Duration(SIXTY_HOURS) * time.Second,
		ImportantMaxNum:        MIN_IMPORTANT_NUM,
		ImportantThreshold:     sdk.NewDecWithPrec(8, 1),
		ImportantVeto:          sdk.NewDecWithPrec(334, 3),
		ImportantParticipation: sdk.NewDecWithPrec(834, 3),
		ImportantPenalty:       sdk.NewDecWithPrec(7, 4),

		NormalDepositPeriod: time.Duration(ONE_DAY) * time.Second,
		NormalMinDeposit:    sdk.Coins{normalMinDeposit},
		NormalVotingPeriod:  time.Duration(TWO_DAYS) * time.Second,
		NormalMaxNum:        MIN_NORMAL_NUM,
		NormalThreshold:     sdk.NewDecWithPrec(667, 3),
		NormalVeto:          sdk.NewDecWithPrec(334, 3),
		NormalParticipation: sdk.NewDecWithPrec(75, 2),
		NormalPenalty:       sdk.NewDecWithPrec(5, 4),
	}
}

func validateParams(p GovParams) sdk.Error {
	if err := validateDepositProcedure(DepositProcedure{
		MaxDepositPeriod: p.CriticalDepositPeriod,
		MinDeposit:       p.CriticalMinDeposit,
	}, CRITICAL); err != nil {
		return err
	}

	if err := validatorVotingProcedure(VotingProcedure{
		VotingPeriod:p.CriticalVotingPeriod,
		MaxNum:p.CriticalMaxNum,
	},CRITICAL,STABLE_CRITIACAL_NUM); err!=nil{
			return  err
	}

	if err := validateTallyingProcedure(TallyingProcedure{
		Threshold:     p.CriticalThreshold,
		Veto:          p.CriticalVeto,
		Participation: p.CriticalParticipation,
		Penalty:       p.CriticalPenalty,
	}, CRITICAL); err !=nil{
		return err
	}

	if err := validateDepositProcedure(DepositProcedure{
		MaxDepositPeriod: p.ImportantDepositPeriod,
		MinDeposit:       p.ImportantMinDeposit,
	}, IMPORTANT); err != nil {
		return err
	}

	if err := validatorVotingProcedure(VotingProcedure{
		VotingPeriod:p.ImportantVotingPeriod,
		MaxNum:p.ImportantMaxNum,
	},IMPORTANT,MIN_IMPORTANT_NUM); err!=nil{
		return  err
	}

	if err := validateTallyingProcedure(TallyingProcedure{
		Threshold:     p.ImportantThreshold,
		Veto:          p.ImportantVeto,
		Participation: p.ImportantParticipation,
		Penalty:       p.ImportantPenalty,
	}, IMPORTANT); err !=nil{
		return err
	}

	if err := validateDepositProcedure(DepositProcedure{
		MaxDepositPeriod: p.NormalDepositPeriod,
		MinDeposit:       p.NormalMinDeposit,
	}, NORMAL); err != nil {
		return err
	}

	if err := validatorVotingProcedure(VotingProcedure{
		VotingPeriod:p.NormalVotingPeriod,
		MaxNum:p.NormalMaxNum,
	},NORMAL,MIN_NORMAL_NUM); err!=nil{
		return  err
	}

	if err := validateTallyingProcedure(TallyingProcedure{
		Threshold:     p.NormalThreshold,
		Veto:          p.NormalVeto,
		Participation: p.NormalParticipation,
		Penalty:       p.NormalPenalty,
	}, NORMAL); err !=nil{
		return err
	}

	return nil
}

//______________________________________________________________________

// get inflation params from the global param store
func (k Keeper) GetParamSet(ctx sdk.Context) GovParams {
	var params GovParams
	k.paramSpace.GetParamSet(ctx, &params)
	return params
}

// set inflation params from the global param store
func (k Keeper) SetParamSet(ctx sdk.Context, params GovParams) {
	k.paramSpace.SetParamSet(ctx, &params)
}

type DepositProcedure struct {
	MinDeposit       sdk.Coins
	MaxDepositPeriod time.Duration
}

type VotingProcedure struct {
	VotingPeriod time.Duration `json:"critical_voting_period"` //  Length of the critical voting period.
	MaxNum       uint64        `json:"normal_max_num"`
}

type TallyingProcedure struct {
	Threshold     sdk.Dec `json:"threshold"`     //  Minimum propotion of Yes votes for proposal to pass. Initial value: 0.5
	Veto          sdk.Dec `json:"veto"`          //  Minimum value of Veto votes to Total votes ratio for proposal to be vetoed. Initial value: 1/3
	Participation sdk.Dec `json:"participation"` //
	Penalty       sdk.Dec `json:"penalty"`       //  Penalty if validator does not vote
}

func validateDepositProcedure(dp DepositProcedure, str string) sdk.Error {
	if dp.MinDeposit[0].Denom != stakeTypes.StakeDenom {
		return sdk.NewError(params.DefaultCodespace, params.CodeInvalidMinDepositDenom, fmt.Sprintf(str+"MinDeposit should be %s!", stakeTypes.StakeDenom))
	}

	LowerBound, _ := sdk.NewDefaultCoinType(stakeTypes.StakeDenomName).ConvertToMinCoin(fmt.Sprintf("%d%s", LOWER_BOUND_AMOUNT, stakeTypes.StakeDenomName))
	UpperBound, _ := sdk.NewDefaultCoinType(stakeTypes.StakeDenomName).ConvertToMinCoin(fmt.Sprintf("%d%s", UPPER_BOUND_AMOUNT, stakeTypes.StakeDenomName))

	if dp.MinDeposit[0].Amount.LT(LowerBound.Amount) || dp.MinDeposit[0].Amount.GT(UpperBound.Amount) {
		return sdk.NewError(params.DefaultCodespace, params.CodeInvalidMinDepositAmount, fmt.Sprintf(str+"MinDepositAmount"+dp.MinDeposit[0].String()+" should be larger than 10iris and less than 10000iris"))
	}

	if dp.MaxDepositPeriod.Seconds() < 20 || dp.MaxDepositPeriod.Seconds() > THREE_DAYS {
		return sdk.NewError(params.DefaultCodespace, params.CodeInvalidDepositPeriod, fmt.Sprintf(str+"MaxDepositPeriod (%s) should be between 20s and %ds", strconv.Itoa(int(dp.MaxDepositPeriod.Seconds())), THREE_DAYS))
	}
	return nil
}

func validatorVotingProcedure(vp VotingProcedure, str string , min_num uint64) sdk.Error {
	if vp.VotingPeriod.Seconds() < 20 || vp.VotingPeriod.Seconds() > THREE_DAYS {
		return sdk.NewError(params.DefaultCodespace, params.CodeInvalidVotingPeriod, fmt.Sprintf(str+"VotingPeriod (%s) should be between 20s and %ds", strconv.Itoa(int(vp.VotingPeriod.Seconds())), THREE_DAYS))
	}

	if vp.MaxNum < min_num {
		return sdk.NewError(params.DefaultCodespace, params.CodeInvalidMaxProposalNum, fmt.Sprintf("The num of Max"+str+"Proposal should be no less than %v.", min_num))
	}
	return nil
}
func validateTallyingProcedure(tp TallyingProcedure, str string) sdk.Error {
	if tp.Threshold.LTE(sdk.ZeroDec()) || tp.Threshold.GTE(sdk.NewDec(1)) {
		return sdk.NewError(params.DefaultCodespace, params.CodeInvalidThreshold, fmt.Sprintf("Invalid "+str+" Threshold ( "+tp.Threshold.String()+" ) should be between 0 and 1"))
	}
	if tp.Participation.LTE(sdk.ZeroDec()) || tp.Participation.GTE(sdk.NewDec(1)) {
		return sdk.NewError(params.DefaultCodespace, params.CodeInvalidParticipation, fmt.Sprintf("Invalid "+str+" participation ( "+tp.Participation.String()+" ) should be between 0 and 1"))
	}
	if tp.Veto.LTE(sdk.ZeroDec()) || tp.Veto.GTE(sdk.NewDec(1)) {
		return sdk.NewError(params.DefaultCodespace, params.CodeInvalidVeto, fmt.Sprintf("Invalid "+str+" Veto ( "+tp.Veto.String()+" ) should be between 0 and 1"))
	}
	if tp.Penalty.LTE(sdk.ZeroDec()) || tp.Penalty.GTE(sdk.NewDec(1)) {
		return sdk.NewError(params.DefaultCodespace, params.CodeInvalidGovernancePenalty, fmt.Sprintf("Invalid "+str+" GovernancePenalty ( "+tp.Penalty.String()+" ) should be between 0 and 1"))
	}
	return nil
}
