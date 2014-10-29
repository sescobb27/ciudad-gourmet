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
    return pattern.MatchString(email)
}

func UserNamesValidator(name string) bool {
    //  Ejemplo: "Nombre Apellido Apellido "
    //  Ejemplo: "Nombre Apellido Apellido"
    //  Ejemplo: "Nombre Apellido"
    //  Ejemplo: "Apellido Apellido "
    //  Ejemplo: "Nombre"
    //  Ejemplo: "Apellido"
    pattern := regexp.MustCompile(`\A([ña-zA-ZÑ]{3,16} {0,1}){1,3}\z`)
    return pattern.MatchString(name)
}

func UniqueNamesValidator(unique_name string) bool {
    pattern := regexp.MustCompile(`\A\w{4,10}\z`)
    return pattern.MatchString(unique_name)
}

func ProductNameValidator(p_name string) bool {
    pattern := regexp.MustCompile(`\A(\w|\s){4,30}\z`)
    return pattern.MatchString(p_name)
}

func TextOnlyValidator(text string) bool {
    pattern := regexp.MustCompile(`\A(\w+|\s)+\z`)
    return pattern.MatchString(text)
}
