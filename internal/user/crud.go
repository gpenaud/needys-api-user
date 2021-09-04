package user

import (
  _   "github.com/go-sql-driver/mysql"
  log "github.com/sirupsen/logrus"
  sql "database/sql"
)

var userLog *log.Entry

func init() {
  userLog = log.WithFields(log.Fields{
    "_file": "internal/user/crud.go",
    "_type": "data",
  })
}

type User struct {
  Id        int
  Firstname string
  Lastname  string
  Address   string
  Phone     string
}

func (n *User) CreateUser(db *sql.DB) (err error) {
  userLog.WithFields(log.Fields{
    "type": "database query",
    "parameter_firstname": n.Firstname,
    "parameter_lastname": n.Lastname,
    "parameter_address": n.Address,
    "parameter_phone": n.Phone,
  }).Debug("INSERT INTO user (name, priority) VALUES ({name}, {priority})")

  r, err := db.Query("INSERT INTO user (firstname, lastname, address, phone) VALUES (?, ?, ?, ?)", n.Firstname, n.Lastname, n.Address, n.Phone)
  r.Close()

  return err
}

func (n *User) GetUser(db *sql.DB) (error) {
  userLog.WithFields(log.Fields{
    "type": "database query",
    "parameter_id": n.Id,
  }).Debug("SELECT * FROM user WHERE id={id}")

  err := db.QueryRow("SELECT * FROM user WHERE id=?", n.Id).Scan(&n.Id, &n.Firstname, &n.Lastname, &n.Address, &n.Phone)
  return err
}

func (n *User) UpdateUser(db *sql.DB) (err error) {
  userLog.WithFields(log.Fields{
    "type": "database query",
    "parameter_firstname": n.Firstname,
    "parameter_lastname": n.Lastname,
    "parameter_address": n.Address,
    "parameter_phone": n.Phone,
  }).Debug("UPDATE user SET address = '{address}', phone = '{phone}' WHERE firstname = '{firstname}' AND lastname = '{lastname}'")

  r, err := db.Query("UPDATE user SET address = ?, phone = ? WHERE firstname = ? AND lastname = ?", n.Address, n.Phone, n.Firstname, n.Lastname)
  r.Close()

  return err
}

func (n *User) DeleteUser(db *sql.DB) error {
  userLog.WithFields(log.Fields{
    "type": "database query",
    "parameter_firstname": n.Firstname,
    "parameter_lastname": n.Lastname,
  }).Debug("DELETE FROM user WHERE name='{name}'")

  r, err := db.Query("DELETE FROM user WHERE firstname = ? AND lastname = ?", n.Firstname, n.Lastname)
  r.Close()

  return err
}

func (n *User) GetUsers(db *sql.DB) ([]User, error) {
  user  := User{}
  users := []User{}

  selDB, err := db.Query("SELECT * FROM user ORDER BY id ASC")

  if err != nil {
    return users, err
  }

  for selDB.Next() {
    var id int
    var firstname, lastname, address, phone string

    err = selDB.Scan(&id, &firstname, &lastname, &address, &phone)
    if err != nil {
      return users, err
    }

    user.Id        = id
    user.Firstname = firstname
    user.Lastname  = lastname
    user.Address   = address
    user.Phone     = phone

    users = append(users, user)
  }

  return users, err
}
