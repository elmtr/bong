package bong

type Period struct {
  ID        string `json:"id" bson:"id"`

  // modifiable
  Day       int    `json:"day" bson:"day"`
  Interval  int    `json:"interval" bson:"interval"`
  Room      string `json:"room" bson:"room"`
  Assigned  bool   `json:"assigned" bson:"assigned"`
  
  // indexable
  GradeID   string `json:"gradeID" bson:"gradeID"`
  SubjectID string `json:"subjectID" bson:"subjectID"`
}