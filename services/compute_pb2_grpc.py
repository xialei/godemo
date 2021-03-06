# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
import grpc

import compute_pb2 as compute__pb2


class ComputeStub(object):
    """Missing associated documentation comment in .proto file"""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.RunTask = channel.unary_unary(
                '/services.Compute/RunTask',
                request_serializer=compute__pb2.TaskRequest.SerializeToString,
                response_deserializer=compute__pb2.TaskResponse.FromString,
                )


class ComputeServicer(object):
    """Missing associated documentation comment in .proto file"""

    def RunTask(self, request, context):
        """Missing associated documentation comment in .proto file"""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_ComputeServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'RunTask': grpc.unary_unary_rpc_method_handler(
                    servicer.RunTask,
                    request_deserializer=compute__pb2.TaskRequest.FromString,
                    response_serializer=compute__pb2.TaskResponse.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'services.Compute', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class Compute(object):
    """Missing associated documentation comment in .proto file"""

    @staticmethod
    def RunTask(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/services.Compute/RunTask',
            compute__pb2.TaskRequest.SerializeToString,
            compute__pb2.TaskResponse.FromString,
            options, channel_credentials,
            call_credentials, compression, wait_for_ready, timeout, metadata)
