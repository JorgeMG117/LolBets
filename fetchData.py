"""
import mwclient
site = mwclient.Site('lol.fandom.com', path='/')

response = site.api('cargoquery',
	limit = 'max',
	tables = "ScoreboardGames=SG",
	fields = "SG.Tournament, SG.DateTime_UTC, SG.Team1, SG.Team2",
	where = "SG.DateTime_UTC >= '2022-06-24 00:00:00'" #Results after Aug 1, 2019
)

response = site.api('cargoquery',
	limit = 'max',
	tables = "MatchSchedule=SG",
	fields = "SG.DateTime_UTC, SG.Team1, SG.Team2",
	where = "SG.DateTime_UTC >= '2022-06-24 00:00:00'" #Results after Aug 1, 2019
)

print(response)
"""

import requests
from bs4 import BeautifulSoup

r = requests.get('https://superliga.lvp.global/')
soup = BeautifulSoup(r.text, 'lxml')
footer_links = soup.find_all(class_="item match round15")
#print(footer_links[0])
hijos = footer_links[0].contents
for child in hijos:
    if child.name:  # Ignoramos los saltos de línea
        print(f'{child.name}')
        print(f'{child.attrs}')

aux = hijos[1]
aux2 = aux.contents
for child in aux2:
    if child.name:  # Ignoramos los saltos de línea
        print(f'{child.name}')
        print(f'{child.attrs}')
        print(child)
#print(aux.string)
"""div_main = soup.div
hijos = div_main.contents
child = hijos[1]
for child in hijos:
    if child.name:  # Ignoramos los saltos de línea
        print(f'{child.name}')
        print(f'{child.attrs}')


hijos = child.contents

for child in hijos:
    if child.name:  # Ignoramos los saltos de línea
        print(f'{child.name}')
        print(f'{child.attrs}')
        print(f'{child}')

child = hijos[5]
print(child)
"""
