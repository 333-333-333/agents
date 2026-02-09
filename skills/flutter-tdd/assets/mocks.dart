// test/helpers/mocks.dart
import 'package:bastet/features/auth/domain/repositories/auth_repository.dart';
import 'package:bastet/features/auth/domain/entities/user.dart';

/// Mock implementation of [AuthRepository] for unit testing.
///
/// Set function fields to control behavior per test.
class MockAuthRepository implements AuthRepository {
  Future<User> Function(String email, String password)? loginFunc;
  Future<void> Function()? logoutFunc;
  Future<User?> Function()? currentUserFunc;

  @override
  Future<User> login(String email, String password) {
    if (loginFunc != null) return loginFunc!(email, password);
    throw UnimplementedError('loginFunc not set');
  }

  @override
  Future<void> logout() {
    if (logoutFunc != null) return logoutFunc!();
    throw UnimplementedError('logoutFunc not set');
  }

  @override
  Future<User?> currentUser() {
    if (currentUserFunc != null) return currentUserFunc!();
    return Future.value(null);
  }
}
