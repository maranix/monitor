final class DirectoryOrFileDoesNotExist implements Exception {
  const DirectoryOrFileDoesNotExist(this.path);

  final String path;

  @override
  String toString() {
    return "$runtimeType:\n\n$path does not exist, Please make sure that the destination is correct.";
  }
}
