// test/features/auth/application/use_cases/login_use_case_test.dart
import 'package:flutter_test/flutter_test.dart';
import 'package:bastet/features/auth/application/use_cases/login_use_case.dart';
import 'package:bastet/features/auth/domain/entities/user.dart';
import '../../helpers/mocks.dart';

void main() {
  late LoginUseCase useCase;
  late MockAuthRepository mockRepo;

  setUp(() {
    mockRepo = MockAuthRepository();
    useCase = LoginUseCase(repository: mockRepo);
  });

  group('LoginUseCase', () {
    test('returns user on successful login', () async {
      final expectedUser = User(id: '1', email: 'a@b.com', name: 'Test');
      mockRepo.loginFunc = (email, password) async => expectedUser;

      final result = await useCase.execute(
        email: 'a@b.com',
        password: 'pass123',
      );

      expect(result.isRight(), isTrue);
      result.fold(
        (failure) => fail('Expected success'),
        (user) => expect(user.email, equals('a@b.com')),
      );
    });

    test('returns failure on invalid credentials', () async {
      mockRepo.loginFunc = (email, password) async {
        throw InvalidCredentialsException();
      };

      final result = await useCase.execute(email: 'a@b.com', password: 'wrong');

      expect(result.isLeft(), isTrue);
    });
  });
}
