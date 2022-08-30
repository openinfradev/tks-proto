# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: common.proto
"""Generated protocol buffer code."""
from google.protobuf.internal import enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.protobuf import timestamp_pb2 as google_dot_protobuf_dot_timestamp__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x0c\x63ommon.proto\x12\x06tks_pb\x1a\x1fgoogle/protobuf/timestamp.proto\"\x14\n\x05\x45rror\x12\x0b\n\x03msg\x18\x01 \x01(\t\"\xba\x02\n\x08\x41ppGroup\x12\x14\n\x0c\x61pp_group_id\x18\x01 \x01(\t\x12\x16\n\x0e\x61pp_group_name\x18\x02 \x01(\t\x12\"\n\x04type\x18\x03 \x01(\x0e\x32\x14.tks_pb.AppGroupType\x12\x12\n\ncluster_id\x18\x04 \x01(\t\x12\x13\n\x0bworkflow_id\x18\x05 \x01(\t\x12&\n\x06status\x18\x06 \x01(\x0e\x32\x16.tks_pb.AppGroupStatus\x12\x13\n\x0bstatus_desc\x18\x07 \x01(\t\x12\x16\n\x0e\x65xternal_label\x18\x08 \x01(\t\x12.\n\ncreated_at\x18\t \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12.\n\nupdated_at\x18\n \x01(\x0b\x32\x1a.google.protobuf.Timestamp\"\xd6\x01\n\x0b\x41pplication\x12\x0e\n\x06\x61pp_id\x18\x01 \x01(\t\x12\x1d\n\x04type\x18\x02 \x01(\x0e\x32\x0f.tks_pb.AppType\x12\x14\n\x0c\x61pp_group_id\x18\x03 \x01(\t\x12\x10\n\x08\x65ndpoint\x18\x04 \x01(\t\x12\x10\n\x08metadata\x18\x05 \x01(\t\x12.\n\ncreated_at\x18\x06 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12.\n\nupdated_at\x18\x07 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\"\xd6\x02\n\x07\x43luster\x12\n\n\x02id\x18\x01 \x01(\t\x12\x0c\n\x04name\x18\x02 \x01(\t\x12.\n\ncreated_at\x18\x03 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12.\n\nupdated_at\x18\x04 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12\x13\n\x0bworkflow_id\x18\x05 \x01(\t\x12%\n\x06status\x18\x06 \x01(\x0e\x32\x15.tks_pb.ClusterStatus\x12\x13\n\x0bstatus_desc\x18\x07 \x01(\t\x12\x13\n\x0b\x63ontract_id\x18\x08 \x01(\t\x12\x0e\n\x06\x63sp_id\x18\t \x01(\t\x12!\n\x04\x63onf\x18\n \x01(\x0b\x32\x13.tks_pb.ClusterConf\x12$\n\napp_groups\x18\x0b \x03(\x0b\x32\x10.tks_pb.AppGroup\x12\x12\n\nkubeconfig\x18\x0c \x01(\t\"y\n\x0e\x43lusterRawConf\x12\x14\n\x0cssh_key_name\x18\x01 \x01(\t\x12\x0e\n\x06region\x18\x02 \x01(\t\x12\x11\n\tnum_of_az\x18\x03 \x01(\x05\x12\x14\n\x0cmachine_type\x18\x04 \x01(\t\x12\x18\n\x10machine_replicas\x18\x05 \x01(\x05\"\x8e\x01\n\x0b\x43lusterConf\x12\x14\n\x0cssh_key_name\x18\x01 \x01(\t\x12\x0e\n\x06region\x18\x02 \x01(\t\x12\x11\n\tnum_of_az\x18\x03 \x01(\x05\x12\x14\n\x0cmachine_type\x18\x04 \x01(\t\x12\x17\n\x0fmin_size_per_az\x18\x05 \x01(\x05\x12\x17\n\x0fmax_size_per_az\x18\x06 \x01(\x05\"\xcb\x02\n\x0b\x41ppServeApp\x12\n\n\x02id\x18\x01 \x01(\t\x12\x0c\n\x04name\x18\x02 \x01(\t\x12\x13\n\x0b\x63ontract_id\x18\x03 \x01(\t\x12\x0f\n\x07version\x18\x04 \x01(\t\x12\x11\n\ttask_type\x18\x05 \x01(\t\x12\x0e\n\x06status\x18\x06 \x01(\t\x12\x0e\n\x06output\x18\x07 \x01(\t\x12\x14\n\x0c\x61rtifact_url\x18\x08 \x01(\t\x12\x11\n\timage_url\x18\t \x01(\t\x12\x14\n\x0c\x65ndpoint_url\x18\n \x01(\t\x12\x19\n\x11target_cluster_id\x18\x0b \x01(\t\x12\x0f\n\x07profile\x18\x0c \x01(\t\x12.\n\ncreated_at\x18\r \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12.\n\nupdated_at\x18\x0e \x01(\x0b\x32\x1a.google.protobuf.Timestamp\"\x17\n\tIDRequest\x12\n\n\x02id\x18\x01 \x01(\t\"J\n\x0eSimpleResponse\x12\x1a\n\x04\x63ode\x18\x01 \x01(\x0e\x32\x0c.tks_pb.Code\x12\x1c\n\x05\x65rror\x18\x02 \x01(\x0b\x32\r.tks_pb.Error\"T\n\x0bIDsResponse\x12\x1a\n\x04\x63ode\x18\x01 \x01(\x0e\x32\x0c.tks_pb.Code\x12\x1c\n\x05\x65rror\x18\x02 \x01(\x0b\x32\r.tks_pb.Error\x12\x0b\n\x03ids\x18\x03 \x03(\t\"R\n\nIDResponse\x12\x1a\n\x04\x63ode\x18\x01 \x01(\x0e\x32\x0c.tks_pb.Code\x12\x1c\n\x05\x65rror\x18\x02 \x01(\x0b\x32\r.tks_pb.Error\x12\n\n\x02id\x18\x03 \x01(\t\"i\n\x0cKeycloakInfo\x12\x12\n\ncluster_id\x18\x01 \x01(\t\x12\r\n\x05realm\x18\x02 \x01(\t\x12\x11\n\tclient_id\x18\x03 \x01(\t\x12\x0e\n\x06secret\x18\x04 \x01(\t\x12\x13\n\x0bprivate_key\x18\x05 \x01(\t*\xb4\x02\n\x04\x43ode\x12\x12\n\x0eOK_UNSPECIFIED\x10\x00\x12\r\n\tCANCELLED\x10\x01\x12\x0b\n\x07UNKNOWN\x10\x02\x12\x14\n\x10INVALID_ARGUMENT\x10\x03\x12\x15\n\x11\x44\x45\x41\x44LINE_EXCEEDED\x10\x04\x12\r\n\tNOT_FOUND\x10\x05\x12\x12\n\x0e\x41LREADY_EXISTS\x10\x06\x12\x15\n\x11PERMISSION_DENIED\x10\x07\x12\x13\n\x0fUNAUTHENTICATED\x10\x10\x12\x16\n\x12RESOURCE_EXHAUSTED\x10\x08\x12\x17\n\x13\x46\x41ILED_PRECONDITION\x10\t\x12\x0b\n\x07\x41\x42ORTED\x10\n\x12\x10\n\x0cOUT_OF_RANGE\x10\x0b\x12\x11\n\rUNIMPLEMENTED\x10\x0c\x12\x0c\n\x08INTERNAL\x10\r\x12\x0f\n\x0bUNAVAILABLE\x10\x0e*\xa0\x01\n\x0e\x41ppGroupStatus\x12\x19\n\x15\x41PP_GROUP_UNSPECIFIED\x10\x00\x12\x18\n\x14\x41PP_GROUP_INSTALLING\x10\x01\x12\x15\n\x11\x41PP_GROUP_RUNNING\x10\x02\x12\x16\n\x12\x41PP_GROUP_DELETING\x10\x04\x12\x15\n\x11\x41PP_GROUP_DELETED\x10\x05\x12\x13\n\x0f\x41PP_GROUP_ERROR\x10\x06*c\n\rClusterStatus\x12\x0f\n\x0bUNSPECIFIED\x10\x00\x12\x0e\n\nINSTALLING\x10\x01\x12\x0b\n\x07RUNNING\x10\x02\x12\x0c\n\x08\x44\x45LETING\x10\x04\x12\x0b\n\x07\x44\x45LETED\x10\x05\x12\t\n\x05\x45RROR\x10\x06*\xb5\x01\n\x07\x41ppType\x12\x12\n\x0e\x45P_UNSPECIFIED\x10\x00\x12\n\n\x06THANOS\x10\x01\x12\x0e\n\nPROMETHEUS\x10\x02\x12\x0b\n\x07GRAFANA\x10\x03\x12\t\n\x05KIALI\x10\x04\x12\n\n\x06KIBANA\x10\x05\x12\x10\n\x0c\x45LASTICSERCH\x10\x06\x12\x11\n\rCLOUD_CONSOLE\x10\x07\x12\x0b\n\x07HORIZON\x10\x08\x12\n\n\x06JAEGER\x10\t\x12\x18\n\x14KUBERNETES_DASHBOARD\x10\n*P\n\x0c\x41ppGroupType\x12\x18\n\x14\x41PP_TYPE_UNSPECIFIED\x10\x00\x12\x07\n\x03LMA\x10\x01\x12\x10\n\x0cSERVICE_MESH\x10\x02\x12\x0b\n\x07LMA_EFK\x10\x03*?\n\x07\x43spType\x12\x17\n\x13\x43SPTYPE_UNSPECIFIED\x10\x00\x12\x07\n\x03\x41WS\x10\x01\x12\x07\n\x03GCP\x10\x02\x12\t\n\x05\x41ZURE\x10\x03\x42*Z(github.com/openinfradev/tks-proto/tks_pbb\x06proto3')

