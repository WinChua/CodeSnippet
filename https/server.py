import  BaseHTTPServer, SimpleHTTPServer
import ssl

httpd = BaseHTTPServer.HTTPServer(("localhost", 443), SimpleHTTPServer.SimpleHTTPRequestHandler)
httpd.socket = ssl.wrap_socket(httpd.socket, keyfile="./cert/domain.key", certfile="./cert/domain.crt", server_side=True)
httpd.serve_forever()

