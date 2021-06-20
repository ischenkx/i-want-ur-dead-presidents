import grpc

import agregator_pb2
import agregator_pb2_grpc

channel = grpc.insecure_channel('localhost:50055')

stub = agregator_pb2_grpc.AgregatorStub(channel)

product = agregator_pb2.Product(id = '1', inn='7813325520')

response = stub.Get(product)

print(response)