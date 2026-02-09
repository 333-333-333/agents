import 'package:fpdart/fpdart.dart';
import 'package:flutter_test/flutter_test.dart';

test('returns Right on success', () async {
  final result = await useCase.execute(validInput);

  expect(result.isRight(), isTrue);
  final value = result.getOrElse((_) => fail('Expected Right'));
  expect(value.id, isNotEmpty);
});

test('returns Left on validation failure', () async {
  final result = await useCase.execute(invalidInput);

  expect(result.isLeft(), isTrue);
  result.fold(
    (failure) => expect(failure, isA<ValidationFailure>()),
    (_) => fail('Expected Left'),
  );
});
