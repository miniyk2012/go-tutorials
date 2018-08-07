package badger

import (
	"github.com/CreatCodeBuild/go-tutorials/graphql-server/core/adapters"
	"github.com/CreatCodeBuild/go-tutorials/graphql-server/core/entities"
	"github.com/dgraph-io/badger"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
)

type Badger struct {
	userDB *badger.DB
}

var _ adapters.User = (*Badger)(nil)

func NewBadger() (*Badger, error) {
	return &Badger{}, nil
}

func (b *Badger) connect() error {
	opts := badger.DefaultOptions
	opts.Dir = "./ggg/db/badger/data"
	opts.ValueDir = "./ggg/db/badger/data"
	db, err := badger.Open(opts)
	if err != nil {
		return errors.WithStack(err)
	}
	b.userDB = db
	return nil
}

func (b *Badger) GetUser(uid string) (entities.User, error) {
	err := b.connect()
	if err != nil {
		return entities.User{}, errors.WithStack(err)
	}
	fmt.Println("2")
	defer b.userDB.Close()

	var user entities.User
	var bytes []byte

	err = b.userDB.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(uid))
		if err != nil {
			return err
		}
		bytes, err = item.Value()
		if err != nil {
			return err
		}
		return nil
	})
	fmt.Println("3")
	if err != nil {
		return entities.User{}, err
	}
	fmt.Println("4")

	err = json.Unmarshal(bytes, &user)
	return user, err
}

func (b *Badger) CreateUser(user entities.User) (entities.User, error) {
	err := b.connect()
	if err != nil {
		return entities.User{}, err
	}
	defer b.userDB.Close()

	bytes, err := json.Marshal(user)
	if err != nil {
		return entities.User{}, err
	}

	err = b.userDB.Update(func(txn *badger.Txn) error {
		err := txn.Set([]byte(user.UID), bytes)
		return err
	})
	return user, err
}
