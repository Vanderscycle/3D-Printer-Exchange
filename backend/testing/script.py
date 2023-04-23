#!/usr/bin/env python3
import pretty_errors
import os
import json
import requests

port: int = 8080
baseUrl: str = f"localhost:{port}"
dummyAcct: dict[str, str] = {"username": "borisxss",
                   "email":"borixs@mcBoris.com",
                   "password":"nottVerySafeInit"}
headers: dict[str, str] = {'Content-type': 'application/json', 'Accept': 'text/plain'}

parseResponse: str = "{'data': {'CreatedAt': '2023-04-22T20:25:34.835144066-07:00', 'UpdatedAt': '2023-04-22T20:25:34.835144066-07:00', 'DeletedAt': None, 'ID': '27b7bf49-5701-41ac-a2c2-43d2cab60745', 'username': 'borisxss', 'email': 'borixs@mcBoris.com', 'password': '$2a$14$5cOwOb8Tml3hd5y0DzzPV.lwoAkcvuhCUhycQMEAQAAdwwe2UH5bG'}, 'message': 'User has created', 'status': 'success'}"
# convert into JSON:
dummyAcctPayload: str = json.dumps(dummyAcct, indent=2)


def ping(endpoint: str) -> None:
    os.system(f'echo "{dummyAcctPayload}"')
    # os.system(f'http GET {baseUrl}/{endpoint}')
    r = requests.post(f'http://{baseUrl}/{endpoint}', json=dummyAcct, headers=headers) # needs to capture the return data
    data = r.json()
    print(data)
    # os.system(f'http GET {baseUrl}/{endpoint}/1')
    # os.system(f'http DELETE {baseUrl}/{endpoint}/1')


if __name__ == "__main__":
    ping("api/user")
