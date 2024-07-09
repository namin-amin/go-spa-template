# GO SPA-SelfHosted Template

- This repo contains the template that has Golag/Echo server which also hosts a Vite/SPA app.
- Run below command to run the app in development mode 

```
make dev
```
- This command installs npm packages then runs the vite app
- Then it start echo server with a proxy that points to the ui app.
- This enables reload fetures in development.



---

- Run Below command to build the app
  ```
  make build
  ```
- This will build the ui app then will build the golang app with ui app embedded in it
- Run 
  ```
  make run-prod
  ```
- This will run  the app in production mode using the embedded files instead of proxy.
----

Envirnment files reading is configured, currently supports
- .env : for production
- .env.dev : for development


NOTE & WARNINGS 🔥
-- 
- Check make file before running. This template used bun for npm package management change it to your preferred one during use(npm/yarn).
- Dont forget to run *go mod tidy* before starting anything
- This implementation used basic envirnment variable **RUNENV** which is checked during server startup to check if to use proxy or embedded file. 
    You can change this to your preffered method instead of this.(Check make command **run-prod**)
- This uses Angular as ui app you can delete and create new SPA app and just copy embed.go and modify it as required.
