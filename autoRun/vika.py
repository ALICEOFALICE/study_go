from vika import Vika
vika = Vika("uskK7MnF6tOf89ammoV851g")
dst = vika.datasheet("https://vika.cn/space/spcNuHrHJgQVZ/workbench/dstXHYNcifpEsJuMjZ/viwvYeltzpZ5Z")
record = dst.records.create({"提交ID":"new record from Python SDK"})
print(GITHUB_EVENT_PATH)