package mongo

import (
	"time"

	"github.com/Tinee/doit/domain"

	"github.com/globalsign/mgo"
)

// ClientInfo defines all the different fields you need to call the database.
type ClientInfo struct {
	DBName   string
	Addr     string
	Username string
	Password string
}

// Client can open a connection to mongo
type Client struct {
	ClientInfo
	s        *mgo.Session
	userRepo UserRepository
}

// NewClient creates a new client.
func NewClient(i ClientInfo) *Client {
	c := &Client{
		ClientInfo: i,
	}

	c.userRepo.c = c

	return c
}

// Open attempts to open a connection between the client and the database.
// It also configure different indexs in the mongo database.
func (c *Client) Open() error {
	s, err := mgo.DialWithInfo(&mgo.DialInfo{
		Addrs:    []string{c.Addr},
		Username: c.Username,
		Password: c.Password,
		Database: c.DBName,
		Timeout:  time.Second * 8,
	})
	if err != nil {
		return err
	}
	// Assigns the main session to the client.
	c.s = s

	// Ensures indexes for the User Collection.
	s.DB(c.DBName).C(userCol).EnsureIndex(mgo.Index{
		Key:    []string{"email"},
		Unique: true,
	})

	return nil
}

// Close attempts to close the main session in the client.
// If for whatever reason it does not have a main session it gracefully handles that.
func (c *Client) Close() error {
	if c.s != nil {
		c.s.Close()
		return nil
	}
	return nil
}

// UserRepository exposes the underlaying credential repository
func (c *Client) UserRepository() domain.UserRepository {
	return &c.userRepo
}
