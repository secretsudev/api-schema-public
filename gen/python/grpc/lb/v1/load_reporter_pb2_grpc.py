# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

from grpc.lb.v1 import load_reporter_pb2 as grpc_dot_lb_dot_v1_dot_load__reporter__pb2


class LoadReporterStub(object):
    """The LoadReporter service.
    """

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.ReportLoad = channel.stream_stream(
                '/grpc.lb.v1.LoadReporter/ReportLoad',
                request_serializer=grpc_dot_lb_dot_v1_dot_load__reporter__pb2.LoadReportRequest.SerializeToString,
                response_deserializer=grpc_dot_lb_dot_v1_dot_load__reporter__pb2.LoadReportResponse.FromString,
                )


class LoadReporterServicer(object):
    """The LoadReporter service.
    """

    def ReportLoad(self, request_iterator, context):
        """Report load from server to lb.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_LoadReporterServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'ReportLoad': grpc.stream_stream_rpc_method_handler(
                    servicer.ReportLoad,
                    request_deserializer=grpc_dot_lb_dot_v1_dot_load__reporter__pb2.LoadReportRequest.FromString,
                    response_serializer=grpc_dot_lb_dot_v1_dot_load__reporter__pb2.LoadReportResponse.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'grpc.lb.v1.LoadReporter', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class LoadReporter(object):
    """The LoadReporter service.
    """

    @staticmethod
    def ReportLoad(request_iterator,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.stream_stream(request_iterator, target, '/grpc.lb.v1.LoadReporter/ReportLoad',
            grpc_dot_lb_dot_v1_dot_load__reporter__pb2.LoadReportRequest.SerializeToString,
            grpc_dot_lb_dot_v1_dot_load__reporter__pb2.LoadReportResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)