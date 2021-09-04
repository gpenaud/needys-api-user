package internal

// -----------------------------------------------------------------------------
// 1. Database initialization (if specified in configuration)
// -----------------------------------------------------------------------------

const dbReset = `DROP TABLE IF EXISTS user`
const dbInit = `
  CREATE TABLE user (
    id INTEGER PRIMARY KEY NOT NULL AUTO_INCREMENT,
    firstname VARCHAR(100),
    lastname VARCHAR(100),
    address VARCHAR(100),
    phone VARCHAR(100)
  ) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
  `
const dbInsertFirst  = `INSERT INTO user (firstname, lastname, address, phone) VALUES ('Guillaume', 'Penaud', '16 sentier de la côte 94370 Sucy-En-Brie', '0666222475');`
const dbInsertSecond = `INSERT INTO user (firstname, lastname, address, phone) VALUES ('Pauline', 'Breniaux', '10 route de Rhye 74210 Mouthier-En-Bresse', '0645124365');`

func (a *Application) InitializeDatabase() (bool, error) {
  var err error

  if _, err = a.DB.Exec(dbReset); err != nil {
    return false, err
  }

  if _, err = a.DB.Exec(dbInit); err != nil {
    return false, err
  }

  if _, err = a.DB.Exec(dbInsertFirst); err != nil {
    return false, err
  }

  if _, err = a.DB.Exec(dbInsertSecond); err != nil {
    return false, err
  }

  return true, err
}
