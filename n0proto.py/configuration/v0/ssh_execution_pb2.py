# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: configuration/v0/ssh_execution.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.protobuf import empty_pb2 as google_dot_protobuf_dot_empty__pb2
from google.protobuf import timestamp_pb2 as google_dot_protobuf_dot_timestamp__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='configuration/v0/ssh_execution.proto',
  package='n0stack.configuration',
  syntax='proto3',
  serialized_options=_b('ZEgithub.com/n0stack/n0stack/n0proto.go/configuration/v0;pconfiguration'),
  serialized_pb=_b('\n$configuration/v0/ssh_execution.proto\x12\x15n0stack.configuration\x1a\x1bgoogle/protobuf/empty.proto\x1a\x1fgoogle/protobuf/timestamp.proto\"\xbc\x01\n\x12SSHExecutionResult\x12\x1c\n\x14virtual_machine_name\x18\x01 \x01(\t\x12\x17\n\x0fkey_secret_name\x18\x02 \x01(\t\x12.\n\nstarted_at\x18\x03 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12/\n\x0b\x66inished_at\x18\x04 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12\x0e\n\x06output\x18\x05 \x01(\t\"\xbe\x01\n\x0cSSHExecution\x12\x0c\n\x04name\x18\x01 \x01(\t\x12I\n\x0b\x61nnotations\x18\x03 \x03(\x0b\x32\x34.n0stack.configuration.SSHExecution.AnnotationsEntry\x12\x10\n\x08\x63ommands\x18\n \x03(\t\x12\x0f\n\x07results\x18\x32 \x03(\t\x1a\x32\n\x10\x41nnotationsEntry\x12\x0b\n\x03key\x18\x01 \x01(\t\x12\r\n\x05value\x18\x02 \x01(\t:\x02\x38\x01\"\x1a\n\x18ListSSHExecutionsRequest\"N\n\x19ListSSHExecutionsResponse\x12\x31\n\x04sshs\x18\x01 \x03(\x0b\x32#.n0stack.configuration.SSHExecution\"&\n\x16GetSSHExecutionRequest\x12\x0c\n\x04name\x18\x01 \x01(\t\"\xc5\x01\n\x18\x41pplySSHExecutionRequest\x12\x0c\n\x04name\x18\x01 \x01(\t\x12U\n\x0b\x61nnotations\x18\x03 \x03(\x0b\x32@.n0stack.configuration.ApplySSHExecutionRequest.AnnotationsEntry\x12\x10\n\x08\x63ommands\x18\n \x03(\t\x1a\x32\n\x10\x41nnotationsEntry\x12\x0b\n\x03key\x18\x01 \x01(\t\x12\r\n\x05value\x18\x02 \x01(\t:\x02\x38\x01\")\n\x19\x44\x65leteSSHExecutionRequest\x12\x0c\n\x04name\x18\x01 \x01(\t\"\x93\x04\n\x16RunSSHExecutionRequest\x12^\n\x06\x66ilter\x18\x01 \x01(\x0b\x32N.n0stack.configuration.RunSSHExecutionRequest.SSHExecutionVirtualMachineFilter\x12\x10\n\x08username\x18\x02 \x01(\t\x12\x17\n\x0fkey_secret_name\x18\x03 \x01(\t\x12\x66\n\x15\x65nvironment_variables\x18\x04 \x03(\x0b\x32G.n0stack.configuration.RunSSHExecutionRequest.EnvironmentVariablesEntry\x1a\xc8\x01\n SSHExecutionVirtualMachineFilter\x12\x0c\n\x04name\x18\x01 \x01(\t\x12h\n\x05label\x18\x02 \x03(\x0b\x32Y.n0stack.configuration.RunSSHExecutionRequest.SSHExecutionVirtualMachineFilter.LabelEntry\x1a,\n\nLabelEntry\x12\x0b\n\x03key\x18\x01 \x01(\t\x12\r\n\x05value\x18\x02 \x01(\t:\x02\x38\x01\x1a;\n\x19\x45nvironmentVariablesEntry\x12\x0b\n\x03key\x18\x01 \x01(\t\x12\r\n\x05value\x18\x02 \x01(\t:\x02\x38\x01\x32\xb6\x04\n\x13SSHExecutionService\x12x\n\x11ListSSHExecutions\x12/.n0stack.configuration.ListSSHExecutionsRequest\x1a\x30.n0stack.configuration.ListSSHExecutionsResponse\"\x00\x12g\n\x0fGetSSHExecution\x12-.n0stack.configuration.GetSSHExecutionRequest\x1a#.n0stack.configuration.SSHExecution\"\x00\x12k\n\x11\x41pplySSHExecution\x12/.n0stack.configuration.ApplySSHExecutionRequest\x1a#.n0stack.configuration.SSHExecution\"\x00\x12`\n\x12\x44\x65leteSSHExecution\x12\x30.n0stack.configuration.DeleteSSHExecutionRequest\x1a\x16.google.protobuf.Empty\"\x00\x12m\n\x0fRunSSHExecution\x12-.n0stack.configuration.RunSSHExecutionRequest\x1a).n0stack.configuration.SSHExecutionResult\"\x00\x42GZEgithub.com/n0stack/n0stack/n0proto.go/configuration/v0;pconfigurationb\x06proto3')
  ,
  dependencies=[google_dot_protobuf_dot_empty__pb2.DESCRIPTOR,google_dot_protobuf_dot_timestamp__pb2.DESCRIPTOR,])




