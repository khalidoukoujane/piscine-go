#!/bin/bash

curl https://learn.zone01oujda.ma/assets/superhero/all.json | jq '.[] |select(.id == 70) | .name'

