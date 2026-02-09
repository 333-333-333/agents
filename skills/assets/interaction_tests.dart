// Common interaction patterns

void main() {
  group('Interactions', () {
    // TAP
    testWidgets('tap a button', (tester) async {
      await tester.pumpApp(MyWidget());
      await tester.tap(find.byType(AppButton));
      await tester.pump(); // Process the tap

      expect(find.text('Tapped'), findsOneWidget);
    });

    // SCROLL
    testWidgets('scroll a list', (tester) async {
      await tester.pumpApp(MyScrollableWidget());
      await tester.drag(find.byType(ListView), const Offset(0, -300));
      await tester.pumpAndSettle();

      expect(find.text('Item 10'), findsOneWidget);
    });

    // ENTER TEXT
    testWidgets('type in a field', (tester) async {
      await tester.pumpApp(MyFormWidget());
      await tester.enterText(find.byType(AppTextField), 'Hello');
      await tester.pump();

      expect(find.text('Hello'), findsOneWidget);
    });

    // LONG PRESS
    testWidgets('long press shows menu', (tester) async {
      await tester.pumpApp(MyContextMenuWidget());
      await tester.longPress(find.text('Hold me'));
      await tester.pumpAndSettle();

      expect(find.text('Delete'), findsOneWidget);
    });
  });
}
