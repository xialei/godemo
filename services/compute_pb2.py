# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: compute.proto

from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor.FileDescriptor(
  name='compute.proto',
  package='services',
  syntax='proto3',
  serialized_options=b'Z\n.;services',
  serialized_pb=b'\n\rcompute.proto\x12\x08services\"\x1a\n\x0bTaskRequest\x12\x0b\n\x03msg\x18\x01 \x01(\t\"\x1e\n\x0cTaskResponse\x12\x0e\n\x06result\x18\x01 \x01(\t2C\n\x07\x43ompute\x12\x38\n\x07RunTask\x12\x15.services.TaskRequest\x1a\x16.services.TaskResponseB\x0cZ\n.;servicesb\x06proto3'
)




_TASKREQUEST = _descriptor.Descriptor(
  name='TaskRequest',
  full_name='services.TaskRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='msg', full_name='services.TaskRequest.msg', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=27,
  serialized_end=53,
)


_TASKRESPONSE = _descriptor.Descriptor(
  name='TaskResponse',
  full_name='services.TaskResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='result', full_name='services.TaskResponse.result', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=55,
  serialized_end=85,
)

DESCRIPTOR.message_types_by_name['TaskRequest'] = _TASKREQUEST
DESCRIPTOR.message_types_by_name['TaskResponse'] = _TASKRESPONSE
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

TaskRequest = _reflection.GeneratedProtocolMessageType('TaskRequest', (_message.Message,), {
  'DESCRIPTOR' : _TASKREQUEST,
  '__module__' : 'compute_pb2'
  # @@protoc_insertion_point(class_scope:services.TaskRequest)
  })
_sym_db.RegisterMessage(TaskRequest)

TaskResponse = _reflection.GeneratedProtocolMessageType('TaskResponse', (_message.Message,), {
  'DESCRIPTOR' : _TASKRESPONSE,
  '__module__' : 'compute_pb2'
  # @@protoc_insertion_point(class_scope:services.TaskResponse)
  })
_sym_db.RegisterMessage(TaskResponse)


DESCRIPTOR._options = None

_COMPUTE = _descriptor.ServiceDescriptor(
  name='Compute',
  full_name='services.Compute',
  file=DESCRIPTOR,
  index=0,
  serialized_options=None,
  serialized_start=87,
  serialized_end=154,
  methods=[
  _descriptor.MethodDescriptor(
    name='RunTask',
    full_name='services.Compute.RunTask',
    index=0,
    containing_service=None,
    input_type=_TASKREQUEST,
    output_type=_TASKRESPONSE,
    serialized_options=None,
  ),
])
_sym_db.RegisterServiceDescriptor(_COMPUTE)

DESCRIPTOR.services_by_name['Compute'] = _COMPUTE

# @@protoc_insertion_point(module_scope)