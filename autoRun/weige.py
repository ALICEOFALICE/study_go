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
dst = vika.datasheet("dstn3TYmEbWgl3EYbQ")

//vika_upadta_dagta =
record = dst.records.create({
                            	"提交者：": env["commits"][0]["author"]["name"],
                            	"邮箱：": env["commits"][0]["author"]["email"],
                            	"提交者用户名：": env["commits"][0]["author"]["username"],
                            	"提交注释：": env["commits"][0]["message"],
                            	"commit_ID": str(env["commits"][0]["id"]),
                            	"提交时间：": env["commits"][0]["timestamp"],
                            	"仓库名称：": env["repository"]["name"],
                            	"仓库简介：": env["repository"]["description"],
                            	"编程语言：": env["repository"]["language"],
                            	"主分支：": env["repository"]["master_branch"],
                            	"链接：": env["repository"]["html_url"],
                            })