_SSHEXECUTIONRESULT = _descriptor.Descriptor(
  name='SSHExecutionResult',
  full_name='n0stack.configuration.SSHExecutionResult',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='virtual_machine_name', full_name='n0stack.configuration.SSHExecutionResult.virtual_machine_name', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='key_secret_name', full_name='n0stack.configuration.SSHExecutionResult.key_secret_name', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='started_at', full_name='n0stack.configuration.SSHExecutionResult.started_at', index=2,
      number=3, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='finished_at', full_name='n0stack.configuration.SSHExecutionResult.finished_at', index=3,
      number=4, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='output', full_name='n0stack.configuration.SSHExecutionResult.output', index=4,
      number=5, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
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
  serialized_start=126,
  serialized_end=314,
)


_SSHEXECUTION_ANNOTATIONSENTRY = _descriptor.Descriptor(
  name='AnnotationsEntry',
  full_name='n0stack.configuration.SSHExecution.AnnotationsEntry',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='key', full_name='n0stack.configuration.SSHExecution.AnnotationsEntry.key', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='value', full_name='n0stack.configuration.SSHExecution.AnnotationsEntry.value', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=_b('8\001'),
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=457,
  serialized_end=507,
)

_SSHEXECUTION = _descriptor.Descriptor(
  name='SSHExecution',
  full_name='n0stack.configuration.SSHExecution',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='name', full_name='n0stack.configuration.SSHExecution.name', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='annotations', full_name='n0stack.configuration.SSHExecution.annotations', index=1,
      number=3, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='commands', full_name='n0stack.configuration.SSHExecution.commands', index=2,
      number=10, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='results', full_name='n0stack.configuration.SSHExecution.results', index=3,
      number=50, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[_SSHEXECUTION_ANNOTATIONSENTRY, ],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=317,
  serialized_end=507,
)


_LISTSSHEXECUTIONSREQUEST = _descriptor.Descriptor(
  name='ListSSHExecutionsRequest',
  full_name='n0stack.configuration.ListSSHExecutionsRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
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
  serialized_start=509,
  serialized_end=535,
)


_LISTSSHEXECUTIONSRESPONSE = _descriptor.Descriptor(
  name='ListSSHExecutionsResponse',
  full_name='n0stack.configuration.ListSSHExecutionsResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='sshs', full_name='n0stack.configuration.ListSSHExecutionsResponse.sshs', index=0,
      number=1, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
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
  serialized_start=537,
  serialized_end=615,
)


_GETSSHEXECUTIONREQUEST = _descriptor.Descriptor(
  name='GetSSHExecutionRequest',
  full_name='n0stack.configuration.GetSSHExecutionRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='name', full_name='n0stack.configuration.GetSSHExecutionRequest.name', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
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
  serialized_start=617,
  serialized_end=655,
)


_APPLYSSHEXECUTIONREQUEST_ANNOTATIONSENTRY = _descriptor.Descriptor(
  name='AnnotationsEntry',
  full_name='n0stack.configuration.ApplySSHExecutionRequest.AnnotationsEntry',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='key', full_name='n0stack.configuration.ApplySSHExecutionRequest.AnnotationsEntry.key', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='value', full_name='n0stack.configuration.ApplySSHExecutionRequest.AnnotationsEntry.value', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=_b('8\001'),
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=457,
  serialized_end=507,
)

