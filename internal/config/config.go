package config

import (
    "errors"
    "log"
    "os"

    "github.com/joho/godotenv"
)

/*
   config - configuration storage
*/

type storage struct {
    port   string
    host   string
    dbConf string
    dbPath string
    secret string
}

var rs storage

func newStorage(port, host, dbConf, dbPath, secret string) storage {
    return storage{
        port:   port,
        host:   host,
        dbConf: dbConf,
        dbPath: dbPath,
        secret: secret,
    }
}

func GetPort() string {
    return rs.port
}

func GetHost() string {
    return rs.host
}

func GetDBConfig() string {
    return rs.dbConf
}

func GetDBPath() string {
    return rs.dbPath
}

func GetSecret() string {
    return rs.secret
}

func IsValid() error {
    // @TODO: Validation
    s := rs

    if len(s.port) == 0 {
        return errors.New("config: cannot read PORT from env")
    }

    if len(s.host) == 0 {
        return errors.New("config: cannot read HOST from env")
    }

    if len(s.dbConf) == 0 {
        return errors.New("config: cannot read DB_CONF from env")
    }

    if len(s.dbPath) == 0 {
        return errors.New("config: cannot read SQLITE_PATH from env")
    }

    if len(s.secret) == 0 {
        return errors.New("config: cannot read SECRET_KEY from env")
    }

    return nil
}

func init() {
    err := godotenv.Load()

    if err != nil {
        log.Fatal(err)
    }

    rs = newStorage(
        os.Getenv("PORT"),
        os.Getenv("HOST"),
        os.Getenv("DB_CONF"),
        os.Getenv("SQLITE_PATH"),
        os.Getenv("SECRET_KEY"),
    )
}
