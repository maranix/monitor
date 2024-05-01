import 'dart:async';
import 'dart:io' as io;

import 'package:monitor/monitor.dart';
import 'package:watcher/watcher.dart' as w;

void main(List<String> arguments) async {
  final cli = MonitorCLI(arguments);
  // TODO: workDir and command exec dir should be separate.
  ///
  /// Separate out path for the entity to monitor and path where command will be
  /// executed. By default it'll be the same as the entity path.
  final workDir = cli.options.path;
  final watch = w.Watcher(workDir);

  io.Process? process;
  Timer? debounceTimer;

  watch.events.where(_modifyEvents).listen((changeEvent) async {
    // Todo: Add Filter capability for files and folders
    ///
    /// By default we are monitoring for changes for all of the child paths.
    /// This leads to a problem where any changes in the private or hidden
    /// entities leads to reloading the command as well. Which is something
    /// that should not happend for eg: .git, .build, in the case of some
    /// tools such as node it'll end monitoring the whole node_modules directory.
    if (changeEvent.path.split('/').any((e) => e.startsWith('.'))) return;

    debounceTimer?.cancel();

    debounceTimer = Timer(const Duration(milliseconds: 200), () async {
      final results = await Future.wait([
        _killProcess(process),
        _spawnProcess(cli.options.command, workDir),
      ]);

      process = results.whereType<io.Process>().firstOrNull;
    });
  });

  await _killProcess(process);
}

bool _modifyEvents(w.WatchEvent change) {
  return change.type == w.ChangeType.MODIFY;
}

Future<void> _killProcess(io.Process? process) async {
  if (process != null) {
    io.Process.killPid(process.pid, io.ProcessSignal.sigquit);
    await process.exitCode;
  }
}

Future<io.Process> _spawnProcess(Command cmd, String workDir) async {
  final process = await cmd.exec(workDir);

  process.stdout.listen(io.stdout.add);
  process.stderr.listen(io.stderr.add);

  return process;
}
