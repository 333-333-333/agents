// test/features/auth/presentation/providers/auth_state_provider_test.dart
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:bastet/features/auth/presentation/providers/auth_state_provider.dart';
import 'package:bastet/features/auth/domain/repositories/auth_repository.dart';
import '../../helpers/mocks.dart';

void main() {
  late ProviderContainer container;
  late MockAuthRepository mockRepo;

  setUp(() {
    mockRepo = MockAuthRepository();
    container = ProviderContainer(
      overrides: [authRepositoryProvider.overrideWithValue(mockRepo)],
    );
  });

  tearDown(() {
    container.dispose();
  });

  group('AuthStateNotifier', () {
    test('initial state is unauthenticated', () {
      final state = container.read(authStateProvider);
      expect(state, isA<Unauthenticated>());
    });

    test('login transitions to authenticated on success', () async {
      final user = User(id: '1', email: 'a@b.com', name: 'Test');
      mockRepo.loginFunc = (_, __) async => user;

      final notifier = container.read(authStateProvider.notifier);
      await notifier.login('a@b.com', 'pass123');

      final state = container.read(authStateProvider);
      expect(state, isA<Authenticated>());
      expect((state as Authenticated).user.email, equals('a@b.com'));
    });

    test('login transitions to error on failure', () async {
      mockRepo.loginFunc = (_, __) async {
        throw InvalidCredentialsException();
      };

      final notifier = container.read(authStateProvider.notifier);
      await notifier.login('a@b.com', 'wrong');

      final state = container.read(authStateProvider);
      expect(state, isA<AuthError>());
    });
  });
}
