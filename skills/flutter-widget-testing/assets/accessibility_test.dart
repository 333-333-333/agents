// Verify semantics and accessibility
import 'package:flutter/material.dart';
import 'package:flutter_test/flutter_test.dart';

void main() {
  group('Accessibility', () {
    testWidgets('button has semantic label', (tester) async {
      await tester.pumpApp(AppButton(label: 'Continuar', onPressed: () {}));

      final semantics = tester.getSemantics(find.byType(AppButton));
      expect(semantics.label, contains('Continuar'));
    });

    testWidgets('meets minimum touch target', (tester) async {
      await tester.pumpApp(AppButton(label: 'Tap', onPressed: () {}));

      final size = tester.getSize(find.byType(AppButton));
      expect(size.width, greaterThanOrEqualTo(48));
      expect(size.height, greaterThanOrEqualTo(48));
    });

    testWidgets('text field announces error to screen reader', (tester) async {
      await tester.pumpApp(const AppEmailField(errorText: 'Campo requerido'));

      // Verify error is in the semantics tree
      expect(tester.getSemantics(find.text('Campo requerido')), isNotNull);
    });
  });
}
