/// Service that authenticates users against the backend API.
///
/// Uses JWT tokens for session management. Tokens are refreshed
/// automatically when they expire within 5 minutes.
///
/// See also:
/// - [AuthRepository] for the domain interface
/// - [TokenStorage] for persistence
class AuthService {
  /// Creates an [AuthService] with the given [httpClient] and [storage].
  AuthService(this.httpClient, this.storage);
}
