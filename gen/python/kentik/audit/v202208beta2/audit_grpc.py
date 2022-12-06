# Generated by the Protocol Buffers compiler. DO NOT EDIT!
# source: kentik/audit/v202208beta2/audit.proto
# plugin: grpclib.plugin.main
import abc
import typing

import grpclib.const
import grpclib.client
if typing.TYPE_CHECKING:
    import grpclib.server

import google.api.annotations_pb2
import google.api.client_pb2
import google.protobuf.timestamp_pb2
import google.protobuf.duration_pb2
import protoc_gen_openapiv2.options.annotations_pb2
import kentik.core.v202012alpha1.annotations_pb2
import kentik.audit.v202208beta2.audit_pb2


class AuditServiceBase(abc.ABC):

    @abc.abstractmethod
    async def CreateAuditEvents(self, stream: 'grpclib.server.Stream[kentik.audit.v202208beta2.audit_pb2.CreateAuditEventsRequest, kentik.audit.v202208beta2.audit_pb2.CreateAuditEventsResponse]') -> None:
        pass

    @abc.abstractmethod
    async def ListAuditEvents(self, stream: 'grpclib.server.Stream[kentik.audit.v202208beta2.audit_pb2.ListAuditEventsRequest, kentik.audit.v202208beta2.audit_pb2.ListAuditEventsResponse]') -> None:
        pass

    @abc.abstractmethod
    async def GetAuditEvent(self, stream: 'grpclib.server.Stream[kentik.audit.v202208beta2.audit_pb2.GetAuditEventRequest, kentik.audit.v202208beta2.audit_pb2.GetAuditEventResponse]') -> None:
        pass

    def __mapping__(self) -> typing.Dict[str, grpclib.const.Handler]:
        return {
            '/kentik.audit.v202208beta2.AuditService/CreateAuditEvents': grpclib.const.Handler(
                self.CreateAuditEvents,
                grpclib.const.Cardinality.UNARY_UNARY,
                kentik.audit.v202208beta2.audit_pb2.CreateAuditEventsRequest,
                kentik.audit.v202208beta2.audit_pb2.CreateAuditEventsResponse,
            ),
            '/kentik.audit.v202208beta2.AuditService/ListAuditEvents': grpclib.const.Handler(
                self.ListAuditEvents,
                grpclib.const.Cardinality.UNARY_UNARY,
                kentik.audit.v202208beta2.audit_pb2.ListAuditEventsRequest,
                kentik.audit.v202208beta2.audit_pb2.ListAuditEventsResponse,
            ),
            '/kentik.audit.v202208beta2.AuditService/GetAuditEvent': grpclib.const.Handler(
                self.GetAuditEvent,
                grpclib.const.Cardinality.UNARY_UNARY,
                kentik.audit.v202208beta2.audit_pb2.GetAuditEventRequest,
                kentik.audit.v202208beta2.audit_pb2.GetAuditEventResponse,
            ),
        }


class AuditServiceStub:

    def __init__(self, channel: grpclib.client.Channel) -> None:
        self.CreateAuditEvents = grpclib.client.UnaryUnaryMethod(
            channel,
            '/kentik.audit.v202208beta2.AuditService/CreateAuditEvents',
            kentik.audit.v202208beta2.audit_pb2.CreateAuditEventsRequest,
            kentik.audit.v202208beta2.audit_pb2.CreateAuditEventsResponse,
        )
        self.ListAuditEvents = grpclib.client.UnaryUnaryMethod(
            channel,
            '/kentik.audit.v202208beta2.AuditService/ListAuditEvents',
            kentik.audit.v202208beta2.audit_pb2.ListAuditEventsRequest,
            kentik.audit.v202208beta2.audit_pb2.ListAuditEventsResponse,
        )
        self.GetAuditEvent = grpclib.client.UnaryUnaryMethod(
            channel,
            '/kentik.audit.v202208beta2.AuditService/GetAuditEvent',
            kentik.audit.v202208beta2.audit_pb2.GetAuditEventRequest,
            kentik.audit.v202208beta2.audit_pb2.GetAuditEventResponse,
        )