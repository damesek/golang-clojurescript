# Gracefully shutdown Clojurescript with fast Go backend
TODO add fasthttp package

In this way x5 better performance (with fasthttp x10) vs non-clustered nodejs. 
I used a *shadow-cljs project example* [https://github.com/shadow-cljs/quickstart-browser] and a *gracefull shutdown Go example* to create that mix.

Clone the repo or just Release RELEASE_X86_64 file and run<BR> 
*I use Mac OSX x86_64, if you use Linux/Windows/RPI maybe better if you rebuild*
```bash
./gocljs
```
-> Don't forget add your public folder with index.html or unzip [public.zip] my shadow-cljs test
 
After this you can see the message: 
```bash
2021/03/03 18:10:11 server started
```
on localhost:8081

When you stop the server with CTRC+C, you will see:
```bash
^C2021/03/03 18:11:32 system call:interrupt
2021/03/03 18:11:32 server stopped
2021/03/03 18:11:32 server exited properly
```

Details:

Test with  
```shell
go run gocljs.go
```
or compile a file to architecture
```shell
go build .
```
We use external files in this example, but I tested without any problem possible compile into binary the css/js/html files or run CLJS with V8 engine from Golang. 

Links more info:
- https://medium.com/@pinkudebnath/graceful-shutdown-of-golang-servers-using-context-and-os-signals-cc1fa2c55e97. That is possible with js/process `(.on js/process "SIGINT" shutdown-gracefully-functions)` and 
- Clojure side too (https://medium.com/helpshift-engineering/achieving-graceful-restarts-of-clojure-services-b3a3b9c1d60d).

Btw "Go- and Java-based webservers have shown to be the most efficient. Clustered NodeJS is reasonably efficient, but will run out of RAM once overloaded." Source: https://stressgrid.com/blog/webserver_benchmark/


