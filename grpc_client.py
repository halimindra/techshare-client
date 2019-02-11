from datetime import datetime
import random
import logging

import grpc

from pkg import tech_share_pb2, tech_share_pb2_grpc


def run():
    starttime = datetime.now()

    with grpc.insecure_channel('localhost:11000') as channel:
        try:
            stub = tech_share_pb2_grpc.TechShareStub(channel)

            print("-------------- GetPerson --------------")
            print(stub.GetPerson(tech_share_pb2.PersonRequest(id=2)))

            print("-------------- ListPeople --------------")
            people = stub.ListPeople(tech_share_pb2.PeopleRequest(limit=2000000))
            for person in people:
                print(person)

            endtime = datetime.now()
            diff = endtime - starttime
            print("Process took %s seconds" % diff.seconds)

        except grpc.RpcError as e:
            print("Error raised: " + e.details())


if __name__ == '__main__':
    logging.basicConfig()
    run()
