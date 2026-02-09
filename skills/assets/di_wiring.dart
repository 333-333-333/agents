// app/di/providers.dart
// Clean Architecture wiring: repository + datasource providers
import 'package:flutter_riverpod/flutter_riverpod.dart';

final authRepositoryProvider = Provider<AuthRepository>((ref) {
  final datasource = ref.read(firebaseAuthDatasourceProvider);
  return AuthRepositoryImpl(datasource);
});

final firebaseAuthDatasourceProvider = Provider<FirebaseAuthDatasource>((ref) {
  return FirebaseAuthDatasource(FirebaseAuth.instance);
});

// features/auth/presentation/providers/auth_providers.dart
// Use case provider wiring
final loginUseCaseProvider = Provider<LoginUseCase>((ref) {
  final repo = ref.read(authRepositoryProvider);
  return LoginUseCase(repo);
});

final loginProvider = AsyncNotifierProvider<LoginNotifier, void>(
  LoginNotifier.new,
);
