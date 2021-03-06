package main

import (
    "flag"
    "fmt"
    "github.com/Danzabar/WatchDog/api"
    "github.com/Danzabar/WatchDog/core"
    "github.com/Danzabar/WatchDog/site"
    "github.com/Danzabar/WatchDog/watcher"
    "github.com/jasonlvhit/gocron"
)

func main() {
    r := flag.Bool("r", false, "When included will run the application")
    m := flag.Bool("migrate", false, "Runs database schema migration tool if included")
    dd := flag.String("driver", "sqlite3", "The database driver to use")
    dc := flag.String("creds", "/tmp/main.db", "The database credentials")
    p := flag.String("port", ":8080", "The port on which this listens")
    w := flag.Bool("w", false, "Performs a watch operation on load if set")
    n := flag.Bool("n", true, "When set to true, enables alerts, otherwise bypasses them")

    flag.Parse()

    core.NewApp(*p, *dd, *dc, *n)

    api.Setup()
    site.Setup("site/templates/")

    fmt.Print(`
     ___          ___                   ___          ___                  _____         ___          ___     
    /__/\        /  /\         ___     /  /\        /__/\                /  /::\       /  /\        /  /\    
   _\_ \:\      /  /::\       /  /\   /  /:/        \  \:\              /  /:/\:\     /  /::\      /  /:/_   
  /__/\ \:\    /  /:/\:\     /  /:/  /  /:/          \__\:\            /  /:/  \:\   /  /:/\:\    /  /:/ /\  
 _\_ \:\ \:\  /  /:/~/::\   /  /:/  /  /:/  ___  ___ /  /::\          /__/:/ \__\:| /  /:/  \:\  /  /:/_/::\ 
/__/\ \:\ \:\/__/:/ /:/\:\ /  /::\ /__/:/  /  /\/__/\  /:/\:\         \  \:\ /  /://__/:/ \__\:\/__/:/__\/\:\
\  \:\ \:\/:/\  \:\/:/__\//__/:/\:\\  \:\ /  /:/\  \:\/:/__\/          \  \:\  /:/ \  \:\ /  /:/\  \:\ /~~/:/
 \  \:\ \::/  \  \::/     \__\/  \:\\  \:\  /:/  \  \::/                \  \:\/:/   \  \:\  /:/  \  \:\  /:/ 
  \  \:\/:/    \  \:\          \  \:\\  \:\/:/    \  \:\                 \  \::/     \  \:\/:/    \  \:\/:/  
   \  \::/      \  \:\          \__\/ \  \::/      \  \:\                 \__\/       \  \::/      \  \::/   
    \__\/        \__\/                 \__\/        \__\/                              \__\/        \__\/    


    `)

    if *m {
        Migrate()
    }

    if *w {
        watcher.Watch()
    }

    if *r {
        gocron.Every(1).Minute().Do(watcher.Watch)
        gocron.Start()
        core.App.Run()
    }
}

func Migrate() {
    core.App.Log.Debug("Starting Migrations")
    site.Migrate()
    core.App.Log.Debug("Finished Migrations")
}
