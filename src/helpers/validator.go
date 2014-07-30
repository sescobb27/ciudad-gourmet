package helpers

import (
        "regexp"
)

const (
        email_key    = "[a-z0-9!#$%&'*+/=?^_`{|}~-]+(?:\\.[a-z0-9!#$%&'*+/=?^_`{|}~-]+)*"
        email_domain = "@(?:[a-z0-9](?:[a-z0-9-]*[a-z0-9])?\\.)+[a-z0-9](?:[a-z0-9-]*[a-z0-9])?"
)

func EmailValidator(email string) bool {
        pattern := regexp.MustCompile(email_key + email_domain)

        if !pattern.MatchString(email) {
                return false
        }
        return true
}

func UserNamesValidator(name string) bool {
        //  Ejemplo: "Nombre Apellido Apellido "
        //  Ejemplo: "Nombre Apellido Apellido"
        //  Ejemplo: "Nombre Apellido"
        //  Ejemplo: "Apellido Apellido "
        //  Ejemplo: "Nombre"
        //  Ejemplo: "Apellido"
        pattern := regexp.MustCompile(`\A([ña-zA-ZÑ]{3,16} {0,1}){1,3}\z`)

        if !pattern.MatchString(name) {
                return false
        }
        return true
}

func UniqueNamesValidator(unique_name string) bool {
        pattern := regexp.MustCompile(`\A\w{4,10}\z`)

        if !pattern.MatchString(unique_name) {
                return false
        }
        return true
}

func ProductNameValidator(p_name string) bool {
        pattern := regexp.MustCompile(`\A(\w|\s){4,30}\z`)

        if !pattern.MatchString(p_name) {
                return false
        }
        return true
}

func TextOnlyValidator(unique_name string) bool {
        pattern := regexp.MustCompile(`\A(\w+|\s)+\z`)

        if !pattern.MatchString(unique_name) {
                return false
        }
        return true
}
