import hashlib
from urllib import parse
import requests


uuid = 'e02f3c18-bf26-11eb-8529-0242ac130003'
# uuid = 'c91e1286-bf28-11eb-8529-0242ac130003'
mac = hashlib.md5(b'user_2_mac').hexdigest()
uname = hashlib.md5(b'user_2').hexdigest()
program = 'Analyzer'
version = '2021.0.0'
platform = 'Windows 10 x64'

delimiter = '&'

uuid = parse.quote_plus(uuid)
mac = parse.quote_plus(mac)
uname = parse.quote_plus(uname)
program = parse.quote_plus(program)
version = parse.quote_plus(version)
platform = parse.quote_plus(platform)

command = 'uuid=' + uuid
command += delimiter + 'mac=' + mac
command += delimiter + 'uname=' + uname
command += delimiter + 'program=' + program
command += delimiter + 'version=' + version
command += delimiter + 'platform=' + platform

print(command)

url = 'http://localhost:8090/us?' + command
response = requests.post(url)
print("all done")
print(response.status_code, response.text)
