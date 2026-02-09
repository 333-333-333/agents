// StreamProvider — real-time data from Firestore, WebSocket, etc.
import 'package:flutter_riverpod/flutter_riverpod.dart';

final chatMessagesProvider =
    StreamProvider.family<List<Message>, String>((ref, chatId) {
  final repo = ref.read(chatRepositoryProvider);
  return repo.watchMessages(chatId);
});
