import 'package:args/args.dart' as args;
import 'package:monitor/src/models.dart';
import 'package:monitor/src/cli_options.dart';

final class MonitorCLI {
  MonitorCLI._({
    required args.ArgParser parser,
    required args.ArgResults argResults,
  })  : _parser = parser,
        _argResults = argResults,
        _options = Options.fromArgResults(argResults);

  /// Main instance of [ArgParser]
  final args.ArgParser _parser;

  /// Results from the provided [ArgParser]
  final args.ArgResults _argResults;

  final Options _options;

  Options get options => _options;

  factory MonitorCLI(
    List<String> arguments, [
    args.ArgParser? argParser,
  ]) {
    final parser = argParser ?? args.ArgParser();

    CLIOptions.bind(parser);
    final results = parser.parse(arguments);

    return MonitorCLI._(
      parser: parser,
      argResults: results,
    );
  }
}
