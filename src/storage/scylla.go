package storage

import (
	"fmt"
	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx/v2"
	"github.com/wzslr321/road_runner/server/users/src/domain"
	pb "github.com/wzslr321/road_runner/server/users/src/proto-gen"
	"log"
	"strconv"
)

type IUserStorage interface {
	SaveUser(user *domain.User) error
	FindUserByUsername(username string) (*domain.User, error)
	UpdateUserByID(req *pb.UpdateUserRequest) error
	DeleteUserByID(id string) error
}

type UserStorage struct {
	cluster *gocql.ClusterConfig
}

func New() *UserStorage {
	cluster := gocql.NewCluster("scylladb:9042")
	cluster.Consistency = gocql.Quorum
	cluster.ProtoVersion = 4
	log.Println("userStorage cluster successfully created")

	session, err := gocqlx.WrapSession(cluster.CreateSession())
	if err != nil {
		// gotta handle it better
		log.Printf("Failed to create scylladb session: %v", err)
	}
	defer session.Close()

	err = tryToCreateKeyspace(&session)
	if err != nil {
		// gotta handle it better
		log.Printf("Failed to create keyspace: %v", err)
	}

	err = tryToCreateTable(&session)
	if err != nil {
		// gotta handle it better
		log.Printf("Failed to create table: %v", err)
	}

	return &UserStorage{cluster: cluster}
}

// make unique or sth
func (s *UserStorage) SaveUser(user *domain.User) error {
	session, err := gocqlx.WrapSession(s.cluster.CreateSession())
	if err != nil {
		return err
	}

	q := fmt.Sprintf("INSERT INTO users.users (id, email, password, role, username) VALUES ('%s', '%s', '%s', %s, '%s')", user.Id, user.Email, user.Password, strconv.Itoa(user.Role), user.Username)
	err = session.Query(q, nil).Exec()
	if err != nil {
		return err
	}

	return nil
}

func (s *UserStorage) UpdateUserByID(req *pb.UpdateUserRequest) error {
	session, err := gocqlx.WrapSession(s.cluster.CreateSession())
	if err != nil {
		return err
	}

	q := fmt.Sprintf("UPDATE users.users SET username='%s', email='%s', password='%s' WHERE id='%s'", req.Username, req.Email, req.Password, req.Id)
	err = session.Query(q, nil).Exec()
	if err != nil {
		return err
	}

	return nil
}

// I think i should accept user model here and then convert it to proto
// but is it worth it tho?
func (s *UserStorage) FindUserByUsername(username string) (*domain.User, error) {
	session, err := gocqlx.WrapSession(s.cluster.CreateSession())
	if err != nil {
		return nil, err
	}

	var users []*domain.User
	q := fmt.Sprintf("SELECT * FROM users.users WHERE username = '%s' ALLOW FILTERING", username)
	err = session.Query(q, nil).Select(&users)
	if err != nil {
		return nil, err
	}

	if len(users) == 0 {
		return nil, fmt.Errorf("user not found")
	}

	return users[0], nil
}

func (s *UserStorage) DeleteUserByID(id string) error {
	session, err := gocqlx.WrapSession(s.cluster.CreateSession())
	if err != nil {
		return err
	}

	q := fmt.Sprintf("DELETE FROM users.users WHERE id = '%s'", id)
	err = session.Query(q, nil).Exec()
	if err != nil {
		return err
	}

	return nil
}

func tryToCreateKeyspace(session *gocqlx.Session) error {
	q := "CREATE KEYSPACE IF NOT EXISTS users WITH replication = {'class': 'SimpleStrategy', 'replication_factor': 1}"
	err := session.Query(q, nil).Exec()
	if err != nil {
		return err
	}
	return nil
}

func tryToCreateTable(session *gocqlx.Session) error {
	q := "CREATE TABLE IF NOT EXISTS users.users (id text, email text, username text, password text, role int, PRIMARY KEY (id))"
	err := session.Query(q, nil).Exec()
	if err != nil {
		return err
	}
	return nil
}
