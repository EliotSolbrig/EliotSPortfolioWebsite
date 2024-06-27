package Util

import (
    "os"
)

func GetHostname() string {
    if os.Getenv("ENV_TYPE") == "dev" {
        return os.Getenv("DEV_HOSTNAME")
    } else {
        if os.Getenv("DEPLOYED") == "true" {
            return os.Getenv("PROD_HOSTNAME")
        } else {
            return os.Getenv("DEV_HOSTNAME")
        }
    }
}
