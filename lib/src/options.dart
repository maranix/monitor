import 'package:args/args.dart' as args;

enum CLIOptionsEnum {
  path('path', 'p'),
  exec('exec', 'e');

  const CLIOptionsEnum(this.name, this.abbr);

  final String name;
  final String abbr;
}

abstract final class CLIOptions {
  static void bind(args.ArgParser parser) {
    parser
      ..addOption(
        CLIOptionsEnum.path.name,
        abbr: CLIOptionsEnum.path.abbr,
        help: "path to the file or directory to monitor",
      )
      ..addOption(
        CLIOptionsEnum.exec.name,
        abbr: CLIOptionsEnum.exec.abbr,
        help: "command to execute when the monitored path is modified",
        valueHelp: "echo \"changes detected\"",
      );
  }
}
