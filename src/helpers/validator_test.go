package helpers

import (
        "testing"
)

func TestEmailValidator(t *testing.T) {

}

func TestUserNamesValidator_Valid(t *testing.T) {
        valid_names := []string{"Simon", "Kiro", "Edgardo", "JuAn",
                "CARLOS", "Luz", "VeryLongName", "veryLONGLONGname",
                "Simon Escobar Benitez", "Edgardo Andres Sierra",
                "Juan Carlos Noreña", "Very Long LongName",
                "pepitoperezperez pepitoperezperez pepitoperezperez",
                "Luz Marina", "MAYUS MAYUS MAYUS", "CaMeL LoNg NaMe"}

        for _, name := range valid_names {
                if !UserNamesValidator(name) {
                        t.Fatalf("%s should be valid", name)
                }
        }
}

func TestUserNamesValidator_Invalid(t *testing.T) {
        invalid_names := []string{"",
                "   ", "Simon1", "1Simon", "_Simon", "Simon_",
                "_Simon_", "Si", "S", "12345", "@simon", "51m0N",
                "simon...", "...simon", "...simon...", "!\"#$&",
                "pepitoperezperez pepitoperezperez pepitoperezperez pepitoperezperez",
                "pepitoperezperez 1234",
                "pepitoperezperez pepitoperezperez pepitoperezperezperez",
                "pepitoperezperez p",
                "pepitoperezperez pe",
                "pepitoperezperez pepitoperezperez p",
                "pepitoperezperez pepitoperezperez pe",
                "p pepitoperezperez",
                "pe pepitoperezperez",
                "<script> alert('fuckyou'); </script>"}

        for _, name := range invalid_names {
                if UserNamesValidator(name) {
                        t.Fatalf("%s should be invalid", name)
                }
        }
}

func TestUniqueNamesValidator(t *testing.T) {

}

func TestProductNameValidator(t *testing.T) {

}

func TestTextOnlyValidator(t *testing.T) {

}
