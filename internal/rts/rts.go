package rts

import (
    "errors"
    "log"
    "os"

    "github.com/joho/godotenv"
)

/*
   RTS - runtime storage.
   Manages configuration in runtime.
*/

type storage struct {
    port   string
    host   string
    dbConf string
    dbPath string
    secret string
}

var runtimeStorage storage

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
    return runtimeStorage.port
}

func GetHost() string {
    return runtimeStorage.host
}

func GetDBConfig() string {
    return runtimeStorage.dbConf
}

func GetDBPath() string {
    return runtimeStorage.dbPath
}

func GetSecret() string {
    return runtimeStorage.secret
}

func IsValid() error {
    // @TODO: Validation
    s := runtimeStorage

    if len(s.port) == 0 {
        return errors.New("rts: cannot read PORT from env")
    }

    if len(s.host) == 0 {
        return errors.New("rts: cannot read HOST from env")
    }

    if len(s.dbConf) == 0 {
        return errors.New("rts: cannot read DB_CONF from env")
    }

    if len(s.dbPath) == 0 {
        return errors.New("rts: cannot read SQLITE_PATH from env")
    }

    if len(s.secret) == 0 {
        return errors.New("rts: cannot read SECRET_KEY from env")
    }

    return nil
}

func init() {
    err := godotenv.Load()

    if err != nil {
        log.Fatal(err)
    }

    runtimeStorage = newStorage(
        os.Getenv("PORT"),
        os.Getenv("HOST"),
        os.Getenv("DB_CONF"),
        os.Getenv("SQLITE_PATH"),
        os.Getenv("SECRET_KEY"),
    )
}
