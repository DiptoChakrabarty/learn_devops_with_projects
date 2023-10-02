# Flask API

## @DiptoChakrabarty

## DESCRIPTION OF PROJECT
```
A simple flask api performing GET and POST request
```

## Steps to run

- Create a virtualenv of the project

- Install dependencies
```
pip3 install -r requirements.txt
```

- Run project
```
python3 app.py
```

- Try the GET request
```
curl http://localhost:8080/
{"Message":"app up and running successfully"}
```

- Try the POST request
```
curl -X POST -H "Content-Type: application/json" \  
    -d '{"name": "ezio", "server": "dchost"}' \
    http://localhost:8080/access

{"Message":"User ezio received access to server dchost"}
```

## ScreenShots -> IF NEEDED