_CODE = DESCRIPTOR.enum_types_by_name['Code']
Code = enum_type_wrapper.EnumTypeWrapper(_CODE)
_APPGROUPSTATUS = DESCRIPTOR.enum_types_by_name['AppGroupStatus']
AppGroupStatus = enum_type_wrapper.EnumTypeWrapper(_APPGROUPSTATUS)
_CLUSTERSTATUS = DESCRIPTOR.enum_types_by_name['ClusterStatus']
ClusterStatus = enum_type_wrapper.EnumTypeWrapper(_CLUSTERSTATUS)
_APPTYPE = DESCRIPTOR.enum_types_by_name['AppType']
AppType = enum_type_wrapper.EnumTypeWrapper(_APPTYPE)
_APPGROUPTYPE = DESCRIPTOR.enum_types_by_name['AppGroupType']
AppGroupType = enum_type_wrapper.EnumTypeWrapper(_APPGROUPTYPE)
_CSPTYPE = DESCRIPTOR.enum_types_by_name['CspType']
CspType = enum_type_wrapper.EnumTypeWrapper(_CSPTYPE)
OK_UNSPECIFIED = 0
CANCELLED = 1
UNKNOWN = 2
INVALID_ARGUMENT = 3
DEADLINE_EXCEEDED = 4
NOT_FOUND = 5
ALREADY_EXISTS = 6
PERMISSION_DENIED = 7
UNAUTHENTICATED = 16
RESOURCE_EXHAUSTED = 8
FAILED_PRECONDITION = 9
ABORTED = 10
OUT_OF_RANGE = 11
UNIMPLEMENTED = 12
INTERNAL = 13
UNAVAILABLE = 14
APP_GROUP_UNSPECIFIED = 0
APP_GROUP_INSTALLING = 1
APP_GROUP_RUNNING = 2
APP_GROUP_DELETING = 4
APP_GROUP_DELETED = 5
APP_GROUP_ERROR = 6
UNSPECIFIED = 0
INSTALLING = 1
RUNNING = 2
DELETING = 4
DELETED = 5
ERROR = 6
EP_UNSPECIFIED = 0
THANOS = 1
PROMETHEUS = 2
GRAFANA = 3
KIALI = 4
KIBANA = 5
ELASTICSERCH = 6
CLOUD_CONSOLE = 7
HORIZON = 8
JAEGER = 9
KUBERNETES_DASHBOARD = 10
APP_TYPE_UNSPECIFIED = 0
LMA = 1
SERVICE_MESH = 2
LMA_EFK = 3
CSPTYPE_UNSPECIFIED = 0
AWS = 1
GCP = 2
AZURE = 3


