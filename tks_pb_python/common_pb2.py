# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: common.proto
"""Generated protocol buffer code."""
from google.protobuf.internal import builder as _builder
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.protobuf import timestamp_pb2 as google_dot_protobuf_dot_timestamp__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x0c\x63ommon.proto\x12\x06tks_pb\x1a\x1fgoogle/protobuf/timestamp.proto\"\x14\n\x05\x45rror\x12\x0b\n\x03msg\x18\x01 \x01(\t\"\xe0\x02\n\x08\x41ppGroup\x12\x14\n\x0c\x61pp_group_id\x18\x01 \x01(\t\x12\x16\n\x0e\x61pp_group_name\x18\x02 \x01(\t\x12\"\n\x04type\x18\x03 \x01(\x0e\x32\x14.tks_pb.AppGroupType\x12\x12\n\ncluster_id\x18\x04 \x01(\t\x12\x13\n\x0bworkflow_id\x18\x05 \x01(\t\x12&\n\x06status\x18\x06 \x01(\x0e\x32\x16.tks_pb.AppGroupStatus\x12\x13\n\x0bstatus_desc\x18\x07 \x01(\t\x12\x16\n\x0e\x65xternal_label\x18\x08 \x01(\t\x12.\n\ncreated_at\x18\t \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12.\n\nupdated_at\x18\n \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12\x0f\n\x07\x63reator\x18\x0b \x01(\t\x12\x13\n\x0b\x64\x65scription\x18\x0c \x01(\t\"\xd6\x01\n\x0b\x41pplication\x12\x0e\n\x06\x61pp_id\x18\x01 \x01(\t\x12\x1d\n\x04type\x18\x02 \x01(\x0e\x32\x0f.tks_pb.AppType\x12\x14\n\x0c\x61pp_group_id\x18\x03 \x01(\t\x12\x10\n\x08\x65ndpoint\x18\x04 \x01(\t\x12\x10\n\x08metadata\x18\x05 \x01(\t\x12.\n\ncreated_at\x18\x06 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12.\n\nupdated_at\x18\x07 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\"\xfc\x02\n\x07\x43luster\x12\n\n\x02id\x18\x01 \x01(\t\x12\x0c\n\x04name\x18\x02 \x01(\t\x12.\n\ncreated_at\x18\x03 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12.\n\nupdated_at\x18\x04 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12\x13\n\x0bworkflow_id\x18\x05 \x01(\t\x12%\n\x06status\x18\x06 \x01(\x0e\x32\x15.tks_pb.ClusterStatus\x12\x13\n\x0bstatus_desc\x18\x07 \x01(\t\x12\x13\n\x0b\x63ontract_id\x18\x08 \x01(\t\x12\x0e\n\x06\x63sp_id\x18\t \x01(\t\x12!\n\x04\x63onf\x18\n \x01(\x0b\x32\x13.tks_pb.ClusterConf\x12$\n\napp_groups\x18\x0b \x03(\x0b\x32\x10.tks_pb.AppGroup\x12\x12\n\nkubeconfig\x18\x0c \x01(\t\x12\x0f\n\x07\x63reator\x18\r \x01(\t\x12\x13\n\x0b\x64\x65scription\x18\x0e \x01(\t\"y\n\x0e\x43lusterRawConf\x12\x14\n\x0cssh_key_name\x18\x01 \x01(\t\x12\x0e\n\x06region\x18\x02 \x01(\t\x12\x11\n\tnum_of_az\x18\x03 \x01(\x05\x12\x14\n\x0cmachine_type\x18\x04 \x01(\t\x12\x18\n\x10machine_replicas\x18\x05 \x01(\x05\"\x8e\x01\n\x0b\x43lusterConf\x12\x14\n\x0cssh_key_name\x18\x01 \x01(\t\x12\x0e\n\x06region\x18\x02 \x01(\t\x12\x11\n\tnum_of_az\x18\x03 \x01(\x05\x12\x14\n\x0cmachine_type\x18\x04 \x01(\t\x12\x17\n\x0fmin_size_per_az\x18\x05 \x01(\x05\x12\x17\n\x0fmax_size_per_az\x18\x06 \x01(\x05\"\xfd\x01\n\x0b\x41ppServeApp\x12\n\n\x02id\x18\x01 \x01(\t\x12\x0c\n\x04name\x18\x02 \x01(\t\x12\x13\n\x0b\x63ontract_id\x18\x03 \x01(\t\x12\x0c\n\x04type\x18\x04 \x01(\t\x12\x10\n\x08\x61pp_type\x18\x05 \x01(\t\x12\x14\n\x0c\x65ndpoint_url\x18\x06 \x01(\t\x12\x19\n\x11target_cluster_id\x18\x07 \x01(\t\x12\x0e\n\x06status\x18\x08 \x01(\t\x12.\n\ncreated_at\x18\t \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12.\n\nupdated_at\x18\n \x01(\x0b\x32\x1a.google.protobuf.Timestamp\"\xd7\x02\n\x0f\x41ppServeAppTask\x12\n\n\x02id\x18\x01 \x01(\t\x12\x18\n\x10\x61pp_serve_app_id\x18\x02 \x01(\t\x12\x0f\n\x07version\x18\x03 \x01(\t\x12\x0e\n\x06status\x18\x04 \x01(\t\x12\x0e\n\x06output\x18\x05 \x01(\t\x12\x14\n\x0c\x61rtifact_url\x18\x06 \x01(\t\x12\x11\n\timage_url\x18\x07 \x01(\t\x12\x17\n\x0f\x65xecutable_path\x18\x08 \x01(\t\x12\x0f\n\x07profile\x18\t \x01(\t\x12\x0c\n\x04port\x18\n \x01(\t\x12\x15\n\rresource_spec\x18\x0b \x01(\t\x12\x15\n\rhelm_revision\x18\x0c \x01(\x05\x12.\n\ncreated_at\x18\r \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12.\n\nupdated_at\x18\x0e \x01(\x0b\x32\x1a.google.protobuf.Timestamp\"i\n\x13\x41ppServeAppCombined\x12*\n\rapp_serve_app\x18\x01 \x01(\x0b\x32\x13.tks_pb.AppServeApp\x12&\n\x05tasks\x18\x02 \x03(\x0b\x32\x17.tks_pb.AppServeAppTask\"\x17\n\tIDRequest\x12\n\n\x02id\x18\x01 \x01(\t\"J\n\x0eSimpleResponse\x12\x1a\n\x04\x63ode\x18\x01 \x01(\x0e\x32\x0c.tks_pb.Code\x12\x1c\n\x05\x65rror\x18\x02 \x01(\x0b\x32\r.tks_pb.Error\"T\n\x0bIDsResponse\x12\x1a\n\x04\x63ode\x18\x01 \x01(\x0e\x32\x0c.tks_pb.Code\x12\x1c\n\x05\x65rror\x18\x02 \x01(\x0b\x32\r.tks_pb.Error\x12\x0b\n\x03ids\x18\x03 \x03(\t\"R\n\nIDResponse\x12\x1a\n\x04\x63ode\x18\x01 \x01(\x0e\x32\x0c.tks_pb.Code\x12\x1c\n\x05\x65rror\x18\x02 \x01(\x0b\x32\r.tks_pb.Error\x12\n\n\x02id\x18\x03 \x01(\t\"i\n\x0cKeycloakInfo\x12\x12\n\ncluster_id\x18\x01 \x01(\t\x12\r\n\x05realm\x18\x02 \x01(\t\x12\x11\n\tclient_id\x18\x03 \x01(\t\x12\x0e\n\x06secret\x18\x04 \x01(\t\x12\x13\n\x0bprivate_key\x18\x05 \x01(\t*\xb4\x02\n\x04\x43ode\x12\x12\n\x0eOK_UNSPECIFIED\x10\x00\x12\r\n\tCANCELLED\x10\x01\x12\x0b\n\x07UNKNOWN\x10\x02\x12\x14\n\x10INVALID_ARGUMENT\x10\x03\x12\x15\n\x11\x44\x45\x41\x44LINE_EXCEEDED\x10\x04\x12\r\n\tNOT_FOUND\x10\x05\x12\x12\n\x0e\x41LREADY_EXISTS\x10\x06\x12\x15\n\x11PERMISSION_DENIED\x10\x07\x12\x13\n\x0fUNAUTHENTICATED\x10\x10\x12\x16\n\x12RESOURCE_EXHAUSTED\x10\x08\x12\x17\n\x13\x46\x41ILED_PRECONDITION\x10\t\x12\x0b\n\x07\x41\x42ORTED\x10\n\x12\x10\n\x0cOUT_OF_RANGE\x10\x0b\x12\x11\n\rUNIMPLEMENTED\x10\x0c\x12\x0c\n\x08INTERNAL\x10\r\x12\x0f\n\x0bUNAVAILABLE\x10\x0e*\xa0\x01\n\x0e\x41ppGroupStatus\x12\x19\n\x15\x41PP_GROUP_UNSPECIFIED\x10\x00\x12\x18\n\x14\x41PP_GROUP_INSTALLING\x10\x01\x12\x15\n\x11\x41PP_GROUP_RUNNING\x10\x02\x12\x16\n\x12\x41PP_GROUP_DELETING\x10\x04\x12\x15\n\x11\x41PP_GROUP_DELETED\x10\x05\x12\x13\n\x0f\x41PP_GROUP_ERROR\x10\x06*c\n\rClusterStatus\x12\x0f\n\x0bUNSPECIFIED\x10\x00\x12\x0e\n\nINSTALLING\x10\x01\x12\x0b\n\x07RUNNING\x10\x02\x12\x0c\n\x08\x44\x45LETING\x10\x04\x12\x0b\n\x07\x44\x45LETED\x10\x05\x12\t\n\x05\x45RROR\x10\x06*\xb5\x01\n\x07\x41ppType\x12\x12\n\x0e\x45P_UNSPECIFIED\x10\x00\x12\n\n\x06THANOS\x10\x01\x12\x0e\n\nPROMETHEUS\x10\x02\x12\x0b\n\x07GRAFANA\x10\x03\x12\t\n\x05KIALI\x10\x04\x12\n\n\x06KIBANA\x10\x05\x12\x10\n\x0c\x45LASTICSERCH\x10\x06\x12\x11\n\rCLOUD_CONSOLE\x10\x07\x12\x0b\n\x07HORIZON\x10\x08\x12\n\n\x06JAEGER\x10\t\x12\x18\n\x14KUBERNETES_DASHBOARD\x10\n*P\n\x0c\x41ppGroupType\x12\x18\n\x14\x41PP_TYPE_UNSPECIFIED\x10\x00\x12\x07\n\x03LMA\x10\x01\x12\x10\n\x0cSERVICE_MESH\x10\x02\x12\x0b\n\x07LMA_EFK\x10\x03*?\n\x07\x43spType\x12\x17\n\x13\x43SPTYPE_UNSPECIFIED\x10\x00\x12\x07\n\x03\x41WS\x10\x01\x12\x07\n\x03GCP\x10\x02\x12\t\n\x05\x41ZURE\x10\x03\x42*Z(github.com/openinfradev/tks-proto/tks_pbb\x06proto3')