_APPLYSSHEXECUTIONREQUEST = _descriptor.Descriptor(
  name='ApplySSHExecutionRequest',
  full_name='n0stack.configuration.ApplySSHExecutionRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='name', full_name='n0stack.configuration.ApplySSHExecutionRequest.name', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='annotations', full_name='n0stack.configuration.ApplySSHExecutionRequest.annotations', index=1,
      number=3, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='commands', full_name='n0stack.configuration.ApplySSHExecutionRequest.commands', index=2,
      number=10, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[_APPLYSSHEXECUTIONREQUEST_ANNOTATIONSENTRY, ],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=658,
  serialized_end=855,
)


_DELETESSHEXECUTIONREQUEST = _descriptor.Descriptor(
  name='DeleteSSHExecutionRequest',
  full_name='n0stack.configuration.DeleteSSHExecutionRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='name', full_name='n0stack.configuration.DeleteSSHExecutionRequest.name', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
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
  serialized_start=857,
  serialized_end=898,
)


_RUNSSHEXECUTIONREQUEST_SSHEXECUTIONVIRTUALMACHINEFILTER_LABELENTRY = _descriptor.Descriptor(
  name='LabelEntry',
  full_name='n0stack.configuration.RunSSHExecutionRequest.SSHExecutionVirtualMachineFilter.LabelEntry',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='key', full_name='n0stack.configuration.RunSSHExecutionRequest.SSHExecutionVirtualMachineFilter.LabelEntry.key', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='value', full_name='n0stack.configuration.RunSSHExecutionRequest.SSHExecutionVirtualMachineFilter.LabelEntry.value', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=_b('8\001'),
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=1327,
  serialized_end=1371,
)

_RUNSSHEXECUTIONREQUEST_SSHEXECUTIONVIRTUALMACHINEFILTER = _descriptor.Descriptor(
  name='SSHExecutionVirtualMachineFilter',
  full_name='n0stack.configuration.RunSSHExecutionRequest.SSHExecutionVirtualMachineFilter',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='name', full_name='n0stack.configuration.RunSSHExecutionRequest.SSHExecutionVirtualMachineFilter.name', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='label', full_name='n0stack.configuration.RunSSHExecutionRequest.SSHExecutionVirtualMachineFilter.label', index=1,
      number=2, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[_RUNSSHEXECUTIONREQUEST_SSHEXECUTIONVIRTUALMACHINEFILTER_LABELENTRY, ],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=1171,
  serialized_end=1371,
)

_RUNSSHEXECUTIONREQUEST_ENVIRONMENTVARIABLESENTRY = _descriptor.Descriptor(
  name='EnvironmentVariablesEntry',
  full_name='n0stack.configuration.RunSSHExecutionRequest.EnvironmentVariablesEntry',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='key', full_name='n0stack.configuration.RunSSHExecutionRequest.EnvironmentVariablesEntry.key', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='value', full_name='n0stack.configuration.RunSSHExecutionRequest.EnvironmentVariablesEntry.value', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=_b('8\001'),
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=1373,
  serialized_end=1432,
)

_RUNSSHEXECUTIONREQUEST = _descriptor.Descriptor(
  name='RunSSHExecutionRequest',
  full_name='n0stack.configuration.RunSSHExecutionRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='filter', full_name='n0stack.configuration.RunSSHExecutionRequest.filter', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='username', full_name='n0stack.configuration.RunSSHExecutionRequest.username', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='key_secret_name', full_name='n0stack.configuration.RunSSHExecutionRequest.key_secret_name', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='environment_variables', full_name='n0stack.configuration.RunSSHExecutionRequest.environment_variables', index=3,
      number=4, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[_RUNSSHEXECUTIONREQUEST_SSHEXECUTIONVIRTUALMACHINEFILTER, _RUNSSHEXECUTIONREQUEST_ENVIRONMENTVARIABLESENTRY, ],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=901,
  serialized_end=1432,
)