_ERROR = DESCRIPTOR.message_types_by_name['Error']
_APPGROUP = DESCRIPTOR.message_types_by_name['AppGroup']
_APPLICATION = DESCRIPTOR.message_types_by_name['Application']
_CLUSTER = DESCRIPTOR.message_types_by_name['Cluster']
_CLUSTERRAWCONF = DESCRIPTOR.message_types_by_name['ClusterRawConf']
_CLUSTERCONF = DESCRIPTOR.message_types_by_name['ClusterConf']
_APPSERVEAPP = DESCRIPTOR.message_types_by_name['AppServeApp']
_IDREQUEST = DESCRIPTOR.message_types_by_name['IDRequest']
_SIMPLERESPONSE = DESCRIPTOR.message_types_by_name['SimpleResponse']
_IDSRESPONSE = DESCRIPTOR.message_types_by_name['IDsResponse']
_IDRESPONSE = DESCRIPTOR.message_types_by_name['IDResponse']
_KEYCLOAKINFO = DESCRIPTOR.message_types_by_name['KeycloakInfo']
Error = _reflection.GeneratedProtocolMessageType('Error', (_message.Message,), {
  'DESCRIPTOR' : _ERROR,
  '__module__' : 'common_pb2'
  # @@protoc_insertion_point(class_scope:tks_pb.Error)
  })
