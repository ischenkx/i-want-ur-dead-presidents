import time
from concurrent import futures

import grpc

import aggregator_pb2_grpc
import aggregator_pb2
from rootine import check_on_company, check_fin_koefs, check_coart, check_scoring, analysis


class AggregatorServicer(aggregator_pb2_grpc.AggregatorServicer):

    def Get(self, request, context):
        inn = request.inn
        try:
            exists, short_name, long_name, ogrn = check_on_company(inn)
            if not exists:
                return aggregator_pb2.Response(id=request.id, inn=inn, overallScore=-1, shortCompanyName='',
                                              fullCompanyName='')

            avg_fin_k = check_fin_koefs(inn)
            lawsuit_cnt = check_coart(inn)
            smart_scores = check_scoring(inn)
            rating = analysis(avg_fin_k, lawsuit_cnt, smart_scores)
        except:
            rating = -1
            short_name, long_name = '', ''

        return aggregator_pb2.Response(id=request.id, inn=inn, overallScore=rating, shortCompanyName=short_name,
                                      fullCompanyName=long_name)


# create a gRPC server
server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))

# use the generated function `add_AggregatorServicer_to_server`
# to add the defined class to the server
aggregator_pb2_grpc.add_AggregatorServicer_to_server(
    AggregatorServicer(), server)

# listen on port 50055
print('Starting server. Listening on port 50055.')
server.add_insecure_port('localhost:50055')
server.start()

# since server.start() will not block,
# a sleep-loop is added to keep alive
try:
    while True:
        time.sleep(86400)
except KeyboardInterrupt:
    server.stop(0)
