// Testing AsyncNotifierProvider
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:flutter_test/flutter_test.dart';

void main() {
  late ProviderContainer container;

  setUp(() {
    container = ProviderContainer(
      overrides: [
        bookingRepositoryProvider.overrideWithValue(MockBookingRepository()),
      ],
    );
  });

  tearDown(() => container.dispose());

  test('fetchBookings returns list', () async {
    // AsyncNotifier: listen for state changes
    final listener = Listener<AsyncValue<List<Booking>>>();
    container.listen(bookingsProvider, listener.call, fireImmediately: true);

    // Initial state is loading
    expect(container.read(bookingsProvider), isA<AsyncLoading>());

    // Wait for async resolution
    await container.read(bookingsProvider.future);

    // Now it should have data
    final state = container.read(bookingsProvider);
    expect(state, isA<AsyncData<List<Booking>>>());
  });
}

/// Helper to capture provider state changes.
class Listener<T> {
  final List<T> values = [];
  void call(T? previous, T next) => values.add(next);
}
