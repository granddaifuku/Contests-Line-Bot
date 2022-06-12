#!/bin/bash

if [ $1 = "enterdb" ]; then
	PGPASSWORD=password psql -h 127.0.0.1 -p 5432 -U postgres test
elif [ $1 = "cleartable" ]; then
	psql -h 127.0.0.1 -p 5432 -f ./config/init.sql -U postgres -W password -d test
else
	echo "I do not know that command."
fi
