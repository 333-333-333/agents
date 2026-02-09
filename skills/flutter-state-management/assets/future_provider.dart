// FutureProvider — read-only async data with family modifier
import 'package:flutter_riverpod/flutter_riverpod.dart';

final caregiverDetailProvider =
    FutureProvider.family<Caregiver, String>((ref, caregiverId) async {
  final repo = ref.read(caregiverRepositoryProvider);
  return repo.getById(caregiverId);
});
