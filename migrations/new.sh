#!/usr/bin/env bash

TS=$(date +%s)

touch sql/$TS\_$1.up.sql sql/$TS\_$1.down.sql
