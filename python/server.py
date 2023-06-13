"""
Author: Dhruva Agrawal
Author E-mail: dhruva_agrawal@outlook.com
"""
from http import HTTPStatus
from http.server import HTTPServer, BaseHTTPRequestHandler

HOST = "127.0.0.1"


class BaseHttpServer(BaseHTTPRequestHandler):
    def do_GET(self):
        self.send_response(HTTPStatus.OK)
        self.send_header("Content-type", "text/html")
        self.end_headers()

        self.wfile.write(
            bytes("<html><body><h1>Hello world!</h1></body></html>", "utf-8")
        )


server = HTTPServer((HOST, 8000), BaseHttpServer)
print("server:", server.server_name, server.server_address)
server.serve_forever()
server.server_close()
