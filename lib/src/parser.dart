import 'package:args/args.dart' as args;
import 'package:monitor/src/models.dart';
import 'package:monitor/src/options.dart';

final class CLIParser {
  CLIParser._({
    required args.ArgParser parser,
    required args.ArgResults argResults,
  })  : _parser = parser,
        _argResults = argResults,
        _options = ParsedOptions.fromArgResults(argResults);

  /// Main instance of [ArgParser]
  final args.ArgParser _parser;

  /// Results from the provided [ArgParser]
  final args.ArgResults _argResults;

  final ParsedOptions _options;

  ParsedOptions get options => _options;

  factory CLIParser(
    List<String> arguments, [
    args.ArgParser? argParser,
  ]) {
    final parser = argParser ?? args.ArgParser();

    CLIOptions.bind(parser);
    final results = parser.parse(arguments);

    return CLIParser._(
      parser: parser,
      argResults: results,
    );
  }
}
