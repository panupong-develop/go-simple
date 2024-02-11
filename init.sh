#!/bin/bash

echo "Initialize Projet!"

# Copy .env-example to .env
cp ./.devcontainer/.env-example ./.devcontainer/.env

# Copy config-example.yml to config.yml
cp ./go-simple/config-example.yaml ./go-simple/config.yaml