_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, globals())
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'common_pb2', globals())
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'Z(github.com/openinfradev/tks-proto/tks_pb'
  _CODE._serialized_start=2390
  _CODE._serialized_end=2698
  _APPGROUPSTATUS._serialized_start=2701
  _APPGROUPSTATUS._serialized_end=2861
  _CLUSTERSTATUS._serialized_start=2863
  _CLUSTERSTATUS._serialized_end=2962
  _APPTYPE._serialized_start=2965
  _APPTYPE._serialized_end=3146
  _APPGROUPTYPE._serialized_start=3148
  _APPGROUPTYPE._serialized_end=3228
  _CSPTYPE._serialized_start=3230
  _CSPTYPE._serialized_end=3293
  _ERROR._serialized_start=57
  _ERROR._serialized_end=77
  _APPGROUP._serialized_start=80
  _APPGROUP._serialized_end=432
  _APPLICATION._serialized_start=435
  _APPLICATION._serialized_end=649
  _CLUSTER._serialized_start=652
  _CLUSTER._serialized_end=1032
  _CLUSTERRAWCONF._serialized_start=1034
  _CLUSTERRAWCONF._serialized_end=1155
  _CLUSTERCONF._serialized_start=1158
  _CLUSTERCONF._serialized_end=1300
  _APPSERVEAPP._serialized_start=1303
  _APPSERVEAPP._serialized_end=1556
  _APPSERVEAPPTASK._serialized_start=1559
  _APPSERVEAPPTASK._serialized_end=1902
  _APPSERVEAPPCOMBINED._serialized_start=1904
  _APPSERVEAPPCOMBINED._serialized_end=2009
  _IDREQUEST._serialized_start=2011
  _IDREQUEST._serialized_end=2034
  _SIMPLERESPONSE._serialized_start=2036
  _SIMPLERESPONSE._serialized_end=2110
  _IDSRESPONSE._serialized_start=2112
  _IDSRESPONSE._serialized_end=2196
  _IDRESPONSE._serialized_start=2198
  _IDRESPONSE._serialized_end=2280
  _KEYCLOAKINFO._serialized_start=2282
  _KEYCLOAKINFO._serialized_end=2387
# @@protoc_insertion_point(module_scope)
