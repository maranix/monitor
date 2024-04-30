import 'dart:io' as io;

import 'package:args/args.dart' as args;
import 'package:monitor/src/exceptions.dart';
import 'package:monitor/src/cli_options.dart';
import 'package:path/path.dart' as p;

typedef ParseResult = ({String path, Command cmd});

enum ExitCodeEnum {
  error(1),
  invalidPath(2),
  processError(127);

  const ExitCodeEnum(this.val);
  final int val;
}

final class Options {
  const Options({
    required this.path,
    required this.command,
  });

  final String path;
  final Command command;

  static bool _isPathValid(path) {
    return io.FileSystemEntity.isFileSync(path) ||
        io.FileSystemEntity.isDirectorySync(path);
  }

  static String _parseAndValidatePath(args.ArgResults results) {
    String? namedParam = results[CLIOptionsEnum.path.name];
    if (namedParam != null && _isPathValid(namedParam)) {
      return namedParam;
    }

    String unnamedParam = results.rest.first;
    if (_isPathValid(unnamedParam)) {
      return unnamedParam;
    }

    throw DirectoryOrFileDoesNotExist(unnamedParam);
  }

  static List<String> _parseCommandStr(args.ArgResults results) {
    List<String>? namedParam = results[CLIOptionsEnum.exec.name];
    if (namedParam != null) {
      return namedParam;
    }

    return results.rest.sublist(1, results.rest.length);
  }

  static ParseResult _fromArgResults(args.ArgResults results) {
    try {
      final path = p.normalize(
        p.absolute(_parseAndValidatePath(results)),
      );
      final command = Command.fromString(_parseCommandStr(results));

      return (path: path, cmd: command);
    } on DirectoryOrFileDoesNotExist catch (e) {
      io.stderr.writeln(e);
      io.exitCode = ExitCodeEnum.invalidPath.val;
    } catch (e) {
      io.stderr.writeln(e);
      io.exitCode = ExitCodeEnum.error.val;
    } finally {
      io.exit(io.exitCode);
    }
  }

  factory Options.fromArgResults(args.ArgResults results) {
    final parsedResult = _fromArgResults(results);

    return Options(
      path: parsedResult.path,
      command: parsedResult.cmd,
    );
  }

  @override
  int get hashCode => Object.hash(path, command);

  @override
  bool operator ==(Object other) {
    if (identical(this, other)) return true;

    return other is Options && path == other.path && command == other.command;
  }

  @override
  String toString() {
    return "ParsedOptions(path: $path, command: $command)";
  }
}

final class Command {
  const Command({
    required this.executable,
    required this.params,
  });

  final String executable;
  final List<String> params;

  Future<io.Process> _run(String workingDirectory) async {
    try {
      final process = await io.Process.start(
        executable,
        params,
        workingDirectory: workingDirectory,
      ).onError(
        (error, stackTrace) {
          throw CouldNotSpawnProcess(
            executable,
            params,
            error: error,
            stackTrace: stackTrace,
          );
        },
      );

      return process;
    } on CouldNotSpawnProcess catch (e) {
      io.stderr.writeln(e);
    } catch (e) {
      io.stderr.writeln(e);
    } finally {
      io.exit(ExitCodeEnum.processError.val);
    }
  }

  Future<io.Process> exec(String workingDirectory) async {
    return await _run(workingDirectory);
  }

  factory Command.fromString(List<String> cmd) {
    return Command(
      executable: cmd.first,
      params: cmd.sublist(1, cmd.length),
    );
  }

  @override
  int get hashCode => Object.hash(executable, params);

  @override
  bool operator ==(Object other) {
    if (identical(this, other)) return true;

    return other is Command &&
        executable == other.executable &&
        params == other.params;
  }

  @override
  String toString() {
    return "Command(executable: $executable, params: $params)";
  }
}
