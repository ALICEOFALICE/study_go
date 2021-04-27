#!/usr/bin/python3.7
# -*- coding: utf-8 -*-
#
# Copyright (C) 2021 桜火, Inc. All Rights Reserved 
#
# @Time    : 2021/4/26 21:35
# @Author  : 桜火
# @Email   : xie@loli.fit
# @File    : aba.py
# @Software: PyCharm
import json
from vika import Vika
import os
json_file = open("/home/runner/work/_temp/_github_workflow/event.json",mode="r")
json_file = json_file.read()
env = json.loads(json_file)
vika = Vika("uskK7MnF6tOf89ammoV851g")
# 通过 datasheetId 来指定要从哪张维格表操作数据。
datasheet = vika.datasheet("dstn3TYmEbWgl3EYbQ", field_key="id")

row = datasheet.records.create({
  "fldp3Jyag9dcT": str(env["commits"][0]["id"]),
  "fldIilczY4AtL": env["repository"]["name"],
  "fld4CxF56Aiqg": env["repository"]["description"],
  "fldcNKW9tOQob": env["commits"][0]["author"]["name"],
  "fldGoBim2N4ru": env["commits"][0]["author"]["username"],
  "fldoDfcbS2FZT": env["commits"][0]["author"]["email"],
  "fldSMyuxmszuo": env["repository"]["language"],
  "fldUJsyEaYrET": env["commits"][0]["message"],
  "fldDPC0Cjm6W3": env["commits"][0]["timestamp"],
  "fldhpQbZKKfaW": env["repository"]["master_branch"],
  "fld3fYMdpSROJ": env["repository"]["html_url"],
})

