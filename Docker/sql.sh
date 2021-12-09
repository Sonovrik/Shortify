#!/bin/bash

user=$1
pass=$2
db_name=$3


if [ \"$user\" = \"\" ] || [ \"$pass\" = \"\" ] || [ \"$db_name\" = \"\" ]; then \
  echo "Enter user, password and db_name"; \
else \
  echo "ALTER USER $user PASSWORD '$pass';
        ALTER USER $user WITH CREATEDB;
        DROP DATABASE $db_name IF EXISTS;
        CREATE DATABASE $db_name;" > initDB.sql; \
  fi;




