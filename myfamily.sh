#! /bin/bash
a=$(curl -s https://01.alem.school/assets/superhero/all.json | jq -r ".[] | select(.id==${HERO_ID})" | jq .connections | jq .relatives)
echo $a | tr -d '"'
