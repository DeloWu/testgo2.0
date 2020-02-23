package models

type Plan struct {
    BaseModel `bson:",inline"`
    PlanRunSuite []string `bson:"planRunSuite" json:"planRunSuite"`
    PlanPeriod string `bson:"planPeriod" json:"planPeriod"`
    PlanCreator string `bson:"planCreator" json:"planCreator"`

}
