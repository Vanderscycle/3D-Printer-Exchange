#!/usr/bin/env python3
import pretty_errors
import os
import json
import requests

port: int = 8080
baseUrl: str = f"localhost:{port}"
dummyAcct: dict[str, str] = {"username": "dasdborisxsss",
                   "email":"bsdsdorixss@mcBoris.com",
                   "password":"nottVerySafeInit"}
headers: dict[str, str] = {'Content-type': 'application/json', 'Accept': 'text/plain'}

# convert into JSON:
dummyAcctPayload: str = json.dumps(dummyAcct, indent=2)


def ping(endpoint: str) -> None:
    os.system(f'echo "{dummyAcctPayload}"')
    # os.system(f'http GET {baseUrl}/{endpoint}')
    r = requests.post(f'http://{baseUrl}/{endpoint}', json=dummyAcct, headers=headers) # needs to capture the return data
    data = r.json()
    print(data['data']['ID']) # TODO: add a try/catch block
    # os.system(f'http GET {baseUrl}/{endpoint}/1')
    # os.system(f'http DELETE {baseUrl}/{endpoint}/1')


if __name__ == "__main__":
    ping("api/user")
