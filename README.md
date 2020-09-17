# golang-clojurescript
Gracefully shutdown Clojurescript (basic shadow-cljs project) with Golang

Based on https://github.com/shadow-cljs/quickstart-browser shadow-cljs basic example. Detailed info about this project you can find on the link.

Test with  '`go run gocljs.go`' or '`go build .`' and run the compiled file. We use external files, but possible compile into binary the css/js/html files. 

This basic example show how to shut down gracefully the cljs projects with Go, based on https://medium.com/@pinkudebnath/graceful-shutdown-of-golang-servers-using-context-and-os-signals-cc1fa2c55e97. That is possible with js/process `(.on js/process "SIGINT" shutdown-gracefully-functions)` and Clojure side too (https://medium.com/helpshift-engineering/achieving-graceful-restarts-of-clojure-services-b3a3b9c1d60d).