_SSHEXECUTIONRESULT.fields_by_name['started_at'].message_type = google_dot_protobuf_dot_timestamp__pb2._TIMESTAMP
_SSHEXECUTIONRESULT.fields_by_name['finished_at'].message_type = google_dot_protobuf_dot_timestamp__pb2._TIMESTAMP
_SSHEXECUTION_ANNOTATIONSENTRY.containing_type = _SSHEXECUTION
_SSHEXECUTION.fields_by_name['annotations'].message_type = _SSHEXECUTION_ANNOTATIONSENTRY
_LISTSSHEXECUTIONSRESPONSE.fields_by_name['sshs'].message_type = _SSHEXECUTION
_APPLYSSHEXECUTIONREQUEST_ANNOTATIONSENTRY.containing_type = _APPLYSSHEXECUTIONREQUEST
_APPLYSSHEXECUTIONREQUEST.fields_by_name['annotations'].message_type = _APPLYSSHEXECUTIONREQUEST_ANNOTATIONSENTRY
_RUNSSHEXECUTIONREQUEST_SSHEXECUTIONVIRTUALMACHINEFILTER_LABELENTRY.containing_type = _RUNSSHEXECUTIONREQUEST_SSHEXECUTIONVIRTUALMACHINEFILTER
_RUNSSHEXECUTIONREQUEST_SSHEXECUTIONVIRTUALMACHINEFILTER.fields_by_name['label'].message_type = _RUNSSHEXECUTIONREQUEST_SSHEXECUTIONVIRTUALMACHINEFILTER_LABELENTRY
_RUNSSHEXECUTIONREQUEST_SSHEXECUTIONVIRTUALMACHINEFILTER.containing_type = _RUNSSHEXECUTIONREQUEST
_RUNSSHEXECUTIONREQUEST_ENVIRONMENTVARIABLESENTRY.containing_type = _RUNSSHEXECUTIONREQUEST
_RUNSSHEXECUTIONREQUEST.fields_by_name['filter'].message_type = _RUNSSHEXECUTIONREQUEST_SSHEXECUTIONVIRTUALMACHINEFILTER
_RUNSSHEXECUTIONREQUEST.fields_by_name['environment_variables'].message_type = _RUNSSHEXECUTIONREQUEST_ENVIRONMENTVARIABLESENTRY
DESCRIPTOR.message_types_by_name['SSHExecutionResult'] = _SSHEXECUTIONRESULT
DESCRIPTOR.message_types_by_name['SSHExecution'] = _SSHEXECUTION
DESCRIPTOR.message_types_by_name['ListSSHExecutionsRequest'] = _LISTSSHEXECUTIONSREQUEST
DESCRIPTOR.message_types_by_name['ListSSHExecutionsResponse'] = _LISTSSHEXECUTIONSRESPONSE
DESCRIPTOR.message_types_by_name['GetSSHExecutionRequest'] = _GETSSHEXECUTIONREQUEST
DESCRIPTOR.message_types_by_name['ApplySSHExecutionRequest'] = _APPLYSSHEXECUTIONREQUEST
DESCRIPTOR.message_types_by_name['DeleteSSHExecutionRequest'] = _DELETESSHEXECUTIONREQUEST
DESCRIPTOR.message_types_by_name['RunSSHExecutionRequest'] = _RUNSSHEXECUTIONREQUEST
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

SSHExecutionResult = _reflection.GeneratedProtocolMessageType('SSHExecutionResult', (_message.Message,), dict(
  DESCRIPTOR = _SSHEXECUTIONRESULT,
  __module__ = 'configuration.v0.ssh_execution_pb2'
  # @@protoc_insertion_point(class_scope:n0stack.configuration.SSHExecutionResult)
  ))
_sym_db.RegisterMessage(SSHExecutionResult)

SSHExecution = _reflection.GeneratedProtocolMessageType('SSHExecution', (_message.Message,), dict(

  AnnotationsEntry = _reflection.GeneratedProtocolMessageType('AnnotationsEntry', (_message.Message,), dict(
    DESCRIPTOR = _SSHEXECUTION_ANNOTATIONSENTRY,
    __module__ = 'configuration.v0.ssh_execution_pb2'
    # @@protoc_insertion_point(class_scope:n0stack.configuration.SSHExecution.AnnotationsEntry)
    ))
  ,
  DESCRIPTOR = _SSHEXECUTION,
  __module__ = 'configuration.v0.ssh_execution_pb2'
  # @@protoc_insertion_point(class_scope:n0stack.configuration.SSHExecution)
  ))
