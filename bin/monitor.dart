import 'dart:io' as io;

import 'package:monitor/monitor.dart';
import 'package:watcher/watcher.dart' as w;

void main(List<String> arguments) async {
  final cli = MonitorCLI(arguments);
  final workDir = cli.options.path;
  final watch = w.Watcher(workDir);

  io.Process? process;
  // TODO: debounce changeEvents and process execution queue.
  //
  // Subsequent event changes should be debounced and only be completed if 
  // there were  no changes in the previous X mili/seconds. This ensures
  // that we're not constantly spawning and killing process.
  //
  // Bug 1: The process doesn't spawn initially;
  // Bug 2: Writing to a file spawns process normally and emits correct.
  //        However, writing again emits 4 events (add, modify, modify and remove)
  //        in no particular order. Hence 4 processes are spawned together
  //        simutaneously.
  watch.events.distinct().listen((changeEvent) async {
    if (process != null) {
      process!.kill(io.ProcessSignal.sigquit);
    }

    process = await cli.options.command.exec(workDir);
    print(process?.pid);
  }).onDone(() => _cleanup(process!));
}

void _cleanup(io.Process process) {
  process.kill(io.ProcessSignal.sigquit);
}
