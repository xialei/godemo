#!/usr/bin/env python
#coding=utf-8

import grpc
from concurrent import futures
import time
import compute_pb2
import compute_pb2_grpc

_HOST = '127.0.0.1'
_PORT = '19999'

class compute(compute_pb2_grpc.ComputeServicer):
    def RunTask(self, request, context):
        msg = request.msg
        print("hi compute")
        return compute_pb2.TaskResponse(result=msg)
    
def serve():
    grpcServer = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    compute_pb2_grpc.add_ComputeServicer_to_server(compute(), grpcServer)
    grpcServer.add_insecure_port(_HOST + ':' + _PORT)
    grpcServer.start()
    try:
        while True:
            time.sleep(60*60*24)
    except KeyboardInterrupt:
        grpcServer.stop(0)

if __name__ == '__main__':
    serve()