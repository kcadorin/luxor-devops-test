from flask import Flask
import requests
import time
import json
import os

app = Flask(__name__)

# Load configuration from environment variables
config = {
    "hosts": json.loads(os.environ.get("OTHER_HOSTS", "[]"))
}

@app.route('/ping')
def ping():
    return 'pong\n'

def check_other_hosts():
    while True:
        for host in config["hosts"]:
            try:
                response = requests.get(f'http://{host}/ping')
                if response.status_code == 200:
                    print(f'Response from {host}: {response.text}')
            except requests.exceptions.RequestException as e:
                print(f'Error connecting to {host}: {str(e)}')
        time.sleep(60)  # check every 1 minute

if __name__ == '__main__':
    # Start the HTTP server
    app.run(debug=True, host='0.0.0.0', port=8000)
    # Start checking other hosts in the background
    check_other_hosts()
