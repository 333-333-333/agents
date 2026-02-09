// features/booking/presentation/providers/booking_filter_provider.dart
// Sync Notifier — local state with mutations
import 'package:flutter_riverpod/flutter_riverpod.dart';

class BookingFilterNotifier extends Notifier<BookingFilter> {
  @override
  BookingFilter build() => const BookingFilter.defaultFilter();

  void setServiceType(ServiceType type) {
    state = state.copyWith(serviceType: type);
  }

  void setDateRange(DateTimeRange range) {
    state = state.copyWith(dateRange: range);
  }

  void reset() => state = const BookingFilter.defaultFilter();
}

final bookingFilterProvider =
    NotifierProvider<BookingFilterNotifier, BookingFilter>(
  BookingFilterNotifier.new,
);
