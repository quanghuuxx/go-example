package model

type Package struct {
	Id                int     `json:"id" bson:"id"`
	ParentId          *int    `json:"parent_id,omitempty" bson:"parent_id"`
	Title             string  `json:"title" bson:"title"`
	FullTitle         string  `json:"full_title" bson:"full_title"`
	Slug              string  `json:"slug" bson:"slug"`
	AssetPhoto        *string `json:"asset_photo,omitempty" bson:"asset_photo"`
	NetPhoto          *string `json:"net_photo,omitempty" bson:"net_photo"`
	LearningProgramId *int    `json:"-" bson:"learning_program_id"`
	SubjectTypeId     *int    `json:"-" bson:"subject_type_id"`
}

func (p *Package) TableName() string {
	return "package"
}

type LearningProgram struct {
	Id         int    `json:"id" bson:"id"`
	Title      string `json:"title" bson:"title"`
	Slug       string `json:"slug" bson:"slug" db:"slug"`
	AssetPhoto string `json:"asset_photo,omitempty" bson:"asset_photo"`
	NetPhoto   string `json:"net_photo,omitempty" bson:"net_photo"`
}

func (p *LearningProgram) TableName() string {
	return "learning_program"
}

type SubjectType struct {
	Id          int    `json:"id" bson:"id"`
	Title       string `json:"title" bson:"title"`
	Slug        string `json:"slug" bson:"slug" db:"slug"`
	Description string `json:"description" bson:"description"`
	Ordinal     int    `json:"ordinal" bson:"ordinal"`
	AssetPhoto  string `json:"asset_photo,omitempty" bson:"asset_photo"`
	NetPhoto    string `json:"net_photo,omitempty" bson:"net_photo"`
}

func (p *SubjectType) TableName() string {
	return "subject_type"
}

type Topic struct {
	Id             int    `json:"id" bson:"id"`
	SubjectId      int    `json:"subject_id" bson:"subject_id"`
	ParentId       *int   `json:"parent_id,omitempty" bson:"parent_id"`
	RootTreeId     *int   `json:"root_tree_id,omitempty" bson:"root_tree_id"`
	Title          string `json:"title" bson:"title"`
	Slug           string `json:"slug" bson:"slug"`
	Ordinal        int    `json:"ordinal" bson:"ordinal"`
	Level          int16  `json:"level" bson:"level"`
	DetailViewType int16  `json:"detail_view_type" bson:"detail_view_type"`
	AccessLevel    int    `json:"access_level" bson:"access_level"`
	Semester       int    `json:"semester" bson:"semester"`
}

func (p *Topic) TableName() string {
	return "topic"
}
