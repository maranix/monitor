import 'dart:io' as io;

import 'package:monitor/src/models.dart';

final class DirectoryOrFileDoesNotExist implements Exception {
  const DirectoryOrFileDoesNotExist(this.path);

  final String path;

  @override
  String toString() {
    return "$runtimeType :\n\n$path does not exist, Please make sure that the destination is correct.";
  }
}

final class CouldNotSpawnProcess extends io.ProcessException {
  CouldNotSpawnProcess(
    String executable,
    List<String> arguments, {
    required this.error,
    required this.stackTrace,
  }) : super(
          executable,
          arguments,
          "",
          ExitCodeEnum.processError.val,
        );

  final Object? error;
  final StackTrace stackTrace;

  @override
  String toString() {
    final message = super.toString();

    return "$message\n\nError: ${error.toString()}\n\nStackTrace:$stackTrace";
  }
}
