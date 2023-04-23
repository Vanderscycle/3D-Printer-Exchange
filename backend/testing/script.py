#!/usr/bin/env python3
import pretty_errors
import os
import json

port: int = 8080
baseUrl: str = f"localhost:{port}"
dummyAcct: dict[str, str] = {"username": "boris",
                   "email":"boris@mcBoris.com",
                   "password":"notVerySafeInit"}
# convert into JSON:
dummyAcctPayload: str = json.dumps(dummyAcct, indent=2)


def ping(endpoint: str) -> None:
    os.system(f'echo "{dummyAcctPayload}" ')
    # os.system(f'http GET {baseUrl}/{endpoint}')
    os.system(f'echo "{dummyAcctPayload}" | http POST {baseUrl}/{endpoint}') # needs to capture the return data
    # os.system(f'http GET {baseUrl}/{endpoint}/1')
    # os.system(f'http DELETE {baseUrl}/{endpoint}/1')


if __name__ == "__main__":
    ping("api/user")
