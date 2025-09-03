# CHAT APP

## Installation

clone this repo:
```bash
git clone https://github.com/Zyprush18/deeepen-golang/chat-app/backend
cd chat-app
```

before you running this app,make sure you have created a new database in postgresql with the database name `chatrealtime`.

after you created a new database,running this app:
```bash
go run main.go
```


you can access auth at:`http://localhost:3000` and for websocket at:`http://localhost:3000/ws/chat`