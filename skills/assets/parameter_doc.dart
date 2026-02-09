/// Fetches a user by their [id] from the remote API.
///
/// If [forceRefresh] is `true`, bypasses the local cache.
/// Defaults to `false`.
Future<User> getUser(String id, {bool forceRefresh = false}) { ... }