_sym_db.RegisterMessage(Error)

AppGroup = _reflection.GeneratedProtocolMessageType('AppGroup', (_message.Message,), {
  'DESCRIPTOR' : _APPGROUP,
  '__module__' : 'common_pb2'
  # @@protoc_insertion_point(class_scope:tks_pb.AppGroup)
  })
_sym_db.RegisterMessage(AppGroup)

Application = _reflection.GeneratedProtocolMessageType('Application', (_message.Message,), {
  'DESCRIPTOR' : _APPLICATION,
  '__module__' : 'common_pb2'
  # @@protoc_insertion_point(class_scope:tks_pb.Application)
  })
_sym_db.RegisterMessage(Application)

Cluster = _reflection.GeneratedProtocolMessageType('Cluster', (_message.Message,), {
  'DESCRIPTOR' : _CLUSTER,
  '__module__' : 'common_pb2'
  # @@protoc_insertion_point(class_scope:tks_pb.Cluster)
  })
_sym_db.RegisterMessage(Cluster)

ClusterRawConf = _reflection.GeneratedProtocolMessageType('ClusterRawConf', (_message.Message,), {
  'DESCRIPTOR' : _CLUSTERRAWCONF,
  '__module__' : 'common_pb2'
  # @@protoc_insertion_point(class_scope:tks_pb.ClusterRawConf)
  })
_sym_db.RegisterMessage(ClusterRawConf)

ClusterConf = _reflection.GeneratedProtocolMessageType('ClusterConf', (_message.Message,), {
  'DESCRIPTOR' : _CLUSTERCONF,
  '__module__' : 'common_pb2'
  # @@protoc_insertion_point(class_scope:tks_pb.ClusterConf)
  })
_sym_db.RegisterMessage(ClusterConf)

