// test/features/auth/domain/value_objects/email_test.dart
import 'package:flutter_test/flutter_test.dart';
import 'package:bastet/features/auth/domain/value_objects/email.dart';

void main() {
  group('Email', () {
    test('accepts valid email', () {
      final email = Email('user@example.com');
      expect(email.value, equals('user@example.com'));
    });

    test('rejects email without @', () {
      expect(() => Email('invalid'), throwsA(isA<InvalidEmailException>()));
    });

    test('rejects empty email', () {
      expect(() => Email(''), throwsA(isA<InvalidEmailException>()));
    });

    test('trims whitespace', () {
      final email = Email('  user@example.com  ');
      expect(email.value, equals('user@example.com'));
    });

    test('equality by value', () {
      expect(Email('a@b.com'), equals(Email('a@b.com')));
    });
  });
}