_sym_db.RegisterMessage(SSHExecution)
_sym_db.RegisterMessage(SSHExecution.AnnotationsEntry)

ListSSHExecutionsRequest = _reflection.GeneratedProtocolMessageType('ListSSHExecutionsRequest', (_message.Message,), dict(
  DESCRIPTOR = _LISTSSHEXECUTIONSREQUEST,
  __module__ = 'configuration.v0.ssh_execution_pb2'
  # @@protoc_insertion_point(class_scope:n0stack.configuration.ListSSHExecutionsRequest)
  ))
_sym_db.RegisterMessage(ListSSHExecutionsRequest)

ListSSHExecutionsResponse = _reflection.GeneratedProtocolMessageType('ListSSHExecutionsResponse', (_message.Message,), dict(
  DESCRIPTOR = _LISTSSHEXECUTIONSRESPONSE,
  __module__ = 'configuration.v0.ssh_execution_pb2'
  # @@protoc_insertion_point(class_scope:n0stack.configuration.ListSSHExecutionsResponse)
  ))
_sym_db.RegisterMessage(ListSSHExecutionsResponse)

GetSSHExecutionRequest = _reflection.GeneratedProtocolMessageType('GetSSHExecutionRequest', (_message.Message,), dict(
  DESCRIPTOR = _GETSSHEXECUTIONREQUEST,
  __module__ = 'configuration.v0.ssh_execution_pb2'
  # @@protoc_insertion_point(class_scope:n0stack.configuration.GetSSHExecutionRequest)
  ))
_sym_db.RegisterMessage(GetSSHExecutionRequest)

ApplySSHExecutionRequest = _reflection.GeneratedProtocolMessageType('ApplySSHExecutionRequest', (_message.Message,), dict(

  AnnotationsEntry = _reflection.GeneratedProtocolMessageType('AnnotationsEntry', (_message.Message,), dict(
    DESCRIPTOR = _APPLYSSHEXECUTIONREQUEST_ANNOTATIONSENTRY,
    __module__ = 'configuration.v0.ssh_execution_pb2'
    # @@protoc_insertion_point(class_scope:n0stack.configuration.ApplySSHExecutionRequest.AnnotationsEntry)
    ))
  ,
  DESCRIPTOR = _APPLYSSHEXECUTIONREQUEST,
  __module__ = 'configuration.v0.ssh_execution_pb2'
  # @@protoc_insertion_point(class_scope:n0stack.configuration.ApplySSHExecutionRequest)
  ))
_sym_db.RegisterMessage(ApplySSHExecutionRequest)
_sym_db.RegisterMessage(ApplySSHExecutionRequest.AnnotationsEntry)

DeleteSSHExecutionRequest = _reflection.GeneratedProtocolMessageType('DeleteSSHExecutionRequest', (_message.Message,), dict(
  DESCRIPTOR = _DELETESSHEXECUTIONREQUEST,
  __module__ = 'configuration.v0.ssh_execution_pb2'
  # @@protoc_insertion_point(class_scope:n0stack.configuration.DeleteSSHExecutionRequest)
  ))
_sym_db.RegisterMessage(DeleteSSHExecutionRequest)

