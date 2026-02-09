// pubspec.yaml dev_dependencies:
//   mocktail: ^1.0.4

import 'package:mocktail/mocktail.dart';
import 'package:bastet/features/auth/domain/repositories/auth_repository.dart';

class MockAuthRepository extends Mock implements AuthRepository {}

// In test:
void main() {
  late MockAuthRepository mockRepo;

  setUp(() {
    mockRepo = MockAuthRepository();
  });

  test('login calls repository', () async {
    when(
      () => mockRepo.login('a@b.com', 'pass123'),
    ).thenAnswer((_) async => testUser);

    final result = await mockRepo.login('a@b.com', 'pass123');
    expect(result, equals(testUser));
    verify(() => mockRepo.login('a@b.com', 'pass123')).called(1);
  });
}
