// features/pet/presentation/providers/pet_list_provider.dart
// AsyncNotifier — async state with CRUD mutations (MOST COMMON pattern)
import 'dart:async';
import 'package:flutter_riverpod/flutter_riverpod.dart';

class PetListNotifier extends AsyncNotifier<List<Pet>> {
  @override
  FutureOr<List<Pet>> build() async {
    final repo = ref.read(petRepositoryProvider);
    return repo.getOwnerPets();
  }

  Future<void> addPet(CreatePetInput input) async {
    final repo = ref.read(petRepositoryProvider);
    state = const AsyncLoading();

    state = await AsyncValue.guard(() async {
      await repo.createPet(input);
      return repo.getOwnerPets(); // Refresh list
    });
  }

  Future<void> deletePet(PetId id) async {
    final repo = ref.read(petRepositoryProvider);

    state = await AsyncValue.guard(() async {
      await repo.deletePet(id);
      return repo.getOwnerPets();
    });
  }
}

final petListProvider =
    AsyncNotifierProvider<PetListNotifier, List<Pet>>(
  PetListNotifier.new,
);
