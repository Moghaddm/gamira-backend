package judgement_result

type Verdict int8
type ValueAssessment int8
type Compatibility int8

const (
	VeryLow  Verdict = 1
	Low      Verdict = 2
	Fair     Verdict = 4
	Moderate Verdict = 8
	Good     Verdict = 16
)

const (
	PoorValue         ValueAssessment = 1
	BelowAverageValue ValueAssessment = 2
	FairValue         ValueAssessment = 4
	GoodValue         ValueAssessment = 8
	GreatValue        ValueAssessment = 16
	ExcellentValue    ValueAssessment = 32
)

const (
	CPUCompatibility     Compatibility = 1
	GPUCompatibility     Compatibility = 2
	MemoryCompatibility  Compatibility = 4
	StorageCompatibility Compatibility = 8
)

type JudgementResult struct {
	ID                         int64               `bson:"_id,omitempty"`
	Score                      int8                `bson:"score"`
	Verdict                    Verdict             `bson:"verdict"`
	Compatibility              []CompatibilitySpec `bson:"compatibility"`
	Strengths                  []string            `bson:"strengths"`
	Weaknesses                 []string            `bson:"weaknesses"`
	ValueAssessment            ValueAssessment     `bson:"value_assessment"`
	UpgradePotential           UpgradePotential    `bson:"upgrade_potential"`
	LongevityYears             int8                `bson:"longevity_years"`
	PersonalizedRecommendation string              `bson:"personalized_recommendation"`
}

type CompatibilitySpec struct {
	Compatibility string `bson:"compatibility"`
	Value         string `bson:"value"` // percentage value
}

type UpgradePotential struct {
	CPU     bool `bson:"cpu"`
	GPU     bool `bson:"gpu"`
	Memory  bool `bson:"memory"`
	Storage bool `bson:"storage"`
}
