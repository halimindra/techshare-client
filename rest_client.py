from datetime import datetime
import logging

import requests


def run():
    starttime = datetime.now()

    r = requests.get("http://localhost:10000/people/1")
    print("-------------- GetPerson --------------")
    print(r.json())

    r = requests.get("http://localhost:10000/people?limit=1000000")
    print("-------------- ListPeople --------------")
    for person in r.json():
        print(person)

    endtime = datetime.now()
    diff = endtime - starttime
    print("Process took %s seconds" % diff.seconds)

if __name__ == '__main__':
    logging.basicConfig()
    run()