#!/bin/bash

 curl https://learn.zone01oujda.ma/assets/superhero/all.json | jq '.[] | select(.id == 170) | .name,.powerstats.power,.appearance.gender'

