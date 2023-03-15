.PHONY: all test clean

run:	
	sudo docker-compose build --parallel --force-rm
	sudo docker-compose up --remove-orphans --force-recreate