RunSSHExecutionRequest = _reflection.GeneratedProtocolMessageType('RunSSHExecutionRequest', (_message.Message,), dict(

  SSHExecutionVirtualMachineFilter = _reflection.GeneratedProtocolMessageType('SSHExecutionVirtualMachineFilter', (_message.Message,), dict(

    LabelEntry = _reflection.GeneratedProtocolMessageType('LabelEntry', (_message.Message,), dict(
      DESCRIPTOR = _RUNSSHEXECUTIONREQUEST_SSHEXECUTIONVIRTUALMACHINEFILTER_LABELENTRY,
      __module__ = 'configuration.v0.ssh_execution_pb2'
      # @@protoc_insertion_point(class_scope:n0stack.configuration.RunSSHExecutionRequest.SSHExecutionVirtualMachineFilter.LabelEntry)
      ))
    ,
    DESCRIPTOR = _RUNSSHEXECUTIONREQUEST_SSHEXECUTIONVIRTUALMACHINEFILTER,
    __module__ = 'configuration.v0.ssh_execution_pb2'
    # @@protoc_insertion_point(class_scope:n0stack.configuration.RunSSHExecutionRequest.SSHExecutionVirtualMachineFilter)
    ))
  ,

  EnvironmentVariablesEntry = _reflection.GeneratedProtocolMessageType('EnvironmentVariablesEntry', (_message.Message,), dict(
    DESCRIPTOR = _RUNSSHEXECUTIONREQUEST_ENVIRONMENTVARIABLESENTRY,
    __module__ = 'configuration.v0.ssh_execution_pb2'
    # @@protoc_insertion_point(class_scope:n0stack.configuration.RunSSHExecutionRequest.EnvironmentVariablesEntry)
    ))
  ,
  DESCRIPTOR = _RUNSSHEXECUTIONREQUEST,
  __module__ = 'configuration.v0.ssh_execution_pb2'
  # @@protoc_insertion_point(class_scope:n0stack.configuration.RunSSHExecutionRequest)
  ))
_sym_db.RegisterMessage(RunSSHExecutionRequest)
_sym_db.RegisterMessage(RunSSHExecutionRequest.SSHExecutionVirtualMachineFilter)
_sym_db.RegisterMessage(RunSSHExecutionRequest.SSHExecutionVirtualMachineFilter.LabelEntry)
_sym_db.RegisterMessage(RunSSHExecutionRequest.EnvironmentVariablesEntry)


DESCRIPTOR._options = None
_SSHEXECUTION_ANNOTATIONSENTRY._options = None
_APPLYSSHEXECUTIONREQUEST_ANNOTATIONSENTRY._options = None
_RUNSSHEXECUTIONREQUEST_SSHEXECUTIONVIRTUALMACHINEFILTER_LABELENTRY._options = None
_RUNSSHEXECUTIONREQUEST_ENVIRONMENTVARIABLESENTRY._options = None

_SSHEXECUTIONSERVICE = _descriptor.ServiceDescriptor(
  name='SSHExecutionService',
  full_name='n0stack.configuration.SSHExecutionService',
  file=DESCRIPTOR,
  index=0,
  serialized_options=None,
  serialized_start=1435,
  serialized_end=2001,
  methods=[
  _descriptor.MethodDescriptor(
    name='ListSSHExecutions',
    full_name='n0stack.configuration.SSHExecutionService.ListSSHExecutions',
    index=0,
    containing_service=None,
    input_type=_LISTSSHEXECUTIONSREQUEST,
    output_type=_LISTSSHEXECUTIONSRESPONSE,
    serialized_options=None,
  ),
  _descriptor.MethodDescriptor(
    name='GetSSHExecution',
    full_name='n0stack.configuration.SSHExecutionService.GetSSHExecution',
    index=1,
    containing_service=None,
    input_type=_GETSSHEXECUTIONREQUEST,
    output_type=_SSHEXECUTION,
    serialized_options=None,
  ),
  _descriptor.MethodDescriptor(
    name='ApplySSHExecution',
    full_name='n0stack.configuration.SSHExecutionService.ApplySSHExecution',
    index=2,
    containing_service=None,
    input_type=_APPLYSSHEXECUTIONREQUEST,
    output_type=_SSHEXECUTION,
    serialized_options=None,
  ),
  _descriptor.MethodDescriptor(
    name='DeleteSSHExecution',
    full_name='n0stack.configuration.SSHExecutionService.DeleteSSHExecution',
    index=3,
    containing_service=None,
    input_type=_DELETESSHEXECUTIONREQUEST,
    output_type=google_dot_protobuf_dot_empty__pb2._EMPTY,
    serialized_options=None,
  ),
  _descriptor.MethodDescriptor(
    name='RunSSHExecution',
    full_name='n0stack.configuration.SSHExecutionService.RunSSHExecution',
    index=4,
    containing_service=None,
    input_type=_RUNSSHEXECUTIONREQUEST,
    output_type=_SSHEXECUTIONRESULT,
    serialized_options=None,
  ),
])
_sym_db.RegisterServiceDescriptor(_SSHEXECUTIONSERVICE)

DESCRIPTOR.services_by_name['SSHExecutionService'] = _SSHEXECUTIONSERVICE

# @@protoc_insertion_point(module_scope)
