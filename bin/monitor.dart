import 'package:monitor/monitor.dart' as monitor;

void main(List<String> arguments) async {
  final monCLI = monitor.CLIParser(arguments);

  final isPathValidated = await monCLI.options.validatePath();

  print(isPathValidated);
}
