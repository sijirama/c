#!/bin/bash

while [[ "$#" -gt 0 ]]; do
  case $1 in
    -s|--start)
      docker run -d -p 27017:27017 --name=rest mongo
  docker ps -a
      ;;
    -r|--remove)
      # Remove Docker container
      docker stop rest
      docker rm rest
  docker ps -a
      ;;
    *)
      echo "Unknown option: $1"
      exit 1
      ;;
  esac
  shift
done


#sh mongo.sh -s // start the container
#sh mongo.sh -r // stop the container
