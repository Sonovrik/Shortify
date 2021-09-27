
COMPOSE_FILE=docker-compose.yml
ENV=.env
PATH_TO_BUILD=./

all:
	docker-compose -f $(PATH_TO_BUILD)$(COMPOSE_FILE) --env-file $(PATH_TO_BUILD)$(ENV) up --build -d

clean:
	docker-compose -f $(PATH_TO_BUILD)$(COMPOSE_FILE) --env-file $(PATH_TO_BUILD)$(ENV) down


fclean: clean
	docker system prune -f -a --volumes

re: fclean all

.PHONY: all clean fclean re