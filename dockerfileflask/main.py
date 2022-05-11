from flask import Flask
import socket

app = Flask(__name__)

@app.route('/')
def index():
    return "<h1> Hello from Flask \n" \
	   "<h1> v2 \n" \
           "<h1>Container name: " + socket.gethostname()

if __name__ == "__main__":
    app.run(host='0.0.0.0', port=80)
