// Consuming async providers — ALWAYS handle all 3 states
import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';

// Option 1: .when() method
class PetListPage extends ConsumerWidget {
  const PetListPage({super.key});

  @override
  Widget build(BuildContext context, WidgetRef ref) {
    final petsAsync = ref.watch(petListProvider);

    return petsAsync.when(
      loading: () => const Center(child: AppLoadingSpinner()),
      error: (error, stack) => AppErrorView(
        message: error.toString(),
        onRetry: () => ref.invalidate(petListProvider),
      ),
      data: (pets) => ListView.builder(
        itemCount: pets.length,
        itemBuilder: (context, index) => PetCard(pet: pets[index]),
      ),
    );
  }
}

// Option 2: Dart 3.0 pattern matching
Widget buildWithPatternMatching(AsyncValue<List<Pet>> petsAsync) {
  return switch (petsAsync) {
    AsyncData(:final value) => PetListView(pets: value),
    AsyncError(:final error) => AppErrorView(message: '$error'),
    _ => const AppLoadingSpinner(),
  };
}
