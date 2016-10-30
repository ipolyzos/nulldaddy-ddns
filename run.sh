#!/bin/bash

glide install
go-wrapper install
go-wrapper run $@