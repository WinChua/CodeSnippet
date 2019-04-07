gencacert:
	openssl req -newkey rsa:2048 -nodes -sha256 -keyout ca.key -x509 -days 365 -out ca.crt

gencert:
	openssl req -newkey rsa:2048 -nodes -sha256 -keyout domain.key -new -out domain.csr

signature:
	openssl x509 -CA ca.crt -CAkey ca.key -in domain.csr -req -days 365 -out domain.crt -CAcreateserial -sha256

check:
	openssl x509 -text -noout -in ca.crt
	openssl x509 -text -noout -in domain.crt


