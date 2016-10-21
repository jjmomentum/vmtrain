
docker build -t queue-monitor .
docker run --net=host queue-monitor monitor --brokers=http://localhost:8081 --topic=reserveation --datamanager=http://localhost:6001

