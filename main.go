package main

import (
    "flag"
    "github.com/julienschmidt/httprouter"
    sql "github.com/sescobb27/ciudad-gourmet/db"
    . "github.com/sescobb27/ciudad-gourmet/handlers"
    "net/http"
    "os"
    "os/signal"
    "syscall"
)

func signalHandler() {
    signalChan := make(chan os.Signal, 1)
    signal.Notify(signalChan,
        syscall.SIGINT,
        syscall.SIGKILL,
        syscall.SIGTERM,
        syscall.SIGQUIT,
        syscall.SIGABRT,
    )
    <-signalChan
    println("Closing DB Connection")
    sql.DB.Close()
    os.Exit(0)
}

func main() {
    seed := flag.Bool("seed", false, "Seed the database")
    restore := flag.Bool("restore", false, "Restore the database")
    flag.Parse()

    if *seed {
        Seed()
        os.Exit(0)
    } else if *restore {
        Restore()
        os.Exit(0)
    }

    router := httprouter.New()

    go signalHandler()

    router.GET("/", Index_Handler)
    router.GET("/categories", Categories_Handler)
    router.GET("/locations", Locations_Handler)
    router.GET("/products", Products_Handler)
    router.GET("/products/find", FindProducts_Handler)

    router.POST("/signin", SignIn_Handler)
    router.POST("/signout", SignOut_Handler)
    router.POST("/signup", SignUp_Handler)
    router.POST("/purchase", Purchase_Handler)

    router.GET("/chefs", Chefs_Handler)
    router.POST("/chefs/product/add", ChefAddProduct_Handler)

    router.Handler("GET", "/images/*filename",
        http.StripPrefix("/images/",
            http.FileServer(http.Dir("resources/images"))))

    router.Handler("GET", "/css/*filename",
        http.StripPrefix("/css/",
            http.FileServer(http.Dir("resources/css"))))

    router.Handler("GET", "/js/*filename",
        http.StripPrefix("/js/",
            http.FileServer(http.Dir("resources/js"))))

    router.Handler("GET", "/catalog/*filename",
        http.StripPrefix("/catalog/",
            http.FileServer(http.Dir("resources/catalog"))))

    http.ListenAndServe(":3000", router)
}
