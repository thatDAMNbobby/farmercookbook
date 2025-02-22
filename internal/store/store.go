package store

type User struct {
	ID          int    `gorm:"primaryKey" json:"id"`
	Email       string `json:"email"`
	Password    string `json:"-"`
	CreatedDate string `json:"createdDate"`
}

type Session struct {
	ID        int    `gorm:"primaryKey" json:"id"`
	SessionID string `json:"sessionId"`
	UserID    int    `json:"userId"`
	User      User   `gorm:"foreignKey:UserID" json:"user"`
}

type Recipe struct {
	Id          int `gorm:"primaryKey" json:"id"`
	Image       string
	Name        string
	Description string
	PrepTime    string
	CookTime    string
	Servings    int
	Ingredients string
	Steps       string
	CreatedDate string
	UserID      int  `json:"userId"`
	Creator     User `gorm:"foreignKey:UserID" json:"user"`
}

type UserStore interface {
	CreateUser(email string, password string) error
	GetUser(email string) (*User, error)
}

type SessionStore interface {
	CreateSession(session *Session) (*Session, error)
	GetUserFromSession(sessionID string, userID string) (*User, error)
}

type RecipeStore interface {
	CreateRecipe(recipe *Recipe) (*Recipe, error)
	GetRecipe(id int) (*Recipe, error)
	GetRecipes(max int) ([]*Recipe, error)
	UpdateRecipe(recipe *Recipe) (*Recipe, error)
	DeleteRecipe(id int) error
}
