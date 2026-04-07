//db schema
type User struct{
	ID uint `gorm:"primaryKey"`
	Name string 
	Email string `gorm:"unique"`
	Password string
	Projects []Project `gorm:"many2many:user_projects;"`
}

type Project struct{
	ID uint `gorm:"primaryKey"`
	Name string 
	Description string 
	Members []User `gorm:"many2many:project_members;"`
	Issues []Issue
}

type Task struct{
	ID uint `gorm:"primaryKey"`
	Name string
	Description string 

	ProjectID uint  
	Project Project `gorm:"foreignKey:ProjectID"`

	AssignedToId uint
	AssignedTo User `gorm:"foreignKey:AssignedToId"`
	createdAt time.Time
	dueDate time.Time
	status string //pending, completed
}
type Issue struct{
	ID uint `gorm:"primaryKey"`
	Title string
	Description string 

	ProjectID uint  
	Project Project `gorm:"foreignKey:ProjectID"`

	RaisedById uint
	RaisedBy User `gorm:"foreignKey:RaisedById"`

	AssignedToId uint
	AssignedTo User `gorm:"foreignKey:AssignedToId"`
	
	createdAt time.Time
	resolvedAt time.Time
	status string //open, closed
	Labels []Label `gorm:"many2many:issue_labels;"`
}

type Label struct{
	ID uint `gorm:"primaryKey"`
	Name string
	Issues []Issue `gorm:"many2many:issue_labels;"`
}