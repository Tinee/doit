package mongo

import (
	"time"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"

	"github.com/Tinee/doit/domain"
)

const userCol = "user"

// UserRepository have access to methods that you can communicate to the mongo instance
type UserRepository struct {
	c *Client
}

// Create inserts a User into the user collection. If any index fails it returns an error.
func (r UserRepository) Create(c domain.User) (*domain.User, error) {
	s := r.c.s.Clone()
	defer s.Close()

	c.ID = bson.NewObjectId()
	c.CreatedAt = time.Now()

	err := s.DB("").C(userCol).Insert(c)

	if err != nil {
		if mgo.IsDup(err) {
			return nil, domain.ErrUserExists
		}
		return nil, domain.ErrInternal
	}

	return &c, nil
}
