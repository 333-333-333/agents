// autoDispose — cleans up state when no widget watches the provider
import 'package:flutter_riverpod/flutter_riverpod.dart';

final bookingDetailProvider =
    FutureProvider.autoDispose.family<Booking, String>((ref, id) async {
  final repo = ref.read(bookingRepositoryProvider);
  return repo.getById(id);
});
