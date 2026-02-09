// Navigation commands available in widgets
import 'package:go_router/go_router.dart';

void navigationExamples(BuildContext context) {
  // Declarative — navigate by path
  context.go('/home');                          // Replace current stack
  context.push('/pet/123');                     // Push onto stack
  context.pop();                               // Go back

  // With query parameters
  context.go('/search?q=golden+retriever');

  // Named routes (optional, use sparingly)
  context.goNamed('petDetail', pathParameters: {'petId': '123'});
}