AppServeApp = _reflection.GeneratedProtocolMessageType('AppServeApp', (_message.Message,), {
  'DESCRIPTOR' : _APPSERVEAPP,
  '__module__' : 'common_pb2'
  # @@protoc_insertion_point(class_scope:tks_pb.AppServeApp)
  })
_sym_db.RegisterMessage(AppServeApp)

IDRequest = _reflection.GeneratedProtocolMessageType('IDRequest', (_message.Message,), {
  'DESCRIPTOR' : _IDREQUEST,
  '__module__' : 'common_pb2'
  # @@protoc_insertion_point(class_scope:tks_pb.IDRequest)
  })
_sym_db.RegisterMessage(IDRequest)

SimpleResponse = _reflection.GeneratedProtocolMessageType('SimpleResponse', (_message.Message,), {
  'DESCRIPTOR' : _SIMPLERESPONSE,
  '__module__' : 'common_pb2'
  # @@protoc_insertion_point(class_scope:tks_pb.SimpleResponse)
  })
_sym_db.RegisterMessage(SimpleResponse)

IDsResponse = _reflection.GeneratedProtocolMessageType('IDsResponse', (_message.Message,), {
  'DESCRIPTOR' : _IDSRESPONSE,
  '__module__' : 'common_pb2'
  # @@protoc_insertion_point(class_scope:tks_pb.IDsResponse)
  })
_sym_db.RegisterMessage(IDsResponse)

IDResponse = _reflection.GeneratedProtocolMessageType('IDResponse', (_message.Message,), {
  'DESCRIPTOR' : _IDRESPONSE,
  '__module__' : 'common_pb2'
  # @@protoc_insertion_point(class_scope:tks_pb.IDResponse)
  })
_sym_db.RegisterMessage(IDResponse)

KeycloakInfo = _reflection.GeneratedProtocolMessageType('KeycloakInfo', (_message.Message,), {
  'DESCRIPTOR' : _KEYCLOAKINFO,
  '__module__' : 'common_pb2'
  # @@protoc_insertion_point(class_scope:tks_pb.KeycloakInfo)
  })
_sym_db.RegisterMessage(KeycloakInfo)

if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'Z(github.com/openinfradev/tks-proto/tks_pb'
  _CODE._serialized_start=1939
  _CODE._serialized_end=2247
  _APPGROUPSTATUS._serialized_start=2250
  _APPGROUPSTATUS._serialized_end=2410
  _CLUSTERSTATUS._serialized_start=2412
  _CLUSTERSTATUS._serialized_end=2511
  _APPTYPE._serialized_start=2514
  _APPTYPE._serialized_end=2695
  _APPGROUPTYPE._serialized_start=2697
  _APPGROUPTYPE._serialized_end=2777
  _CSPTYPE._serialized_start=2779
  _CSPTYPE._serialized_end=2842
  _ERROR._serialized_start=57
  _ERROR._serialized_end=77
  _APPGROUP._serialized_start=80
  _APPGROUP._serialized_end=394
  _APPLICATION._serialized_start=397
  _APPLICATION._serialized_end=611
  _CLUSTER._serialized_start=614
  _CLUSTER._serialized_end=956
  _CLUSTERRAWCONF._serialized_start=958
  _CLUSTERRAWCONF._serialized_end=1079
  _CLUSTERCONF._serialized_start=1082
  _CLUSTERCONF._serialized_end=1224
  _APPSERVEAPP._serialized_start=1227
  _APPSERVEAPP._serialized_end=1558
  _IDREQUEST._serialized_start=1560
  _IDREQUEST._serialized_end=1583
  _SIMPLERESPONSE._serialized_start=1585
  _SIMPLERESPONSE._serialized_end=1659
  _IDSRESPONSE._serialized_start=1661
  _IDSRESPONSE._serialized_end=1745
  _IDRESPONSE._serialized_start=1747
  _IDRESPONSE._serialized_end=1829
  _KEYCLOAKINFO._serialized_start=1831
  _KEYCLOAKINFO._serialized_end=1936
# @@protoc_insertion_point(module_scope)
