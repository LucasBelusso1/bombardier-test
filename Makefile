# Health
bomb-standard-health:
	bombardier -c 125 -n 100000 http://localhost:8080/health

bomb-gorilla-health:
	bombardier -c 125 -n 100000 http://localhost:8081/health

bomb-gin-health:
	bombardier -c 125 -n 100000 http://localhost:8082/health

bomb-chi-health:
	bombardier -c 125 -n 100000 http://localhost:8083/health

# withBodyAndHeader
bomb-standard-withBodyAndHeader:
	bombardier -c 125 -n 10000 -H 'Content-Type: application/json' -H 'x-api-key:0c4a5e30-2468-4084-9705-12e2ff13e7ae' -m POST -b '{"message":"Hello world"}' -l http://localhost:8080/withBodyAndHeader

bomb-gorilla-withBodyAndHeader:
	bombardier -c 125 -n 10000 -H 'Content-Type: application/json' -H 'x-api-key:0c4a5e30-2468-4084-9705-12e2ff13e7ae' -m POST -b '{"message":"Hello world"}' -l http://localhost:8081/withBodyAndHeader

bomb-gin-withBodyAndHeader:
	bombardier -c 125 -n 10000 -H 'Content-Type: application/json' -H 'x-api-key:0c4a5e30-2468-4084-9705-12e2ff13e7ae' -m POST -b '{"message":"Hello world"}' -l http://localhost:8082/withBodyAndHeader

bomb-chi-withBodyAndHeader:
	bombardier -c 125 -n 10000 -H 'Content-Type: application/json' -H 'x-api-key:0c4a5e30-2468-4084-9705-12e2ff13e7ae' -m POST -b '{"message":"Hello world"}' -l http://localhost:8083/withBodyAndHeader

