import 'dart:io' as io;

import 'package:args/args.dart' as args;
import 'package:path/path.dart' as p;

final class ParsedOptions {
  const ParsedOptions({
    required this.path,
    required this.command,
  });

  final String path;
  final Command command;

  Future<bool> validatePath() async {
    if (await io.FileSystemEntity.isFile(path) ||
        await io.FileSystemEntity.isDirectory(path)) {
      return true;
    }

    return false;
  }

  // TODO: handle positioned paramters
  //
  // Running `monitor "echo something"` should watch changes in the current dir
  // and reload `echo something` command on each modifications.
  //
  // For now we support flag based paramters:
  //
  // monitor -p "some/path" -e "echo something"
  //
  // or
  //
  // monitor -e "echo something" for pwd.
  factory ParsedOptions.fromArgResult(args.ArgResults results) {
    final command = Command.fromArgResult(results);
    final path = p.normalize(p.absolute(results["path"]));

    return ParsedOptions(
      path: path,
      command: command,
    );
  }

  @override
  int get hashCode => Object.hash(path, command);

  @override
  bool operator ==(Object other) {
    if (identical(this, other)) return true;

    return other is ParsedOptions &&
        path == other.path &&
        command == other.command;
  }
}

final class Command {
  const Command({
    required this.executable,
    required this.params,
  });

  final String executable;
  final List<String> params;

  factory Command.fromArgResult(args.ArgResults results) {
    final parsedCommand = (results['exec'] as String).split(' ');

    return Command(
      executable: parsedCommand.first,
      params: parsedCommand.sublist(1, parsedCommand.length),
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
}
